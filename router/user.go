package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, app *initialize.Application) {
	userHandler := handler.NewUserHandler(app)

	users := router.Group("/users")
	{
		users.PUT("/change-profile", handler.GinHandlerWrapper(userHandler.UpdateUserData))
		users.GET("/coupons", handler.GinHandlerWrapper(userHandler.GetCoupons))
	}
}
