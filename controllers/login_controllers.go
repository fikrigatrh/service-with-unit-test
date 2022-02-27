package controllers

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	uc usecase.LoginUcInterface
}

func NewLoginController(r *gin.RouterGroup, uc usecase.LoginUcInterface) {
	handler := &LoginController{
		uc: uc,
	}

	r.POST("/login", handler.Login)
}

func (a LoginController) Login(c *gin.Context) {
	var encrpytData models.EncryptData
	err := c.ShouldBindJSON(&encrpytData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "1111",
			"responseMessage": "Ops, Error when bind json from body",
		})
		fmt.Printf("[login Controller] error when encode data enkripsi : %v\n", err)
		return
	}


	token, err := a.uc.LoginUser(encrpytData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "1111",
			"responseMessage": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "0000",
		"responseMessage": "Success",
		"data" : token,
	})
}