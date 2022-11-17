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
