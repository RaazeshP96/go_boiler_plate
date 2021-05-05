// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "prototype2/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserService) Create(user *domain.User) (*domain.User, error) {
	ret := _m.Called(user)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *UserService) FindAll() ([]domain.User, error) {
	ret := _m.Called()

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: user
func (_m *UserService) Validate(user *domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateAge provides a mock function with given fields: user
func (_m *UserService) ValidateAge(user *domain.User) bool {
	ret := _m.Called(user)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.User) bool); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
