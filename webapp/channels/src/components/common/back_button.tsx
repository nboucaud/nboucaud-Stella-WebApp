// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import classNames from 'classnames';
import React from 'react';
import {Link} from 'react-router-dom';

type Props = {

    /**
     * URL to return to
     */
    url: string;

    className?: string;

    /**
     * onClick handler when user clicks back button
     */
    onClick?: React.EventHandler<React.MouseEvent>;
}

const BackButton = ({url, className, onClick}: Props): JSX.Element => {
    return (
        <div className={classNames('signup-header', className)}>
            <Link
                data-testid='back_button'
                onClick={onClick}
                to={url}
            >
                <svg
                    width='9'
                    height='15'
                    viewBox='0 0 9 15'
                    fill='none'
                    xmlns='http://www.w3.org/2000/svg'
                >
                    <path
                        d='M7.64331 15L0 7.5L7.64331 0L9 1.33125L2.71338 7.5L9 13.6688L7.64331 15Z'
                        fill='white'
                    />
                </svg>
            </Link>
        </div>
    );
};
BackButton.defaultProps = {
    url: '/',
};

export default BackButton;
