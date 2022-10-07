// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	cart "capstone/happyApp/features/cart"

	mock "github.com/stretchr/testify/mock"
)

// CartData is an autogenerated mock type for the DataInterface type
type CartData struct {
	mock.Mock
}

// CheckCommunity provides a mock function with given fields: prodid
func (_m *CartData) CheckCommunity(prodid int) (int, string, error) {
	ret := _m.Called(prodid)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(prodid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int) string); ok {
		r1 = rf(prodid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(prodid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CheckMember provides a mock function with given fields: userid, communityid
func (_m *CartData) CheckMember(userid int, communityid int) (int, string, error) {
	ret := _m.Called(userid, communityid)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(userid, communityid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int, int) string); ok {
		r1 = rf(userid, communityid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(userid, communityid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CheckStock provides a mock function with given fields: _a0
func (_m *CartData) CheckStock(_a0 []int) ([]int, string, error) {
	ret := _m.Called(_a0)

	var r0 []int
	if rf, ok := ret.Get(0).(func([]int) []int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func([]int) string); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]int) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteCart provides a mock function with given fields: core
func (_m *CartData) DeleteCart(core cart.CoreHistory) (string, error) {
	ret := _m.Called(core)

	var r0 string
	if rf, ok := ret.Get(0).(func(cart.CoreHistory) string); ok {
		r0 = rf(core)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cart.CoreHistory) error); ok {
		r1 = rf(core)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFromCart provides a mock function with given fields: userid, cartid
func (_m *CartData) DeleteFromCart(userid int, cartid int) (string, error) {
	ret := _m.Called(userid, cartid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(userid, cartid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userid, cartid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommunity provides a mock function with given fields: communityid
func (_m *CartData) GetCommunity(communityid int) (cart.CoreCommunity, string, error) {
	ret := _m.Called(communityid)

	var r0 cart.CoreCommunity
	if rf, ok := ret.Get(0).(func(int) cart.CoreCommunity); ok {
		r0 = rf(communityid)
	} else {
		r0 = ret.Get(0).(cart.CoreCommunity)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int) string); ok {
		r1 = rf(communityid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(communityid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTotalTransaction provides a mock function with given fields: trasacid
func (_m *CartData) GetTotalTransaction(trasacid int) (int, string, error) {
	ret := _m.Called(trasacid)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(trasacid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int) string); ok {
		r1 = rf(trasacid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(trasacid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserRole provides a mock function with given fields: Userid, communityid
func (_m *CartData) GetUserRole(Userid int, communityid int) (string, error) {
	ret := _m.Called(Userid, communityid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(Userid, communityid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(Userid, communityid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertIntoCart provides a mock function with given fields: userid, productid
func (_m *CartData) InsertIntoCart(userid int, productid int) (string, error) {
	ret := _m.Called(userid, productid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(userid, productid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userid, productid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertIntoTransaction provides a mock function with given fields: core
func (_m *CartData) InsertIntoTransaction(core cart.CoreHistory) (int, string, error) {
	ret := _m.Called(core)

	var r0 int
	if rf, ok := ret.Get(0).(func(cart.CoreHistory) int); ok {
		r0 = rf(core)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(cart.CoreHistory) string); ok {
		r1 = rf(core)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(cart.CoreHistory) error); ok {
		r2 = rf(core)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListHistoryProduct provides a mock function with given fields: communityid
func (_m *CartData) ListHistoryProduct(communityid int) ([]cart.CoreProductResponse, string, error) {
	ret := _m.Called(communityid)

	var r0 []cart.CoreProductResponse
	if rf, ok := ret.Get(0).(func(int) []cart.CoreProductResponse); ok {
		r0 = rf(communityid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.CoreProductResponse)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int) string); ok {
		r1 = rf(communityid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(communityid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectCartList provides a mock function with given fields: userid, communityid
func (_m *CartData) SelectCartList(userid int, communityid int) ([]cart.CoreCart, string, error) {
	ret := _m.Called(userid, communityid)

	var r0 []cart.CoreCart
	if rf, ok := ret.Get(0).(func(int, int) []cart.CoreCart); ok {
		r0 = rf(userid, communityid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.CoreCart)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int, int) string); ok {
		r1 = rf(userid, communityid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(userid, communityid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectCommunity provides a mock function with given fields: communityid
func (_m *CartData) SelectCommunity(communityid int) (cart.CoreCommunity, string, error) {
	ret := _m.Called(communityid)

	var r0 cart.CoreCommunity
	if rf, ok := ret.Get(0).(func(int) cart.CoreCommunity); ok {
		r0 = rf(communityid)
	} else {
		r0 = ret.Get(0).(cart.CoreCommunity)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int) string); ok {
		r1 = rf(communityid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(communityid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateHistory provides a mock function with given fields: _a0
func (_m *CartData) UpdateHistory(_a0 cart.CoreHistory) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(cart.CoreHistory) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cart.CoreHistory) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStock provides a mock function with given fields: _a0
func (_m *CartData) UpdateStock(_a0 []int) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func([]int) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCartData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartData creates a new instance of CartData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartData(t mockConstructorTestingTNewCartData) *CartData {
	mock := &CartData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}