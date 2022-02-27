package models

import "gorm.io/gorm"

func InitTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&User{}, &AboutUsDb{})
	if err != nil {
		return
	}

}
