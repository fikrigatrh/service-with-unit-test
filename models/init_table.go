package models

import "gorm.io/gorm"

func InitTable(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&User{}, &AboutUsDb{}, &ExpeditionSchedule{})
	if err != nil {
		return
	}

}
