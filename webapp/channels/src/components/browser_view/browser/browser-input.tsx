// src/components/UrlInput.tsx
import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { setTabUrl } from "./browser-state";

const UrlInput: React.FC = () => {
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
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                value={url}
                onChange={(e) => setUrl(e.target.value)}
                placeholder="Enter URL"
                required
            />
            <button type="submit">Add URL</button>
        </form>
    );
};

export default UrlInput;
