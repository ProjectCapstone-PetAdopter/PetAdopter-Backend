// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"
	domain "petadopter/domain"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userID
func (_m *UserUseCase) Delete(userID int) int {
	ret := _m.Called(userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetProfile provides a mock function with given fields: id
func (_m *UserUseCase) GetProfile(id int) (map[string]interface{}, int) {
	ret := _m.Called(id)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(int) map[string]interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// GetProfileID provides a mock function with given fields: userid
func (_m *UserUseCase) GetProfileID(userid int) (map[string]interface{}, int) {
	ret := _m.Called(userid)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(int) map[string]interface{}); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// Login provides a mock function with given fields: userdata, token
func (_m *UserUseCase) Login(userdata domain.User, token *oauth2.Token) (map[string]interface{}, int) {
	ret := _m.Called(userdata, token)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(domain.User, *oauth2.Token) map[string]interface{}); ok {
		r0 = rf(userdata, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(domain.User, *oauth2.Token) int); ok {
		r1 = rf(userdata, token)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: newuser, cost, token, ui
func (_m *UserUseCase) RegisterUser(newuser domain.User, cost int, token *oauth2.Token, ui domain.UserInfo) (map[string]interface{}, int) {
	ret := _m.Called(newuser, cost, token, ui)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(domain.User, int, *oauth2.Token, domain.UserInfo) map[string]interface{}); ok {
		r0 = rf(newuser, cost, token, ui)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(domain.User, int, *oauth2.Token, domain.UserInfo) int); ok {
		r1 = rf(newuser, cost, token, ui)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: newuser, userid, cost, form
func (_m *UserUseCase) UpdateUser(newuser domain.User, userid int, cost int, form *multipart.FileHeader) int {
	ret := _m.Called(newuser, userid, cost, form)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.User, int, int, *multipart.FileHeader) int); ok {
		r0 = rf(newuser, userid, cost, form)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUseCase(t mockConstructorTestingTNewUserUseCase) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
