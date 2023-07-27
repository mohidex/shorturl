package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mohidex/shorturl/config"
	"github.com/redis/go-redis/v9"
)

var (
	conf          *config.RedisConfig
	redisOnce     sync.Once
	redisInstance *RedisClient
)

// RedisClient represents a Redis client.
type RedisClient struct {
	client *redis.Client
}

// init loads the configuration once during package initialization.
func init() {
	conf = config.LoadRedisConfig()
}

// NewRedisClient creates a new RedisClient instance with the provided Redis configuration.
func NewRedisClient() (*RedisClient, error) {
	addr := conf.RedisAddr
	password := conf.RedisPassword
	db := 0 // Default to DB 0

	options := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}

	client := redis.NewClient(options)

	// Test the Redis connection
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	fmt.Println("Successfully connected to Redis!")

	return &RedisClient{
		client: client,
	}, nil
}

// GetRedisClient returns the singleton RedisClient instance.
func GetRedisClient() (*RedisClient, error) {
	var err error
	redisOnce.Do(func() {
		redisInstance, err = NewRedisClient()
	})
	return redisInstance, err
}

// Get retrieves a value from Redis using the given key.
func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

// Set sets a key-value pair in Redis with an expiration time.
func (r *RedisClient) Set(key, value string, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}
