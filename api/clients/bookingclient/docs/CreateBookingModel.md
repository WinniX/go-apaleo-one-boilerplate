# CreateBookingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAccount** | Pointer to [**CreatePaymentAccountModel**](CreatePaymentAccountModel.md) |  | [optional] 
**Booker** | [**BookerModel**](BookerModel.md) |  | 
**Comment** | Pointer to **string** | Additional information and comments | [optional] 
**BookerComment** | Pointer to **string** | Additional information and comments by the booker | [optional] 
**Reservations** | [**[]CreateReservationModel**](CreateReservationModel.md) | List of reservations to create | 
**TransactionReference** | Pointer to **string** | The reference of a payment transaction. This should be set when a payment transaction has been initiated and should be used to complete the transaction upon reservation creation. | [optional] 

## Methods

### NewCreateBookingModel

`func NewCreateBookingModel(booker BookerModel, reservations []CreateReservationModel, ) *CreateBookingModel`

NewCreateBookingModel instantiates a new CreateBookingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBookingModelWithDefaults

`func NewCreateBookingModelWithDefaults() *CreateBookingModel`

NewCreateBookingModelWithDefaults instantiates a new CreateBookingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentAccount

`func (o *CreateBookingModel) GetPaymentAccount() CreatePaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *CreateBookingModel) GetPaymentAccountOk() (*CreatePaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *CreateBookingModel) SetPaymentAccount(v CreatePaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *CreateBookingModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetBooker

`func (o *CreateBookingModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *CreateBookingModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *CreateBookingModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.


### GetComment

`func (o *CreateBookingModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateBookingModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateBookingModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateBookingModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetBookerComment

`func (o *CreateBookingModel) GetBookerComment() string`

GetBookerComment returns the BookerComment field if non-nil, zero value otherwise.

### GetBookerCommentOk

`func (o *CreateBookingModel) GetBookerCommentOk() (*string, bool)`

GetBookerCommentOk returns a tuple with the BookerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookerComment

`func (o *CreateBookingModel) SetBookerComment(v string)`

SetBookerComment sets BookerComment field to given value.

### HasBookerComment

`func (o *CreateBookingModel) HasBookerComment() bool`

HasBookerComment returns a boolean if a field has been set.

### GetReservations

`func (o *CreateBookingModel) GetReservations() []CreateReservationModel`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *CreateBookingModel) GetReservationsOk() (*[]CreateReservationModel, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *CreateBookingModel) SetReservations(v []CreateReservationModel)`

SetReservations sets Reservations field to given value.


### GetTransactionReference

`func (o *CreateBookingModel) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *CreateBookingModel) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *CreateBookingModel) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *CreateBookingModel) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


