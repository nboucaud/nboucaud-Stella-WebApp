// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make sharedchannel-mocks`.

package sharedchannel

import (
	filestore "github.com/mattermost/mattermost/server/v8/platform/shared/filestore"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost/server/public/model"

	request "github.com/mattermost/mattermost/server/public/shared/request"
)

// MockAppIface is an autogenerated mock type for the AppIface type
type MockAppIface struct {
	mock.Mock
}

// AddUserToChannel provides a mock function with given fields: c, user, channel, skipTeamMemberIntegrityCheck
func (_m *MockAppIface) AddUserToChannel(c request.CTX, user *model.User, channel *model.Channel, skipTeamMemberIntegrityCheck bool) (*model.ChannelMember, *model.AppError) {
	ret := _m.Called(c, user, channel, skipTeamMemberIntegrityCheck)

	var r0 *model.ChannelMember
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.User, *model.Channel, bool) (*model.ChannelMember, *model.AppError)); ok {
		return rf(c, user, channel, skipTeamMemberIntegrityCheck)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.User, *model.Channel, bool) *model.ChannelMember); ok {
		r0 = rf(c, user, channel, skipTeamMemberIntegrityCheck)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ChannelMember)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.User, *model.Channel, bool) *model.AppError); ok {
		r1 = rf(c, user, channel, skipTeamMemberIntegrityCheck)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// AddUserToTeamByTeamId provides a mock function with given fields: c, teamId, user
func (_m *MockAppIface) AddUserToTeamByTeamId(c request.CTX, teamId string, user *model.User) *model.AppError {
	ret := _m.Called(c, teamId, user)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, *model.User) *model.AppError); ok {
		r0 = rf(c, teamId, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// CreateChannelWithUser provides a mock function with given fields: c, channel, userId
func (_m *MockAppIface) CreateChannelWithUser(c request.CTX, channel *model.Channel, userId string) (*model.Channel, *model.AppError) {
	ret := _m.Called(c, channel, userId)

	var r0 *model.Channel
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Channel, string) (*model.Channel, *model.AppError)); ok {
		return rf(c, channel, userId)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Channel, string) *model.Channel); ok {
		r0 = rf(c, channel, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Channel)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Channel, string) *model.AppError); ok {
		r1 = rf(c, channel, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CreatePost provides a mock function with given fields: c, post, channel, triggerWebhooks, setOnline
func (_m *MockAppIface) CreatePost(c request.CTX, post *model.Post, channel *model.Channel, triggerWebhooks bool, setOnline bool) (*model.Post, *model.AppError) {
	ret := _m.Called(c, post, channel, triggerWebhooks, setOnline)

	var r0 *model.Post
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post, *model.Channel, bool, bool) (*model.Post, *model.AppError)); ok {
		return rf(c, post, channel, triggerWebhooks, setOnline)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post, *model.Channel, bool, bool) *model.Post); ok {
		r0 = rf(c, post, channel, triggerWebhooks, setOnline)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Post, *model.Channel, bool, bool) *model.AppError); ok {
		r1 = rf(c, post, channel, triggerWebhooks, setOnline)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CreateUploadSession provides a mock function with given fields: c, us
func (_m *MockAppIface) CreateUploadSession(c request.CTX, us *model.UploadSession) (*model.UploadSession, *model.AppError) {
	ret := _m.Called(c, us)

	var r0 *model.UploadSession
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.UploadSession) (*model.UploadSession, *model.AppError)); ok {
		return rf(c, us)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.UploadSession) *model.UploadSession); ok {
		r0 = rf(c, us)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UploadSession)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.UploadSession) *model.AppError); ok {
		r1 = rf(c, us)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// DeletePost provides a mock function with given fields: c, postID, deleteByID
func (_m *MockAppIface) DeletePost(c request.CTX, postID string, deleteByID string) (*model.Post, *model.AppError) {
	ret := _m.Called(c, postID, deleteByID)

	var r0 *model.Post
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) (*model.Post, *model.AppError)); ok {
		return rf(c, postID, deleteByID)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) *model.Post); ok {
		r0 = rf(c, postID, deleteByID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string, string) *model.AppError); ok {
		r1 = rf(c, postID, deleteByID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// DeleteReactionForPost provides a mock function with given fields: c, reaction
func (_m *MockAppIface) DeleteReactionForPost(c request.CTX, reaction *model.Reaction) *model.AppError {
	ret := _m.Called(c, reaction)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Reaction) *model.AppError); ok {
		r0 = rf(c, reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// FileReader provides a mock function with given fields: path
func (_m *MockAppIface) FileReader(path string) (filestore.ReadCloseSeeker, *model.AppError) {
	ret := _m.Called(path)

	var r0 filestore.ReadCloseSeeker
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (filestore.ReadCloseSeeker, *model.AppError)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) filestore.ReadCloseSeeker); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filestore.ReadCloseSeeker)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(path)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOrCreateDirectChannel provides a mock function with given fields: c, userId, otherUserId, channelOptions
func (_m *MockAppIface) GetOrCreateDirectChannel(c request.CTX, userId string, otherUserId string, channelOptions ...model.ChannelOption) (*model.Channel, *model.AppError) {
	_va := make([]interface{}, len(channelOptions))
	for _i := range channelOptions {
		_va[_i] = channelOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, c, userId, otherUserId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.Channel
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string, ...model.ChannelOption) (*model.Channel, *model.AppError)); ok {
		return rf(c, userId, otherUserId, channelOptions...)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string, string, ...model.ChannelOption) *model.Channel); ok {
		r0 = rf(c, userId, otherUserId, channelOptions...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Channel)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string, string, ...model.ChannelOption) *model.AppError); ok {
		r1 = rf(c, userId, otherUserId, channelOptions...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfileImage provides a mock function with given fields: user
func (_m *MockAppIface) GetProfileImage(user *model.User) ([]byte, bool, *model.AppError) {
	ret := _m.Called(user)

	var r0 []byte
	var r1 bool
	var r2 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User) ([]byte, bool, *model.AppError)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*model.User) []byte); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) bool); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(*model.User) *model.AppError); ok {
		r2 = rf(user)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// InvalidateCacheForUser provides a mock function with given fields: userID
func (_m *MockAppIface) InvalidateCacheForUser(userID string) {
	_m.Called(userID)
}

// MentionsToTeamMembers provides a mock function with given fields: c, message, teamID
func (_m *MockAppIface) MentionsToTeamMembers(c request.CTX, message string, teamID string) model.UserMentionMap {
	ret := _m.Called(c, message, teamID)

	var r0 model.UserMentionMap
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) model.UserMentionMap); ok {
		r0 = rf(c, message, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.UserMentionMap)
		}
	}

	return r0
}

// NotifySharedChannelUserUpdate provides a mock function with given fields: user
func (_m *MockAppIface) NotifySharedChannelUserUpdate(user *model.User) {
	_m.Called(user)
}

// OnSharedChannelsAttachmentSyncMsg provides a mock function with given fields: fi, post, rc
func (_m *MockAppIface) OnSharedChannelsAttachmentSyncMsg(fi *model.FileInfo, post *model.Post, rc *model.RemoteCluster) error {
	ret := _m.Called(fi, post, rc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.FileInfo, *model.Post, *model.RemoteCluster) error); ok {
		r0 = rf(fi, post, rc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnSharedChannelsProfileImageSyncMsg provides a mock function with given fields: user, rc
func (_m *MockAppIface) OnSharedChannelsProfileImageSyncMsg(user *model.User, rc *model.RemoteCluster) error {
	ret := _m.Called(user, rc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.RemoteCluster) error); ok {
		r0 = rf(user, rc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnSharedChannelsSyncMsg provides a mock function with given fields: msg, rc
func (_m *MockAppIface) OnSharedChannelsSyncMsg(msg *model.SyncMsg, rc *model.RemoteCluster) (model.SyncResponse, error) {
	ret := _m.Called(msg, rc)

	var r0 model.SyncResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.SyncMsg, *model.RemoteCluster) (model.SyncResponse, error)); ok {
		return rf(msg, rc)
	}
	if rf, ok := ret.Get(0).(func(*model.SyncMsg, *model.RemoteCluster) model.SyncResponse); ok {
		r0 = rf(msg, rc)
	} else {
		r0 = ret.Get(0).(model.SyncResponse)
	}

	if rf, ok := ret.Get(1).(func(*model.SyncMsg, *model.RemoteCluster) error); ok {
		r1 = rf(msg, rc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchChannelModerationsForChannel provides a mock function with given fields: c, channel, channelModerationsPatch
func (_m *MockAppIface) PatchChannelModerationsForChannel(c request.CTX, channel *model.Channel, channelModerationsPatch []*model.ChannelModerationPatch) ([]*model.ChannelModeration, *model.AppError) {
	ret := _m.Called(c, channel, channelModerationsPatch)

	var r0 []*model.ChannelModeration
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Channel, []*model.ChannelModerationPatch) ([]*model.ChannelModeration, *model.AppError)); ok {
		return rf(c, channel, channelModerationsPatch)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Channel, []*model.ChannelModerationPatch) []*model.ChannelModeration); ok {
		r0 = rf(c, channel, channelModerationsPatch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelModeration)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Channel, []*model.ChannelModerationPatch) *model.AppError); ok {
		r1 = rf(c, channel, channelModerationsPatch)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// PermanentDeleteChannel provides a mock function with given fields: c, channel
func (_m *MockAppIface) PermanentDeleteChannel(c request.CTX, channel *model.Channel) *model.AppError {
	ret := _m.Called(c, channel)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Channel) *model.AppError); ok {
		r0 = rf(c, channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// SaveReactionForPost provides a mock function with given fields: c, reaction
func (_m *MockAppIface) SaveReactionForPost(c request.CTX, reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	ret := _m.Called(c, reaction)

	var r0 *model.Reaction
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Reaction) (*model.Reaction, *model.AppError)); ok {
		return rf(c, reaction)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Reaction) *model.Reaction); ok {
		r0 = rf(c, reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Reaction)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Reaction) *model.AppError); ok {
		r1 = rf(c, reaction)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SendEphemeralPost provides a mock function with given fields: c, userId, post
func (_m *MockAppIface) SendEphemeralPost(c request.CTX, userId string, post *model.Post) *model.Post {
	ret := _m.Called(c, userId, post)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(request.CTX, string, *model.Post) *model.Post); ok {
		r0 = rf(c, userId, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	return r0
}

// UpdatePost provides a mock function with given fields: c, post, safeUpdate
func (_m *MockAppIface) UpdatePost(c request.CTX, post *model.Post, safeUpdate bool) (*model.Post, *model.AppError) {
	ret := _m.Called(c, post, safeUpdate)

	var r0 *model.Post
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post, bool) (*model.Post, *model.AppError)); ok {
		return rf(c, post, safeUpdate)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Post, bool) *model.Post); ok {
		r0 = rf(c, post, safeUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Post, bool) *model.AppError); ok {
		r1 = rf(c, post, safeUpdate)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewMockAppIface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAppIface creates a new instance of MockAppIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAppIface(t mockConstructorTestingTNewMockAppIface) *MockAppIface {
	mock := &MockAppIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
