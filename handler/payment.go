package handler

import (
	"final-project-backend/domain"
	"final-project-backend/initialize"
	"final-project-backend/repository"
	"final-project-backend/service"
	"final-project-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	paymentHandler struct {
		s service.PaymentService
	}

	PaymentHandler interface {
		GetAllPayment(c *gin.Context) *domain.Response
	}
)

func NewPaymentHandler(app *initialize.Application) PaymentHandler {
	return &paymentHandler{
		s: service.NewPaymentService(
			app.DB,
			repository.NewPaymentRepository(app.DB),
		),
	}
}

func (h *paymentHandler) GetAllPayment(c *gin.Context) *domain.Response {
	response := h.s.GetAllPayment()
	if response.Err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, response.Err)
	}

	return response
}
