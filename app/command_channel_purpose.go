// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"github.com/mattermost/platform/model"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
)

type PurposeProvider struct {
}

const (
	CMD_PURPOSE = "purpose"
)

func init() {
	RegisterCommandProvider(&PurposeProvider{})
}

func (me *PurposeProvider) GetTrigger() string {
	return CMD_PURPOSE
}

func (me *PurposeProvider) GetCommand(T goi18n.TranslateFunc) *model.Command {
	return &model.Command{
		Trigger:          CMD_PURPOSE,
		AutoComplete:     true,
		AutoCompleteDesc: T("api.command_channel_purpose.desc"),
		AutoCompleteHint: T("api.command_channel_purpose.hint"),
		DisplayName:      T("api.command_channel_purpose.name"),
	}
}

func (me *PurposeProvider) DoCommand(args *model.CommandArgs, message string) *model.CommandResponse {
	channel, err := GetChannel(args.ChannelId)
	if err != nil {
		return &model.CommandResponse{Text: args.T("api.command_channel_purpose.channel.app_error"), ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL}
	}

	if channel.Type == model.CHANNEL_OPEN && !SessionHasPermissionToChannel(args.Session, args.ChannelId, model.PERMISSION_MANAGE_PUBLIC_CHANNEL_PROPERTIES) {
		return &model.CommandResponse{Text: args.T("api.command_channel_purpose.permission.app_error"), ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL}
	}

	if channel.Type == model.CHANNEL_PRIVATE && !SessionHasPermissionToChannel(args.Session, args.ChannelId, model.PERMISSION_MANAGE_PRIVATE_CHANNEL_PROPERTIES) {
		return &model.CommandResponse{Text: args.T("api.command_channel_purpose.permission.app_error"), ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL}
	}

	if channel.Type == model.CHANNEL_GROUP || channel.Type == model.CHANNEL_DIRECT {
		return &model.CommandResponse{Text: args.T("api.command_channel_purpose.direct_group.app_error"), ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL}
	}

	if len(message) == 0 {
		return &model.CommandResponse{Text: args.T("api.command_channel_purpose.message.app_error"), ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL}
	}

	patch := &model.ChannelPatch{
		Name:        new(string),
		DisplayName: new(string),
		Header:      new(string),
		Purpose:     new(string),
	}

	*patch.Name = channel.Name
	*patch.DisplayName = channel.DisplayName
	*patch.Header = channel.Header
	*patch.Purpose = message

	_, err = PatchChannel(channel, patch, args.UserId)
	if err != nil {
		return &model.CommandResponse{Text: args.T("api.command_channel_purpose.update_channel.app_error"), ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL}
	}

	return &model.CommandResponse{}
}
