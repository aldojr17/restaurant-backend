// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "final-project-backend/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// AddOrDeleteMenuFavorite provides a mock function with given fields: payload
func (_m *UserService) AddOrDeleteMenuFavorite(payload *domain.UserFavorite) *domain.Response {
	ret := _m.Called(payload)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.UserFavorite) *domain.Response); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// GetCoupons provides a mock function with given fields: user_id
func (_m *UserService) GetCoupons(user_id string) *domain.Response {
	ret := _m.Called(user_id)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(string) *domain.Response); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// GetProfile provides a mock function with given fields: user_id
func (_m *UserService) GetProfile(user_id string) *domain.Response {
	ret := _m.Called(user_id)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(string) *domain.Response); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// UpdateUserData provides a mock function with given fields: payload
func (_m *UserService) UpdateUserData(payload *domain.UserProfile) *domain.Response {
	ret := _m.Called(payload)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.UserProfile) *domain.Response); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
