// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	login "capstone/happyApp/features/login"

	mock "github.com/stretchr/testify/mock"
)

// DataLogin is an autogenerated mock type for the DataInterface type
type DataLogin struct {
	mock.Mock
}

// LoginUser provides a mock function with given fields: email
func (_m *DataLogin) LoginUser(email string) (login.Core, error) {
	ret := _m.Called(email)

	var r0 login.Core
	if rf, ok := ret.Get(0).(func(string) login.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(login.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDataLogin interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataLogin creates a new instance of DataLogin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataLogin(t mockConstructorTestingTNewDataLogin) *DataLogin {
	mock := &DataLogin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
