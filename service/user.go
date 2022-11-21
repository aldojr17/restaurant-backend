package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type (
	UserService interface {
		GetCoupons(user_id string) *domain.Response
		GetProfile(user_id string) *domain.Response

		AddMenuFavorite(payload *domain.UserFavorite) *domain.Response

		UpdateUserData(payload *domain.UserProfile) *domain.Response
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
		"full_name":       payload.FullName,
		"phone":           payload.Phone,
		"address":         payload.Address,
		"profile_picture": payload.ProfilePicture,
	}

	return s.userRepo.UpdateUserData(payload.UserId, data)
}

func (s *userService) GetCoupons(user_id string) *domain.Response {
	return s.couponRepo.GetCouponOwnedByUser(user_id)
}

func (s *userService) GetProfile(user_id string) *domain.Response {
	return s.userRepo.GetUserById(user_id)
}

func (s *userService) AddMenuFavorite(payload *domain.UserFavorite) *domain.Response {
	return s.userRepo.AddMenuFavorite(payload)
}
