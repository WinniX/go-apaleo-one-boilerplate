# PaymentAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | The account number (e.g. masked credit card number or last 4 digits) | [optional] 
**AccountHolder** | Pointer to **string** | The account holder (e.g. card holder) | [optional] 
**ExpiryMonth** | Pointer to **string** | The credit card&#39;s expiration month | [optional] 
**ExpiryYear** | Pointer to **string** | The credit card&#39;s expiration year | [optional] 
**PaymentMethod** | Pointer to **string** | The payment method (e.g. visa) | [optional] 
**PayerEmail** | Pointer to **string** | The email address of the shopper / customer | [optional] 
**PayerReference** | Pointer to **string** | The payer reference. It is used to make recurring captures and its usage is allowed only in the scope of the booking.  For the reason above this is a write-only field. | [optional] 
**IsVirtual** | Pointer to **bool** | Indicates if the payment account is a virtual credit card. If not specified it defaults to &#39;false&#39; | [optional] 
**IsActive** | **bool** | Indicates if the payment account can be used for capturing payments. A payment account is active, when it has a valid payer reference set | 
**InactiveReason** | Pointer to **string** | A reason why account is inactive | [optional] 

## Methods

### NewPaymentAccountModel

`func NewPaymentAccountModel(isActive bool, ) *PaymentAccountModel`

NewPaymentAccountModel instantiates a new PaymentAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAccountModelWithDefaults

`func NewPaymentAccountModelWithDefaults() *PaymentAccountModel`

NewPaymentAccountModelWithDefaults instantiates a new PaymentAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *PaymentAccountModel) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentAccountModel) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentAccountModel) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentAccountModel) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountHolder

`func (o *PaymentAccountModel) GetAccountHolder() string`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *PaymentAccountModel) GetAccountHolderOk() (*string, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *PaymentAccountModel) SetAccountHolder(v string)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *PaymentAccountModel) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *PaymentAccountModel) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *PaymentAccountModel) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *PaymentAccountModel) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *PaymentAccountModel) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *PaymentAccountModel) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *PaymentAccountModel) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *PaymentAccountModel) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *PaymentAccountModel) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentAccountModel) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentAccountModel) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentAccountModel) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentAccountModel) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPayerEmail

`func (o *PaymentAccountModel) GetPayerEmail() string`

GetPayerEmail returns the PayerEmail field if non-nil, zero value otherwise.

### GetPayerEmailOk

`func (o *PaymentAccountModel) GetPayerEmailOk() (*string, bool)`

GetPayerEmailOk returns a tuple with the PayerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerEmail

`func (o *PaymentAccountModel) SetPayerEmail(v string)`

SetPayerEmail sets PayerEmail field to given value.

### HasPayerEmail

`func (o *PaymentAccountModel) HasPayerEmail() bool`

HasPayerEmail returns a boolean if a field has been set.

### GetPayerReference

`func (o *PaymentAccountModel) GetPayerReference() string`

GetPayerReference returns the PayerReference field if non-nil, zero value otherwise.

### GetPayerReferenceOk

`func (o *PaymentAccountModel) GetPayerReferenceOk() (*string, bool)`

GetPayerReferenceOk returns a tuple with the PayerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerReference

`func (o *PaymentAccountModel) SetPayerReference(v string)`

SetPayerReference sets PayerReference field to given value.

### HasPayerReference

`func (o *PaymentAccountModel) HasPayerReference() bool`

HasPayerReference returns a boolean if a field has been set.

### GetIsVirtual

`func (o *PaymentAccountModel) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *PaymentAccountModel) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *PaymentAccountModel) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *PaymentAccountModel) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetIsActive

`func (o *PaymentAccountModel) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PaymentAccountModel) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PaymentAccountModel) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetInactiveReason

`func (o *PaymentAccountModel) GetInactiveReason() string`

GetInactiveReason returns the InactiveReason field if non-nil, zero value otherwise.

### GetInactiveReasonOk

`func (o *PaymentAccountModel) GetInactiveReasonOk() (*string, bool)`

GetInactiveReasonOk returns a tuple with the InactiveReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveReason

`func (o *PaymentAccountModel) SetInactiveReason(v string)`

SetInactiveReason sets InactiveReason field to given value.

### HasInactiveReason

`func (o *PaymentAccountModel) HasInactiveReason() bool`

HasInactiveReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


