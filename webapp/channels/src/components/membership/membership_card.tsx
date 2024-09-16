// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {createContext, useContext, type ReactNode} from 'react';
import {useIntl} from 'react-intl';

import CtaButton from './membership_card_cta_button';
import Info from './membership_card_info';
import Pricing from './membership_card_pricing';
import PricingContainer from './membership_card_pricing_container';
import Title from './membership_card_title';
import TopFeatures from './membership_card_top_features';

export type MembershipPlan = {
    title: string;
    titleSubtext?: string;
    info: string;
    topFeatures: string[];
    pricing: string;
    pricingSubtext?: string;
    isBilledAnnually?: boolean;
    isCurrentPlan?: boolean;
}

export type Props = {
    membershipPlan: MembershipPlan;
    color: string;
    children: ReactNode;
};

const MembershipCardContext = createContext<MembershipPlan | undefined>(undefined);

export const useMembershipCardContext = () => {
    const membershipPlan = useContext(MembershipCardContext);
    if (membershipPlan === undefined) {
        throw new Error('useMembdershipCardContext() called outside of MembershipCardContext provider');
    }
    return membershipPlan;
};

const MembershipCard = ({membershipPlan, color, children}: Props) => {
    const {formatMessage} = useIntl();

    return (
        <MembershipCardContext.Provider value={membershipPlan}>
            <div className='membership-card-container'>
                {membershipPlan.isCurrentPlan && <div className='current-plan'>
                    <svg
                        width='19'
                        height='19'
                        viewBox='0 0 19 19'
                        fill='none'
                        xmlns='http://www.w3.org/2000/svg'
                    >
                        <g clipPath='url(#clip0_386_125)'>
                            <path
                                d='M9.34151 1.05566C7.83025 1.05566 6.42598 1.43682 5.1287 2.19914C3.85817 2.93471 2.85512 3.93776 2.11955 5.20829C1.35723 6.50556 0.976074 7.90984 0.976074 9.4211C0.976074 10.9324 1.35723 12.3366 2.11955 13.6339C2.85512 14.9044 3.85817 15.9075 5.1287 16.6431C6.42598 17.4054 7.83025 17.7865 9.34151 17.7865C10.8528 17.7865 12.257 17.4054 13.5543 16.6431C14.8249 15.9075 15.8279 14.9044 16.5635 13.6339C17.3258 12.3366 17.7069 10.9324 17.7069 9.4211C17.7069 7.90984 17.3258 6.50556 16.5635 5.20829C15.8279 3.93776 14.8249 2.93471 13.5543 2.19914C12.257 1.43682 10.8528 1.05566 9.34151 1.05566ZM7.87706 13.333L6.69346 12.1695L4.32626 9.80226L5.50986 8.61866L7.87706 10.9859L13.1932 5.64963L14.3768 6.83323L7.87706 13.333Z'
                                fill='#3DB887'
                            />
                        </g>
                        <defs>
                            <clipPath>
                                <rect
                                    width='17.832'
                                    height='17.832'
                                    fill='white'
                                    transform='translate(0.425537 0.505127)'
                                />
                            </clipPath>
                        </defs>
                    </svg>

                    {formatMessage({id: 'membership_card.current_plan', defaultMessage: 'current plan'})}
                </div>}
                <div
                    className='membership-card alt a11y__region'
                    style={{'--strip-color': color} as React.CSSProperties}
                >
                    {children}
                </div>
            </div>
        </MembershipCardContext.Provider>
    );
};

MembershipCard.Title = Title;
MembershipCard.Info = Info;
MembershipCard.Pricing = Pricing;
MembershipCard.PricingContainer = PricingContainer;
MembershipCard.CtaButton = CtaButton;
MembershipCard.TopFeatures = TopFeatures;

export default MembershipCard;
