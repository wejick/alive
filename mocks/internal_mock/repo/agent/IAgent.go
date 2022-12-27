// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/wejick/alive/internal/model"
)

// IAgent is an autogenerated mock type for the IAgent type
type IAgent struct {
	mock.Mock
}

// AddAgent provides a mock function with given fields: _a0
func (_m *IAgent) AddAgent(_a0 model.Agent) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Agent) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAgentIDToSetActive provides a mock function with given fields: lastPingThreshold
func (_m *IAgent) GetAgentIDToSetActive(lastPingThreshold int) ([]int64, error) {
	ret := _m.Called(lastPingThreshold)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(int) []int64); ok {
		r0 = rf(lastPingThreshold)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(lastPingThreshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgentIDToSetUnhealthy provides a mock function with given fields: lastPingThreshold
func (_m *IAgent) GetAgentIDToSetUnhealthy(lastPingThreshold int) ([]int64, error) {
	ret := _m.Called(lastPingThreshold)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(int) []int64); ok {
		r0 = rf(lastPingThreshold)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(lastPingThreshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgents provides a mock function with given fields: agentIDs
func (_m *IAgent) GetAgents(agentIDs []string) []model.Agent {
	ret := _m.Called(agentIDs)

	var r0 []model.Agent
	if rf, ok := ret.Get(0).(func([]string) []model.Agent); ok {
		r0 = rf(agentIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Agent)
		}
	}

	return r0
}

// Ping provides a mock function with given fields: agentID
func (_m *IAgent) Ping(agentID string) error {
	ret := _m.Called(agentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(agentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAgentStatus provides a mock function with given fields: agentIDs, status
func (_m *IAgent) SetAgentStatus(agentIDs []string, status model.AgentStatus) error {
	ret := _m.Called(agentIDs, status)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, model.AgentStatus) error); ok {
		r0 = rf(agentIDs, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIAgent interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAgent creates a new instance of IAgent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAgent(t mockConstructorTestingTNewIAgent) *IAgent {
	mock := &IAgent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}