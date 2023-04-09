package services

import (
	"context"
	"log"
	"time"

	"github.com/mohidex/shorturl/database"
	"github.com/mohidex/shorturl/models"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

func Search4ShortUrl(shortUrl string) (string, error) {
	val, err := RedisGetString(shortUrl)

	if err == redis.Nil {
		log.Println("Not found in Redis/Cache")
		destUrl, err := FindUrlFromDB(shortUrl)
		if err == nil {
			_ = RedisSetString(shortUrl, destUrl)
			return destUrl, err
		}
	}
	return val, err
}

func RedisGetString(key string) (string, error) {
	rdb := database.GetRedis()
	val, err := rdb.Get(ctx, key).Result()
	return val, err
}

func RedisSetString(key, val string) error {
	rdb := database.GetRedis()
	if err := rdb.Set(ctx, key, val, 30*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func FindUrlFromDB(key string) (string, error) {
	var url models.ShortUrl
	db := database.GetDB()

	if result := db.Where("short_url=?", key).First(&url); result.Error != nil {
		return "", result.Error
	}
	return url.DestUrl, nil
}
