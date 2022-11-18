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
	orderHandler struct {
		s service.OrderService
	}

	OrderHandler interface {
		GetAllOrders(c *gin.Context) *domain.Response
	}
)

func NewOrderHandler(app *initialize.Application) OrderHandler {
	return &orderHandler{
		s: service.NewOrderService(
			app.DB,
			repository.NewOrderRepository(app.DB),
		),
	}
}

func (h *orderHandler) GetAllOrders(c *gin.Context) *domain.Response {
	user_id, exists := c.Get(domain.USER_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrUnauthorized)
	}

	data, err := h.s.GetAllOrders(newOrderPageableRequest(c.Request), user_id.(string))
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(data, 0, nil)
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

	p.Search[util.SEARCH_BY_NAME] = queryLikeParamOrNull(r, util.SEARCH_BY_NAME)

	return p
}
