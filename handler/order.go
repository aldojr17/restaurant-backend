package handler

import (
	"final-project-backend/domain"
	"final-project-backend/initialize"
	"final-project-backend/repository"
	"final-project-backend/service"
	"final-project-backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	orderHandler struct {
		s service.OrderService
	}

	OrderHandler interface {
		GetAllUserOrders(c *gin.Context) *domain.Response
		GetAllOrders(c *gin.Context) *domain.Response
		UpdateOrderStatus(c *gin.Context) *domain.Response
		CreateOrder(oc *gin.Context) *domain.Response
		CreateOrderDetails(c *gin.Context) *domain.Response
	}
)

func NewOrderHandler(app *initialize.Application) OrderHandler {
	return &orderHandler{
		s: service.NewOrderService(
			app.DB,
			repository.NewOrderRepository(app.DB),
			repository.NewCouponRepository(app.DB),
		),
	}
}

func (h *orderHandler) GetAllUserOrders(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	data, err := h.s.GetAllUserOrders(newOrderPageableRequest(c.Request), user_id.(string))
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(data, 0, nil)
}

func (h *orderHandler) GetAllOrders(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	data, err := h.s.GetAllOrders(newOrderPageableRequest(c.Request))
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(data, 0, nil)
}

func (h *orderHandler) UpdateOrderStatus(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.OrderStatusPayload)

	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrMenuIdRequired)
	} else {
		param.Id = id
	}

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.UpdateOrderStatus(param)
}

func (h *orderHandler) CreateOrder(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.OrderPayload)
	param.UserId = user_id.(string)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.CreateOrder(param)
}

func (h *orderHandler) CreateOrderDetails(c *gin.Context) *domain.Response {
	_, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	param := new(domain.OrderDetails)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.CreateOrderDetails(param)
}

func newOrderPageableRequest(r *http.Request) *domain.PageableRequest {
	p := &domain.PageableRequest{}
	p.Page = util.PageFromQueryParam(r)
	p.Limit = util.LimitFromQueryParam(r)
	p.Sort_by = util.SortValueFromQueryParam(r)
	p.Type = "order"

	if p.Sort_by == "" {
		p.Sort_by = util.DEFAULT_SORT_BY_ORDER
	}

	p.Desceding = util.SortDirectionFromQueryParam(r)
	p.Search = map[string]interface{}{}
	p.Filters = map[string]interface{}{}

	p.Search[util.SEARCH_BY_NAME] = queryLikeParamOrNull(r, util.SEARCH_BY_NAME)
	p.Filters[util.FILTER_BY_CATEGORY] = queryParamOrNull(r, util.FILTER_BY_CATEGORY)

	return p
}
