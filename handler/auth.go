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
	authHandler struct {
		s service.AuthService
	}

	AuthHandler interface {
		Register(c *gin.Context) *domain.Response
		Login(c *gin.Context) *domain.Response
	}
)

func NewAuthHandler(app *initialize.Application) AuthHandler {
	return &authHandler{
		s: service.NewAuthService(
			app.DB,
			repository.NewUserRepository(app.DB),
		),
	}
}

func (h *authHandler) Register(c *gin.Context) *domain.Response {
	param := new(domain.AuthPayload)
	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.Register(param)
}

func (h *authHandler) Login(c *gin.Context) *domain.Response {
	param := new(domain.AuthPayload)
	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.Login(param)
}
