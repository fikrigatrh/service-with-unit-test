package usecase

import "bitbucket.org/service-ekspedisi/models"

type AboutUsUcInterface interface {
	AddAbout(v models.AboutUs) (models.AboutUs, error)
	GetAll() ([]models.AboutUs, error)
	GetById(id int) (models.AboutUs, error)
	UpdateData(id int, v models.AboutUs) (models.AboutUs, error)
	DeleteData(id []string) error
}

type ExpeditionUcInterface interface {
	AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	GetById(id int) (models.ExpeditionSchedule, error)
	GetAll() ([]models.ExpeditionSchedule, error)
	Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	DeleteData(id []string) error
}

type UserUcInterface interface {
	AddUser(v models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	UpdateData(id int, v models.User) (models.User, error)
	DeleteData(id []string) error
}