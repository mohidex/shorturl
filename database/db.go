package database

import (
	"context"
	"fmt"
	"log"

	"github.com/mohidex/shorturl/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
)

func InitDB() {
	dbConfig := config.GetEnv().DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)

	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err == nil {
		DB = db
		log.Println("Successfully connected to the database")
	} else {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return DB
}

func InitRedis() {

	redisConfig := config.GetEnv().Redis

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.RedisDb,
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

	if err := RedisClient.Ping(context.TODO()).Err(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected to Redis")
	}
}

func GetRedis() *redis.Client {
	return RedisClient
}
