package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string `json:"title"`
	UrlImage string `json:"url_image"`
	Content  string `json:"content"`
}
