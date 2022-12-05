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
	gameHandler struct {
		s service.GameService
	}

	GameHandler interface {
		GetAllQuestions(c *gin.Context) *domain.Response
	}
)

func NewGameHandler(app *initialize.Application) GameHandler {
	return &gameHandler{
		s: service.NewGameService(
			app.DB,
			repository.NewGameRepository(app.DB),
		),
	}
}

func (h *gameHandler) GetAllQuestions(c *gin.Context) *domain.Response {
	_, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	response := h.s.GetAllQuestions()
	if response.Err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, response.Err)
	}

	return response
}
