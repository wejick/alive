// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/wejick/alive/internal/model"
)

// Itest is an autogenerated mock type for the Itest type
type Itest struct {
	mock.Mock
}

// AddTest provides a mock function with given fields: _a0
func (_m *Itest) AddTest(_a0 model.Test) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Test) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTest provides a mock function with given fields: IDs, agent, rows, offset
func (_m *Itest) GetTest(IDs []string, agent string, rows int, offset int) ([]model.Test, error) {
	ret := _m.Called(IDs, agent, rows, offset)

	var r0 []model.Test
	if rf, ok := ret.Get(0).(func([]string, string, int, int) []model.Test); ok {
		r0 = rf(IDs, agent, rows, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Test)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, string, int, int) error); ok {
		r1 = rf(IDs, agent, rows, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalTest provides a mock function with given fields:
func (_m *Itest) GetTotalTest() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewItest interface {
	mock.TestingT
	Cleanup(func())
}

// NewItest creates a new instance of Itest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItest(t mockConstructorTestingTNewItest) *Itest {
	mock := &Itest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
