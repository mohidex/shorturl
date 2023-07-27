package database

import (
	"fmt"
	"log"

	"github.com/mohidex/shorturl/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
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
