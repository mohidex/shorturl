package models

import (
	"reflect"

	"github.com/mohidex/shorturl/database"
	"gorm.io/gorm"
)

type ShortUrl struct {
	gorm.Model
	ShortUrl string `gorm:"size:255;not null;unique" json:"short_url"`
	DestUrl  string `gorm:"size:512;not null;unique" json:"dest_url"`
	Valid    bool   `gorm:"not null;default:true" json:"is_valid"`
}

func (url *ShortUrl) IsEmpty() bool {
	return reflect.DeepEqual(url, ShortUrl{})
}

func (url *ShortUrl) Save() (*ShortUrl, error) {
	db := database.GetDB()
	if result := db.Create(&url); result.Error != nil {
		return &ShortUrl{}, result.Error
	}
	return url, nil
}
