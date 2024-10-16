import React from "react";
import { useSelector } from "react-redux";
import UrlIframe from "./browser-iframe";
import BrowserBody from "../browser_body";

const ContentDisplay: React.FC = () => {
    const { tabs, activeTabIndex } = useSelector(
        (state: any) => state.urlManager
    );

    return (
        <div className="tab-content-browser">
            {tabs.map((tab: any, index: number) => (
                <div
                    key={index}
                    style={{
                        display: activeTabIndex === index ? "flex" : "none",
                        width: "100%",
                        height: "100%",
                    }} // Hide inactive tabs
                >
                    {tab.url ? <UrlIframe url={tab.url} /> : <BrowserBody />}
                </div>
            ))}
        </div>
    );
};

export default ContentDisplay;
