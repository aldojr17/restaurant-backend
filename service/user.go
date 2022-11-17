package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type (
	UserService interface {
		UpdateUserData(payload *domain.UserProfile) *domain.Response
		GetCoupons(user_id string) *domain.Response
	}

	userService struct {
		db         *gorm.DB
		userRepo   repository.UserRepository
		couponRepo repository.CouponRepository
	}
)

func NewUserService(db *gorm.DB, userRepo repository.UserRepository, couponRepo repository.CouponRepository) UserService {
	return &userService{
		db:         db,
		userRepo:   userRepo,
		couponRepo: couponRepo,
	}
}

func (s *userService) UpdateUserData(payload *domain.UserProfile) *domain.Response {
	data := map[string]interface{}{
		"full_name": payload.FullName,
		"phone":     payload.Phone,
		"address":   payload.Address,
	}

	return s.userRepo.UpdateUserData(payload.UserId, data)
}

func (s *userService) GetCoupons(user_id string) *domain.Response {
	return s.couponRepo.GetCouponOwnedByUser(user_id)
}
