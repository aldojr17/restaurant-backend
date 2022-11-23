package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, app *initialize.Application) {
	categoryHandler := handler.NewCategoryHandler(app)
	router.GET("/categories", handler.GinHandlerWrapper(categoryHandler.GetAllCategory))
}
