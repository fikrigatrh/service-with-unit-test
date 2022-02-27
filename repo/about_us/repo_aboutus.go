package about_us

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/models/contract"
	"bitbucket.org/service-ekspedisi/repo"
	"errors"
	"gorm.io/gorm"
)

type AboutUsRepoStruct struct {
	db  *gorm.DB
	log *log.LogCustom
}

func NewAboutUsRepo(db *gorm.DB, log *log.LogCustom) repo.AboutUsRepoInterface {
	return &AboutUsRepoStruct{db, log}
}

func (a AboutUsRepoStruct) AddAbout(v models.AboutUsDb) (models.AboutUsDb, error) {
	tx := a.db.Begin()
	err := a.db.Debug().Create(&v).Error
	if err != nil {
		tx.Rollback()
		return models.AboutUsDb{}, errors.New(contract.ErrCannotSaveToDB)
	}

	tx.Commit()
	return v, err
}

func (a AboutUsRepoStruct) GetAll() ([]models.AboutUsRequest, error) {
	var v []models.AboutUsRequest
	err := a.db.Debug().Find(&v).Error
	if err != nil {
		return nil, err
	}
	return v, err
}

func (a AboutUsRepoStruct) GetById(id int) (models.AboutUsRequest, error) {
	var v models.AboutUsRequest
	err := a.db.Debug().Where("id = ?", id).First(&v).Error
	if err != nil {
		return models.AboutUsRequest{}, errors.New(contract.ErrDataNotFound)
	}
	return v, err
}

func (a AboutUsRepoStruct) UpdateData(id int, v models.AboutUsRequest) (models.AboutUsRequest, error) {
	tx := a.db.Begin()
	err := a.db.Debug().Model(&models.AboutUsRequest{}).Where("id = ?", id).Updates(v).Error
	if err != nil {
		tx.Rollback()
		return models.AboutUsRequest{}, err
	}

	tx.Commit()
	return v, err

}

func (a AboutUsRepoStruct) DeleteData(id []string) error {
	v := models.AboutUsRequest{}

	tx := a.db.Begin()
	err := a.db.Debug().Where("id in (?)", id).Delete(&v).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}
