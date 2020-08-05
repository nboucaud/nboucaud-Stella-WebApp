// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/utils"
)

type AutoChannelCreator struct {
	a                  *App
	team               *model.Team
	Fuzzy              bool
	DisplayNameLen     utils.Range
	DisplayNameCharset string
	NameLen            utils.Range
	NameCharset        string
	ChannelType        string
}

func NewAutoChannelCreator(a *App, team *model.Team) *AutoChannelCreator {
	return &AutoChannelCreator{
		a:                  a,
		team:               team,
		Fuzzy:              false,
		DisplayNameLen:     CHANNEL_DISPLAY_NAME_LEN,
		DisplayNameCharset: utils.ALPHANUMERIC,
		NameLen:            CHANNEL_NAME_LEN,
		NameCharset:        utils.LOWERCASE,
		ChannelType:        CHANNEL_TYPE,
	}
}

func (cfg *AutoChannelCreator) createRandomChannel() (*model.Channel, error) {
	var displayName string
	if cfg.Fuzzy {
		displayName = utils.FuzzName()
	} else {
		displayName = utils.RandomName(cfg.NameLen, cfg.NameCharset)
	}
	name := utils.RandomName(cfg.NameLen, cfg.NameCharset)

	channel := &model.Channel{
		TeamId:      cfg.team.Id,
		DisplayName: displayName,
		Name:        name,
		Type:        cfg.ChannelType,
	}

	channel, err := cfg.a.CreateChannel(channel, false)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (cfg *AutoChannelCreator) CreateTestChannels(num utils.Range) ([]*model.Channel, error) {
	numChannels := utils.RandIntFromRange(num)
	channels := make([]*model.Channel, numChannels)

	for i := 0; i < numChannels; i++ {
		var err error
		channels[i], err = cfg.createRandomChannel()
		if err != nil {
			return nil, err
		}
	}

	return channels, nil
}
