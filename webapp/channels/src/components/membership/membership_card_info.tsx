// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import {useMembershipCardContext} from './membership_card';

const Info = () => {
    const {info} = useMembershipCardContext();

    return (
        <div className='plan-info'>
            {info}
        </div>
    );
};

export default Info;
