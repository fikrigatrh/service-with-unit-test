package application

import (
	"bitbucket.org/service-ekspedisi/config/db"
	"bitbucket.org/service-ekspedisi/config/env"
	"bitbucket.org/service-ekspedisi/controllers"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo/user_repo"
	"bitbucket.org/service-ekspedisi/usecase/user_usecase"
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
	models.InitTable(dbBase.DB)

	// init db log
	dBaseLog := db.NewDB(env.Config, true)
	dBaseLog.AutoMigrate(models.Logs{})

	// repository
	repoUser := user_repo.NewUserRepo(dbBase.DB)

	//usecase
	ucUser :=  user_usecase.NewUserUsecase(repoUser)

	newRoute := router.Group("api/v1")

	//newRoute.Use(middlewares.TokenAuthMiddleware())

	// controller
	controllers.NewUserController(newRoute,ucUser)

	if err := router.Run(env.Config.ServiceHost + ":" + env.Config.Port); err != nil {
		log.Fatal("error starting server", err)
	}
}
