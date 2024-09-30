// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import BrowserBody from './browser_body';
import BrowserHeader from './browser_header';
import BrowserSearchSection from './browser_search_section';

const BrowserView = () => {
    return (
        <div style={{marginLeft: '-1px', display: 'flex', flexDirection: 'column', height: '100%'}}>
            <BrowserHeader/>
            <BrowserSearchSection/>
            <BrowserBody/>
        </div>
    );
};

export default BrowserView;
