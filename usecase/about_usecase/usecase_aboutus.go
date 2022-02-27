package about_usecase

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/usecase"
	"strconv"
)

type AboutUsUsecaseStruct struct {
	repo repo.AboutUsRepoInterface
	log  *log.LogCustom
}

func NewAboutUsUsecase(repo repo.AboutUsRepoInterface, log *log.LogCustom) usecase.AboutUsUcInterface {
	return &AboutUsUsecaseStruct{
		repo: repo,
		log:  log,
	}
}

func (a AboutUsUsecaseStruct) AddAbout(v models.AboutUsRequest) (models.AboutUsRequest, error) {
	about, err := a.repo.AddAbout(v)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	return about, nil
}

func (a AboutUsUsecaseStruct) GetAll() ([]models.AboutUsRequest, error) {
	about, err := a.repo.GetAll()
	if err != nil {
		return []models.AboutUsRequest{}, err
	}

	return about, nil
}

func (a AboutUsUsecaseStruct) GetById(id int) (models.AboutUsRequest, error) {
	about, err := a.repo.GetById(id)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	return about, nil
}

func (a AboutUsUsecaseStruct) UpdateData(id int, v models.AboutUsRequest) (models.AboutUsRequest, error) {
	about, err := a.repo.UpdateData(id, v)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	return about, nil
}

func (a AboutUsUsecaseStruct) DeleteData(id []string) error {
	for _, s := range id {
		idRes, _ := strconv.Atoi(s)
		_, err := a.repo.GetById(idRes)
		if err != nil {
			a.log.Error(err, "usecase error when get data by id", "", nil, idRes, nil)
			return err
		}
	}

	err := a.repo.DeleteData(id)
	if err != nil {
		return err
	}

	return nil
}
