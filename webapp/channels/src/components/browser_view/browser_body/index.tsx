// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, { useState } from "react";

import CameraOutlineIcon from "@mattermost/compass-icons/components/camera-outline";

import "./browser_body.scss";

import logoImage from "../../../images/infogito.png";
import backgroundImage from "../../../images/internet_browser_background.png";
import robotImage from "../../../images/internet_browser_robot.png";
import { useDispatch, useSelector } from "react-redux";
import { setTabUrl } from "../browser/browser-state";

const BrowserBody = () => {
    const [url, setUrl] = useState("");
    const dispatch = useDispatch();
    const { activeTabIndex } = useSelector((state: any) => state.urlManager);
    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (url) {
            dispatch(setTabUrl(activeTabIndex, url));
            setUrl("");
        }
    };
    return (
        <div id="BrowserBody" className="BrowserBody">
            <img className="background" src={backgroundImage} loading="lazy" />

            <div className="BrowserBody__content">
                <div className="searchbar">
                    <i className="icon icon-magnify" />
                    <form className="form-input" onSubmit={handleSubmit}>
                        <input
                            type="text"
                            value={url}
                            onChange={(e) => setUrl(e.target.value)}
                            placeholder="Search google or type a URL"
                            required
                        />
                    </form>

                    <i className="icon icon-microphone" />
                    <CameraOutlineIcon className="icon" color="black" />
                </div>

                <img className="logo" src={logoImage} />
                <img className="robot" src={robotImage} />
            </div>
        </div>
    );
};

export default BrowserBody;
