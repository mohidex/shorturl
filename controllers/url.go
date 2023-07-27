package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/shorturl/database"
	"github.com/mohidex/shorturl/models"
	"github.com/mohidex/shorturl/services"
	"github.com/mohidex/shorturl/utils"
	"gorm.io/gorm"
)

type UrlController struct{}

func (uc *UrlController) APIGenerateShortUrl(c *gin.Context) {
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

	url := models.ShortUrl{
		ShortUrl: shortUrl,
		DestUrl:  input.OriginalUrl,
	}
	pgClient, _ := database.GetPostgresClient()

	savedUrl, err := url.Save(pgClient)
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
		"url": savedUrl,
	})
}

func (uc *UrlController) GetShortUrl(c *gin.Context) {
	shortUrl := c.Param("url")

	if utils.EmptyString(shortUrl) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The link you given is broken",
		})
		return
	}

	destUrl, err := services.Search4ShortUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortUrl": shortUrl,
		"fullUrl":  destUrl,
	})
}
