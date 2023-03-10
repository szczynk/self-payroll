// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "self-payrol/model"

	mock "github.com/stretchr/testify/mock"

	request "self-payrol/request"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// DestroyUser provides a mock function with given fields: ctx, id
func (_m *UserUsecase) DestroyUser(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditUser provides a mock function with given fields: ctx, id, req
func (_m *UserUsecase) EditUser(ctx context.Context, id int, req *request.UserRequest) (*model.User, error) {
	ret := _m.Called(ctx, id, req)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *request.UserRequest) (*model.User, error)); ok {
		return rf(ctx, id, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *request.UserRequest) *model.User); ok {
		r0 = rf(ctx, id, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *request.UserRequest) error); ok {
		r1 = rf(ctx, id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchUser provides a mock function with given fields: ctx, limit, offset
func (_m *UserUsecase) FetchUser(ctx context.Context, limit int, offset int) ([]*model.User, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []*model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]*model.User, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*model.User); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserUsecase) GetByID(ctx context.Context, id int) (*model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreUser provides a mock function with given fields: ctx, req
func (_m *UserUsecase) StoreUser(ctx context.Context, req *request.UserRequest) (*model.User, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.UserRequest) (*model.User, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *request.UserRequest) *model.User); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *request.UserRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithdrawSalary provides a mock function with given fields: ctx, req
func (_m *UserUsecase) WithdrawSalary(ctx context.Context, req *request.WithdrawRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.WithdrawRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
