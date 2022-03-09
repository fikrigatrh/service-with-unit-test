package expedition_schedule_rp

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/models/contract"
	"bitbucket.org/service-ekspedisi/repo"
	"errors"
	"gorm.io/gorm"
)

type ExpeditionRepoStruct struct {
	db  *gorm.DB
	log *log.LogCustom
}

func NewExpeditionRepo(db *gorm.DB, log *log.LogCustom) repo.ExpeditionRepoInterface {
	return &ExpeditionRepoStruct{db: db, log: log}
}

func (e ExpeditionRepoStruct) AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error) {
	tx := e.db.Begin()
	err := e.db.Debug().Create(&v).Error
	if err != nil {
		tx.Rollback()
		return models.ExpeditionSchedule{}, errors.New(contract.ErrCannotSaveToDB)
	}

	tx.Commit()
	return v, err
}

func (e ExpeditionRepoStruct) GetById(id int) (models.ExpeditionSchedule, error) {
	var v models.ExpeditionSchedule
	err := e.db.Debug().First(&v, id).Error
	if err != nil {
		return models.ExpeditionSchedule{}, errors.New(contract.ErrDataNotFound)
	}

	return v, err
}

func (e ExpeditionRepoStruct) GetAll() ([]models.ExpeditionSchedule, error) {
	var v []models.ExpeditionSchedule
	err := e.db.Debug().Find(&v).Error
	if err != nil {
		return []models.ExpeditionSchedule{}, err
	}

	return v, err

}

func (e ExpeditionRepoStruct) Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error) {
	tx := e.db.Begin()
	err := e.db.Debug().Model(&models.ExpeditionSchedule{}).Where("id = ?", id).Updates(v).Error
	if err != nil {
		tx.Rollback()
		return models.ExpeditionSchedule{}, err
	}

	tx.Commit()
	return v, err
}

func (e ExpeditionRepoStruct) DeleteData(id []string) error {
	var v models.ExpeditionSchedule

	tx := e.db.Begin()
	err := e.db.Debug().Model(&models.ExpeditionSchedule{}).Where("id in (?)", id).Delete(&v).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err

}

func (e ExpeditionRepoStruct) GetByRoute(rf, rd string) ([]models.ExpeditionSchedule, error) {
	var v []models.ExpeditionSchedule
	err := e.db.Debug().Where("route_from = ? and route_destination = ? and to_date(eta, 'YYYY-MM-DD') >= current_date", rf, rd).Find(&v).Error
	if err != nil {
		return []models.ExpeditionSchedule{}, err
	}

	return v, err
}

func (e ExpeditionRepoStruct) GetKotaById(id int) (models.KotaKab, error) {
	var v models.KotaKab
	err := e.db.Debug().Table("kota_kabs").First(&v, id).Error
	if err != nil {
		return models.KotaKab{}, errors.New(contract.ErrDataNotFound)
	}

	return v, err

}
