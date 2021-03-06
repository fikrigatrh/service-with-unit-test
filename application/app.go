package application

import (
	"bitbucket.org/service-ekspedisi/config/db"
	"bitbucket.org/service-ekspedisi/config/env"
	log2 "bitbucket.org/service-ekspedisi/config/log"
	"bitbucket.org/service-ekspedisi/controllers"
	"bitbucket.org/service-ekspedisi/middlewares"
	"bitbucket.org/service-ekspedisi/models"
	"bitbucket.org/service-ekspedisi/repo/about_us"
	"bitbucket.org/service-ekspedisi/repo/blog"
	"bitbucket.org/service-ekspedisi/repo/expedition_schedule_rp"
	"bitbucket.org/service-ekspedisi/repo/login_repo"
	"bitbucket.org/service-ekspedisi/repo/user_repo"
	"bitbucket.org/service-ekspedisi/usecase/about_usecase"
	blog2 "bitbucket.org/service-ekspedisi/usecase/blog"
	error2 "bitbucket.org/service-ekspedisi/usecase/error"
	"bitbucket.org/service-ekspedisi/usecase/expedition_schedule_uc"
	"bitbucket.org/service-ekspedisi/usecase/login_usecase"
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

	logCustom := log2.NewLogCustom(env.Config)
	if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
		log.Println("main : init contract")
	}

	dbBase := db.NewDB(env.Config, false).DB
	fmt.Println(dbBase)
	//dbBase.Debug().Migrator().DropTable(models.ExpeditionSchedule{})
	err := dbBase.Debug().AutoMigrate(models.ExpeditionSchedule{})
	if err != nil {
		log.Println("main: cannot auto migrate remapping response code")
		return
	}
	//dbBase.AutoMigrate(nil)
	models.InitTable(dbBase)

	// init db log
	logDb := log2.NewLogDbCustom(dbBase)
	logCustom.LogDb = logDb

	// repository
	repoUser := user_repo.NewUserRepo(dbBase)
	repoLogin := login_repo.NewLoginRepo(dbBase)
	aboutRepo := about_us.NewAboutUsRepo(dbBase, logCustom)
	esRepo := expedition_schedule_rp.NewExpeditionRepo(dbBase, logCustom)
	blogRepo := blog.NewBlogRepo(dbBase, logCustom)

	errorUc := error2.NewErrorHandlerUsecase()
	//usecase

	ucLogin := login_usecase.NewLoginUsecase(repoLogin)
	ucUser := user_usecase.NewUserUsecase(repoUser)
	abtUc := about_usecase.NewAboutUsUsecase(aboutRepo, logCustom)
	esUc := expedition_schedule_uc.NewEsUc(esRepo, logCustom)
	blogUc := blog2.NewBlogUc(blogRepo, logCustom)

	router.Use(middlewares.CORSMiddleware())
	newRoute := router.Group("api/v1")

	//newRoute.Use(middlewares.TokenAuthMiddleware())
	middlewares.NewErrorHandler(newRoute, errorUc, logCustom)

	// controller

	controllers.NewUserController(newRoute, ucUser, errorUc, logCustom)
	controllers.NewAboutUsController(newRoute, abtUc, errorUc, logCustom)
	controllers.NewExpeditionController(newRoute, esUc, errorUc, logCustom)
	controllers.NewBlogController(newRoute, blogUc, errorUc, logCustom)
	controllers.NewDaerahController(newRoute, dbBase)
	router.Use(middlewares.TokenAuthMiddlewareCustom(repoLogin))
	controllers.NewLoginController(newRoute, ucLogin)

	if err := router.Run(env.Config.Host + ":" + env.Config.Port); err != nil {
		log.Fatal("error starting server", err)
	}
}
