package models

import "gorm.io/gorm"

func InitTable(db *gorm.DB) {
	db.Debug().AutoMigrate(&User{})
	db.Debug().AutoMigrate(&Auth{})
}