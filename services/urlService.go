package services

import (
	"log"
	"time"

	"github.com/mohidex/shorturl/database"
	"github.com/mohidex/shorturl/models"
	"github.com/redis/go-redis/v9"
)

func Search4ShortUrl(shortUrl string) (string, error) {

	redisClient, err := database.GetRedisClient()
	if err != nil {
		log.Panic("Error getting Redis client:", err)
	}

	val, err := redisClient.Get(shortUrl)

	if err == redis.Nil {
		log.Println("Not found in Redis/Cache")
		destUrl, err := FindUrlFromDB(shortUrl)
		if err == nil {
			_ = redisClient.Set(shortUrl, destUrl, 30*time.Minute)
			return destUrl, err
		}
	}
	return val, err
}

func FindUrlFromDB(key string) (string, error) {
	var url models.ShortUrl
	pgClient, err := database.GetPostgresClient()
	if err != nil {
		panic(err)
	}

	if result := pgClient.DB.Where("short_url=?", key).First(&url); result.Error != nil {
		return "", result.Error
	}
	return url.DestUrl, nil
}
