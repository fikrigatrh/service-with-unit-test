package controllers

import (
	"bitbucket.org/service-ekspedisi/usecase"
	"github.com/gin-gonic/gin"
)

type AboutUsController struct {
	uc usecase.AboutUsRepoInterface
}

func NewAboutUsController(r *gin.RouterGroup, uc usecase.AboutUsRepoInterface) {
	handler := &AboutUsController{
		uc: uc,
	}

	r.POST("/add-about-us", handler.AddAboutUs)
	r.GET("/all-about-us", handler.GetAboutUs)
	r.PUT("/update-about-us", handler.UpdateAboutUs)
	r.DELETE("/delete-about-us", handler.DeleteAboutUs)
}

func (a AboutUsController) GetAboutUs(c *gin.Context) {

}

func (a AboutUsController) AddAboutUs(c *gin.Context) {

}

func (a AboutUsController) UpdateAboutUs(c *gin.Context) {

}

func (a AboutUsController) DeleteAboutUs(c *gin.Context) {

}
