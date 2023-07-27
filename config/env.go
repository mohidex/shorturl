package config

import "os"

var env = Env{
	Debug:        true,
	ServerPort:   "5000",
	MaxIdleConns: 50,
	MaxOpenConns: 100,

	DB: DatabaseConfig{
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	},
}
