import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { LuPlus } from "react-icons/lu";
import { LiaCompass } from "react-icons/lia";
import UrlTabs from "../browser/browser-tab";
import "./browser-header.scss";
import { addbrowserTab } from "../browser/browser-state";

const BrowserHeader = () => {
    const dispatch = useDispatch();
    const { urls, activeTabIndex, isBrowserTabActive } = useSelector(
        (state: any) => state.urlManager
    );

    return (
        <div className="browser-header">
            {/* Default BrowserTab (Infogito) */}
            {/* <div className={`browser-tab`}>
                <LiaCompass className="icon" /> <span>Infogito</span>
            </div> */}

            <UrlTabs />

            {/* <LuPlus
                onClick={() => dispatch(addbrowserTab(idnumber))}
                className="icon"
            /> */}
        </div>
    );
};

export default BrowserHeader;
