package expedition_schedule_uc

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/usecase"
)

type EsUcStruct struct {
	repo repo.ExpeditionRepoInterface
	log  *log.LogCustom
}

func NewEsUc(repo repo.ExpeditionRepoInterface, log *log.LogCustom) usecase.ExpeditionUcInterface {
	return &EsUcStruct{
		repo: repo,
		log:  log,
	}
}

func (e EsUcStruct) AddEs(v models.ExpeditionSchedule) (models.ExpeditionSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (e EsUcStruct) GetById(id int) (models.ExpeditionSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (e EsUcStruct) GetAll() ([]models.ExpeditionSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (e EsUcStruct) Update(id int, v models.ExpeditionSchedule) (models.ExpeditionSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (e EsUcStruct) DeleteData(id []string) error {
	//TODO implement me
	panic("implement me")
}
