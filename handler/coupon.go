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
	couponHandler struct {
		s service.CouponService
	}

	CouponHandler interface {
		CreateCoupon(c *gin.Context) *domain.Response
		DeleteCoupon(c *gin.Context) *domain.Response
		GetAllCoupon(c *gin.Context) *domain.Response
		GetCoupon(c *gin.Context) *domain.Response
		UpdateCoupon(c *gin.Context) *domain.Response
	}
)

func NewCouponHandler(app *initialize.Application) CouponHandler {
	return &couponHandler{
		s: service.NewCouponService(
			app.DB,
			repository.NewCouponRepository(app.DB),
		),
	}
}

func (h *couponHandler) CreateCoupon(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	param := new(domain.Coupon)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.CreateCoupon(param)
}

func (h *couponHandler) DeleteCoupon(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	id := c.Param("id")

	return h.s.DeleteCoupon(id)
}

func (h *couponHandler) GetAllCoupon(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	return h.s.GetAllCoupon()
}

func (h *couponHandler) GetCoupon(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	id := c.Param("id")

	return h.s.GetCoupon(id)
}

func (h *couponHandler) UpdateCoupon(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	id := c.Param("id")

	param := new(domain.Coupon)
	param.Id = id

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.UpdateCoupon(param)
}
