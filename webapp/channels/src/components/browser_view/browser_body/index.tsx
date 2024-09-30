// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import CameraOutlineIcon from '@mattermost/compass-icons/components/camera-outline';

import './browser_body.scss';

import logoImage from '../../../images/infogito.png';
import backgroundImage from '../../../images/internet_browser_background.png';
import robotImage from '../../../images/internet_browser_robot.png';

const BrowserBody = () => {
    return (
        <div
            id='BrowserBody'
            className='BrowserBody'
        >
            <img
                className='background'
                src={backgroundImage}
                loading='lazy'
            />

            <div className='BrowserBody__content'>
                <div className='searchbar'>
                    <i className='icon icon-magnify'/>
                    <input placeholder='Search google or type a URL'/>
                    <i className='icon icon-microphone'/>
                    <CameraOutlineIcon
                        className='icon'
                        color='black'
                    />
                </div>

                <img
                    className='logo'
                    src={logoImage}
                />
                <img
                    className='robot'
                    src={robotImage}
                />
            </div>
        </div>
    );
};

export default BrowserBody;
