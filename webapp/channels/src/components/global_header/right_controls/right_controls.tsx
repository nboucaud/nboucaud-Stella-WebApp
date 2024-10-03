// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useSelector} from 'react-redux';
import {Link, useLocation} from 'react-router-dom';
import styled from 'styled-components';

import type {ProductIdentifier} from '@mattermost/types/products';

import {isCurrentUserGuestUser} from 'mattermost-redux/selectors/entities/users';

import StatusDropdown from 'components/status_dropdown';
import {OnboardingTourSteps, OnboardingTourStepsForGuestUsers} from 'components/tours';
import {
    CustomizeYourExperienceTour,
    useShowOnboardingTutorialStep,
} from 'components/tours/onboarding_tour';

import Pluggable from 'plugins/pluggable';
import {isChannels} from 'utils/products';

import type {GlobalState} from 'types/store';

import PlanUpgradeButton from './plan_upgrade_button';
import PreviousGPTsButton from './previous_gpts_button/previous_gpts_button';
import SavedPostsButton from './saved_posts_button/saved_posts_button';
import SettingsButton from './settings_button';

const RightControlsContainer = styled.div`
    display: flex;
    align-items: center;
    height: 40px;
    flex-shrink: 0;
    position: relative;
    flex-basis: 30%;
    justify-content: flex-end;

    > * + * {
        margin-left: 8px;
    }
`;

const StyledCustomizeYourExperienceTour = styled.div`
    display: flex;
    align-items: center;
    height: 100%
`;

const StyledStatusDropdown = styled.div`
    margin-left: 8px;
`;

export type Props = {
    productId?: ProductIdentifier;
}

const RightControls = ({productId = null}: Props): JSX.Element => {
    // guest validation to see which point the messaging tour tip starts
    const isGuestUser = useSelector((state: GlobalState) => isCurrentUserGuestUser(state));
    const tourStep = isGuestUser ? OnboardingTourStepsForGuestUsers.CUSTOMIZE_EXPERIENCE : OnboardingTourSteps.CUSTOMIZE_EXPERIENCE;
    const {pathname} = useLocation();
    const urlParts = pathname.split('/');
    const teamName = urlParts.length > 1 ? `/${urlParts[1]}` : '';

    const showCustomizeTip = useShowOnboardingTutorialStep(tourStep);

    return (
        <RightControlsContainer
            id={'RightControlsContainer'}
        >
            <Link
                className='btn btn-xs btn-primary'
                to={`${teamName}/membership`}
            >{'Membership'}</Link>
            <PlanUpgradeButton/>
            {isChannels(productId) ? (
                <>
                    <PreviousGPTsButton/>
                    <SavedPostsButton/>
                </>
            ) : (
                <Pluggable
                    pluggableName={'Product'}
                    subComponentName={'headerRightComponent'}
                    pluggableId={productId}
                />
            )}
            <StyledCustomizeYourExperienceTour id='CustomizeYourExperienceTour'>
                {
                    isChannels(productId) ? (
                        <>
                            <SettingsButton/>
                            {showCustomizeTip && <CustomizeYourExperienceTour/>}
                        </>
                    ) : null
                }
                <StyledStatusDropdown>
                    <StatusDropdown/>
                </StyledStatusDropdown>
            </StyledCustomizeYourExperienceTour>
        </RightControlsContainer>
    );
};

export default RightControls;
