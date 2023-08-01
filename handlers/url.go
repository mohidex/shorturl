package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/shorturl/db"
	"github.com/mohidex/shorturl/models"
	"github.com/mohidex/shorturl/utils"
	"golang.org/x/sync/errgroup"
)

type ShortURLHandler struct {
	CacheDB      db.ShortURLDB
	PersistantDB db.ShortURLPersistantDB
}

func (h *ShortURLHandler) APIGenerateShortUrl(c *gin.Context) {
	var input models.UrlInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	longURL := input.OriginalUrl
	shortUrl, err := utils.GenerateShortLink(longURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	url := models.ShortURL{
		ShortUrl: shortUrl,
		DestUrl:  longURL,
	}

	// Create a bool channel to receive the result of the PostgreSQL save operation
	pgSaveResult := make(chan bool, 1)

	// Use Goroutines to save to both Redis and PostgreSQL simultaneously
	ctx := context.Background()

	go func() {
		if err := h.PersistantDB.SetLongURL(ctx, url.ShortUrl, url.DestUrl); err != nil {
			log.Printf("Failed to save short URL to PostgreSQL: %s\n", err)
			pgSaveResult <- false // Signal failure to the bool channel
			return
		}
		pgSaveResult <- true // Signal success to the bool channel
	}()

	// Wait for the PostgreSQL save operation result
	if success := <-pgSaveResult; success {
		// PostgreSQL save operation was successful, proceed to save to Redis
		go func() {
			if err := h.CacheDB.SetLongURL(ctx, url.ShortUrl, url.DestUrl); err != nil {
				log.Printf("Failed to save long URL to Redis: %s\n", err)
			}
		}()
	} else {
		// PostgreSQL save operation failed, handle the error here if needed
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url":  url.DestUrl,
		"code": url.ShortUrl,
	})
}

func (h *ShortURLHandler) GetShortUrl(c *gin.Context) {
	shortCode := c.Param("url")

	if utils.EmptyString(shortCode) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The link you given is broken",
		})
		return
	}

	// Use an errgroup to handle multiple goroutines and their potential errors
	g, ctx := errgroup.WithContext(c.Request.Context())

	// Create a variable to store the fetched long URL
	var longURL string

	// Try fetching the long URL from Redis first
	g.Go(func() error {
		var err error
		longURL, err = h.CacheDB.GetLongURL(ctx, shortCode)
		return err
	})

	// Wait for the result from Redis cache response
	if err := g.Wait(); err != nil {
		// Redis query failed or URL not found in Redis, try fetching from postgreSQL database
		var err error
		longURL, err = h.PersistantDB.GetLongURL(context.Background(), shortCode)
		if err != nil {
			// If URL not found in PostgreSQL as well, return a 404
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		}
		// since url found in the database which is not in cache. So setting it in cache
		_ = h.CacheDB.SetLongURL(context.Background(), shortCode, longURL)

	}

	// Return the long URL to the client
	c.Redirect(http.StatusFound, longURL)
}
