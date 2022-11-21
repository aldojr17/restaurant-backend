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
		UpdateReview(review *domain.Review) *domain.Response
		GetReview(user_id string, menu_id int) *domain.Response
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

func (repo *reviewRepository) UpdateReview(review *domain.Review) *domain.Response {
	if err := repo.db.Where("id = ?", review.Id).Updates(review).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseReviewUpdated, 0, nil)
}

func (repo *reviewRepository) GetReview(user_id string, menu_id int) *domain.Response {
	review := new(domain.Review)

	if err := repo.db.Where("user_id = ? AND menu_id = ?", user_id, menu_id).First(&review).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(review, 0, nil)
}
