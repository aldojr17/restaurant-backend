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
		CreateGame(c *gin.Context) *domain.Response
		GetLeaderboards(c *gin.Context) *domain.Response
		GetHistoryGame(c *gin.Context) *domain.Response
	}
)

func NewGameHandler(app *initialize.Application) GameHandler {
	return &gameHandler{
		s: service.NewGameService(
			app.DB,
			repository.NewGameRepository(app.DB),
			repository.NewCouponRepository(app.DB),
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

func (h *gameHandler) CreateGame(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.GamePayload)
	param.UserId = user_id.(string)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	response := h.s.CreateGame(param)
	if response.Err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, response.Err)
	}

	return response
}

func (h *gameHandler) GetLeaderboards(c *gin.Context) *domain.Response {
	_, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	response := h.s.GetLeaderboards()
	if response.Err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, response.Err)
	}

	return response
}

func (h *gameHandler) GetHistoryGame(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	response := h.s.GetHistoryGame(user_id.(string))
	if response.Err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, response.Err)
	}

	return response
}
