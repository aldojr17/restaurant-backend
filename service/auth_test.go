package service

import (
	"final-project-backend/domain"
	"final-project-backend/mocks"
	"final-project-backend/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	s := SetupMockDb()

	userRepo := mocks.NewUserRepository(t)
	service := NewAuthService(s.db, userRepo)

	email := "admin@email.com"

	payload := new(domain.AuthPayload)
	payload.Email = email
	payload.Password = "1234"

	user := new(domain.User)
	user.Id = util.GenerateUUID()
	user.Email = email
	user.Password = "$2a$04$LoG0NFPuaobuUM0X9j4PGOU8PNxlcdaO7CHckjOcNPPotR6EKNWQq"

	userRepo.On("GetUserByEmail", email).Return(util.SetResponse(user, 0, nil))

	response := service.Login(payload)
	assert.Nil(t, response.Err)
}

func TestLoginErrorUserNotFound(t *testing.T) {
	s := SetupMockDb()

	userRepo := mocks.NewUserRepository(t)
	service := NewAuthService(s.db, userRepo)

	email := "test@test.com"

	payload := new(domain.AuthPayload)
	payload.Email = email
	payload.Password = "1234"

	userRepo.On("GetUserByEmail", email).Return(util.SetResponse(nil, 0, util.NewError("error")))

	response := service.Login(payload)
	assert.NotNil(t, response.Err)
}

func TestLoginErrorPassword(t *testing.T) {
	s := SetupMockDb()

	userRepo := mocks.NewUserRepository(t)
	service := NewAuthService(s.db, userRepo)

	email := "test@test.com"

	payload := new(domain.AuthPayload)
	payload.Email = email
	payload.Password = "1234"

	user := new(domain.User)
	user.Id = util.GenerateUUID()
	user.Email = email
	user.Password = "1234"

	userRepo.On("GetUserByEmail", email).Return(util.SetResponse(user, 0, nil))

	response := service.Login(payload)
	assert.NotNil(t, response.Err)
}
