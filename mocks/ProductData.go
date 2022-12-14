// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	product "capstone/happyApp/features/product"

	mock "github.com/stretchr/testify/mock"
)

// DataProduct is an autogenerated mock type for the DataInterface type
type DataProduct struct {
	mock.Mock
}

// DelProduct provides a mock function with given fields: idProduct, userId
func (_m *DataProduct) DelProduct(idProduct int, userId int) int {
	ret := _m.Called(idProduct, userId)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(idProduct, userId)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InsertProduct provides a mock function with given fields: _a0, _a1
func (_m *DataProduct) InsertProduct(_a0 product.ProductCore, _a1 int) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(product.ProductCore, int) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SelectProduct provides a mock function with given fields: idProduct, userId
func (_m *DataProduct) SelectProduct(idProduct int, userId int) (product.Comu, product.ProductCore, error) {
	ret := _m.Called(idProduct, userId)

	var r0 product.Comu
	if rf, ok := ret.Get(0).(func(int, int) product.Comu); ok {
		r0 = rf(idProduct, userId)
	} else {
		r0 = ret.Get(0).(product.Comu)
	}

	var r1 product.ProductCore
	if rf, ok := ret.Get(1).(func(int, int) product.ProductCore); ok {
		r1 = rf(idProduct, userId)
	} else {
		r1 = ret.Get(1).(product.ProductCore)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(idProduct, userId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectProductComu provides a mock function with given fields: idComu, userId
func (_m *DataProduct) SelectProductComu(idComu int, userId int) (product.Comu, []product.ProductCore, error) {
	ret := _m.Called(idComu, userId)

	var r0 product.Comu
	if rf, ok := ret.Get(0).(func(int, int) product.Comu); ok {
		r0 = rf(idComu, userId)
	} else {
		r0 = ret.Get(0).(product.Comu)
	}

	var r1 []product.ProductCore
	if rf, ok := ret.Get(1).(func(int, int) []product.ProductCore); ok {
		r1 = rf(idComu, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]product.ProductCore)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(idComu, userId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdtProduct provides a mock function with given fields: _a0, _a1
func (_m *DataProduct) UpdtProduct(_a0 product.ProductCore, _a1 int) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(product.ProductCore, int) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewDataProduct interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataProduct creates a new instance of DataProduct. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataProduct(t mockConstructorTestingTNewDataProduct) *DataProduct {
	mock := &DataProduct{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
