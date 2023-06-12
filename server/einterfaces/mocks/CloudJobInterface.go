// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// CloudJobInterface is an autogenerated mock type for the CloudJobInterface type
type CloudJobInterface struct {
	mock.Mock
}

// MakeScheduler provides a mock function with given fields:
func (_m *CloudJobInterface) MakeScheduler() model.Scheduler {
	ret := _m.Called()

	var r0 model.Scheduler
	if rf, ok := ret.Get(0).(func() model.Scheduler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Scheduler)
		}
	}

	return r0
}

// MakeWorker provides a mock function with given fields:
func (_m *CloudJobInterface) MakeWorker() model.Worker {
	ret := _m.Called()

	var r0 model.Worker
	if rf, ok := ret.Get(0).(func() model.Worker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Worker)
		}
	}

	return r0
}

type mockConstructorTestingTNewCloudJobInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCloudJobInterface creates a new instance of CloudJobInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCloudJobInterface(t mockConstructorTestingTNewCloudJobInterface) *CloudJobInterface {
	mock := &CloudJobInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
