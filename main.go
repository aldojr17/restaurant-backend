package main

import (
	"final-project-backend/initialize"
	"final-project-backend/logger"
	"final-project-backend/middleware"
	"final-project-backend/router"

	"github.com/gin-gonic/gin"
)

type Api struct {
	App *initialize.Application
}

func main() {
	app := initialize.App()
	api := &Api{
		App: app,
	}
	api.StartApi()
}

func (a *Api) StartApi() {
	r := gin.New()

	r.NoRoute(middleware.NoRoute())

	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(logger.GetGinCustomLogger))
	r.Use(middleware.LatencyMiddleware())
	r.Use(middleware.AuthMiddleware())

	router.Routes(r, a.App)

	r.Run()
}
