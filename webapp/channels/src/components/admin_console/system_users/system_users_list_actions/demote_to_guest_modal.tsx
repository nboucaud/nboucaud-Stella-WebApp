// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useState} from 'react';
import {FormattedMessage} from 'react-intl';
import {useDispatch} from 'react-redux';

import type {ServerError} from '@mattermost/types/errors';
import type {UserProfile} from '@mattermost/types/users';

import {demoteUserToGuest} from 'mattermost-redux/actions/users';

import ConfirmModal from 'components/confirm_modal';

type Props = {
    user: UserProfile;
    onHide: () => void;
    onError: (error: ServerError) => void;
}

export default function DemoteToGuestModal({user, onHide, onError}: Props) {
    const [show, setShow] = useState(true);
    const dispatch = useDispatch();

    async function confirm() {
        const {error} = await dispatch(demoteUserToGuest(user.id));
        if (error) {
            onError(error);
        }
        close();
    }

    function close() {
        setShow(false);
        onHide();
    }

    const title = (
        <FormattedMessage
            id='demote_to_user_modal.title'
            defaultMessage='Demote User {username} to Guest'
            values={{
                username: user.username,
            }}
        />
    );

    const message = (
        <FormattedMessage
            id='demote_to_user_modal.desc'
            defaultMessage={'This action demotes the user {username} to a guest. It will restrict the user\'s ability to join public channels and interact with users outside of the channels they are currently members of. Are you sure you want to demote user {username} to guest?'}
            values={{
                username: user.username,
            }}
        />
    );

    const demoteGuestButton = (
        <FormattedMessage
            id='demote_to_user_modal.demote'
            defaultMessage='Demote'
        />
    );

    return (
        <ConfirmModal
            show={show}
            title={title}
            message={message}
            confirmButtonClass='btn btn-danger'
            confirmButtonText={demoteGuestButton}
            onConfirm={confirm}
            onCancel={close}
        />
    );
}
