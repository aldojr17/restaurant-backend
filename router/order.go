package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, app *initialize.Application) {
	orderHandler := handler.NewOrderHandler(app)
	router.POST("/orders", handler.GinHandlerWrapper(orderHandler.CreateOrder))
}
