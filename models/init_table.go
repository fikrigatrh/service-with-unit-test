package models

import "gorm.io/gorm"

func InitTable(db *gorm.DB) {
	//err := db.Debug().Migrator().DropTable(&KotaKab{}, &Provinsi{})
	//if err != nil {
	//	return
	//}
	err := db.Debug().AutoMigrate(&User{}, &AboutUsDb{}, &ExpeditionSchedule{}, &Auth{}, &Blog{}, &Provinsi{}, &KotaKab{})
	if err != nil {
		return
	}

}
