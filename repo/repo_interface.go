package repo

import "bitbucket.org/service-ekspedisi/models"

type AboutUsRepoInterface interface {
	AddAbout(v models.AboutUs) (models.AboutUs, error)
	GetAll() ([]models.AboutUs, error)
	GetById(id int) (models.AboutUs, error)
	UpdateData(id int, v models.AboutUs) (models.AboutUs, error)
	DeleteData(id []string) error
}

type ExpeditionRepoInterface interface {
	AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	GetById(id int) (models.ExpeditionSchedule, error)
	GetAll() ([]models.ExpeditionSchedule, error)
	Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error)
	DeleteData(id []string) error
}
