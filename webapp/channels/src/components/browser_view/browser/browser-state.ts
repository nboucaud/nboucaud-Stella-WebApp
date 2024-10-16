// src/store/urlManager.ts
import { AnyAction } from 'redux';

// Action Types
export const ADD_BROWSER_TAB = 'ADD_BROWSER_TAB';
export const REMOVE_BROWSER_TAB = 'REMOVE_BROWSER_TAB';
export const SET_ACTIVE_BROWSER_TAB = 'SET_ACTIVE_BROWSER_TAB';
export const SET_TAB_URL = 'SET_TAB_URL';

// Action Interfaces
interface AddTabAction {
    type: typeof ADD_BROWSER_TAB;
}

interface RemoveTabAction {
    type: typeof REMOVE_BROWSER_TAB;
    payload: number; // index of the tab to remove
}

interface SetActiveTabAction {
    type: typeof SET_ACTIVE_BROWSER_TAB;
    payload: number; // index of the active tab
}

interface SetTabUrlAction {
    type: typeof SET_TAB_URL;
    payload: {
        index: number;
        url: string; // URL to set in the tab
    };
}

// Action Creators
export const addbrowserTab = (id:number): AddTabAction & { payload: { id: number; title: string; uniqueId?: string } } => ({
    type: ADD_BROWSER_TAB,
    payload: { id: Date.now(), title: "New Tab", uniqueId: "" }, 
});
export const removebrowserTab = (index: number): RemoveTabAction => ({
    type: REMOVE_BROWSER_TAB,
    payload: index,
});

export const setActivebrowserTab = (index: number): SetActiveTabAction => ({
    type: SET_ACTIVE_BROWSER_TAB,
    payload: index,
});

export const setTabUrl = (index: number, url: string): SetTabUrlAction => ({
    type: SET_TAB_URL,
    payload: { index, url },
});

// State Interface
interface Tab {
    id: number; // Unique identifier for the tab
    url: string | null; // URL or null to indicate BrowserBody
}

interface UrlManagerState {
    tabs: Tab[]; // Array of tabs
    activeTabIndex: number; // Active tab index
}

// Initial State
const initialState: UrlManagerState = {
    tabs: [{ id:0, url: null }], 
    activeTabIndex: 0,
};

// Reducer
const urlManagerReducer = (state = initialState, action: AnyAction): UrlManagerState => {
    switch (action.type) {
        case ADD_BROWSER_TAB:
            const newTabId = state.tabs?.length ?? 0; 
            return {
                ...state,
                tabs: [...state.tabs, { id: newTabId, url: null }],
                activeTabIndex: newTabId,
            };

        case REMOVE_BROWSER_TAB:
            if (state.tabs.length === 1) return state; // Prevent removing the last tab
            return {
                ...state,
                tabs: state.tabs.filter((_, idx) => idx !== action.payload),
                activeTabIndex: Math.max(0, state.activeTabIndex - 1), // Adjust active tab index if needed
            };
        case SET_ACTIVE_BROWSER_TAB:
            return {
                ...state,
                activeTabIndex: action.payload, // Update the active tab index
            };
        case SET_TAB_URL:
            return {
                ...state,
                tabs: state.tabs.map((tab, idx) =>
                    idx === action.payload.index ? { ...tab, url: action.payload.url } : tab
                ),
            };
        default:
            return state; // Return unchanged state for unknown action types
    }
};

export default urlManagerReducer;
