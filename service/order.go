package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"

	"gorm.io/gorm"
)

type (
	OrderService interface {
		GetAllUserOrders(pageable util.Pageable, user_id string) (*util.Page, error)
		GetAllOrders(pageable util.Pageable) (*util.Page, error)
		UpdateOrderStatus(order *domain.OrderStatusPayload) *domain.Response
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

func (s *orderService) GetAllUserOrders(pageable util.Pageable, user_id string) (*util.Page, error) {
	return s.orderRepo.GetAllUserOrders(pageable, user_id)
}

func (s *orderService) GetAllOrders(pageable util.Pageable) (*util.Page, error) {
	return s.orderRepo.GetAllOrders(pageable)
}

func (s *orderService) UpdateOrderStatus(order *domain.OrderStatusPayload) *domain.Response {
	return s.orderRepo.UpdateOrderStatus(order)
}
