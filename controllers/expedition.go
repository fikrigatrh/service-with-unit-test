package controllers

import (
	"bitbucket.org/service-ekspedisi/usecase"
	"github.com/gin-gonic/gin"
)

type ExpeditionController struct {
	uc usecase.ExpeditionUcInterface
}

func NewExpeditionController(r *gin.RouterGroup, uc usecase.ExpeditionUcInterface) {
	handler := &ExpeditionController{uc: uc}

	r.GET("/expeditions", handler.GetAll)
	r.GET("/expedition/:id", handler.GetById)
	r.POST("/expedition", handler.Create)
	r.PUT("/expedition/:id", handler.Update)
	r.DELETE("/expedition/:id", handler.Delete)
}

func (e ExpeditionController) GetAll(c *gin.Context) {

}

func (e ExpeditionController) GetById(c *gin.Context) {

}

func (e ExpeditionController) Create(c *gin.Context) {

}

func (e ExpeditionController) Update(c *gin.Context) {

}

func (e ExpeditionController) Delete(c *gin.Context) {

}
