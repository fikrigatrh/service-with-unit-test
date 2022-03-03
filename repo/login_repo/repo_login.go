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

//Once a user row in the auth table
func (a LoginRepoStruct) DeleteAuthData(givenUuid string) (int, error) {
	au := &models.Auth{}
	deleted := a.db.Debug().Where("auth_uuid = ?", givenUuid).Delete(&au)
	if deleted.Error != nil {
		return 0, deleted.Error
	}
	fmt.Println("Delete data from database success")
	return 0, nil
}

func (a LoginRepoStruct) GetAuthByEmailAndAuthID(email string, authUUID string) (*models.Auth, error) {
	data := models.Auth{}
	err := a.db.Debug().Raw("SELECT * FROM public.auths where email = ? and auth_uuid = ? ;", email, authUUID).Scan(&data).Error
	if err != nil {
		fmt.Println("[LoginStruct.GetAuthByUsernameAndAuthID] Error when GetAuthByEmailAndAuthID Repo")
		return nil, err
	}

	return &data, nil
}