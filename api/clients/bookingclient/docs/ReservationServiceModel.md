# ReservationServiceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | [**EmbeddedServiceModel**](EmbeddedServiceModel.md) |  | 
**ServiceDate** | **string** | The date this service is delivered | 
**Count** | **int32** | The count of booked services | 
**Amount** | [**AmountModel**](AmountModel.md) |  | 
**BookedAsExtra** | **bool** | Whether this service is already booked as extra | 

## Methods

### NewReservationServiceModel

`func NewReservationServiceModel(service EmbeddedServiceModel, serviceDate string, count int32, amount AmountModel, bookedAsExtra bool, ) *ReservationServiceModel`

NewReservationServiceModel instantiates a new ReservationServiceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationServiceModelWithDefaults

`func NewReservationServiceModelWithDefaults() *ReservationServiceModel`

NewReservationServiceModelWithDefaults instantiates a new ReservationServiceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ReservationServiceModel) GetService() EmbeddedServiceModel`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ReservationServiceModel) GetServiceOk() (*EmbeddedServiceModel, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ReservationServiceModel) SetService(v EmbeddedServiceModel)`

SetService sets Service field to given value.


### GetServiceDate

`func (o *ReservationServiceModel) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *ReservationServiceModel) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *ReservationServiceModel) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetCount

`func (o *ReservationServiceModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReservationServiceModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReservationServiceModel) SetCount(v int32)`

SetCount sets Count field to given value.


### GetAmount

`func (o *ReservationServiceModel) GetAmount() AmountModel`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReservationServiceModel) GetAmountOk() (*AmountModel, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReservationServiceModel) SetAmount(v AmountModel)`

SetAmount sets Amount field to given value.


### GetBookedAsExtra

`func (o *ReservationServiceModel) GetBookedAsExtra() bool`

GetBookedAsExtra returns the BookedAsExtra field if non-nil, zero value otherwise.

### GetBookedAsExtraOk

`func (o *ReservationServiceModel) GetBookedAsExtraOk() (*bool, bool)`

GetBookedAsExtraOk returns a tuple with the BookedAsExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookedAsExtra

`func (o *ReservationServiceModel) SetBookedAsExtra(v bool)`

SetBookedAsExtra sets BookedAsExtra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


