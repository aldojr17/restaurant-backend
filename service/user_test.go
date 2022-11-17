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
	couponRepo := mocks.NewCouponRepository(t)
	service := NewUserService(s.db, userRepo, couponRepo)

	userId := "aca0702f-df5a-4fa2-af22-596f90edaef8"
	address := "test"
	fullname := "testtt"
	phone := "081234567891"

	payload := new(domain.UserProfile)
	payload.UserId = userId
	payload.Address = address
	payload.FullName = fullname
	payload.Phone = phone

	user := new(domain.UserResponse)
	user.Id = userId

	data := map[string]interface{}{
		"full_name": fullname,
		"phone":     phone,
		"address":   address,
	}

	userRepo.On("UpdateUserData", userId, data).Return(util.SetResponse(user, 0, nil))

	response := service.UpdateUserData(payload)
	assert.Nil(t, response.Err)
}

func TestGetCoupons(t *testing.T) {
	s := SetupMockDb()

	userRepo := mocks.NewUserRepository(t)
	couponRepo := mocks.NewCouponRepository(t)
	service := NewUserService(s.db, userRepo, couponRepo)

	userId := "aca0702f-df5a-4fa2-af22-596f90edaef8"

	couponRepo.On("GetCouponOwnedByUser", userId).Return(util.SetResponse(nil, 0, nil))

	response := service.GetCoupons(userId)
	assert.Nil(t, response.Err)
}
