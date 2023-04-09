package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohidex/shorturl/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		url := new(controllers.UrlController)
		v1.POST("/generate", url.APIGenerateShortUrl)
		v1.GET("/get/:url", url.GetShortUrl)
	}
	return router

}
