package repo

import "bitbucket.org/service-ekspedisi/models"

type AboutUsRepoInterface interface {
	AddAbout(v models.AboutUs) (models.AboutUs, error)
	GetAll() ([]models.AboutUs, error)
	GetById(id int) (models.AboutUs, error)
	UpdateData(id int, v models.AboutUs) (models.AboutUs, error)
	DeleteData(id []string) error
}
