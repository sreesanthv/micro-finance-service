// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/sreesanthv/micro-finance-service/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// AccountService is an autogenerated mock type for the AccountService type
type AccountService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *AccountService) Create(_a0 context.Context, _a1 *domain.Account) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Account) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Account) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *AccountService) Delete(_a0 context.Context, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *AccountService) Get(_a0 context.Context, _a1 int) (*domain.Account, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Account
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: _a0, _a1
func (_m *AccountService) GetByEmail(_a0 context.Context, _a1 string) (*domain.Account, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveTransaction provides a mock function with given fields: _a0, _a1
func (_m *AccountService) SaveTransaction(_a0 context.Context, _a1 *domain.Transaction) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Transaction) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *AccountService) Update(_a0 context.Context, _a1 *domain.Account) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Account) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
