package about_usecase

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/usecase"
)

type AboutUsUsecaseStruct struct {
	repo repo.AboutUsRepoInterface
}

func NewAboutUsUsecase(repo repo.AboutUsRepoInterface) usecase.AboutUsUcInterface {
	return &AboutUsUsecaseStruct{
		repo: repo,
	}
}

func (a AboutUsUsecaseStruct) AddAbout(v models.AboutUs) (models.AboutUs, error) {
	//TODO implement me
	panic("implement me")
}

func (a AboutUsUsecaseStruct) GetAll() ([]models.AboutUs, error) {
	//TODO implement me
	panic("implement me")
}

func (a AboutUsUsecaseStruct) GetById(id int) (models.AboutUs, error) {
	//TODO implement me
	panic("implement me")
}

func (a AboutUsUsecaseStruct) UpdateData(id int, v models.AboutUs) (models.AboutUs, error) {
	//TODO implement me
	panic("implement me")
}

func (a AboutUsUsecaseStruct) DeleteData(id []string) error {
	//TODO implement me
	panic("implement me")
}
