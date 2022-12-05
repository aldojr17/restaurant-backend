package handler

import (
	"errors"
	"final-project-backend/initialize"
	"final-project-backend/mocks"
	"final-project-backend/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewCategoryHandler(t *testing.T) {
	app := initialize.App()
	categoryHandler := NewCategoryHandler(app)

	assert.NotNil(t, categoryHandler)
}

func TestGetAllCategory(t *testing.T) {
	s := mocks.NewCategoryService(t)
	h := categoryHandler{
		s: s,
	}

	s.On("GetAllCategory").Return(util.SetResponse(nil, 0, nil))

	router := gin.Default()
	router.GET("/categories", GinHandlerWrapper(h.GetAllCategory))

	req := httptest.NewRequest("GET", "/categories", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
}

func TestGetAllCategoryError(t *testing.T) {
	s := mocks.NewCategoryService(t)
	h := categoryHandler{
		s: s,
	}

	s.On("GetAllCategory").Return(util.SetResponse(nil, 0, errors.New("error")))

	router := gin.Default()
	router.GET("/categories", GinHandlerWrapper(h.GetAllCategory))

	req := httptest.NewRequest("GET", "/categories", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
}
