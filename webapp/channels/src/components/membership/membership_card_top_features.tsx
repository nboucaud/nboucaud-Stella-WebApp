// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useIntl} from 'react-intl';

import {useMembershipCardContext} from './membership_card';

const TopFeatures = () => {
    const {formatMessage} = useIntl();
    const {topFeatures} = useMembershipCardContext();

    return (
        <div className='top-features'>
            <strong>{formatMessage({id: 'membership_card.top_features.title', defaultMessage: 'Top features'})}</strong>
            <ul>
                {topFeatures.map((feature) => (
                    <li key={feature}>
                        {feature}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TopFeatures;
