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
	reviewHandler struct {
		s service.ReviewService
	}

	ReviewHandler interface {
		AddReview(c *gin.Context) *domain.Response
	}
)

func NewReviewHandler(app *initialize.Application) ReviewHandler {
	return &reviewHandler{
		s: service.NewReviewService(
			app.DB,
			repository.NewReviewRepository(app.DB),
		),
	}
}

func (h *reviewHandler) AddReview(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.Review)
	param.UserId = user_id.(string)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.AddReview(param)
}
