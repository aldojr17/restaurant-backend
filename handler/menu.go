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
	menuHandler struct {
		s service.MenuService
	}

	MenuHandler interface {
		GetAllMenus(c *gin.Context) *domain.Response
		CreateMenu(c *gin.Context) *domain.Response
		UpdateMenu(c *gin.Context) *domain.Response
		DeleteMenu(c *gin.Context) *domain.Response
		GetMenuDetail(c *gin.Context) *domain.Response
	}
)

func NewMenuHandler(app *initialize.Application) MenuHandler {
	return &menuHandler{
		s: service.NewMenuService(
			app.DB,
			repository.NewMenuRepository(app.DB),
		),
	}
}

func (h *menuHandler) GetAllMenus(c *gin.Context) *domain.Response {
	data, err := h.s.GetAllMenus(newMenuPageableRequest(c.Request))
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(data, 0, nil)
}

func (h *menuHandler) CreateMenu(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	param := new(domain.MenuPayload)

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.CreateMenu(param)
}

func (h *menuHandler) UpdateMenu(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	param := new(domain.MenuPayload)
	var menu_id int

	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrMenuIdRequired)
	} else {
		menu_id = id
	}

	if err := param.Validate(c); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return h.s.UpdateMenu(param, menu_id)
}

func (h *menuHandler) DeleteMenu(c *gin.Context) *domain.Response {
	role_id, exists := c.Get(domain.ROLE_ID)
	if !exists {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	if role_id.(int) != 0 {
		return util.SetResponse(nil, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	var menu_id int

	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrMenuIdRequired)
	} else {
		menu_id = id
	}

	return h.s.DeleteMenu(menu_id)
}

func (h *menuHandler) GetMenuDetail(c *gin.Context) *domain.Response {
	var menu_id int

	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrMenuIdRequired)
	} else {
		menu_id = id
	}

	return h.s.GetMenuDetail(menu_id)
}

func newMenuPageableRequest(r *http.Request) *domain.PageableRequest {
	p := &domain.PageableRequest{}
	p.Page = util.PageFromQueryParam(r)
	p.Limit = util.LimitFromQueryParam(r)
	p.Sort_by = util.SortValueFromQueryParam(r)
	p.Type = "menu"

	if p.Sort_by == "" {
		p.Sort_by = util.DEFAULT_SORT_BY
	}

	p.Desceding = util.SortDirectionFromQueryParam(r)
	p.Search = map[string]interface{}{}
	p.Filters = map[string]interface{}{}

	p.Search[util.SEARCH_BY_NAME] = queryLikeParamOrNull(r, util.SEARCH_BY_NAME)
	p.Filters[util.FILTER_BY_CATEGORY] = queryParamOrNull(r, util.FILTER_BY_CATEGORY)

	return p
}
