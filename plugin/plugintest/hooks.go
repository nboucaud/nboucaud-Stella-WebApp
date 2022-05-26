// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make plugin-mocks`.

package plugintest

import (
	io "io"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost-server/v6/model"

	plugin "github.com/mattermost/mattermost-server/v6/plugin"
)

// Hooks is an autogenerated mock type for the Hooks type
type Hooks struct {
	mock.Mock
}

// ChannelHasBeenCreated provides a mock function with given fields: c, channel
func (_m *Hooks) ChannelHasBeenCreated(c *plugin.Context, channel *model.Channel) {
	_m.Called(c, channel)
}

// ExecuteCommand provides a mock function with given fields: c, args
func (_m *Hooks) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	ret := _m.Called(c, args)

	var r0 *model.CommandResponse
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.CommandArgs) *model.CommandResponse); ok {
		r0 = rf(c, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CommandResponse)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.CommandArgs) *model.AppError); ok {
		r1 = rf(c, args)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// FileWillBeUploaded provides a mock function with given fields: c, info, file, output
func (_m *Hooks) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	ret := _m.Called(c, info, file, output)

	var r0 *model.FileInfo
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.FileInfo, io.Reader, io.Writer) *model.FileInfo); ok {
		r0 = rf(c, info, file, output)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfo)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.FileInfo, io.Reader, io.Writer) string); ok {
		r1 = rf(c, info, file, output)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// Implemented provides a mock function with given fields:
func (_m *Hooks) Implemented() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MessageHasBeenPosted provides a mock function with given fields: c, post
func (_m *Hooks) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	_m.Called(c, post)
}

// MessageHasBeenUpdated provides a mock function with given fields: c, newPost, oldPost
func (_m *Hooks) MessageHasBeenUpdated(c *plugin.Context, newPost *model.Post, oldPost *model.Post) {
	_m.Called(c, newPost, oldPost)
}

// MessageWillBePosted provides a mock function with given fields: c, post
func (_m *Hooks) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	ret := _m.Called(c, post)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.Post) *model.Post); ok {
		r0 = rf(c, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.Post) string); ok {
		r1 = rf(c, post)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MessageWillBeUpdated provides a mock function with given fields: c, newPost, oldPost
func (_m *Hooks) MessageWillBeUpdated(c *plugin.Context, newPost *model.Post, oldPost *model.Post) (*model.Post, string) {
	ret := _m.Called(c, newPost, oldPost)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.Post, *model.Post) *model.Post); ok {
		r0 = rf(c, newPost, oldPost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.Post, *model.Post) string); ok {
		r1 = rf(c, newPost, oldPost)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// OnActivate provides a mock function with given fields:
func (_m *Hooks) OnActivate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnCloudLimitsUpdated provides a mock function with given fields: limits
func (_m *Hooks) OnCloudLimitsUpdated(limits *model.ProductLimits) {
	_m.Called(limits)
}

// OnConfigurationChange provides a mock function with given fields:
func (_m *Hooks) OnConfigurationChange() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnDeactivate provides a mock function with given fields:
func (_m *Hooks) OnDeactivate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnInstall provides a mock function with given fields: c, event
func (_m *Hooks) OnInstall(c *plugin.Context, event model.OnInstallEvent) error {
	ret := _m.Called(c, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(*plugin.Context, model.OnInstallEvent) error); ok {
		r0 = rf(c, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnPluginClusterEvent provides a mock function with given fields: c, ev
func (_m *Hooks) OnPluginClusterEvent(c *plugin.Context, ev model.PluginClusterEvent) {
	_m.Called(c, ev)
}

// OnSendDailyTelemetry provides a mock function with given fields:
func (_m *Hooks) OnSendDailyTelemetry() {
	_m.Called()
}

// OnWebSocketConnect provides a mock function with given fields: webConnID, userID
func (_m *Hooks) OnWebSocketConnect(webConnID string, userID string) {
	_m.Called(webConnID, userID)
}

// OnWebSocketDisconnect provides a mock function with given fields: webConnID, userID
func (_m *Hooks) OnWebSocketDisconnect(webConnID string, userID string) {
	_m.Called(webConnID, userID)
}

// ReactionHasBeenAdded provides a mock function with given fields: c, reaction
func (_m *Hooks) ReactionHasBeenAdded(c *plugin.Context, reaction *model.Reaction) {
	_m.Called(c, reaction)
}

// ReactionHasBeenRemoved provides a mock function with given fields: c, reaction
func (_m *Hooks) ReactionHasBeenRemoved(c *plugin.Context, reaction *model.Reaction) {
	_m.Called(c, reaction)
}

// RunDataRetention provides a mock function with given fields: nowTime, batchSize
func (_m *Hooks) RunDataRetention(nowTime int64, batchSize int64) (int64, error) {
	ret := _m.Called(nowTime, batchSize)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(nowTime, batchSize)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(nowTime, batchSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServeHTTP provides a mock function with given fields: c, w, r
func (_m *Hooks) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	_m.Called(c, w, r)
}

// UserHasBeenCreated provides a mock function with given fields: c, user
func (_m *Hooks) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	_m.Called(c, user)
}

// UserHasJoinedChannel provides a mock function with given fields: c, channelMember, actor
func (_m *Hooks) UserHasJoinedChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	_m.Called(c, channelMember, actor)
}

// UserHasJoinedTeam provides a mock function with given fields: c, teamMember, actor
func (_m *Hooks) UserHasJoinedTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	_m.Called(c, teamMember, actor)
}

// UserHasLeftChannel provides a mock function with given fields: c, channelMember, actor
func (_m *Hooks) UserHasLeftChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	_m.Called(c, channelMember, actor)
}

// UserHasLeftTeam provides a mock function with given fields: c, teamMember, actor
func (_m *Hooks) UserHasLeftTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	_m.Called(c, teamMember, actor)
}

// UserHasLoggedIn provides a mock function with given fields: c, user
func (_m *Hooks) UserHasLoggedIn(c *plugin.Context, user *model.User) {
	_m.Called(c, user)
}

// UserWillLogIn provides a mock function with given fields: c, user
func (_m *Hooks) UserWillLogIn(c *plugin.Context, user *model.User) string {
	ret := _m.Called(c, user)

	var r0 string
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.User) string); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WebSocketMessageHasBeenPosted provides a mock function with given fields: webConnID, userID, req
func (_m *Hooks) WebSocketMessageHasBeenPosted(webConnID string, userID string, req *model.WebSocketRequest) {
	_m.Called(webConnID, userID, req)
}
