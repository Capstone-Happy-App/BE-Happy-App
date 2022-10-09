// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	community "capstone/happyApp/features/community"

	mock "github.com/stretchr/testify/mock"
)

// UsecaseInterface is an autogenerated mock type for the UsecaseInterface type
type UsecaseInterface struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: _a0
func (_m *UsecaseInterface) AddComment(_a0 community.CoreComment) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(community.CoreComment) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(community.CoreComment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddNewCommunity provides a mock function with given fields: userid, core
func (_m *UsecaseInterface) AddNewCommunity(userid int, core community.CoreCommunity) (string, error) {
	ret := _m.Called(userid, core)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, community.CoreCommunity) string); ok {
		r0 = rf(userid, core)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, community.CoreCommunity) error); ok {
		r1 = rf(userid, core)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommunityFeed provides a mock function with given fields: userid, communityid
func (_m *UsecaseInterface) GetCommunityFeed(userid int, communityid int) (community.CoreCommunity, string, error) {
	ret := _m.Called(userid, communityid)

	var r0 community.CoreCommunity
	if rf, ok := ret.Get(0).(func(int, int) community.CoreCommunity); ok {
		r0 = rf(userid, communityid)
	} else {
		r0 = ret.Get(0).(community.CoreCommunity)
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

// GetDetailFeed provides a mock function with given fields: feedid
func (_m *UsecaseInterface) GetDetailFeed(feedid int) (community.CoreFeed, string, error) {
	ret := _m.Called(feedid)

	var r0 community.CoreFeed
	if rf, ok := ret.Get(0).(func(int) community.CoreFeed); ok {
		r0 = rf(feedid)
	} else {
		r0 = ret.Get(0).(community.CoreFeed)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(int) string); ok {
		r1 = rf(feedid)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(feedid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetListCommunity provides a mock function with given fields:
func (_m *UsecaseInterface) GetListCommunity() ([]community.CoreCommunity, string, error) {
	ret := _m.Called()

	var r0 []community.CoreCommunity
	if rf, ok := ret.Get(0).(func() []community.CoreCommunity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]community.CoreCommunity)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetListCommunityWithParam provides a mock function with given fields: param
func (_m *UsecaseInterface) GetListCommunityWithParam(param string) ([]community.CoreCommunity, string, error) {
	ret := _m.Called(param)

	var r0 []community.CoreCommunity
	if rf, ok := ret.Get(0).(func(string) []community.CoreCommunity); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]community.CoreCommunity)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetMembers provides a mock function with given fields: communityid
func (_m *UsecaseInterface) GetMembers(communityid int) ([]string, string, error) {
	ret := _m.Called(communityid)

	var r0 []string
	if rf, ok := ret.Get(0).(func(int) []string); ok {
		r0 = rf(communityid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// JoinCommunity provides a mock function with given fields: userid, communityid
func (_m *UsecaseInterface) JoinCommunity(userid int, communityid int) (string, error) {
	ret := _m.Called(userid, communityid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(userid, communityid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userid, communityid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Leave provides a mock function with given fields: userid, communityid
func (_m *UsecaseInterface) Leave(userid int, communityid int) (string, error) {
	ret := _m.Called(userid, communityid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(userid, communityid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userid, communityid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostFeed provides a mock function with given fields: _a0
func (_m *UsecaseInterface) PostFeed(_a0 community.CoreFeed) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(community.CoreFeed) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(community.CoreFeed) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCommunity provides a mock function with given fields: userid, core
func (_m *UsecaseInterface) UpdateCommunity(userid int, core community.CoreCommunity) (string, error) {
	ret := _m.Called(userid, core)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, community.CoreCommunity) string); ok {
		r0 = rf(userid, core)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, community.CoreCommunity) error); ok {
		r1 = rf(userid, core)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsecaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecaseInterface creates a new instance of UsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecaseInterface(t mockConstructorTestingTNewUsecaseInterface) *UsecaseInterface {
	mock := &UsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
