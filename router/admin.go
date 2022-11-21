package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, app *initialize.Application) {
	orderHandler := handler.NewOrderHandler(app)
	menuHandler := handler.NewMenuHandler(app)
	couponHandler := handler.NewCouponHandler(app)

	users := router.Group("/admin")
	{
		users.GET("/orders", handler.GinHandlerWrapper(orderHandler.GetAllOrders))
		users.PUT("/orders/:id", handler.GinHandlerWrapper(orderHandler.UpdateOrderStatus))

		users.POST("/menus", handler.GinHandlerWrapper(menuHandler.CreateMenu))
		users.PUT("/menus/:id", handler.GinHandlerWrapper(menuHandler.UpdateMenu))
		users.DELETE("/menus/:id", handler.GinHandlerWrapper(menuHandler.DeleteMenu))

		users.POST("/coupons", handler.GinHandlerWrapper(couponHandler.CreateCoupon))
		users.DELETE("/coupons/:id", handler.GinHandlerWrapper(couponHandler.DeleteCoupon))
	}
}
