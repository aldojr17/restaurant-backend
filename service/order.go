package service

import (
	"final-project-backend/repository"
	"final-project-backend/util"

	"gorm.io/gorm"
)

type (
	OrderService interface {
		GetAllOrders(pageable util.Pageable, user_id string) (*util.Page, error)
	}

	orderService struct {
		db        *gorm.DB
		orderRepo repository.OrderRepository
	}
)

func NewOrderService(db *gorm.DB, orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		db:        db,
		orderRepo: orderRepo,
	}
}

func (s *orderService) GetAllOrders(pageable util.Pageable, user_id string) (*util.Page, error) {
	return s.orderRepo.GetAllOrders(pageable, user_id)
}
