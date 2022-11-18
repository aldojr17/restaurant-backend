package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, app *initialize.Application) {
	userHandler := handler.NewUserHandler(app)
	orderHandler := handler.NewOrderHandler(app)

	users := router.Group("/users")
	{
		users.PUT("/change-profile", handler.GinHandlerWrapper(userHandler.UpdateUserData))
		users.GET("/coupons", handler.GinHandlerWrapper(userHandler.GetCoupons))
		users.GET("/profile", handler.GinHandlerWrapper(userHandler.GetProfile))
		users.POST("/favorites", handler.GinHandlerWrapper(userHandler.AddMenuFavorite))
		users.GET("/orders", handler.GinHandlerWrapper(orderHandler.GetAllOrders))
	}
}
