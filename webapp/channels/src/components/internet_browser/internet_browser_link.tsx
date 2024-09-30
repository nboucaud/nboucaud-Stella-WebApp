// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {} from 'react';
import {useIntl} from 'react-intl';
import {Link, useRouteMatch} from 'react-router-dom';

import InternetIcon from './internet_icon';

import './internet_browser_link.scss';

const GlobalThreadsLink = () => {
    const {formatMessage} = useIntl();
    const {url} = useRouteMatch();

    return (
        <ul className='SidebarInternetBrowser NavGroupContent nav nav-pills__container'>
            <li
                id={'sidebar-internet-browser-button'}
                className={'SidebarChannel'}
                tabIndex={-1}
            >
                <Link
                    to={`${url}/browser`}
                    id='sidebarItem_internet_browser'
                    draggable='false'
                    className={'SidebarLink sidebar-item'}
                    tabIndex={0}
                >
                    <span className='icon'>
                        <InternetIcon/>
                    </span>
                    <div className='SidebarChannelLinkLabel_wrapper'>
                        <span className='SidebarChannelLinkLabel sidebar-item__name'>
                            {formatMessage({id: 'globalInternet.sidebarLink', defaultMessage: 'Internet'})}
                        </span>
                    </div>
                </Link>
            </li>
        </ul>
    );
};

export default GlobalThreadsLink;
