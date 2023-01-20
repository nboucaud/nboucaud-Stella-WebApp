// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// CloudInterface is an autogenerated mock type for the CloudInterface type
type CloudInterface struct {
	mock.Mock
}

// BootstrapSelfHostedSignup provides a mock function with given fields: req
func (_m *CloudInterface) BootstrapSelfHostedSignup(req model.BootstrapSelfHostedSignupRequest) (*model.BootstrapSelfHostedSignupResponse, error) {
	ret := _m.Called(req)

	var r0 *model.BootstrapSelfHostedSignupResponse
	if rf, ok := ret.Get(0).(func(model.BootstrapSelfHostedSignupRequest) *model.BootstrapSelfHostedSignupResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BootstrapSelfHostedSignupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.BootstrapSelfHostedSignupRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangeSubscription provides a mock function with given fields: userID, subscriptionID, subscriptionChange
func (_m *CloudInterface) ChangeSubscription(userID string, subscriptionID string, subscriptionChange *model.SubscriptionChange) (*model.Subscription, error) {
	ret := _m.Called(userID, subscriptionID, subscriptionChange)

	var r0 *model.Subscription
	if rf, ok := ret.Get(0).(func(string, string, *model.SubscriptionChange) *model.Subscription); ok {
		r0 = rf(userID, subscriptionID, subscriptionChange)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.SubscriptionChange) error); ok {
		r1 = rf(userID, subscriptionID, subscriptionChange)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckCWSConnection provides a mock function with given fields: userId
func (_m *CloudInterface) CheckCWSConnection(userId string) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfirmCustomerPayment provides a mock function with given fields: userID, confirmRequest
func (_m *CloudInterface) ConfirmCustomerPayment(userID string, confirmRequest *model.ConfirmPaymentMethodRequest) error {
	ret := _m.Called(userID, confirmRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *model.ConfirmPaymentMethodRequest) error); ok {
		r0 = rf(userID, confirmRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfirmSelfHostedSignup provides a mock function with given fields: req, requesterEmail
func (_m *CloudInterface) ConfirmSelfHostedSignup(req model.SelfHostedConfirmPaymentMethodRequest, requesterEmail string) (*model.SelfHostedSignupConfirmResponse, error) {
	ret := _m.Called(req, requesterEmail)

	var r0 *model.SelfHostedSignupConfirmResponse
	if rf, ok := ret.Get(0).(func(model.SelfHostedConfirmPaymentMethodRequest, string) *model.SelfHostedSignupConfirmResponse); ok {
		r0 = rf(req, requesterEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SelfHostedSignupConfirmResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.SelfHostedConfirmPaymentMethodRequest, string) error); ok {
		r1 = rf(req, requesterEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfirmSelfHostedSignupLicenseApplication provides a mock function with given fields:
func (_m *CloudInterface) ConfirmSelfHostedSignupLicenseApplication() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCustomerPayment provides a mock function with given fields: userID
func (_m *CloudInterface) CreateCustomerPayment(userID string) (*model.StripeSetupIntent, error) {
	ret := _m.Called(userID)

	var r0 *model.StripeSetupIntent
	if rf, ok := ret.Get(0).(func(string) *model.StripeSetupIntent); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StripeSetupIntent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCustomerSelfHostedSignup provides a mock function with given fields: req, requesterEmail
func (_m *CloudInterface) CreateCustomerSelfHostedSignup(req model.SelfHostedCustomerForm, requesterEmail string) (*model.SelfHostedSignupCustomerResponse, error) {
	ret := _m.Called(req, requesterEmail)

	var r0 *model.SelfHostedSignupCustomerResponse
	if rf, ok := ret.Get(0).(func(model.SelfHostedCustomerForm, string) *model.SelfHostedSignupCustomerResponse); ok {
		r0 = rf(req, requesterEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SelfHostedSignupCustomerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.SelfHostedCustomerForm, string) error); ok {
		r1 = rf(req, requesterEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdateSubscriptionHistoryEvent provides a mock function with given fields: userID, userCount
func (_m *CloudInterface) CreateOrUpdateSubscriptionHistoryEvent(userID string, userCount int) (*model.SubscriptionHistory, error) {
	ret := _m.Called(userID, userCount)

	var r0 *model.SubscriptionHistory
	if rf, ok := ret.Get(0).(func(string, int) *model.SubscriptionHistory); ok {
		r0 = rf(userID, userCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SubscriptionHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(userID, userCount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudCustomer provides a mock function with given fields: userID
func (_m *CloudInterface) GetCloudCustomer(userID string) (*model.CloudCustomer, error) {
	ret := _m.Called(userID)

	var r0 *model.CloudCustomer
	if rf, ok := ret.Get(0).(func(string) *model.CloudCustomer); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CloudCustomer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudLimits provides a mock function with given fields: userID
func (_m *CloudInterface) GetCloudLimits(userID string) (*model.ProductLimits, error) {
	ret := _m.Called(userID)

	var r0 *model.ProductLimits
	if rf, ok := ret.Get(0).(func(string) *model.ProductLimits); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProductLimits)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudProduct provides a mock function with given fields: userID, productID
func (_m *CloudInterface) GetCloudProduct(userID string, productID string) (*model.Product, error) {
	ret := _m.Called(userID, productID)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(string, string) *model.Product); ok {
		r0 = rf(userID, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudProducts provides a mock function with given fields: userID, includeLegacyProducts
func (_m *CloudInterface) GetCloudProducts(userID string, includeLegacyProducts bool) ([]*model.Product, error) {
	ret := _m.Called(userID, includeLegacyProducts)

	var r0 []*model.Product
	if rf, ok := ret.Get(0).(func(string, bool) []*model.Product); ok {
		r0 = rf(userID, includeLegacyProducts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(userID, includeLegacyProducts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInvoicePDF provides a mock function with given fields: userID, invoiceID
func (_m *CloudInterface) GetInvoicePDF(userID string, invoiceID string) ([]byte, string, error) {
	ret := _m.Called(userID, invoiceID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(userID, invoiceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(userID, invoiceID)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(userID, invoiceID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetInvoicesForSubscription provides a mock function with given fields: userID
func (_m *CloudInterface) GetInvoicesForSubscription(userID string) ([]*model.Invoice, error) {
	ret := _m.Called(userID)

	var r0 []*model.Invoice
	if rf, ok := ret.Get(0).(func(string) []*model.Invoice); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Invoice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLicenseExpandStatus provides a mock function with given fields: userID, token
func (_m *CloudInterface) GetLicenseExpandStatus(userID string, token string) (*model.SubscriptionExpandStatus, error) {
	ret := _m.Called(userID, token)

	var r0 *model.SubscriptionExpandStatus
	if rf, ok := ret.Get(0).(func(string, string) *model.SubscriptionExpandStatus); ok {
		r0 = rf(userID, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SubscriptionExpandStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLicenseRenewalStatus provides a mock function with given fields: userID, token
func (_m *CloudInterface) GetLicenseRenewalStatus(userID string, token string) error {
	ret := _m.Called(userID, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSelfHostedInvoicePDF provides a mock function with given fields: invoiceID
func (_m *CloudInterface) GetSelfHostedInvoicePDF(invoiceID string) ([]byte, string, error) {
	ret := _m.Called(invoiceID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(invoiceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(invoiceID)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(invoiceID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSelfHostedInvoices provides a mock function with given fields:
func (_m *CloudInterface) GetSelfHostedInvoices() ([]*model.Invoice, error) {
	ret := _m.Called()

	var r0 []*model.Invoice
	if rf, ok := ret.Get(0).(func() []*model.Invoice); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Invoice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSelfHostedProducts provides a mock function with given fields: userID
func (_m *CloudInterface) GetSelfHostedProducts(userID string) ([]*model.Product, error) {
	ret := _m.Called(userID)

	var r0 []*model.Product
	if rf, ok := ret.Get(0).(func(string) []*model.Product); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscription provides a mock function with given fields: userID
func (_m *CloudInterface) GetSubscription(userID string) (*model.Subscription, error) {
	ret := _m.Called(userID)

	var r0 *model.Subscription
	if rf, ok := ret.Get(0).(func(string) *model.Subscription); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleLicenseChange provides a mock function with given fields:
func (_m *CloudInterface) HandleLicenseChange() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InvalidateCaches provides a mock function with given fields:
func (_m *CloudInterface) InvalidateCaches() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestCloudTrial provides a mock function with given fields: userID, subscriptionID, newValidBusinessEmail
func (_m *CloudInterface) RequestCloudTrial(userID string, subscriptionID string, newValidBusinessEmail string) (*model.Subscription, error) {
	ret := _m.Called(userID, subscriptionID, newValidBusinessEmail)

	var r0 *model.Subscription
	if rf, ok := ret.Get(0).(func(string, string, string) *model.Subscription); ok {
		r0 = rf(userID, subscriptionID, newValidBusinessEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(userID, subscriptionID, newValidBusinessEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelfHostedSignupAvailable provides a mock function with given fields:
func (_m *CloudInterface) SelfHostedSignupAvailable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCloudCustomer provides a mock function with given fields: userID, customerInfo
func (_m *CloudInterface) UpdateCloudCustomer(userID string, customerInfo *model.CloudCustomerInfo) (*model.CloudCustomer, error) {
	ret := _m.Called(userID, customerInfo)

	var r0 *model.CloudCustomer
	if rf, ok := ret.Get(0).(func(string, *model.CloudCustomerInfo) *model.CloudCustomer); ok {
		r0 = rf(userID, customerInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CloudCustomer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *model.CloudCustomerInfo) error); ok {
		r1 = rf(userID, customerInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCloudCustomerAddress provides a mock function with given fields: userID, address
func (_m *CloudInterface) UpdateCloudCustomerAddress(userID string, address *model.Address) (*model.CloudCustomer, error) {
	ret := _m.Called(userID, address)

	var r0 *model.CloudCustomer
	if rf, ok := ret.Get(0).(func(string, *model.Address) *model.CloudCustomer); ok {
		r0 = rf(userID, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CloudCustomer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *model.Address) error); ok {
		r1 = rf(userID, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateBusinessEmail provides a mock function with given fields: userID, email
func (_m *CloudInterface) ValidateBusinessEmail(userID string, email string) error {
	ret := _m.Called(userID, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
