package models

import (
	"reflect"

	"gorm.io/gorm"
)

type ShortURL struct {
	gorm.Model
	ShortUrl string `gorm:"size:255;not null;unique" json:"short_url"`
	DestUrl  string `gorm:"size:512;not null;unique" json:"dest_url"`
	Valid    bool   `gorm:"not null;default:true" json:"is_valid"`
}

func (url *ShortURL) IsEmpty() bool {
	return reflect.DeepEqual(url, ShortURL{})
}
