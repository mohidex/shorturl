package db

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
	redisInstance *RedisDB
)

// RedisDB represents a Redis client.
type RedisDB struct {
	client *redis.Client
}

// init loads the configuration once during package initialization.
func init() {
	conf = config.LoadRedisConfig()
}

// NewRedisDB creates a new RedisDB instance with the provided Redis configuration.
func NewRedisDB() (*RedisDB, error) {
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

	return &RedisDB{
		client: client,
	}, nil
}

// GetRedisDB returns the singleton RedisDB instance.
func GetRedisDB() (*RedisDB, error) {
	var err error
	redisOnce.Do(func() {
		redisInstance, err = NewRedisDB()
	})
	return redisInstance, err
}

// Set sets a key-value pair in Redis with an expiration time.
func (r *RedisDB) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value from Redis using the given key.
func (r *RedisDB) GetLongURL(ctx context.Context, shortCode string) (string, error) {
	select {
	case <-ctx.Done():
		// If the context is done, return immediately.
		return "", ctx.Err()
	default:
		// Get Full Url from redis by short urlCode
		return r.client.Get(ctx, shortCode).Result()
	}
}

// Set sets a shortCode-longURL pair in Redis with an expiration time.
func (r *RedisDB) SetLongURL(ctx context.Context, shortCode, longURL string) error {
	select {
	case <-ctx.Done():
		// If the context is done, return immediately.
		return ctx.Err()
	default:
		// Set urlCode and corresponding FullURL pair to Redis
		return r.Set(ctx, shortCode, longURL, 30*time.Minute)
	}
}
