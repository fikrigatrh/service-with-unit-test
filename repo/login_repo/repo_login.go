package login_repo

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"fmt"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type LoginRepoStruct struct {
	db *gorm.DB
}

func NewLoginRepo(db *gorm.DB) repo.LoginRepoInterface {
	return &LoginRepoStruct{db}
}

func (a LoginRepoStruct) LoginUser(email string) (models.User, error) {
	var data models.User
	tx := a.db.Begin()
	err := a.db.Debug().Where("email = ? and status = 'active'",email).Find(&data).Error
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()

	fmt.Println("INI DATA NYA HEHE",data)
	return data, err
}

func (a LoginRepoStruct) CreateAuth(authFix models.Auth) (models.Auth, error) {
	authFix.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	tx := a.db.Begin()
	err := a.db.Debug().Create(&authFix).Error
	if err != nil {
		tx.Rollback()
		return models.Auth{}, err
	}
	return authFix, nil
}