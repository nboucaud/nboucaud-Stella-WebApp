// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from "react";
import { useSelector } from "react-redux";
import UrlIframe from "./browser/browser-iframe";
import BrowserBody from "./browser_body";
import BrowserHeader from "./browser_header";
import BrowserSearchSection from "./browser_search_section";
import UrlTabs from "./browser/browser-tab";
import ContentDisplay from "./browser/contentDisplay";

const BrowserView = () => {
    // Access the URLs, active tab index, and browser tab active state from the Redux store
    const { urls, activeTabIndex, isBrowserTabActive } = useSelector(
        (state: any) => state.urlManager
    );

    return (
        <div
            style={{
                marginLeft: "-1px",
                display: "flex",
                flexDirection: "column",
                height: "100%",
            }}
        >
            <BrowserHeader />
            <BrowserSearchSection />

            <ContentDisplay />
        </div>
    );
};

export default BrowserView;
