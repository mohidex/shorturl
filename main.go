package main

import (
	"github.com/mohidex/shorturl/database"
	"github.com/mohidex/shorturl/models"
	"github.com/mohidex/shorturl/server"
)

func main() {
	database.InitDB()
	AutoMigrate()
	server.Init()
}

func AutoMigrate() {
	db := database.GetDB()
	db.AutoMigrate(&models.ShortUrl{})
}
