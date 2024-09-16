// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import {useMembershipCardContext} from './membership_card';

const Title = () => {
    const {title, titleSubtext} = useMembershipCardContext();
    return (
        <div className='title-container'>
            <strong
                className='title'
                role='heading'
            >{title}</strong>
            {titleSubtext && <strong className='title-subtext'>{titleSubtext}</strong>}
        </div>
    );
};

export default Title;
