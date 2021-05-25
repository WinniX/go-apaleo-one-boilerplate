# BookReservationServiceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | The id of the service you want to book | 
**Count** | Pointer to **int32** | The number of services to book for each service date. It defaults to the service offer count when not specified. | [optional] 
**Amount** | Pointer to [**MonetaryValueModel**](MonetaryValueModel.md) |  | [optional] 
**Dates** | Pointer to [**[]Date**](Date.md) | The optional dates you want to book the service for; if not specified the default service pattern will be used (e.g. whole stay). | [optional] 

## Methods

### NewBookReservationServiceModel

`func NewBookReservationServiceModel(serviceId string, ) *BookReservationServiceModel`

NewBookReservationServiceModel instantiates a new BookReservationServiceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookReservationServiceModelWithDefaults

`func NewBookReservationServiceModelWithDefaults() *BookReservationServiceModel`

NewBookReservationServiceModelWithDefaults instantiates a new BookReservationServiceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *BookReservationServiceModel) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *BookReservationServiceModel) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *BookReservationServiceModel) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetCount

`func (o *BookReservationServiceModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BookReservationServiceModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BookReservationServiceModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BookReservationServiceModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAmount

`func (o *BookReservationServiceModel) GetAmount() MonetaryValueModel`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BookReservationServiceModel) GetAmountOk() (*MonetaryValueModel, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BookReservationServiceModel) SetAmount(v MonetaryValueModel)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BookReservationServiceModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDates

`func (o *BookReservationServiceModel) GetDates() []Date`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *BookReservationServiceModel) GetDatesOk() (*[]Date, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *BookReservationServiceModel) SetDates(v []Date)`

SetDates sets Dates field to given value.

### HasDates

`func (o *BookReservationServiceModel) HasDates() bool`

HasDates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


