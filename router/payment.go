package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine, app *initialize.Application) {
	paymentHandler := handler.NewPaymentHandler(app)
	router.GET("/payments", handler.GinHandlerWrapper(paymentHandler.GetAllPayment))
}
