package database

import (
	"fmt"
	"sync"

	"github.com/mohidex/shorturl/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbConf           *config.PostgresConfig
	postgresOnce     sync.Once
	postgresInstance *PostgresClient
)

func init() {
	dbConf = config.LoadPostgresConfig()
}

// PostgresClient represents a PostgreSQL client using GORM.
type PostgresClient struct {
	DB *gorm.DB
}

func NewPostgresClient() (*PostgresClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConf.Host, dbConf.User, dbConf.Password, dbConf.Name, dbConf.Port, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	fmt.Println("Connected PostgreSQL database successfully!")
	return &PostgresClient{
		DB: db,
	}, nil
}

func GetPostgresClient() (*PostgresClient, error) {
	var err error
	postgresOnce.Do(func() {
		postgresInstance, err = NewPostgresClient()
	})
	return postgresInstance, err
}

func (pgc *PostgresClient) PerformMigrations(model interface{}) error {
	// Apply your database migrations here.
	if err := pgc.DB.AutoMigrate(model); err != nil {
		return fmt.Errorf("failed to perform migrations: %w", err)
	}
	return nil
}
