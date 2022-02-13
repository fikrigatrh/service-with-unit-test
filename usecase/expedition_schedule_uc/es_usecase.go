package expedition_schedule_uc

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/usecase"
)

type EsUcStruct struct {
	repo repo.ExpeditionRepoInterface
}

func NewEsUc(repo repo.ExpeditionRepoInterface) usecase.ExpeditionUcInterface {
	return &EsUcStruct{repo: repo}
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
