package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	ReviewRepository interface {
		AddReview(review *domain.Review) *domain.Response
	}

	reviewRepository struct {
		db *gorm.DB
	}
)

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{
		db: db,
	}
}

func (repo *reviewRepository) AddReview(review *domain.Review) *domain.Response {
	if err := repo.db.Create(&review).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseReviewAdded, 0, nil)
}
