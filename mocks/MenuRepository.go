// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "final-project-backend/domain"

	mock "github.com/stretchr/testify/mock"

	util "final-project-backend/util"
)

// MenuRepository is an autogenerated mock type for the MenuRepository type
type MenuRepository struct {
	mock.Mock
}

// AddMenuOption provides a mock function with given fields: options
func (_m *MenuRepository) AddMenuOption(options *[]domain.MenuOption) *domain.Response {
	ret := _m.Called(options)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*[]domain.MenuOption) *domain.Response); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// CreateMenu provides a mock function with given fields: menu
func (_m *MenuRepository) CreateMenu(menu *domain.MenuPayload) *domain.Response {
	ret := _m.Called(menu)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.MenuPayload) *domain.Response); ok {
		r0 = rf(menu)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// DeleteMenu provides a mock function with given fields: menu_id
func (_m *MenuRepository) DeleteMenu(menu_id int) *domain.Response {
	ret := _m.Called(menu_id)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(int) *domain.Response); ok {
		r0 = rf(menu_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// GetAllMenus provides a mock function with given fields: pageable
func (_m *MenuRepository) GetAllMenus(pageable util.Pageable) (*util.Page, error) {
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

// GetMenu provides a mock function with given fields: menu_id
func (_m *MenuRepository) GetMenu(menu_id int) *domain.Response {
	ret := _m.Called(menu_id)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(int) *domain.Response); ok {
		r0 = rf(menu_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// UpdateMenu provides a mock function with given fields: menu, menu_id
func (_m *MenuRepository) UpdateMenu(menu *domain.MenuPayload, menu_id int) *domain.Response {
	ret := _m.Called(menu, menu_id)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.MenuPayload, int) *domain.Response); ok {
		r0 = rf(menu, menu_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// UpdateMenuOption provides a mock function with given fields: option
func (_m *MenuRepository) UpdateMenuOption(option *domain.MenuOption) *domain.Response {
	ret := _m.Called(option)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(*domain.MenuOption) *domain.Response); ok {
		r0 = rf(option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

// UpdateMenuRating provides a mock function with given fields: menu_id, data
func (_m *MenuRepository) UpdateMenuRating(menu_id int, data map[string]interface{}) *domain.Response {
	ret := _m.Called(menu_id, data)

	var r0 *domain.Response
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) *domain.Response); ok {
		r0 = rf(menu_id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	return r0
}

type mockConstructorTestingTNewMenuRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenuRepository creates a new instance of MenuRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenuRepository(t mockConstructorTestingTNewMenuRepository) *MenuRepository {
	mock := &MenuRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
