// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useIntl} from 'react-intl';

import {useMembershipCardContext} from './membership_card';

const Pricing = () => {
    const {formatMessage} = useIntl();
    const {pricing, pricingSubtext, isBilledAnnually} = useMembershipCardContext();

    return (
        <div className='pricing'>
            <strong className='pricing__price'>{pricing}</strong>
            {pricingSubtext && <p className='pricing__subtext'>{pricingSubtext}</p>}
            {isBilledAnnually && <p className='pricing__billed_anually'>
                {formatMessage({id: 'membership_card.billed_anually', defaultMessage: '(billed annually)'})}
            </p>}
        </div>
    );
};

export default Pricing;
