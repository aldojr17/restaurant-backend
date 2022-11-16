package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, app *initialize.Application) {
	authHandler := handler.NewAuthHandler(app)

	router.POST("/register", handler.GinHandlerWrapper(authHandler.Register))
	router.POST("/login", handler.GinHandlerWrapper(authHandler.Login))
}
