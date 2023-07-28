package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/shorturl/db"
	"github.com/mohidex/shorturl/models"
	"github.com/mohidex/shorturl/utils"
	"gorm.io/gorm"
)

type ShortURLHandler struct {
	RedisDB *db.RedisDB
	PgDB    *db.PostgresDB
}

func (h *ShortURLHandler) APIGenerateShortUrl(c *gin.Context) {
	var input models.UrlInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	originalUrl := input.OriginalUrl
	shortUrl, err := utils.GenerateShortLink(originalUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	url := models.ShortURL{
		ShortUrl: shortUrl,
		DestUrl:  input.OriginalUrl,
	}

	err = h.PgDB.SetLongURL(context.Background(), url.ShortUrl, url.DestUrl)

	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": url,
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

	if longUrl, err := h.RedisDB.GetLongURL(context.Background(), shortCode); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"shortUrl": shortCode,
			"fullUrl":  longUrl,
		})
		return
	}

	longUrl, err := h.PgDB.GetLongURL(context.Background(), shortCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	// since url found in the database which is not in cache. So setting it in cache
	_ = h.RedisDB.SetLongURL(context.Background(), shortCode, longUrl)

	c.JSON(http.StatusOK, gin.H{
		"shortUrl": shortCode,
		"fullUrl":  longUrl,
	})
}
