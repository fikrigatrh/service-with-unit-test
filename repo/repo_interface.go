package repo

import "bitbucket.org/service-ekspedisi/models"

type AboutUsRepoInterface interface {
	AddAbout(v models.AboutUsDb) (models.AboutUsDb, error)
	GetAll() ([]models.AboutUsRequest, error)
	GetById(id int) (models.AboutUsRequest, error)
	UpdateData(id int, v models.AboutUsRequest) (models.AboutUsRequest, error)
	DeleteData(id []string) error
}

type ExpeditionRepoInterface interface {
	AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	GetById(id int) (models.ExpeditionSchedule, error)
	GetAll() ([]models.ExpeditionSchedule, error)
	Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	DeleteData(id []string) error
	GetByRoute(string string) ([]models.ExpeditionSchedule, error)
}

type UserRepoInterface interface {
	AddUser(v models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	UpdateData(id int, v models.User) (models.User, error)
	DeleteData(id []string) error
}
