package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"
	"net/http"

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
		menuRepo   repository.MenuRepository
	}
)

func NewUserService(db *gorm.DB, userRepo repository.UserRepository, couponRepo repository.CouponRepository, menuRepo repository.MenuRepository) UserService {
	return &userService{
		db:         db,
		userRepo:   userRepo,
		couponRepo: couponRepo,
		menuRepo:   menuRepo,
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
	if response := s.menuRepo.GetMenu(payload.MenuId); response.Err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrMenuNotFound)
	}

	return s.userRepo.AddMenuFavorite(payload)
}
