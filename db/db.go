package db

import (
	"context"

	"github.com/mohidex/shorturl/models"
)

type ShortURLDB interface {
	SetLongURL(ctx context.Context, shortCode, longURL string) error
	GetLongURL(ctx context.Context, shortCode string) (string, error)
}

type ShortURLPersistantDB interface {
	SaveShortURL(ctx context.Context, shortURL *models.ShortURL) (*models.ShortURL, error)
	SetLongURL(ctx context.Context, shortCode, longURL string) error
	GetLongURL(ctx context.Context, shortCode string) (string, error)
}
