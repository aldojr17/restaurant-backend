package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	CategoryRepository interface {
		GetAllCategory() *domain.Response
	}

	categoryRepository struct {
		db *gorm.DB
	}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (repo *categoryRepository) GetAllCategory() *domain.Response {
	categories := new(domain.Categories)

	if err := repo.db.Model(&domain.Category{}).Find(&categories).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(categories, 0, nil)
}
