package models

import "gorm.io/gorm"

func InitTable(db *gorm.DB) {
	//err := db.Debug().Migrator().DropTable(&ExpeditionSchedule{})
	//if err != nil {
	//	return
	//}
	errs := db.Debug().AutoMigrate(&User{}, &AboutUsDb{}, &ExpeditionSchedule{}, &Auth{}, &Blog{}, &Provinsi{}, &KotaKab{})
	if errs != nil {
		return
	}

}
