// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "final-project-backend/domain"

	mock "github.com/stretchr/testify/mock"
)

// CouponService is an autogenerated mock type for the CouponService type
type CouponService struct {
	mock.Mock
}

// CreateCoupon provides a mock function with given fields: coupon
func (_m *CouponService) CreateCoupon(coupon *domain.Coupon) *domain.Response {
	ret := _m.Called(coupon)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.Coupon) *domain.Response); ok {
		r0 = rf(coupon)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// DeleteCoupon provides a mock function with given fields: id
func (_m *CouponService) DeleteCoupon(id string) *domain.Response {
	ret := _m.Called(id)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(string) *domain.Response); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// GetAllCoupon provides a mock function with given fields:
func (_m *CouponService) GetAllCoupon() *domain.Response {
	ret := _m.Called()

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func() *domain.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// GetCoupon provides a mock function with given fields: id
func (_m *CouponService) GetCoupon(id string) *domain.Response {
	ret := _m.Called(id)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(string) *domain.Response); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// UpdateCoupon provides a mock function with given fields: coupon
func (_m *CouponService) UpdateCoupon(coupon *domain.Coupon) *domain.Response {
	ret := _m.Called(coupon)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.Coupon) *domain.Response); ok {
		r0 = rf(coupon)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

type mockConstructorTestingTNewCouponService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCouponService creates a new instance of CouponService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCouponService(t mockConstructorTestingTNewCouponService) *CouponService {
	mock := &CouponService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
