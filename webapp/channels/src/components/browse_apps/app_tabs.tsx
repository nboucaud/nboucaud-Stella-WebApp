import React from "react";
import { useSelector, useDispatch } from "react-redux";
import {
    addTab,
    removeTab,
    setActiveTab,
    setActiveApp,
} from "../../packages/mattermost-redux/src/actions/tabAction";
import iconlayer from "./layers-app-icon.svg";
import "./app_tabs.scss";

interface Tab {
    id: string;
    title: string;
    uniqueId: string; // Ensure you have uniqueId in the Tab interface
}

interface AppState {
    tabs: {
        tabs: Tab[];
        activeTab: string | null;
        activeApp: string | null;
    };
}

const Tabs: React.FC = () => {
    const dispatch = useDispatch();

    const { tabs, activeTab, activeApp } = useSelector(
        (state: AppState) => state.tabs
    );

    const handleTabClick = (uniqueId: string, id: string) => {
        dispatch(setActiveTab(uniqueId));
        dispatch(setActiveApp(id));
    };

    const handleCloseTab = (uniqueId: string, e: React.MouseEvent) => {
        e.stopPropagation(); // Prevent closing tab from triggering the tab click
        dispatch(removeTab(uniqueId));
    };

    console.log("Tabs:", tabs, "Active Tab:", activeTab); // More descriptive logging

    return (
        <div className="tabs-container">
            {tabs.map((tab: Tab) => (
                <div
                    onClick={() => handleTabClick(tab.uniqueId, tab.id)} // Use uniqueId
                    key={tab.uniqueId} // Use uniqueId as key
                    className={`app-tab-wrap ${
                        activeTab === tab.uniqueId ? "active-tab" : ""
                    }`}
                >
                    <span className="app-heading-text">
                        <img src={iconlayer} alt="icon" />
                        {tab.title} #{tab.uniqueId.slice(-2)}
                    </span>
                    <button
                        className="app-close-tag"
                        onClick={(e) => handleCloseTab(tab.uniqueId, e)} // Use uniqueId here too
                    >
                        x
                    </button>
                </div>
            ))}
        </div>
    );
};

export default Tabs;
