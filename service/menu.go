package service

import (
	"final-project-backend/repository"
	"final-project-backend/util"

	"gorm.io/gorm"
)

type (
	MenuService interface {
		GetAllMenus(pageable util.Pageable) (*util.Page, error)
	}

	menuService struct {
		db       *gorm.DB
		menuRepo repository.MenuRepository
	}
)

func NewMenuService(db *gorm.DB, menuRepo repository.MenuRepository) MenuService {
	return &menuService{
		db:       db,
		menuRepo: menuRepo,
	}
}

func (s *menuService) GetAllMenus(pageable util.Pageable) (*util.Page, error) {
	return s.menuRepo.GetAllMenus(pageable)
}
