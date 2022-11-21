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
		GetCoupons(c *gin.Context) *domain.Response
		GetProfile(c *gin.Context) *domain.Response

		AddMenuFavorite(c *gin.Context) *domain.Response

		UpdateUserData(c *gin.Context) *domain.Response
	}
)

func NewUserHandler(app *initialize.Application) UserHandler {
	return &userHandler{
		s: service.NewUserService(
			app.DB,
			repository.NewUserRepository(app.DB),
			repository.NewCouponRepository(app.DB),
			repository.NewMenuRepository(app.DB),
		),
	}
}

func (h *userHandler) UpdateUserData(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.UserProfile)
	param.UserId = user_id.(string)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.UpdateUserData(param)
}

func (h *userHandler) GetCoupons(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	return h.s.GetCoupons(user_id.(string))
}

func (h *userHandler) GetProfile(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	return h.s.GetProfile(user_id.(string))
}

func (h *userHandler) AddMenuFavorite(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.UserFavorite)
	param.UserId = user_id.(string)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.AddMenuFavorite(param)
}
