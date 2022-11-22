package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type (
	CategoryService interface {
		GetAllCategory() *domain.Response
	}

	categoryService struct {
		db   *gorm.DB
		repo repository.CategoryRepository
	}
)

func NewCategoryService(db *gorm.DB, repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		db:   db,
		repo: repo,
	}
}

func (s *categoryService) GetAllCategory() *domain.Response {
	return s.repo.GetAllCategory()
}
