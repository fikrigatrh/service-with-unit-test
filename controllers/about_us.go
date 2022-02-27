package controllers

import (
	"bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/models/contract"
	"bitbucket.org/service-ekspedisi/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type AboutUsController struct {
	uc   usecase.AboutUsUcInterface
	errH usecase.ErrorHandlerUsecase
	logC *log.LogCustom
}

func NewAboutUsController(r *gin.RouterGroup, uc usecase.AboutUsUcInterface, errH usecase.ErrorHandlerUsecase, logC *log.LogCustom) {
	handler := &AboutUsController{
		uc:   uc,
		errH: errH,
		logC: logC,
	}

	r.POST("/add-about-us", handler.AddAboutUs)
	r.GET("/all-about-us", handler.GetAllAboutUs)
	r.PUT("/update-about-us", handler.UpdateAboutUs)
	r.DELETE("/delete-about-us", handler.DeleteAboutUs)
	r.GET("about/", handler.GetAboutUsById)
}

func (a AboutUsController) GetAllAboutUs(c *gin.Context) {
	var aboutUsList []models.AboutUsRequest

	aboutUsList, err := a.uc.GetAll()
	if err != nil {
		a.logC.Error(err, "controller: get all abaout usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(aboutUsList)
}

func (a AboutUsController) AddAboutUs(c *gin.Context) {
	var req models.AboutUsRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		a.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := a.errH.ValidateRequest(req)
	if err != nil {
		a.logC.Error(err, "controller: Validate request data", "", nil, req, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	about, err := a.uc.AddAbout(req)
	if err != nil {
		a.logC.Error(err, "controller: add about us usecase", "", nil, req, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(about)
}

func (a AboutUsController) UpdateAboutUs(c *gin.Context) {

}

func (a AboutUsController) DeleteAboutUs(c *gin.Context) {
	id := c.Query("id")
	idRes := strings.Split(id, ",")

	err := a.uc.DeleteData(idRes)
	if err != nil {
		a.logC.Error(err, "controller: error when delete data submissions", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(nil)
}

func (a AboutUsController) GetAboutUsById(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		a.logC.Error(err, "controller: convert string to int in param", "", nil, nil, nil)
		c.Error(errors.New(contract.ErrBadRequest))
		c.Abort()
		return
	}

	result, err := a.uc.GetById(idInt)
	if err != nil {
		a.logC.Error(err, "controller: get by id about us usecase", "", nil, idInt, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(result)
}

func responseSuccess(data interface{}) gin.H {
	return gin.H{
		"responseCode":    "0000",
		"responseMessage": "Success",
		"data":            data,
	}
}
