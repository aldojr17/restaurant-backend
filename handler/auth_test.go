package handler

import (
	"final-project-backend/domain"
	"final-project-backend/initialize"
	"final-project-backend/mocks"
	"final-project-backend/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthHandler(t *testing.T) {
	app := initialize.App()
	authHandler := NewAuthHandler(app)

	assert.NotNil(t, authHandler)
}

func TestRegister(t *testing.T) {
	payload := &domain.RegisterPayload{
		FullName: "test",
		Email:    "test@test.com",
		Password: "1234",
	}

	jsonBody := `{
		"fullname":"test",
		"email":"test@test.com",
		"password":"1234"
		}`

	body := strings.NewReader(jsonBody)

	s := mocks.NewAuthService(t)
	h := authHandler{
		s: s,
	}

	s.On("Register", payload).Return(util.SetResponse(nil, 0, nil))

	router := gin.Default()
	router.POST("/register", GinHandlerWrapper(h.Register))

	req := httptest.NewRequest("POST", "/register", body)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
}

func TestRegisterError(t *testing.T) {
	payload := &domain.RegisterPayload{
		FullName: "test",
		Email:    "test@test",
		Password: "1234",
	}

	jsonBody := `{
		"fullname":"test",
		"email":"test@test",
		"password":"1234"
		}`

	body := strings.NewReader(jsonBody)

	s := mocks.NewAuthService(t)
	h := authHandler{
		s: s,
	}

	s.On("Register", payload).Return(util.SetResponse(nil, http.StatusBadRequest, util.NewError("error")))

	router := gin.Default()
	router.POST("/register", GinHandlerWrapper(h.Register))

	req := httptest.NewRequest("POST", "/register", body)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
}

func TestLogin(t *testing.T) {
	payload := &domain.AuthPayload{
		Email:    "test@test.com",
		Password: "1234",
	}

	jsonBody := `{
		"email":"test@test.com",
		"password":"1234"
		}`

	body := strings.NewReader(jsonBody)

	s := mocks.NewAuthService(t)
	h := authHandler{
		s: s,
	}

	s.On("Login", payload).Return(util.SetResponse(domain.LoginResponse{}, 0, nil))

	router := gin.Default()
	router.POST("/login", GinHandlerWrapper(h.Login))

	req := httptest.NewRequest("POST", "/login", body)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
}
