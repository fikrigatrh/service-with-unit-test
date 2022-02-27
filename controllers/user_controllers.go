package controllers

import (
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	uc usecase.UserUcInterface
}

func NewUserController(r *gin.RouterGroup, uc usecase.UserUcInterface) {
	handler := &UserController{
		uc: uc,
	}

	r.POST("/add-user", handler.AddUser)
	r.GET("/all-user", handler.GetAllUser)
	r.GET("/user/:id", handler.GetUserByID)
	r.PUT("/update-user/:id", handler.UpdateUser)
	r.DELETE("/delete-user", handler.DeleteUser)
}

func (a UserController) AddUser(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": "Something error",
		})
		return
	}

	_, err = a.uc.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Registered successfully",
	})
}

func (a UserController) GetAllUser(c *gin.Context) {

	allDataUser, err := a.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": err,
		})
		return
	}

	responseSuccess(c, allDataUser)
}

func (a UserController) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := a.uc.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": err,
		})
		return
	}

	responseSuccess(c, data)
}

func (a UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	data, err := a.uc.UpdateData(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": err,
		})
		return
	}

	responseSuccess(c, data)
}

func (a UserController) DeleteUser(c *gin.Context) {
	id := c.Query("id")

	idRes := strings.Split(id, ",")

	err := a.uc.DeleteData(idRes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "1111",
			"responseMessage": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Successfully deleted",
	})
}
