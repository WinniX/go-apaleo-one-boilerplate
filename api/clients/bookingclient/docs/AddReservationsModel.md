# AddReservationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservations** | [**[]CreateReservationModel**](CreateReservationModel.md) | List of reservations to add to the existing booking | 
**TransactionReference** | Pointer to **string** | The reference of a payment transaction. This should be set when a payment transaction has been initiated and should be used to complete the transaction upon reservation creation. | [optional] 

## Methods

### NewAddReservationsModel

`func NewAddReservationsModel(reservations []CreateReservationModel, ) *AddReservationsModel`

NewAddReservationsModel instantiates a new AddReservationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddReservationsModelWithDefaults

`func NewAddReservationsModelWithDefaults() *AddReservationsModel`

NewAddReservationsModelWithDefaults instantiates a new AddReservationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservations

`func (o *AddReservationsModel) GetReservations() []CreateReservationModel`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *AddReservationsModel) GetReservationsOk() (*[]CreateReservationModel, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *AddReservationsModel) SetReservations(v []CreateReservationModel)`

SetReservations sets Reservations field to given value.


### GetTransactionReference

`func (o *AddReservationsModel) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *AddReservationsModel) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *AddReservationsModel) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *AddReservationsModel) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


