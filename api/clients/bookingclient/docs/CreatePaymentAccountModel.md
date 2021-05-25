# CreatePaymentAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | The account number (e.g. masked credit card number or last 4 digits) | [optional] 
**AccountHolder** | Pointer to **string** | The account holder (e.g. card holder) | [optional] 
**ExpiryMonth** | Pointer to **string** | The credit card&#39;s expiration month | [optional] 
**ExpiryYear** | Pointer to **string** | The credit card&#39;s expiration year | [optional] 
**PaymentMethod** | Pointer to **string** | The payment method (e.g. visa) | [optional] 
**PayerEmail** | Pointer to **string** | The email address of the shopper / customer | [optional] 
**PayerReference** | Pointer to **string** | The reference used to uniquely identify the shopper (e.g. user ID or account ID). Used for recurring payments | [optional] 
**TransactionReference** | Pointer to **string** | The reference of a payment transaction. This should be set when a payment transaction has been initiated and should be used to complete the transaction upon reservation creation. - &lt;b&gt;DEPRECATED: This property will be removed 17.07.2021. Use &#x60;TransactionReference&#x60; on the booking/reservation model instead&lt;/b&gt; | [optional] 
**IsVirtual** | Pointer to **bool** | Indicates if the payment account is a virtual credit card. If not specified it defaults to &#39;false&#39; | [optional] 
**InactiveReason** | Pointer to **string** | A reason why account is inactive when PayerReference was not provided | [optional] 

## Methods

### NewCreatePaymentAccountModel

`func NewCreatePaymentAccountModel() *CreatePaymentAccountModel`

NewCreatePaymentAccountModel instantiates a new CreatePaymentAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentAccountModelWithDefaults

`func NewCreatePaymentAccountModelWithDefaults() *CreatePaymentAccountModel`

NewCreatePaymentAccountModelWithDefaults instantiates a new CreatePaymentAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *CreatePaymentAccountModel) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreatePaymentAccountModel) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreatePaymentAccountModel) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreatePaymentAccountModel) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountHolder

`func (o *CreatePaymentAccountModel) GetAccountHolder() string`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *CreatePaymentAccountModel) GetAccountHolderOk() (*string, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *CreatePaymentAccountModel) SetAccountHolder(v string)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *CreatePaymentAccountModel) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetExpiryMonth

`func (o *CreatePaymentAccountModel) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CreatePaymentAccountModel) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CreatePaymentAccountModel) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.

### HasExpiryMonth

`func (o *CreatePaymentAccountModel) HasExpiryMonth() bool`

HasExpiryMonth returns a boolean if a field has been set.

### GetExpiryYear

`func (o *CreatePaymentAccountModel) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CreatePaymentAccountModel) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CreatePaymentAccountModel) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.

### HasExpiryYear

`func (o *CreatePaymentAccountModel) HasExpiryYear() bool`

HasExpiryYear returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CreatePaymentAccountModel) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreatePaymentAccountModel) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreatePaymentAccountModel) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CreatePaymentAccountModel) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPayerEmail

`func (o *CreatePaymentAccountModel) GetPayerEmail() string`

GetPayerEmail returns the PayerEmail field if non-nil, zero value otherwise.

### GetPayerEmailOk

`func (o *CreatePaymentAccountModel) GetPayerEmailOk() (*string, bool)`

GetPayerEmailOk returns a tuple with the PayerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerEmail

`func (o *CreatePaymentAccountModel) SetPayerEmail(v string)`

SetPayerEmail sets PayerEmail field to given value.

### HasPayerEmail

`func (o *CreatePaymentAccountModel) HasPayerEmail() bool`

HasPayerEmail returns a boolean if a field has been set.

### GetPayerReference

`func (o *CreatePaymentAccountModel) GetPayerReference() string`

GetPayerReference returns the PayerReference field if non-nil, zero value otherwise.

### GetPayerReferenceOk

`func (o *CreatePaymentAccountModel) GetPayerReferenceOk() (*string, bool)`

GetPayerReferenceOk returns a tuple with the PayerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerReference

`func (o *CreatePaymentAccountModel) SetPayerReference(v string)`

SetPayerReference sets PayerReference field to given value.

### HasPayerReference

`func (o *CreatePaymentAccountModel) HasPayerReference() bool`

HasPayerReference returns a boolean if a field has been set.

### GetTransactionReference

`func (o *CreatePaymentAccountModel) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *CreatePaymentAccountModel) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *CreatePaymentAccountModel) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *CreatePaymentAccountModel) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.

### GetIsVirtual

`func (o *CreatePaymentAccountModel) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *CreatePaymentAccountModel) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *CreatePaymentAccountModel) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *CreatePaymentAccountModel) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetInactiveReason

`func (o *CreatePaymentAccountModel) GetInactiveReason() string`

GetInactiveReason returns the InactiveReason field if non-nil, zero value otherwise.

### GetInactiveReasonOk

`func (o *CreatePaymentAccountModel) GetInactiveReasonOk() (*string, bool)`

GetInactiveReasonOk returns a tuple with the InactiveReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveReason

`func (o *CreatePaymentAccountModel) SetInactiveReason(v string)`

SetInactiveReason sets InactiveReason field to given value.

### HasInactiveReason

`func (o *CreatePaymentAccountModel) HasInactiveReason() bool`

HasInactiveReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


