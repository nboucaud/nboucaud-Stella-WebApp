// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import React from 'react';

import SuggestionStore from 'stores/suggestion_store.jsx';
import ChannelStore from 'stores/channel_store.jsx';
import UserStore from 'stores/user_store.jsx';
import * as Utils from 'utils/utils.jsx';
import Client from 'utils/web_client.jsx';

import {FormattedMessage} from 'react-intl';
import Suggestion from './suggestion.jsx';

const MaxUserSuggestions = 40;

class AtMentionSuggestion extends Suggestion {
    render() {
        const isSelection = this.props.isSelection;
        const user = this.props.item;

        // fake user to represent a bar in the suggestion list
        if (user.is_separator) {
            return (
                <div className='suggestion-separator'>
                    <hr className='separator__hr'/>
                </div>
            );
        }

        let username;
        let description;
        let icon;
        if (user.username === 'all') {
            username = 'all';
            description = (
                <FormattedMessage
                    id='suggestion.mention.all'
                    defaultMessage='Notifies everyone in the channel, use in {townsquare} to notify the whole team'
                    values={{
                        townsquare: ChannelStore.getByName('town-square').display_name
                    }}
                />
            );
            icon = <i className='mention__image fa fa-users fa-2x'/>;
        } else if (user.username === 'channel') {
            username = 'channel';
            description = (
                <FormattedMessage
                    id='suggestion.mention.channel'
                    defaultMessage='Notifies everyone in the channel'
                />
            );
            icon = <i className='mention__image fa fa-users fa-2x'/>;
        } else {
            username = user.username;

            if ((user.first_name || user.last_name) && user.nickname) {
                description = ` - ${Utils.getFullName(user)} (${user.nickname})`;
            } else if (user.nickname) {
                description = ` - (${user.nickname})`;
            } else if (user.first_name || user.last_name) {
                description = ` - ${Utils.getFullName(user)}`;
            }

            icon = (
                <img
                    className='mention__image'
                    src={Client.getUsersRoute() + '/' + user.id + '/image?time=' + user.update_at}
                />
            );
        }

        let className = 'mentions__name';
        if (isSelection) {
            className += ' suggestion--selected';
        }

        return (
            <div
                className={className}
                onClick={this.handleClick}
            >
                <div className='pull-left'>
                    {icon}
                </div>
                <div className='pull-left mention--align'>
                    <span>
                        {'@' + username}
                    </span>
                    <span className='mention__fullname'>
                        {description}
                    </span>
                </div>
            </div>
        );
    }
}

function filterUsersByPrefix(users, prefix, limit) {
    const filtered = [];

    for (const id of Object.keys(users)) {
        if (filtered.length >= limit) {
            break;
        }

        const user = users[id];

        if (user.delete_at > 0) {
            continue;
        }

        if (user.username.startsWith(prefix) ||
            (user.first_name && user.first_name.toLowerCase().startsWith(prefix)) ||
            (user.last_name && user.last_name.toLowerCase().startsWith(prefix)) ||
            (user.nickname && user.nickname.toLowerCase().startsWith(prefix))) {
            filtered.push(user);
        }
    }

    return filtered;
}

export default class AtMentionProvider {
    handlePretextChanged(suggestionId, pretext) {
        const captured = (/@([a-z0-9\-\._]*)$/i).exec(pretext.toLowerCase());
        if (captured) {
            const prefix = captured[1];

            // Group users into members and nonmembers of the channel.
            const users = UserStore.getActiveOnlyProfiles(true);
            const channelMembers = {};
            const extra = ChannelStore.getCurrentExtraInfo();
            for (let i = 0; i < extra.members.length; i++) {
                const id = extra.members[i].id;
                if (users[id]) {
                    channelMembers[id] = users[id];
                    Reflect.deleteProperty(users, id);
                }
            }
            const channelNonmembers = users;

            // Filter users by prefix.
            const filteredMembers = filterUsersByPrefix(
                    channelMembers, prefix, MaxUserSuggestions);
            const filteredNonmembers = filterUsersByPrefix(
                    channelNonmembers, prefix, MaxUserSuggestions - filteredMembers.length);
            const filteredSpecialMentions = pretext.startsWith('/msg') ?
                [] : ['channel', 'all'].filter((item) => {
                    // eslint-disable-line no-use-before-define
                    return item.startsWith(prefix);
                }).map((name) => {
                    return {username: name};
                });

            // Sort users by username.
            [filteredMembers, filteredNonmembers].forEach((items) => {
                items.sort((a, b) => {
                    const aPrefix = a.username.startsWith(prefix);
                    const bPrefix = b.username.startsWith(prefix);

                    if (aPrefix === bPrefix) {
                        return a.username.localeCompare(b.username);
                    } else if (aPrefix) {
                        return -1;
                    }

                    return 1;
                });
            });

            // Make separators to separate channel members, special mentions and channel nonmembers.
            const separator1 = (filteredMembers.length > 0 && filteredSpecialMentions.length > 0) ?
                [{username: '-1', is_separator: true}] : [];
            const separator2 = ((filteredSpecialMentions.length > 0) && (filteredNonmembers.length > 0)) ?
                [{username: '-2', is_separator: true}] : [];

            const filtered = filteredMembers.concat(separator1).concat(filteredSpecialMentions).concat(separator2).concat(filteredNonmembers);

            const mentions = filtered.map((user) => '@' + user.username);

            SuggestionStore.addSuggestions(suggestionId, mentions, filtered, AtMentionSuggestion, captured[0]);
        }
    }
}
