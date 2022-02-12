package middlewares

import (
	"bitbucket.org/service-ekspedisi/auth"
	"bitbucket.org/service-ekspedisi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResponseErrorCustom{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
