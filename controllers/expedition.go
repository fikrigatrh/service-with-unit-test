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

type ExpeditionController struct {
	uc   usecase.ExpeditionUcInterface
	errH usecase.ErrorHandlerUsecase
	logC *log.LogCustom
}

func NewExpeditionController(r *gin.RouterGroup, uc usecase.ExpeditionUcInterface, errH usecase.ErrorHandlerUsecase, logC *log.LogCustom) {
	handler := &ExpeditionController{
		uc:   uc,
		errH: errH,
		logC: logC,
	}

	r.GET("/expeditions", handler.GetAll)
	r.GET("/expedition", handler.GetById)
	r.POST("/expedition", handler.Create)
	r.PUT("/expedition/:id", handler.Update)
	r.DELETE("/expedition", handler.Delete)
	r.GET("/expedition/get-route", handler.GetByRoute)
}

func (e ExpeditionController) GetAll(c *gin.Context) {

	expeditions, err := e.uc.GetAll()
	if err != nil {
		e.logC.Error(err, "controller: c get all usecase", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(c, expeditions)
}

func (e ExpeditionController) GetById(c *gin.Context) {
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		e.logC.Error(err, "controller: convert string to int in param", "", nil, nil, nil)
		c.Error(errors.New(contract.ErrBadRequest))
		c.Abort()
		return
	}

	result, err := e.uc.GetById(idInt)
	if err != nil {
		e.logC.Error(err, "controller: get by id about us usecase", "", nil, idInt, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(c, result)
}

func (e ExpeditionController) Create(c *gin.Context) {
	var req models.ExpeditionSchedule

	err := c.ShouldBindJSON(&req)
	if err != nil {
		e.logC.Error(err, "controller: c bindjson", "", nil, nil, nil)
		c.JSON(400, err.Error())
		c.Abort()
		return
	}

	fieldErr, err := e.errH.ValidateRequest(req)
	if err != nil {
		e.logC.Error(err, "controller: Validate request data", "", nil, req, nil)
		c.Error(err).SetMeta(models.ErrMeta{
			ServiceCode: models.ServiceCode,
			FieldErr:    fieldErr,
		})
		c.Abort()
		return
	}

	result, err := e.uc.AddEs(req)
	if err != nil {
		e.logC.Error(err, "controller: add about us usecase", "", nil, req, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(c, result)
}

func (e ExpeditionController) Update(c *gin.Context) {

}

func (e ExpeditionController) Delete(c *gin.Context) {
	id := c.Query("id")
	idRes := strings.Split(id, ",")

	err := e.uc.DeleteData(idRes)
	if err != nil {
		e.logC.Error(err, "controller: error when delete data submissions", "", nil, nil, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(c, nil)
}

func (e ExpeditionController) GetByRoute(c *gin.Context) {
	var req models.ExpeditionSchedule

	req.Route = c.Query("route")

	result, err := e.uc.GetByRoute(req.Route)
	if err != nil {
		e.logC.Error(err, "controller: get by id about us usecase", "", nil, req, nil)
		c.Error(err)
		c.Abort()
		return
	}

	responseSuccess(c, result)
}
