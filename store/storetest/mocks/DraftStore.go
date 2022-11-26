// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// DraftStore is an autogenerated mock type for the DraftStore type
type DraftStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userID, channelID, rootID
func (_m *DraftStore) Delete(userID string, channelID string, rootID string) error {
	ret := _m.Called(userID, channelID, rootID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(userID, channelID, rootID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: userID, channelID, rootID, inclDeleted
func (_m *DraftStore) Get(userID string, channelID string, rootID string, inclDeleted bool) (*model.Draft, error) {
	ret := _m.Called(userID, channelID, rootID, inclDeleted)

	var r0 *model.Draft
	if rf, ok := ret.Get(0).(func(string, string, string, bool) *model.Draft); ok {
		r0 = rf(userID, channelID, rootID, inclDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Draft)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, bool) error); ok {
		r1 = rf(userID, channelID, rootID, inclDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDraftsForUser provides a mock function with given fields: userID, teamID
func (_m *DraftStore) GetDraftsForUser(userID string, teamID string) ([]*model.Draft, error) {
	ret := _m.Called(userID, teamID)

	var r0 []*model.Draft
	if rf, ok := ret.Get(0).(func(string, string) []*model.Draft); ok {
		r0 = rf(userID, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Draft)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: d
func (_m *DraftStore) Save(d *model.Draft) (*model.Draft, error) {
	ret := _m.Called(d)

	var r0 *model.Draft
	if rf, ok := ret.Get(0).(func(*model.Draft) *model.Draft); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Draft)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Draft) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: d
func (_m *DraftStore) Update(d *model.Draft) (*model.Draft, error) {
	ret := _m.Called(d)

	var r0 *model.Draft
	if rf, ok := ret.Get(0).(func(*model.Draft) *model.Draft); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Draft)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Draft) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
