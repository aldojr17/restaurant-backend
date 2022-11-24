package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"

	"gorm.io/gorm"
)

type (
	MenuService interface {
		GetAllMenus(pageable util.Pageable) (*util.Page, error)
		CreateMenu(menu *domain.MenuPayload) *domain.Response
		UpdateMenu(menu *domain.MenuPayload, menu_id int) *domain.Response
		DeleteMenu(menu_id int) *domain.Response
		GetMenuDetail(menu_id int) *domain.Response
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

func (s *menuService) CreateMenu(menu *domain.MenuPayload) *domain.Response {
	return s.menuRepo.CreateMenu(menu)
}

func (s *menuService) UpdateMenu(menu *domain.MenuPayload, menu_id int) *domain.Response {
	return s.menuRepo.UpdateMenu(menu, menu_id)
}

func (s *menuService) DeleteMenu(menu_id int) *domain.Response {
	return s.menuRepo.DeleteMenu(menu_id)
}

func (s *menuService) GetMenuDetail(menu_id int) *domain.Response {
	return s.menuRepo.GetMenu(menu_id)
}
