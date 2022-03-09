package expedition_schedule_uc

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/models/contract"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/usecase"
	"bitbucket.org/service-ekspedisi/utils"
	"errors"
	"strconv"
	"strings"
	"time"
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
	var tempRoute string
	arr := []string{v.RouteFrom, v.RouteDestination}
	for _, val := range arr {
		atoi, err := strconv.Atoi(val)
		if err != nil {
			return models.ExpeditionSchedule{}, errors.New(contract.ErrBadRequest)
		}
		res, err := e.repo.GetKotaById(atoi)
		if err != nil {
			return models.ExpeditionSchedule{}, err
		}
		tempRoute += res.Nama + "-"
	}

	routeArr := strings.Split(tempRoute, "-")
	v.Route = routeArr[0] + "-" + routeArr[1]

	if !utils.CheckDate(v.Eta, v.Etd, v.Closing) {
		return models.ExpeditionSchedule{}, errors.New(contract.ErrBadRequest)
	}

	layoutISO := "2006-01-02"
	closingTemp, err := time.Parse(layoutISO, v.Closing)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}
	etdTemp, err := time.Parse(layoutISO, v.Etd)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}
	etaTemp, err := time.Parse(layoutISO, v.Eta)
	if err != nil {
		return models.ExpeditionSchedule{}, err
	}

	if closingTemp.Unix() > etdTemp.Unix() || etdTemp.Unix() > etaTemp.Unix() || closingTemp.Unix() > etaTemp.Unix() {
		return models.ExpeditionSchedule{}, errors.New(contract.ErrBadRequest)
	}

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
	all, err := e.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for i, _ := range all {
		all[i].RouteFrom = ""
		all[i].RouteDestination = ""

		monthEta := utils.ConvertMonth(all[i].Eta)
		etaTemp := strings.Split(all[i].Eta, "-")
		str := etaTemp[2] + "-" + monthEta + "-" + etaTemp[0]
		etaTempArr := strings.Split(str, "-")
		eta := strings.Join(etaTempArr, " ")
		all[i].Eta = eta

		monthEtd := utils.ConvertMonth(all[i].Etd)
		etdTemp := strings.Split(all[i].Etd, "-")
		strEtd := etdTemp[2] + "-" + monthEtd + "-" + etdTemp[0]
		etdTempArr := strings.Split(strEtd, "-")
		etd := strings.Join(etdTempArr, " ")
		all[i].Etd = etd

		monthClosing := utils.ConvertMonth(all[i].Closing)
		closingTemp := strings.Split(all[i].Closing, "-")
		strClosing := closingTemp[2] + "-" + monthClosing + "-" + closingTemp[0]
		cosingTempArr := strings.Split(strClosing, "-")
		closing := strings.Join(cosingTempArr, " ")
		all[i].Closing = closing
		//etdTemp := strings.Split(all[i].Etd, "-")
		//closingTemp := strings.Split(all[i].Closing, "-")
	}

	return all, nil
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
	routeTemp := strings.ToUpper(route)
	routeArr := strings.Split(routeTemp, "-")
	es, err := e.repo.GetByRoute(routeArr[0], routeArr[1])
	if err != nil {
		return []models.ExpeditionSchedule{}, err
	}

	for i, _ := range es {
		es[i].RouteFrom = ""
		es[i].RouteDestination = ""

		monthEta := utils.ConvertMonth(es[i].Eta)
		etaTemp := strings.Split(es[i].Eta, "-")
		str := etaTemp[2] + "-" + monthEta + "-" + etaTemp[0]
		etaTempArr := strings.Split(str, "-")
		eta := strings.Join(etaTempArr, " ")
		es[i].Eta = eta

		monthEtd := utils.ConvertMonth(es[i].Etd)
		etdTemp := strings.Split(es[i].Etd, "-")
		strEtd := etdTemp[2] + "-" + monthEtd + "-" + etdTemp[0]
		etdTempArr := strings.Split(strEtd, "-")
		etd := strings.Join(etdTempArr, " ")
		es[i].Etd = etd

		monthClosing := utils.ConvertMonth(es[i].Closing)
		closingTemp := strings.Split(es[i].Closing, "-")
		strClosing := closingTemp[2] + "-" + monthClosing + "-" + closingTemp[0]
		cosingTempArr := strings.Split(strClosing, "-")
		closing := strings.Join(cosingTempArr, " ")
		es[i].Closing = closing
	}

	return es, nil
}
