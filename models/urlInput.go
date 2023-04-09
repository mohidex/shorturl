package models

type UrlInput struct {
	OriginalUrl string `json:"original_url" binding:"required"`
}
