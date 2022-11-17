package service

import (
	"final-project-backend/domain"
	"final-project-backend/mocks"
	"final-project-backend/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserData(t *testing.T) {
	s := SetupMockDb()

	userRepo := mocks.NewUserRepository(t)
	service := NewUserService(s.db, userRepo)

	email := "admin@email.com"
	username := "test"
	fullname := "testtt"
	phone := "081234567891"

	payload := new(domain.UserProfile)
	payload.Email = email
	payload.Username = username
	payload.FullName = fullname
	payload.Phone = phone

	user := new(domain.UserResponse)
	user.Email = email

	data := map[string]interface{}{
		"username":  username,
		"full_name": fullname,
		"phone":     phone,
	}

	userRepo.On("UpdateUserData", email, data).Return(util.SetResponse(user, 0, nil))

	response := service.UpdateUserData(payload)
	assert.Nil(t, response.Err)
}
