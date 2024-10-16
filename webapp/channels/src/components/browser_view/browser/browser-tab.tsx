// src/components/UrlTabs.tsx
import React from "react";
import { useSelector, useDispatch } from "react-redux";
import {
    setActivebrowserTab,
    addbrowserTab,
    removebrowserTab,
} from "./browser-state";
import "./browsertabs.scss";
import { LuPlus } from "react-icons/lu";

const UrlTabs: React.FC = () => {
    const { tabs, activeTabIndex } = useSelector(
        (state: any) => state.urlManager
    );
    const dispatch = useDispatch();
    const state = useSelector((state: any) => state);
    const idnumber = state.tabs.length;
    return (
        <div className="tabs-wrapper">
            {tabs.map((tab: any, index: any) => (
                <div
                    key={tab.id}
                    className={
                        activeTabIndex === index
                            ? "active-browser custom-tabs"
                            : "custom-tabs"
                    }
                >
                    <button
                        style={{
                            width: "80%",
                            textAlign: "left",
                            overflow: "hidden",
                        }}
                        onClick={() => dispatch(setActivebrowserTab(index))}
                    >
                        {tab.url ? tab.url : "New Tab"}
                    </button>
                    <button
                        style={{ width: "20%" }}
                        onClick={() => dispatch(removebrowserTab(index))}
                    >
                        x
                    </button>
                </div>
            ))}
            <LuPlus
                onClick={() => dispatch(addbrowserTab(idnumber))}
                className="icon"
            />
        </div>
    );
};

export default UrlTabs;
