// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {injectIntl, type IntlShape} from 'react-intl';

export type Props = {
    intl: IntlShape;
};

const MembershipHeader = ({intl}: Props) => {
    return (
        <div
            role='banner'
            className='membership-header alt a11y__region'
        >
            <div className='membership-header__title'>
                <strong
                    role='heading'
                    className='heading'
                >{intl.formatMessage({id: 'membership.heading', defaultMessage: 'Select a plan'})}</strong>
                <span className='subheading'>
                    {intl.formatMessage({id: 'membership.subheading', defaultMessage: 'Choose a plan to get started'})}
                </span>
            </div>
        </div>
    );
};

export default injectIntl(MembershipHeader);
