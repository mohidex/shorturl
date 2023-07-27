// config/config.go

package config

import (
	"os"
)

// Config holds the configuration data for the application.
type RedisConfig struct {
	RedisAddr     string
	RedisPassword string
	// Add other configuration variables as needed
}

// LoadConfig loads the configuration from environment variables.
func LoadRedisConfig() *RedisConfig {
	return &RedisConfig{
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASS", ""),
		// Add other configuration variables as needed
	}
}

// getEnv returns the value of an environment variable or a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
