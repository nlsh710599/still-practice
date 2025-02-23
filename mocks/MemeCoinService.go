// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	result "github.com/nlsh710599/still-practice/internal/result"
	mock "github.com/stretchr/testify/mock"
)

// MemeCoinService is an autogenerated mock type for the MemeCoinService type
type MemeCoinService struct {
	mock.Mock
}

// CreateMemeCoin provides a mock function with given fields: ctx, name, description
func (_m *MemeCoinService) CreateMemeCoin(ctx context.Context, name string, description string) error {
	ret := _m.Called(ctx, name, description)

	if len(ret) == 0 {
		panic("no return value specified for CreateMemeCoin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, name, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMemeCoin provides a mock function with given fields: ctx, id
func (_m *MemeCoinService) DeleteMemeCoin(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMemeCoin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMemeCoinById provides a mock function with given fields: ctx, id
func (_m *MemeCoinService) GetMemeCoinById(ctx context.Context, id uint) (*result.GetMemeCoinResult, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetMemeCoinById")
	}

	var r0 *result.GetMemeCoinResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*result.GetMemeCoinResult, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *result.GetMemeCoinResult); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*result.GetMemeCoinResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PokeMemeCoin provides a mock function with given fields: ctx, id
func (_m *MemeCoinService) PokeMemeCoin(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for PokeMemeCoin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMemeCoin provides a mock function with given fields: ctx, id, description
func (_m *MemeCoinService) UpdateMemeCoin(ctx context.Context, id uint, description string) error {
	ret := _m.Called(ctx, id, description)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMemeCoin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) error); ok {
		r0 = rf(ctx, id, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMemeCoinService creates a new instance of MemeCoinService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMemeCoinService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MemeCoinService {
	mock := &MemeCoinService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
