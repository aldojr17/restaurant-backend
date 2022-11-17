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

func TestNewUserHandler(t *testing.T) {
	app := initialize.App()
	userHandler := NewUserHandler(app)

	assert.NotNil(t, userHandler)
}

func TestUpdateUserData(t *testing.T) {
	payload := &domain.UserProfile{
		UserId:   "aca0702f-df5a-4fa2-af22-596f90edaef8",
		Address:  "test",
		FullName: "testtt",
		Phone:    "081234567891",
	}

	jsonBody := `{
		"address":"test",
		"full_name":"testtt",
		"phone":"081234567891"
		}`

	body := strings.NewReader(jsonBody)

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	req, _ := http.NewRequest("POST", "/users/change-profile", body)

	c.Request = req
	c.Set("user_id", "aca0702f-df5a-4fa2-af22-596f90edaef8")

	s := mocks.NewUserService(t)
	h := userHandler{
		s: s,
	}

	s.On("UpdateUserData", payload).Return(util.SetResponse(nil, 0, nil))

	res := h.UpdateUserData(c)
	assert.Nil(t, res.Err)
}

func TestUpdateUserDataErrorEmail(t *testing.T) {
	jsonBody := `{
		"address":"test",
		"full_name":"testtt",
		"phone":"081234567891"
		}`

	body := strings.NewReader(jsonBody)

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	req, _ := http.NewRequest("POST", "/users/change-profile", body)

	c.Request = req

	s := mocks.NewUserService(t)
	h := userHandler{
		s: s,
	}

	res := h.UpdateUserData(c)
	assert.NotNil(t, res.Err)
}

func TestGetCoupons(t *testing.T) {
	payload := "aca0702f-df5a-4fa2-af22-596f90edaef8"

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	req, _ := http.NewRequest("GET", "/users/coupons", nil)

	c.Request = req
	c.Set("user_id", "aca0702f-df5a-4fa2-af22-596f90edaef8")

	s := mocks.NewUserService(t)
	h := userHandler{
		s: s,
	}

	s.On("GetCoupons", payload).Return(util.SetResponse(nil, 0, nil))

	res := h.GetCoupons(c)
	assert.Nil(t, res.Err)
}

func TestGetCouponsError(t *testing.T) {
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	req, _ := http.NewRequest("GET", "/users/coupons", nil)

	c.Request = req

	s := mocks.NewUserService(t)
	h := userHandler{
		s: s,
	}

	res := h.GetCoupons(c)
	assert.NotNil(t, res.Err)
}
