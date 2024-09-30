// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import MembershipCard, {type MembershipPlan} from './membership_card';
import MembershipHeader from './membership_header';

const studentPlan: MembershipPlan = {
    title: 'Student',
    info: 'Increased productivity for small teams',
    pricing: '$10.99',
    isCurrentPlan: true,
    topFeatures: [
        'Unlimited workspace teams',
        'Unlimited playbooks and runs',
        'Voice calls and screen share',
        'Full message and file history',
        'SSO with Gitlab',
    ],
};

const gitoPlan: MembershipPlan = {
    title: 'Gito',
    info: 'Scalable solutions for growing teams',
    pricing: '$14.99',
    pricingSubtext: 'USD per seat/month',
    isBilledAnnually: true,
    topFeatures: [
        'Custom user groups',
        'SSO with SAML 2.0, including Okta, OneLogin, and ADFS',
        'SSO support with AD/LDAP, Google, O365, OpenID',
        'Full message and file history',
        'Guest access with MFA enforcement',
    ],
};

const gitoPlusPlan: MembershipPlan = {
    title: 'Gito+',
    info: 'Improved solutions for performing teams',
    pricing: '$20',
    pricingSubtext: 'USD per seat/month',
    isBilledAnnually: true,
    topFeatures: [
        'AD/LDAP group sync',
        'Advanced mobile security via ID-only push notifications',
        'Advanced roles and permissions',
        'Advanced compliance management',
        'Playbook analytics dashboard',
    ],
};

const gitoLinkPlan: MembershipPlan = {
    title: 'Gito Link',
    info: 'Improved solutions for performing teams',
    pricing: '$7.99',
    pricingSubtext: 'USD per seat/month',
    isBilledAnnually: true,
    topFeatures: [
        'AD/LDAP group sync',
        'Advanced mobile security via ID-only push notifications',
        'Advanced roles and permissions',
        'Advanced compliance management',
        'Playbook analytics dashboard',
    ],
};

const gitoSuitePlan: MembershipPlan = {
    title: 'GitoSuite',
    titleSubtext: '(Enterprise)',
    info: 'Administration, security, and compliance for large teams',
    pricing: '',
    topFeatures: [
        'AD/LDAP group sync',
        'Advanced mobile security via ID-only push notifications',
        'Advanced roles and permissions',
        'Advanced compliance management',
        'Playbook analytics dashboard',
    ],
};

const Membership = () => {
    return (
        <div id='membership'>
            <MembershipHeader/>
            <div className='content'>
                <div className='membership-plans'>
                    <div className='hint'>
                        <div>
                            {'Looking for custom solutions?'} <a href='#'>{'Contact Us'}</a>
                        </div>
                    </div>
                    <MembershipCard
                        membershipPlan={studentPlan}
                        color='#3DB887'
                    >
                        <MembershipCard.Title/>
                        <MembershipCard.Info/>
                        <MembershipCard.Pricing/>
                        <MembershipCard.CtaButton disabled={true}>
                            {'Downgrade'}
                        </MembershipCard.CtaButton>
                        <MembershipCard.TopFeatures/>
                    </MembershipCard>
                    <MembershipCard
                        membershipPlan={gitoPlan}
                        color='#02B9B5'
                    >
                        <MembershipCard.Title/>
                        <MembershipCard.Info/>
                        <MembershipCard.Pricing/>
                        <MembershipCard.CtaButton>
                            {'Upgrade'}
                        </MembershipCard.CtaButton>
                        <MembershipCard.TopFeatures/>
                    </MembershipCard>
                    <MembershipCard
                        membershipPlan={gitoPlusPlan}
                        color='#7C99D2'
                    >
                        <MembershipCard.Title/>
                        <MembershipCard.Info/>
                        <MembershipCard.Pricing/>
                        <MembershipCard.CtaButton>
                            {'Upgrade'}
                        </MembershipCard.CtaButton>
                        <MembershipCard.TopFeatures/>
                    </MembershipCard>
                    <MembershipCard
                        membershipPlan={gitoLinkPlan}
                        color='#FFBC1F'
                    >
                        <MembershipCard.Title/>
                        <MembershipCard.Info/>
                        <MembershipCard.Pricing/>
                        <MembershipCard.CtaButton>
                            {'Upgrade'}
                        </MembershipCard.CtaButton>
                        <MembershipCard.TopFeatures/>
                    </MembershipCard>
                    <MembershipCard
                        membershipPlan={gitoSuitePlan}
                        color='#1E325C'
                    >
                        <MembershipCard.Title/>
                        <MembershipCard.Info/>
                        <MembershipCard.PricingContainer>
                            <svg
                                width='110'
                                height='90'
                                viewBox='0 0 110 90'
                                fill='none'
                                xmlns='http://www.w3.org/2000/svg'
                            >
                                <g clipPath='url(#clip0_386_208)'>
                                    <mask
                                        style={{maskType: 'luminance'}}
                                        maskUnits='userSpaceOnUse'
                                        x='0'
                                        y='0'
                                        width='110'
                                        height='90'
                                    >
                                        <path
                                            d='M109.956 0.548096H0.705566V89.7318H109.956V0.548096Z'
                                            fill='white'
                                        />
                                    </mask>
                                    <g mask='url(#mask0_386_208)'>
                                        <path
                                            d='M109.108 32.4206C109.676 32.9867 109.965 33.6965 109.965 34.5502V86.7485C109.965 87.6021 109.676 88.312 109.108 88.8781C108.54 89.4442 107.827 89.7317 106.971 89.7317H3.91589C3.0592 89.7317 2.31073 89.4442 1.67047 88.8781C1.03021 88.312 0.705566 87.6021 0.705566 86.7485V34.5502C0.705566 33.6965 1.03021 32.9867 1.67047 32.4206C2.31073 31.8544 3.0592 31.5669 3.91589 31.5669H106.962C107.818 31.5669 108.531 31.8544 109.099 32.4206H109.108Z'
                                            fill='#818698'
                                        />
                                        <path
                                            d='M10.327 54.5796H6.69287C5.6919 54.4358 5.12378 53.9416 4.97949 53.088V43.0778C5.12378 42.0894 5.6919 41.5862 6.69287 41.5862H10.327C11.319 41.5862 11.8962 42.0894 12.0404 43.0778V53.088C11.8962 53.9416 11.328 54.4358 10.327 54.5796ZM10.327 66.5127H6.69287C5.6919 66.3689 5.12378 65.8747 4.97949 65.021V61.6154C5.12378 60.627 5.6919 60.1238 6.69287 60.1238H10.327C11.319 60.1238 11.8962 60.627 12.0404 61.6154V65.021C11.8962 65.8747 11.328 66.3689 10.327 66.5127ZM10.327 78.0145H6.69287C5.6919 77.8707 5.12378 77.3765 4.97949 76.5228V73.1172C5.12378 72.2636 5.6919 71.7693 6.69287 71.6256H10.327C11.319 71.7693 11.8962 72.2636 12.0404 73.1172V76.5228C11.8962 77.3765 11.328 77.8707 10.327 78.0145ZM20.3728 54.5796H16.5222C15.5213 54.4358 15.0253 53.9416 15.0253 53.088V43.0778C15.0253 42.0894 15.5213 41.5862 16.5222 41.5862H20.3728C21.2295 41.5862 21.7254 42.0894 21.8698 43.0778V53.088C21.7254 53.9416 21.2295 54.4358 20.3728 54.5796ZM20.3728 66.5127H16.5222C15.5213 66.3689 15.0253 65.8747 15.0253 65.021V61.6154C15.0253 60.627 15.5213 60.1238 16.5222 60.1238H20.3728C21.2295 60.1238 21.7254 60.627 21.8698 61.6154V65.021C21.7254 65.8747 21.2295 66.3689 20.3728 66.5127ZM20.3728 78.0145H16.5222C15.5213 77.8707 15.0253 77.3765 15.0253 76.5228V73.1172C15.0253 72.2636 15.5213 71.7693 16.5222 71.6256H20.3728C21.2295 71.7693 21.7254 72.2636 21.8698 73.1172V76.5228C21.7254 77.3765 21.2295 77.8707 20.3728 78.0145ZM96.4826 53.088V43.0778C96.3383 42.0894 95.7702 41.5862 94.7692 41.5862H91.1351C90.1341 41.5862 89.566 42.0894 89.4217 43.0778V53.088C89.566 53.9416 90.1341 54.4358 91.1351 54.5796H94.7692C95.7612 54.4358 96.3383 53.9416 96.4826 53.088ZM94.7692 66.5127H91.1351C90.1341 66.3689 89.566 65.8747 89.4217 65.021V61.6154C89.566 60.627 90.1341 60.1238 91.1351 60.1238H94.7692C95.7612 60.1238 96.3383 60.627 96.4826 61.6154V65.021C96.3383 65.8747 95.7702 66.3689 94.7692 66.5127ZM94.7692 78.0145H91.1351C90.1341 77.8707 89.566 77.3765 89.4217 76.5228V73.1172C89.566 72.2636 90.1341 71.7693 91.1351 71.6256H94.7692C95.7612 71.7693 96.3383 72.2636 96.4826 73.1172V76.5228C96.3383 77.3765 95.7702 77.8707 94.7692 78.0145ZM104.815 54.5796H100.964C99.9635 54.4358 99.4674 53.9416 99.4674 53.088V43.0778C99.4674 42.0894 99.9635 41.5862 100.964 41.5862H104.815C105.672 41.5862 106.168 42.0894 106.312 43.0778V53.088C106.168 53.9416 105.672 54.4358 104.815 54.5796ZM104.815 66.5127H100.964C99.9635 66.3689 99.4674 65.8747 99.4674 65.021V61.6154C99.4674 60.627 99.9635 60.1238 100.964 60.1238H104.815C105.672 60.1238 106.168 60.627 106.312 61.6154V65.021C106.168 65.8747 105.672 66.3689 104.815 66.5127ZM104.815 78.0145H100.964C99.9635 77.8707 99.4674 77.3765 99.4674 76.5228V73.1172C99.4674 72.2636 99.9635 71.7693 100.964 71.6256H104.815C105.672 71.7693 106.168 72.2636 106.312 73.1172V76.5228C106.168 77.3765 105.672 77.8707 104.815 78.0145Z'
                                            fill='#EBEBEF'
                                        />
                                        <path
                                            d='M85.1569 4.38508V87.3866C85.0127 88.8063 84.5167 89.5881 83.66 89.7319H28.7148C27.7138 89.5881 27.1457 88.8063 27.0015 87.3866V4.38508C27.1457 2.96533 27.7138 2.18357 28.7148 2.03979H83.66C84.5167 2.18357 85.0127 2.96533 85.1569 4.38508Z'
                                            fill='#BABEC9'
                                        />
                                        <path
                                            d='M39.4009 46.4834C39.2566 47.4808 38.7606 47.975 37.904 47.975H33.1967C32.1957 47.975 31.6997 47.4808 31.6997 46.4834V34.7659C31.6997 33.7775 32.1957 33.2743 33.1967 33.2743H37.904C38.7606 33.2743 39.2566 33.7775 39.4009 34.7659V46.4834ZM37.904 60.5461C38.7606 60.5461 39.2566 60.0518 39.4009 59.0544V55.6488C39.2566 54.7952 38.7606 54.3009 37.904 54.1572H33.1967C32.1957 54.3009 31.6997 54.7952 31.6997 55.6488V59.0544C31.6997 60.0518 32.1957 60.5461 33.1967 60.5461H37.904ZM37.904 72.9015C38.7606 72.9015 39.2566 72.4073 39.4009 71.4098V68.0042C39.2566 67.1506 38.7606 66.6564 37.904 66.5126H33.1967C32.1957 66.6564 31.6997 67.1506 31.6997 68.0042V71.4098C31.6997 72.4073 32.1957 72.9015 33.1967 72.9015H37.904ZM37.904 84.1966C38.7606 84.0528 39.2566 83.5586 39.4009 82.7049V79.2993C39.2566 78.4457 38.7606 77.9516 37.904 77.8077H33.1967C32.1957 77.9516 31.6997 78.4457 31.6997 79.2993V82.7049C31.6997 83.5586 32.1957 84.0528 33.1967 84.1966H37.904ZM49.0229 47.975C50.0148 47.975 50.5919 47.4808 50.7362 46.4834V34.7659C50.5919 33.7775 50.0238 33.2743 49.0229 33.2743H44.3156C43.4589 33.2743 42.9629 33.7775 42.8187 34.7659V46.4834C42.9629 47.4808 43.4589 47.975 44.3156 47.975H49.0229ZM49.0229 60.5461C50.0148 60.5461 50.5919 60.0518 50.7362 59.0544V55.6488C50.5919 54.7952 50.0238 54.3009 49.0229 54.1572H44.3156C43.4589 54.3009 42.9629 54.7952 42.8187 55.6488V59.0544C42.9629 60.0518 43.4589 60.5461 44.3156 60.5461H49.0229ZM67.6265 47.975C68.4832 47.975 68.9792 47.4808 69.1234 46.4834V34.7659C68.9792 33.7775 68.4832 33.2743 67.6265 33.2743H62.9192C61.9183 33.2743 61.3501 33.7775 61.2058 34.7659V46.4834C61.3501 47.4808 61.9183 47.975 62.9192 47.975H67.6265ZM67.6265 60.5461C68.4832 60.5461 68.9792 60.0518 69.1234 59.0544V55.6488C68.9792 54.7952 68.4832 54.3009 67.6265 54.1572H62.9192C61.9183 54.3009 61.3501 54.7952 61.2058 55.6488V59.0544C61.3501 60.0518 61.9183 60.5461 62.9192 60.5461H67.6265ZM78.7453 47.975C79.7373 47.975 80.2423 47.4808 80.2423 46.4834V34.7659C80.2423 33.7775 79.7373 33.2743 78.7453 33.2743H74.0381C73.1814 33.2743 72.6854 33.7775 72.5412 34.7659V46.4834C72.6854 47.4808 73.1814 47.975 74.0381 47.975H78.7453ZM39.4009 25.2321C39.2566 26.2295 38.7606 26.7237 37.904 26.7237H33.1967C32.1957 26.7237 31.6997 26.2295 31.6997 25.2321V13.5145C31.6997 12.5262 32.1957 12.0229 33.1967 12.0229H37.904C38.7606 12.0229 39.2566 12.5262 39.4009 13.5145V25.2321ZM49.0229 26.7237C50.0148 26.7237 50.5919 26.2295 50.7362 25.2321V13.5145C50.5919 12.5262 50.0238 12.0229 49.0229 12.0229H44.3156C43.4589 12.0229 42.9629 12.5262 42.8187 13.5145V25.2321C42.9629 26.2295 43.4589 26.7237 44.3156 26.7237H49.0229ZM67.6265 26.7237C68.4832 26.7237 68.9792 26.2295 69.1234 25.2321V13.5145C68.9792 12.5262 68.4832 12.0229 67.6265 12.0229H62.9192C61.9183 12.0229 61.3501 12.5262 61.2058 13.5145V25.2321C61.3501 26.2295 61.9183 26.7237 62.9192 26.7237H67.6265ZM78.7453 26.7237C79.7373 26.7237 80.2423 26.2295 80.2423 25.2321V13.5145C80.2423 12.5262 79.7373 12.0229 78.7453 12.0229H74.0381C73.1814 12.0229 72.6854 12.5262 72.5412 13.5145V25.2321C72.6854 26.2295 73.1814 26.7237 74.0381 26.7237H78.7453ZM78.7453 60.5461C79.7373 60.5461 80.2423 60.0518 80.2423 59.0544V55.6488C80.2423 54.7952 79.7373 54.3009 78.7453 54.1572H74.0381C73.1814 54.3009 72.6854 54.7952 72.5412 55.6488V59.0544C72.6854 60.0518 73.1814 60.5461 74.0381 60.5461H78.7453ZM78.7453 72.9015C79.7373 72.9015 80.2423 72.4073 80.2423 71.4098V68.0042C80.2423 67.1506 79.7373 66.6564 78.7453 66.5126H74.0381C73.1814 66.6564 72.6854 67.1506 72.5412 68.0042V71.4098C72.6854 72.4073 73.1814 72.9015 74.0381 72.9015H78.7453ZM78.7453 84.1966C79.7373 84.0528 80.2423 83.5586 80.2423 82.7049V79.2993C80.2423 78.4457 79.7373 77.9516 78.7453 77.8077H74.0381C73.1814 77.9516 72.6854 78.4457 72.5412 79.2993V82.7049C72.6854 83.5586 73.1814 84.0528 74.0381 84.1966H78.7453Z'
                                            fill='#818698'
                                        />
                                        <path
                                            d='M67.2026 66.9348H44.9644V89.7323H67.2026V66.9348Z'
                                            fill='#E0E9EF'
                                        />
                                        <path
                                            d='M85.1568 6.29898H27.2178C26.2169 6.15521 25.6486 5.661 25.5044 4.80734V2.03974C25.6486 1.0513 26.2169 0.548096 27.2178 0.548096H85.1568C86.0135 0.548096 86.5095 1.0513 86.6538 2.03974V4.80734C86.5095 5.661 86.0135 6.15521 85.1568 6.29898ZM67.4099 64.1673H44.7483C43.7474 64.3111 43.2514 64.8053 43.2514 65.6589V68.2199C43.2514 69.2173 43.7474 69.7115 44.7483 69.7115H55.4343V89.7408H56.7148V69.7115H67.4009C68.3928 69.7115 68.97 69.2173 69.1142 68.2199V65.6589C68.97 64.8053 68.4019 64.3111 67.4009 64.1673H67.4099Z'
                                            fill='#5A6072'
                                        />
                                    </g>
                                </g>
                                <defs>
                                    <clipPath id='clip0_386_208'>
                                        <rect
                                            width='109.25'
                                            height='89.1837'
                                            fill='white'
                                            transform='translate(0.705566 0.548096)'
                                        />
                                    </clipPath>
                                </defs>
                            </svg>
                        </MembershipCard.PricingContainer>
                        <MembershipCard.CtaButton>
                            {'Contact Sales'}
                        </MembershipCard.CtaButton>
                        <MembershipCard.TopFeatures/>
                    </MembershipCard>
                </div>
            </div>
        </div>
    );
};

export default Membership;
