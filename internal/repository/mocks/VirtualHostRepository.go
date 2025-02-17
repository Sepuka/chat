// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/sepuka/chat/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// VirtualHostRepository is an autogenerated mock type for the VirtualHostRepository type
type VirtualHostRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m VirtualHostRepository) Add(_a0 *domain.Pool, _a1 *domain.Client) (*domain.VirtualHost, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.VirtualHost
	if rf, ok := ret.Get(0).(func(*domain.Pool, *domain.Client) *domain.VirtualHost); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.VirtualHost)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Pool, *domain.Client) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByContainerId provides a mock function with given fields: _a0
func (_m VirtualHostRepository) GetByContainerId(_a0 string) (*domain.VirtualHost, error) {
	ret := _m.Called(_a0)

	var r0 *domain.VirtualHost
	if rf, ok := ret.Get(0).(func(string) *domain.VirtualHost); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.VirtualHost)
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

// GetUsersHosts provides a mock function with given fields: _a0
func (_m VirtualHostRepository) GetUsersHosts(_a0 *domain.Client) ([]*domain.VirtualHost, error) {
	ret := _m.Called(_a0)

	var r0 []*domain.VirtualHost
	if rf, ok := ret.Get(0).(func(*domain.Client) []*domain.VirtualHost); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.VirtualHost)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: host
func (_m VirtualHostRepository) Update(host *domain.VirtualHost) error {
	ret := _m.Called(host)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.VirtualHost) error); ok {
		r0 = rf(host)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
