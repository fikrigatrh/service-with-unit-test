package models

import "gorm.io/gorm"

func InitTable(db *gorm.DB) {
	//err := db.Debug().Migrator().DropTable(&User{}, &Logs{})
	//if err != nil {
	//	return
	//}
	errs := db.Debug().AutoMigrate(&Users{})
	if errs != nil {
		return
	}

}
