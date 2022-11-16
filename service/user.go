package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type (
	UserService interface {
		UpdateUserData(payload *domain.UserProfile) *domain.Response
	}

	userService struct {
		db       *gorm.DB
		userRepo repository.UserRepository
	}
)

func NewUserService(db *gorm.DB, userRepo repository.UserRepository) UserService {
	return &userService{
		db:       db,
		userRepo: userRepo,
	}
}

func (s *userService) UpdateUserData(payload *domain.UserProfile) *domain.Response {
	data := map[string]interface{}{
		"username":  payload.Username,
		"full_name": payload.FullName,
		"phone":     payload.Phone,
	}

	return s.userRepo.UpdateUserData(payload.Email, data)
}
