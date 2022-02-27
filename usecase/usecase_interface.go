package usecase

import "bitbucket.org/service-ekspedisi/models"

type AboutUsUcInterface interface {
	AddAbout(v models.AboutUsRequest) (models.AboutUsDb, error)
	GetAll() ([]models.AboutUsRequest, error)
	GetById(id int) (models.AboutUsRequest, error)
	UpdateData(id int, v models.AboutUsRequest) (models.AboutUsRequest, error)
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

type ErrorHandlerUsecase interface {
	ResponseError(error interface{}) (int, interface{})
	ValidateRequest(error interface{}) (string, error)
}
