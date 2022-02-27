package user_repo

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"gorm.io/gorm"
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
	err := a.db.Debug().Find(&v).Error
	if err != nil {
		return nil, err
	}
	return v, err
}

func (a UserRepoStruct) GetById(id int) (models.User, error) {
	var v models.User
	err := a.db.Debug().Where("id = ?", id).First(&v).Error
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
	v := models.User{}

	tx := a.db.Begin()
	err := a.db.Debug().Where("id in (?)", id).Delete(&v).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}
