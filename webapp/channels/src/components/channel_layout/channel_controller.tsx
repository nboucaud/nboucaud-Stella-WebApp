// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import classNames from "classnames";
import React, { lazy, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import { cleanUpStatusAndProfileFetchingPoll } from "mattermost-redux/actions/status_profile_polling";
import { getIsUserStatusesConfigEnabled } from "mattermost-redux/selectors/entities/common";
import {
    amphion,
    code,
    dawg,
    imageeditor,
    losslesscut,
    lumen,
    motion,
    notable,
    omniclip,
    playbook,
    plexocore,
    viewerjs,
} from "../browse_apps/apps_iframe";
import { addVisibleUsersInCurrentChannelToStatusPoll } from "actions/status_actions";

import Pluggable from "plugins/pluggable";
import { Constants } from "utils/constants";
import { isInternetExplorer, isEdge } from "utils/user_agent";
import IframeContainer from "components/browse_apps/iframe_container";
import { makeAsyncComponent } from "components/async_load";
import CenterChannel from "components/channel_layout/center_channel";
import LoadingScreen from "components/loading_screen";
import ProductResults from "components/shop/product_results";
import Sidebar from "components/sidebar";
import CRTPostsChannelResetWatcher from "components/threading/channel_threads/posts_channel_reset_watcher";
import UnreadsStatusHandler from "components/unreads_status_handler";

const ProductNoticesModal = makeAsyncComponent(
    "ProductNoticesModal",
    lazy(() => import("components/product_notices_modal"))
);
const ResetStatusModal = makeAsyncComponent(
    "ResetStatusModal",
    lazy(() => import("components/reset_status_modal"))
);

const BODY_CLASS_FOR_CHANNEL = ["app__body", "channel-view"];
import notesIcon from "../../images/notes-icon.png";
import storeIcon from "../../images/store.png";

type Props = {
    shouldRenderCenterChannel: boolean;
};

export default function ChannelController(props: Props) {
    const [openShop, setOpenShop] = useState(false);

    const enabledUserStatuses = useSelector(getIsUserStatusesConfigEnabled);
    const dispatch = useDispatch();

    useEffect(() => {
        const isMsBrowser = isInternetExplorer() || isEdge();
        const { navigator } = window;
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        const platform =
            navigator?.userAgentData?.platform ||
            navigator?.platform ||
            "unknown";
        document.body.classList.add(
            ...getClassnamesForBody(platform, isMsBrowser)
        );

        return () => {
            document.body.classList.remove(...BODY_CLASS_FOR_CHANNEL);

            // This cleans up the status and profile setInterval of fetching poll we use to batch requests
            // when fetching statuses and profiles for a list of users.
            dispatch(cleanUpStatusAndProfileFetchingPoll());
        };
    }, []);

    useEffect(() => {
        let loadStatusesIntervalId: NodeJS.Timeout;
        if (enabledUserStatuses) {
            loadStatusesIntervalId = setInterval(() => {
                dispatch(addVisibleUsersInCurrentChannelToStatusPoll());
            }, Constants.STATUS_INTERVAL);
        }

        return () => {
            clearInterval(loadStatusesIntervalId);
        };
    }, [enabledUserStatuses]);

    //app tabs and default screen switch

    const iframeMap: { [key: string]: React.FC } = {
        vscode: code,
        notes: notable,
        motion: motion,
        viewerjs: viewerjs,
        imageeditor: imageeditor,
        lumen: lumen,
        dawg: dawg,
        amphion: amphion,
        losslesscut: losslesscut,
        omniclip: omniclip,
        playbook: playbook,
        plexocore: plexocore,
    };
    interface AppState {
        tabs: {
            activeApp: string | null;
        };
    }
    const activeApp = useSelector((state: AppState) => state.tabs.activeApp);
    const ActiveIframe = activeApp ? iframeMap[activeApp] : null;
    return (
        <>
            <CRTPostsChannelResetWatcher />
            <Sidebar />
            <div
                id="channel_view"
                className="channel-view"
                data-testid="channel_view"
            >
                <UnreadsStatusHandler />
                <ProductNoticesModal />
                <div
                    className={classNames("container-fluid channel-view-inner")}
                >
                    {openShop ? (
                        <ProductResults />
                    ) : (
                        <>
                            {props.shouldRenderCenterChannel ? (
                                <CenterChannel />
                            ) : (
                                <LoadingScreen centered={true} />
                            )}
                            <Pluggable pluggableName="Root" />
                            <ResetStatusModal />
                        </>
                    )}
                </div>
            </div>
            <div className="home-screen-wrapper__sidebar right">
                <div className="top">
                    <button>
                        <img src={notesIcon} />
                    </button>
                </div>
                <div className="bottom">
                    <button onClick={() => setOpenShop(!openShop)}>
                        <img src={storeIcon} />
                    </button>
                </div>
            </div>
        </>
    );
}

export function getClassnamesForBody(
    platform: Window["navigator"]["platform"],
    isMsBrowser = false
) {
    const bodyClass = [...BODY_CLASS_FOR_CHANNEL];

    // OS Detection
    if (platform === "Win32" || platform === "Win64") {
        bodyClass.push("os--windows");
    } else if (platform === "MacIntel" || platform === "MacPPC") {
        bodyClass.push("os--mac");
    }

    // IE Detection
    if (isMsBrowser) {
        bodyClass.push("browser--ie");
    }

    return bodyClass;
}
