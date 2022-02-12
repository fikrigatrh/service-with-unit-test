package about_us

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"gorm.io/gorm"
)

type AboutUsRepoStruct struct {
	db *gorm.DB
}

func NewAboutUsRepo(db *gorm.DB) repo.AboutUsRepoInterface {
	return &AboutUsRepoStruct{db}
}

func (a AboutUsRepoStruct) AddAbout(v models.AboutUs) (models.AboutUs, error) {
	tx := a.db.Begin()
	err := a.db.Debug().Create(&v).Error
	if err != nil {
		tx.Rollback()
		return models.AboutUs{}, err
	}

	tx.Commit()
	return v, err
}

func (a AboutUsRepoStruct) GetAll() ([]models.AboutUs, error) {
	var v []models.AboutUs
	err := a.db.Debug().Find(&v).Error
	if err != nil {
		return nil, err
	}
	return v, err
}

func (a AboutUsRepoStruct) GetById(id int) (models.AboutUs, error) {
	var v models.AboutUs
	err := a.db.Debug().Where("id = ?", id).First(&v).Error
	if err != nil {
		return models.AboutUs{}, err
	}
	return v, err
}

func (a AboutUsRepoStruct) UpdateData(id int, v models.AboutUs) (models.AboutUs, error) {
	tx := a.db.Begin()
	err := a.db.Debug().Model(&models.AboutUs{}).Where("id = ?", id).Updates(v).Error
	if err != nil {
		tx.Rollback()
		return models.AboutUs{}, err
	}

	tx.Commit()
	return v, err

}

func (a AboutUsRepoStruct) DeleteData(id []string) error {
	v := models.AboutUs{}

	tx := a.db.Begin()
	err := a.db.Debug().Where("id in (?)", id).Delete(&v).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}
