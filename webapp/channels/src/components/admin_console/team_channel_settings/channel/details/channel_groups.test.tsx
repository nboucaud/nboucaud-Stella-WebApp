// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {shallow} from 'enzyme';

import {TestHelper} from 'utils/test_helper';
import {Channel} from '@mattermost/types/channels';

import {ChannelGroups} from './channel_groups';

describe('admin_console/team_channel_settings/channel/ChannelGroups', () => {
    test('should match snapshot', () => {
        const groups = [TestHelper.getGroupMock({
            id: '123',
            display_name: 'DN',
            member_count: 3,
        })];

        const testChannel: Partial<Channel> & {team_name: string} = {
            id: '123',
            team_name: 'team',
            type: 'O',
            group_constrained: false,
            name: 'DN',
        };
        const wrapper = shallow(
            <ChannelGroups
                synced={true}
                onAddCallback={jest.fn()}
                onGroupRemoved={jest.fn()}
                removedGroups={[]}
                groups={groups}
                channel={testChannel}
                totalGroups={1}
                setNewGroupRole={jest.fn()}
                isDisabled={false}
            />,
        );
        expect(wrapper).toMatchSnapshot();
    });
});
