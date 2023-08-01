package runner

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/shorturl/config"
	"github.com/mohidex/shorturl/db"
	"github.com/mohidex/shorturl/handlers"
	"github.com/mohidex/shorturl/models"
)

var (
	pgInstance    *db.PostgresDB
	redisInstance *db.RedisDB
)

func init() {
	var err error
	pgInstance, err = db.GetPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	redisInstance, err = db.GetRedisDB()
	if err != nil {
		log.Fatal(err)
	}

	pgInstance.PerformMigrations(&models.ShortURL{})

}

func Run() {
	serverPort := config.GetEnv().ServerPort
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(handlers.HealthHandler)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		urlHandler := &handlers.ShortURLHandler{
			CacheDB:      redisInstance,
			PersistantDB: pgInstance,
		}

		v1.POST("/generate", urlHandler.APIGenerateShortUrl)
		v1.GET("/get/:url", urlHandler.GetShortUrl)
	}

	serverAddr := fmt.Sprintf(":%s", serverPort)
	router.Run(serverAddr)
}
