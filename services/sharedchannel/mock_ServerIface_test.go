// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make sharedchannel-mocks`.

package sharedchannel

import (
	mlog "github.com/mattermost/mattermost-server/v5/shared/mlog"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost-server/v5/model"

	remotecluster "github.com/mattermost/mattermost-server/v5/services/remotecluster"

	store "github.com/mattermost/mattermost-server/v5/store"
)

// MockServerIface is an autogenerated mock type for the ServerIface type
type MockServerIface struct {
	mock.Mock
}

// AddClusterLeaderChangedListener provides a mock function with given fields: listener
func (_m *MockServerIface) AddClusterLeaderChangedListener(listener func()) string {
	ret := _m.Called(listener)

	var r0 string
	if rf, ok := ret.Get(0).(func(func()) string); ok {
		r0 = rf(listener)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Config provides a mock function with given fields:
func (_m *MockServerIface) Config() *model.Config {
	ret := _m.Called()

	var r0 *model.Config
	if rf, ok := ret.Get(0).(func() *model.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Config)
		}
	}

	return r0
}

// GetLogger provides a mock function with given fields:
func (_m *MockServerIface) GetLogger() mlog.LoggerIFace {
	ret := _m.Called()

	var r0 mlog.LoggerIFace
	if rf, ok := ret.Get(0).(func() mlog.LoggerIFace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mlog.LoggerIFace)
		}
	}

	return r0
}

// GetRemoteClusterService provides a mock function with given fields:
func (_m *MockServerIface) GetRemoteClusterService() *remotecluster.Service {
	ret := _m.Called()

	var r0 *remotecluster.Service
	if rf, ok := ret.Get(0).(func() *remotecluster.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*remotecluster.Service)
		}
	}

	return r0
}

// GetStore provides a mock function with given fields:
func (_m *MockServerIface) GetStore() store.Store {
	ret := _m.Called()

	var r0 store.Store
	if rf, ok := ret.Get(0).(func() store.Store); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.Store)
		}
	}

	return r0
}

// IsLeader provides a mock function with given fields:
func (_m *MockServerIface) IsLeader() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveClusterLeaderChangedListener provides a mock function with given fields: id
func (_m *MockServerIface) RemoveClusterLeaderChangedListener(id string) {
	_m.Called(id)
}
