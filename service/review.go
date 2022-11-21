package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	ReviewService interface {
		AddReview(payload *domain.Review) *domain.Response
	}

	reviewService struct {
		db       *gorm.DB
		repo     repository.ReviewRepository
		menuRepo repository.MenuRepository
	}
)

func NewReviewService(db *gorm.DB, repo repository.ReviewRepository, menuRepo repository.MenuRepository) ReviewService {
	return &reviewService{
		db:       db,
		repo:     repo,
		menuRepo: menuRepo,
	}
}

func (s *reviewService) AddReview(payload *domain.Review) *domain.Response {
	if response := s.menuRepo.GetMenu(payload.MenuId); response.Err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrMenuNotFound)
	}

	if response := s.repo.GetReview(payload.UserId, payload.MenuId); response.Err == nil {
		payload.Id = response.Data.(*domain.Review).Id
		return s.repo.UpdateReview(payload)
	}

	return s.repo.AddReview(payload)
}
