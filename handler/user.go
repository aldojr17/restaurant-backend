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
	userHandler struct {
		s service.UserService
	}

	UserHandler interface {
		UpdateUserData(c *gin.Context) *domain.Response
	}
)

func NewUserHandler(app *initialize.Application) UserHandler {
	return &userHandler{
		s: service.NewUserService(
			app.DB,
			repository.NewUserRepository(app.DB),
		),
	}
}

func (h *userHandler) UpdateUserData(c *gin.Context) *domain.Response {
	email, exists := c.Get(domain.EMAIL)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.UserProfile)
	param.Email = email.(string)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.UpdateUserData(param)
}
