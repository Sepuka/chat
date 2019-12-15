// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/sepuka/chat/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// ClientRepository is an autogenerated mock type for the ClientRepository type
type ClientRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m ClientRepository) Add(_a0 string, _a1 domain.ClientSource) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.ClientSource) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByLogin provides a mock function with given fields: _a0
func (_m ClientRepository) GetByLogin(_a0 string) (*domain.Client, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Client
	if rf, ok := ret.Get(0).(func(string) *domain.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
