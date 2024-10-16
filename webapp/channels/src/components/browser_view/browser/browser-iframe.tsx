import React from "react";

interface UrlIframeProps {
    url: string;
}

const UrlIframe: React.FC<UrlIframeProps> = ({ url }) => {
    return (
        <div className="iframe-container">
            <iframe
                src={url}
                width="100%"
                height="100%"
                title={`URL Iframe - ${url}`}
                sandbox="allow-scripts allow-same-origin allow-popups allow-forms"
                style={{ border: "none" }}
            />
        </div>
    );
};

export default UrlIframe;
