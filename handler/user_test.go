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
		Email:    "test@test.com",
		Username: "test",
		FullName: "testtt",
		Phone:    "081234567891",
	}

	jsonBody := `{
		"username":"test",
		"full_name":"testtt",
		"phone":"081234567891"
		}`

	body := strings.NewReader(jsonBody)

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	req, _ := http.NewRequest("POST", "/users/change-profile", body)

	c.Request = req
	c.Set("email", "test@test.com")

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
		"username":"test",
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
