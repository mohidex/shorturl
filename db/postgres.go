package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/mohidex/shorturl/config"
	"github.com/mohidex/shorturl/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbConf           *config.PostgresConfig
	postgresOnce     sync.Once
	postgresInstance *PostgresDB
)

func init() {
	dbConf = config.LoadPostgresConfig()
}

// PostgresDB represents a PostgreSQL client using GORM.
type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConf.Host, dbConf.User, dbConf.Password, dbConf.Name, dbConf.Port, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	fmt.Println("Connected PostgreSQL database successfully!")
	return &PostgresDB{
		DB: db,
	}, nil
}

func GetPostgresDB() (*PostgresDB, error) {
	var err error
	postgresOnce.Do(func() {
		postgresInstance, err = NewPostgresDB()
	})
	return postgresInstance, err
}

func (p *PostgresDB) PerformMigrations(model interface{}) error {
	// Apply your database migrations here.
	if err := p.DB.AutoMigrate(model); err != nil {
		return fmt.Errorf("failed to perform migrations: %w", err)
	}
	return nil
}

func (p *PostgresDB) SaveShortURL(ctx context.Context, shortURL *models.ShortURL) (*models.ShortURL, error) {
	if result := p.DB.Create(&shortURL); result.Error != nil {
		return &models.ShortURL{}, result.Error
	}
	return shortURL, nil
}

func (p *PostgresDB) GetLongURL(ctx context.Context, shortCode string) (string, error) {
	var url models.ShortURL
	if result := p.DB.Where("short_url=?", shortCode).First(&url); result.Error != nil {
		return "", result.Error
	}
	return url.DestUrl, nil

}

func (p *PostgresDB) SetLongURL(ctx context.Context, shortCode, longURL string) error {
	shortURL := models.ShortURL{
		ShortUrl: shortCode,
		DestUrl:  longURL,
	}
	_, err := p.SaveShortURL(ctx, &shortURL)
	return err
}
