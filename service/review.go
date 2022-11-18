package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type (
	ReviewService interface {
		AddReview(payload *domain.Review) *domain.Response
	}

	reviewService struct {
		db   *gorm.DB
		repo repository.ReviewRepository
	}
)

func NewReviewService(db *gorm.DB, repo repository.ReviewRepository) ReviewService {
	return &reviewService{
		db:   db,
		repo: repo,
	}
}

func (s *reviewService) AddReview(payload *domain.Review) *domain.Response {
	return s.repo.AddReview(payload)
}
