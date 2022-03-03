package middlewares

import (
	"bitbucket.org/service-ekspedisi/auth"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResponseCustomErr{
				ResponseCode:    "4011000",
				ResponseMessage: "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func TokenAuthMiddlewareCustom(loginRepo repo.LoginRepoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		errs := errors.New("1")
		err := auth.TokenValidCustom(c.Request,loginRepo)
		if err != nil || err == errs{
			c.JSON(http.StatusUnauthorized, models.ResponseCustomErr{
				ResponseCode:    "4011000",
				ResponseMessage: "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
