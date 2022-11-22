package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine, app *initialize.Application) {
	menuHandler := handler.NewMenuHandler(app)
	categoryHandler := handler.NewCategoryHandler(app)
	router.GET("/menus", handler.GinHandlerWrapper(menuHandler.GetAllMenus))
	router.GET("/menus/categories", handler.GinHandlerWrapper(categoryHandler.GetAllCategory))
}
