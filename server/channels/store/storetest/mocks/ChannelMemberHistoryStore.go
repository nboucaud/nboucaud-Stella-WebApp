// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// ChannelMemberHistoryStore is an autogenerated mock type for the ChannelMemberHistoryStore type
type ChannelMemberHistoryStore struct {
	mock.Mock
}

// DeleteOrphanedRows provides a mock function with given fields: limit
func (_m *ChannelMemberHistoryStore) DeleteOrphanedRows(limit int) (int64, error) {
	ret := _m.Called(limit)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOrphanedRows")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int64, error)); ok {
		return rf(limit)
	}
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelsLeftSince provides a mock function with given fields: userID, since
func (_m *ChannelMemberHistoryStore) GetChannelsLeftSince(userID string, since int64) ([]string, error) {
	ret := _m.Called(userID, since)

	if len(ret) == 0 {
		panic("no return value specified for GetChannelsLeftSince")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64) ([]string, error)); ok {
		return rf(userID, since)
	}
	if rf, ok := ret.Get(0).(func(string, int64) []string); ok {
		r0 = rf(userID, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(userID, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersInChannelDuring provides a mock function with given fields: startTime, endTime, channelID
func (_m *ChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelID string) ([]*model.ChannelMemberHistoryResult, error) {
	ret := _m.Called(startTime, endTime, channelID)

	if len(ret) == 0 {
		panic("no return value specified for GetUsersInChannelDuring")
	}

	var r0 []*model.ChannelMemberHistoryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64, string) ([]*model.ChannelMemberHistoryResult, error)); ok {
		return rf(startTime, endTime, channelID)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, string) []*model.ChannelMemberHistoryResult); ok {
		r0 = rf(startTime, endTime, channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelMemberHistoryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64, string) error); ok {
		r1 = rf(startTime, endTime, channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogJoinEvent provides a mock function with given fields: userID, channelID, joinTime
func (_m *ChannelMemberHistoryStore) LogJoinEvent(userID string, channelID string, joinTime int64) error {
	ret := _m.Called(userID, channelID, joinTime)

	if len(ret) == 0 {
		panic("no return value specified for LogJoinEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(userID, channelID, joinTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogLeaveEvent provides a mock function with given fields: userID, channelID, leaveTime
func (_m *ChannelMemberHistoryStore) LogLeaveEvent(userID string, channelID string, leaveTime int64) error {
	ret := _m.Called(userID, channelID, leaveTime)

	if len(ret) == 0 {
		panic("no return value specified for LogLeaveEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(userID, channelID, leaveTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *ChannelMemberHistoryStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	ret := _m.Called(endTime, limit)

	if len(ret) == 0 {
		panic("no return value specified for PermanentDeleteBatch")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64) (int64, error)); ok {
		return rf(endTime, limit)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(endTime, limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteBatchForRetentionPolicies provides a mock function with given fields: now, globalPolicyEndTime, limit, cursor
func (_m *ChannelMemberHistoryStore) PermanentDeleteBatchForRetentionPolicies(now int64, globalPolicyEndTime int64, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	ret := _m.Called(now, globalPolicyEndTime, limit, cursor)

	if len(ret) == 0 {
		panic("no return value specified for PermanentDeleteBatchForRetentionPolicies")
	}

	var r0 int64
	var r1 model.RetentionPolicyCursor
	var r2 error
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error)); ok {
		return rf(now, globalPolicyEndTime, limit, cursor)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) int64); ok {
		r0 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64, int64, model.RetentionPolicyCursor) model.RetentionPolicyCursor); ok {
		r1 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r1 = ret.Get(1).(model.RetentionPolicyCursor)
	}

	if rf, ok := ret.Get(2).(func(int64, int64, int64, model.RetentionPolicyCursor) error); ok {
		r2 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewChannelMemberHistoryStore creates a new instance of ChannelMemberHistoryStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChannelMemberHistoryStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChannelMemberHistoryStore {
	mock := &ChannelMemberHistoryStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
