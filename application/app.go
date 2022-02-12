package application

import (
	"bitbucket.org/service-ekspedisi/config/db"
	"bitbucket.org/service-ekspedisi/config/env"
	"bitbucket.org/service-ekspedisi/middlewares"
	"bitbucket.org/service-ekspedisi/models"
	"fmt"
	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
	"log"
)

func StartApp() {
	router := gin.New()
	router.Use(gin.Recovery())

	env.NewEnv(".env")

	if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
		log.Println("main : init contract")
	}

	dbBase := db.NewDB(env.Config, false)
	fmt.Println(dbBase)
	//dbBase.AutoMigrate(nil)

	// init db log
	dBaseLog := db.NewDB(env.Config, true)
	dBaseLog.AutoMigrate(models.Logs{})

	// repository

	//usecase

	newRoute := router.Group("")

	newRoute.Use(middlewares.TokenAuthMiddleware())

	// controller

	if err := router.Run(env.Config.ServiceHost + ":" + env.Config.Port); err != nil {
		log.Fatal("error starting server", err)
	}
}
