package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string `json:"title" validate:"required"`
	UrlImage string `json:"url_image" validate:"required"`
	Content  string `json:"content" validate:"required"`
}
