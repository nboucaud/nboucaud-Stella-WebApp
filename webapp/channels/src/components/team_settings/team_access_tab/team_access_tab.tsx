// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useState} from 'react';
import type {ChangeEvent} from 'react';
import {useIntl} from 'react-intl';

import {RefreshIcon} from '@mattermost/compass-icons/components';
import type {Team} from '@mattermost/types/teams';

import Input from 'components/widgets/inputs/input/input';
import BaseSettingItem from 'components/widgets/modals/components/base_setting_item';
import CheckboxSettingItem from 'components/widgets/modals/components/checkbox_setting_item';
import ModalSection from 'components/widgets/modals/components/modal_section';

import OpenInvite from './open_invite';

import type {PropsFromRedux, OwnProps} from '.';

import './team_access_tab.scss';

type Props = PropsFromRedux & OwnProps;

const AccessTab = (props: Props) => {
    const [inviteId, setInviteId] = useState<Team['invite_id']>(props.team?.invite_id ?? '');
    const [allowedDomains, setAllowedDomains] = useState<Team['allowed_domains']>(props.team?.allowed_domains ?? '');
    const [showAllowedDomains, setShowAllowedDomains] = useState<boolean>(false);
    const [serverError, setServerError] = useState<string>('');
    const {formatMessage} = useIntl();

    const handleAllowedDomainsSubmit = async () => {
        const {error} = await props.actions.patchTeam({
            id: props.team?.id,
            allowed_domains: allowedDomains,
        });
        if (error) {
            setServerError(error.message);
        }
    };

    const updateAllowedDomains = (e: ChangeEvent<HTMLInputElement>) => setAllowedDomains(e.target.value);

    const handleRegenerateInviteId = async () => {
        const {data, error} = await props.actions.regenerateTeamInviteId(props.team?.id || '');

        if (data?.invite_id) {
            setInviteId(data.invite_id);
            return;
        }

        if (error) {
            setServerError(error.message);
        }
    };

    let inviteSection;
    if (props.canInviteTeamMembers) {
        const inviteSectionInput = (
            <div id='teamInviteContainer' >
                <Input
                    id='teamInviteId'
                    className='form-control'
                    type='text'
                    value={inviteId}
                    maxLength={32}
                />
                <button
                    id='regenerateButton'
                    className='btn btn-tertiary'
                    onClick={handleRegenerateInviteId}
                >
                    <RefreshIcon/>
                    {formatMessage({id: 'general_tab.regenerate', defaultMessage: 'Regenerate'})}
                </button>
            </div>
        );

        // inviteSection = (
        //     <SettingItemMax
        //         submit={this.handleInviteIdSubmit}
        //         serverError={serverError}
        //         clientError={clientError}
        //         saveButtonText={localizeMessage('general_tab.regenerate', 'Regenerate')}
        //     />
        // );

        inviteSection = (
            <BaseSettingItem
                className='access-invite-section'
                title={{id: 'general_tab.codeTitle', defaultMessage: 'Invite Code'}}
                description={{id: 'general_tab.codeLongDesc', defaultMessage: 'The Invite Code is part of the unique team invitation link which is sent to members you’re inviting to this team. Regenerating the code creates a new invitation link and invalidates the previous link.'}}
                content={inviteSectionInput}
                descriptionAboveContent={true}
            />
        );
    }

    const allowedDomainsSectionInput = (
        <div
            id='allowedDomainsSetting'
            className='form-group'
        >
            <CheckboxSettingItem
                css={{marginBottom: '16px'}}
                inputFieldData={{title: {id: 'general_tab.allowedDomains', defaultMessage: 'Allow only users with a specific email domain to join this team'}, name: 'name'}}
                inputFieldValue={showAllowedDomains}
                handleChange={(checked) => setShowAllowedDomains(checked)}
            />
            {showAllowedDomains &&
                <input
                    id='allowedDomains'
                    className='form-control'
                    type='text'
                    onChange={updateAllowedDomains}
                    value={allowedDomains}
                    placeholder={formatMessage({id: 'general_tab.AllowedDomainsExample', defaultMessage: 'corp.mattermost.com, mattermost.com'})}
                    aria-label={formatMessage({id: 'general_tab.allowedDomains.ariaLabel', defaultMessage: 'Allowed Domains'})}
                />
            }
        </div>
    );

    // const allowedDomainsSection = (
    //     <SettingItemMax
    //         submit={this.handleAllowedDomainsSubmit}
    //         serverError={serverError}
    //         clientError={clientError}
    //     />
    // );

    const allowedDomainsSection = (
        <BaseSettingItem
            className='access-allowed-domains-section'
            title={{id: 'general_tab.allowedDomainsTitle', defaultMessage: 'Users with a specific email domain'}}
            description={{id: 'general_tab.allowedDomainsInfo', defaultMessage: 'When enabled, users can only join the team if their email matches a specific domain (e.g. "mattermost.org")'}}
            content={allowedDomainsSectionInput}
        />
    );

    // todo sinan: check title font size is same as figma
    // todo sinan: descriptions are placed above content. Waiting an input from Matt
    return (
        <ModalSection
            content={
                <div className='user-settings'>
                    {props.team?.group_constrained ? undefined : allowedDomainsSection}
                    <div className='divider-light'/>
                    <OpenInvite
                        teamId={props.team?.id}
                        isGroupConstrained={props.team?.group_constrained}
                        allowOpenInvite={props.team?.allow_open_invite}
                        patchTeam={props.actions.patchTeam}
                    />
                    <div className='divider-light'/>
                    {props.team?.group_constrained ? undefined : inviteSection}
                </div>
            }
        />
    );
};
export default AccessTab;
