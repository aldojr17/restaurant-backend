package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	OrderService interface {
		GetAllUserOrders(pageable util.Pageable, user_id string) (*util.Page, error)
		GetAllOrders(pageable util.Pageable) (*util.Page, error)
		UpdateOrderStatus(order *domain.OrderStatusPayload) *domain.Response
		CreateOrder(order *domain.OrderPayload) *domain.Response
		// CreateOrderDetails(orders *domain.OrderDetails) *domain.Response
	}

	orderService struct {
		db         *gorm.DB
		orderRepo  repository.OrderRepository
		couponRepo repository.CouponRepository
	}
)

func NewOrderService(db *gorm.DB, orderRepo repository.OrderRepository, couponRepo repository.CouponRepository) OrderService {
	return &orderService{
		db:         db,
		orderRepo:  orderRepo,
		couponRepo: couponRepo,
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

func (s *orderService) CreateOrder(order *domain.OrderPayload) *domain.Response {
	if order.CouponId == nil {
		response := s.orderRepo.CreateOrder(order)

		if response.Err != nil {
			return response
		}

		for index := range order.Orders {
			order.Orders[index].OrderId = response.Data.(*domain.OrderPayload).Id
		}

		if response := s.orderRepo.CreateOrderDetails(&order.Orders); response.Err != nil {
			return response
		}

		return response
	}

	if response := s.couponRepo.GetValidCoupon(order.UserId, *order.CouponId); response.Err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrCouponInvalid)
	}

	response := s.orderRepo.CreateOrder(order)

	if response.Err != nil {
		return response
	}

	for index := range order.Orders {
		order.Orders[index].OrderId = response.Data.(*domain.OrderPayload).Id
	}

	if response := s.orderRepo.CreateOrderDetails(&order.Orders); response.Err != nil {
		return response
	}

	if err := s.couponRepo.ReduceQty(order.UserId, *order.CouponId); err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return response
}

// func (s *orderService) CreateOrderDetails(orders *domain.OrderDetails) *domain.Response {
// 	return s.orderRepo.CreateOrderDetails(orders)
// }
