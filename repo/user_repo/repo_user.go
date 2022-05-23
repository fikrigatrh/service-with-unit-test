package user_repo

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"gorm.io/gorm"
	"time"
)

type UserRepoStruct struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepoInterface {
	return &UserRepoStruct{db}
}

func (a UserRepoStruct) AddUser(v models.User) (models.User, error) {
	tx := a.db.Begin()
	err := a.db.Debug().Create(&v).Error
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	return v, err
}

func (a UserRepoStruct) GetAll() ([]models.User, error) {
	var v []models.User
	err := a.db.Debug().Where("status = active").Find(&v).Error
	if err != nil {
		return nil, err
	}
	return v, err
}

func (a UserRepoStruct) GetById(id int) (models.User, error) {
	var v models.User
	err := a.db.Debug().Where("id = ? and status = active", id).First(&v).Error
	if err != nil {
		return models.User{}, err
	}
	return v, err
}

func (a UserRepoStruct) UpdateData(id int, v models.User) (models.User, error) {
	tx := a.db.Begin()
	err := a.db.Debug().Model(&models.User{}).Where("id = ?", id).Updates(v).Error
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	return v, err

}

func (a UserRepoStruct) DeleteData(id []string) error {
	currentTime := time.Now()

	tx := a.db.Begin()
	err := a.db.Debug().Model(&models.User{}).Where("id in (?)", id).Updates(map[string]interface{}{"status": "inactive", "deleted_at": currentTime.Format("2006-01-02 15:04:05")}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}
