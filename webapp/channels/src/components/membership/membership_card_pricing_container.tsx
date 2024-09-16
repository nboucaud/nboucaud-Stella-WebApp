// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {type PropsWithChildren, type ReactNode} from 'react';

const PricingContainer = ({children}: PropsWithChildren<ReactNode>) => {
    return (
        <div className='pricing'>
            {children}
        </div>
    );
};

export default PricingContainer;
