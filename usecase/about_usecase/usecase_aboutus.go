package about_usecase

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"bitbucket.org/service-ekspedisi/usecase"
	"encoding/json"
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

func (a AboutUsUsecaseStruct) AddAbout(req models.AboutUsRequest) (models.AboutUsRequest, error) {

	misiArr, err := json.Marshal(req.Misi)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	resMisi := string(misiArr)

	perusahaanArr, err := json.Marshal(req.PerusahaanRekanan)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	resPerusahaan := string(perusahaanArr)

	sosmedArr, err := json.Marshal(req.SocialMedia)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	resSosmed := string(sosmedArr)

	v := models.AboutUsDb{
		Profil:            req.Profil,
		Visi:              req.Visi,
		Misi:              resMisi,
		Motto:             req.Motto,
		PerusahaanRekanan: resPerusahaan,
		SocialMedia:       resSosmed,
	}

	about, err := a.repo.AddAbout(v)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	var result models.AboutUsRequest
	err = json.Unmarshal([]byte(about.Misi), &result.Misi)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	err = json.Unmarshal([]byte(about.PerusahaanRekanan), &result.PerusahaanRekanan)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	err = json.Unmarshal([]byte(about.SocialMedia), &result.SocialMedia)
	if err != nil {
		return models.AboutUsRequest{}, err
	}

	result.Profil = about.Profil
	result.Motto = about.Motto
	result.Visi = about.Visi

	return result, nil
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
