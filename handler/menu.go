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
	menuHandler struct {
		s service.MenuService
	}

	MenuHandler interface {
		GetAllMenus(c *gin.Context) *domain.Response
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
	data, err := h.s.GetAllMenu(newMenuPageableRequest(c.Request))
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(data, 0, nil)
}

func newMenuPageableRequest(r *http.Request) *domain.PageableRequest {
	p := &domain.PageableRequest{}
	p.Page = util.PageFromQueryParam(r)
	p.Limit = util.LimitFromQueryParam(r)
	p.Sort_by = util.SortValueFromQueryParam(r)

	if p.Sort_by == "" {
		p.Sort_by = util.DEFAULT_SORT_BY
	}

	p.Desceding = util.SortDirectionFromQueryParam(r)
	p.Search = map[string]interface{}{}

	p.Search[util.SEARCH_BY_NAME] = queryLikeParamOrNull(r, util.SEARCH_BY_NAME)

	return p
}
