// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "final-project-backend/domain"

	mock "github.com/stretchr/testify/mock"

	util "final-project-backend/util"
)

// OrderService is an autogenerated mock type for the OrderService type
type OrderService struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: order
func (_m *OrderService) CreateOrder(order *domain.OrderPayload) *domain.Response {
	ret := _m.Called(order)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.OrderPayload) *domain.Response); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// GetAllOrders provides a mock function with given fields: pageable
func (_m *OrderService) GetAllOrders(pageable util.Pageable) (*util.Page, error) {
	ret := _m.Called(pageable)

	var r0 *util.Page
	if rf, ok := ret.Get(0).(func(util.Pageable) *util.Page); ok {
		r0 = rf(pageable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.Page)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(util.Pageable) error); ok {
		r1 = rf(pageable)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUserOrders provides a mock function with given fields: pageable, user_id
func (_m *OrderService) GetAllUserOrders(pageable util.Pageable, user_id string) (*util.Page, error) {
	ret := _m.Called(pageable, user_id)

	var r0 *util.Page
	if rf, ok := ret.Get(0).(func(util.Pageable, string) *util.Page); ok {
		r0 = rf(pageable, user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.Page)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(util.Pageable, string) error); ok {
		r1 = rf(pageable, user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrderStatus provides a mock function with given fields: order
func (_m *OrderService) UpdateOrderStatus(order *domain.OrderStatusPayload) *domain.Response {
	ret := _m.Called(order)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.OrderStatusPayload) *domain.Response); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

type mockConstructorTestingTNewOrderService interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderService creates a new instance of OrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderService(t mockConstructorTestingTNewOrderService) *OrderService {
	mock := &OrderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
