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
	categoryHandler struct {
		s service.CategoryService
	}

	CategoryHandler interface {
		GetAllCategory(c *gin.Context) *domain.Response
	}
)

func NewCategoryHandler(app *initialize.Application) CategoryHandler {
	return &categoryHandler{
		s: service.NewCategoryService(
			app.DB,
			repository.NewCategoryRepository(app.DB),
		),
	}
}

func (h *categoryHandler) GetAllCategory(c *gin.Context) *domain.Response {
	response := h.s.GetAllCategory()
	if response.Err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, response.Err)
	}

	return response
}
