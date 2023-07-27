package main

import (
	"github.com/mohidex/shorturl/database"
	"github.com/mohidex/shorturl/models"
	"github.com/mohidex/shorturl/server"
)

func main() {
	AutoMigrate()
	server.Init()
}

func AutoMigrate() {
	pgClient, err := database.GetPostgresClient()
	if err != nil {
		panic(err)
	}
	pgClient.PerformMigrations(&models.ShortUrl{})
}
