/*
 * apaleo Booking API
 *
 * Resources and methods to manage guest journeys.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bookingclient

import (
	"encoding/json"
)

// PaymentAccountModel struct for PaymentAccountModel
type PaymentAccountModel struct {
	// The account number (e.g. masked credit card number or last 4 digits)
	AccountNumber *string `json:"accountNumber,omitempty"`
	// The account holder (e.g. card holder)
	AccountHolder *string `json:"accountHolder,omitempty"`
	// The credit card's expiration month
	ExpiryMonth *string `json:"expiryMonth,omitempty"`
	// The credit card's expiration year
	ExpiryYear *string `json:"expiryYear,omitempty"`
	// The payment method (e.g. visa)
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// The email address of the shopper / customer
	PayerEmail *string `json:"payerEmail,omitempty"`
	// The payer reference. It is used to make recurring captures and its usage is allowed only in the scope of the booking.  For the reason above this is a write-only field.
	PayerReference *string `json:"payerReference,omitempty"`
	// Indicates if the payment account is a virtual credit card. If not specified it defaults to 'false'
	IsVirtual *bool `json:"isVirtual,omitempty"`
	// Indicates if the payment account can be used for capturing payments. A payment account is active, when it has a valid payer reference set
	IsActive bool `json:"isActive"`
	// A reason why account is inactive
	InactiveReason *string `json:"inactiveReason,omitempty"`
}

// NewPaymentAccountModel instantiates a new PaymentAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAccountModel(isActive bool) *PaymentAccountModel {
	this := PaymentAccountModel{}
	this.IsActive = isActive
	return &this
}

// NewPaymentAccountModelWithDefaults instantiates a new PaymentAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAccountModelWithDefaults() *PaymentAccountModel {
	this := PaymentAccountModel{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *PaymentAccountModel) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetAccountHolder() string {
	if o == nil || o.AccountHolder == nil {
		var ret string
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetAccountHolderOk() (*string, bool) {
	if o == nil || o.AccountHolder == nil {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasAccountHolder() bool {
	if o != nil && o.AccountHolder != nil {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given string and assigns it to the AccountHolder field.
func (o *PaymentAccountModel) SetAccountHolder(v string) {
	o.AccountHolder = &v
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetExpiryMonth() string {
	if o == nil || o.ExpiryMonth == nil {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetExpiryMonthOk() (*string, bool) {
	if o == nil || o.ExpiryMonth == nil {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasExpiryMonth() bool {
	if o != nil && o.ExpiryMonth != nil {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *PaymentAccountModel) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetExpiryYear() string {
	if o == nil || o.ExpiryYear == nil {
		var ret string
		return ret
	}
	return *o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetExpiryYearOk() (*string, bool) {
	if o == nil || o.ExpiryYear == nil {
		return nil, false
	}
	return o.ExpiryYear, true
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasExpiryYear() bool {
	if o != nil && o.ExpiryYear != nil {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given string and assigns it to the ExpiryYear field.
func (o *PaymentAccountModel) SetExpiryYear(v string) {
	o.ExpiryYear = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetPaymentMethod() string {
	if o == nil || o.PaymentMethod == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetPaymentMethodOk() (*string, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *PaymentAccountModel) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPayerEmail returns the PayerEmail field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetPayerEmail() string {
	if o == nil || o.PayerEmail == nil {
		var ret string
		return ret
	}
	return *o.PayerEmail
}

// GetPayerEmailOk returns a tuple with the PayerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetPayerEmailOk() (*string, bool) {
	if o == nil || o.PayerEmail == nil {
		return nil, false
	}
	return o.PayerEmail, true
}

// HasPayerEmail returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasPayerEmail() bool {
	if o != nil && o.PayerEmail != nil {
		return true
	}

	return false
}

// SetPayerEmail gets a reference to the given string and assigns it to the PayerEmail field.
func (o *PaymentAccountModel) SetPayerEmail(v string) {
	o.PayerEmail = &v
}

// GetPayerReference returns the PayerReference field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetPayerReference() string {
	if o == nil || o.PayerReference == nil {
		var ret string
		return ret
	}
	return *o.PayerReference
}

// GetPayerReferenceOk returns a tuple with the PayerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetPayerReferenceOk() (*string, bool) {
	if o == nil || o.PayerReference == nil {
		return nil, false
	}
	return o.PayerReference, true
}

// HasPayerReference returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasPayerReference() bool {
	if o != nil && o.PayerReference != nil {
		return true
	}

	return false
}

// SetPayerReference gets a reference to the given string and assigns it to the PayerReference field.
func (o *PaymentAccountModel) SetPayerReference(v string) {
	o.PayerReference = &v
}

// GetIsVirtual returns the IsVirtual field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetIsVirtual() bool {
	if o == nil || o.IsVirtual == nil {
		var ret bool
		return ret
	}
	return *o.IsVirtual
}

// GetIsVirtualOk returns a tuple with the IsVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetIsVirtualOk() (*bool, bool) {
	if o == nil || o.IsVirtual == nil {
		return nil, false
	}
	return o.IsVirtual, true
}

// HasIsVirtual returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasIsVirtual() bool {
	if o != nil && o.IsVirtual != nil {
		return true
	}

	return false
}

// SetIsVirtual gets a reference to the given bool and assigns it to the IsVirtual field.
func (o *PaymentAccountModel) SetIsVirtual(v bool) {
	o.IsVirtual = &v
}

// GetIsActive returns the IsActive field value
func (o *PaymentAccountModel) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *PaymentAccountModel) SetIsActive(v bool) {
	o.IsActive = v
}

// GetInactiveReason returns the InactiveReason field value if set, zero value otherwise.
func (o *PaymentAccountModel) GetInactiveReason() string {
	if o == nil || o.InactiveReason == nil {
		var ret string
		return ret
	}
	return *o.InactiveReason
}

// GetInactiveReasonOk returns a tuple with the InactiveReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAccountModel) GetInactiveReasonOk() (*string, bool) {
	if o == nil || o.InactiveReason == nil {
		return nil, false
	}
	return o.InactiveReason, true
}

// HasInactiveReason returns a boolean if a field has been set.
func (o *PaymentAccountModel) HasInactiveReason() bool {
	if o != nil && o.InactiveReason != nil {
		return true
	}

	return false
}

// SetInactiveReason gets a reference to the given string and assigns it to the InactiveReason field.
func (o *PaymentAccountModel) SetInactiveReason(v string) {
	o.InactiveReason = &v
}

func (o PaymentAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountNumber != nil {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if o.AccountHolder != nil {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	if o.ExpiryMonth != nil {
		toSerialize["expiryMonth"] = o.ExpiryMonth
	}
	if o.ExpiryYear != nil {
		toSerialize["expiryYear"] = o.ExpiryYear
	}
	if o.PaymentMethod != nil {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if o.PayerEmail != nil {
		toSerialize["payerEmail"] = o.PayerEmail
	}
	if o.PayerReference != nil {
		toSerialize["payerReference"] = o.PayerReference
	}
	if o.IsVirtual != nil {
		toSerialize["isVirtual"] = o.IsVirtual
	}
	if true {
		toSerialize["isActive"] = o.IsActive
	}
	if o.InactiveReason != nil {
		toSerialize["inactiveReason"] = o.InactiveReason
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentAccountModel struct {
	value *PaymentAccountModel
	isSet bool
}

func (v NullablePaymentAccountModel) Get() *PaymentAccountModel {
	return v.value
}

func (v *NullablePaymentAccountModel) Set(val *PaymentAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAccountModel(val *PaymentAccountModel) *NullablePaymentAccountModel {
	return &NullablePaymentAccountModel{value: val, isSet: true}
}

func (v NullablePaymentAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


