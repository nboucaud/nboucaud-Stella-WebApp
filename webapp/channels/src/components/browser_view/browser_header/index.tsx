// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import "./browser-header.scss";

import { LuPlus } from "react-icons/lu";
import { LiaCompass } from "react-icons/lia";

const BrowserHeader = () => {
    return (
        <div className="browser-header">
            <div className="browser-tab"><LiaCompass className="icon"/> <span>Infogito</span></div>
            <LuPlus className="icon"/>
        </div>
    );
};

export default BrowserHeader;
