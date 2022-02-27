package expedition_schedule_uc

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/usecase"
	"strconv"
	"strings"
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
	v.Route = strings.ToUpper(v.Route)
	es, err := e.repo.AddEs(v)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}

	return es, nil
}

func (e EsUcStruct) GetById(id int) (models.ExpeditionSchedule, error) {
	es, err := e.repo.GetById(id)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}

	return es, nil
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
	for _, s := range id {
		idRes, _ := strconv.Atoi(s)
		_, err := e.repo.GetById(idRes)
		if err != nil {
			e.log.Error(err, "usecase error when get data by id", "", nil, idRes, nil)
			return err
		}
	}

	err := e.repo.DeleteData(id)
	if err != nil {
		return err
	}

	return nil
}

func (e EsUcStruct) GetByRoute(route string) ([]models.ExpeditionSchedule, error) {
	es, err := e.repo.GetByRoute(route)
	if err != nil {
		return []models.ExpeditionSchedule{}, err
	}

	return es, nil
}
