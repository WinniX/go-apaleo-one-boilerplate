# RateRestrictionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLengthOfStay** | Pointer to **int32** | The minimum length of stay in order to book the rate. If at least this number  of time slices are covered by the stay duration the rate will be offered. | [optional] 
**MaxLengthOfStay** | Pointer to **int32** | The maximum length of stay in order to book the rate. If not more than this number  of time slices are covered by the stay duration the rate will be offered. | [optional] 
**Closed** | **bool** | Whether the rate can be booked for a stay-through reservation | 
**ClosedOnArrival** | **bool** | Whether the rate can be booked on the reservation&#39;s arrival date | 
**ClosedOnDeparture** | **bool** | Whether the rate can be booked on the reservation&#39;s departure date | 

## Methods

### NewRateRestrictionsModel

`func NewRateRestrictionsModel(closed bool, closedOnArrival bool, closedOnDeparture bool, ) *RateRestrictionsModel`

NewRateRestrictionsModel instantiates a new RateRestrictionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateRestrictionsModelWithDefaults

`func NewRateRestrictionsModelWithDefaults() *RateRestrictionsModel`

NewRateRestrictionsModelWithDefaults instantiates a new RateRestrictionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinLengthOfStay

`func (o *RateRestrictionsModel) GetMinLengthOfStay() int32`

GetMinLengthOfStay returns the MinLengthOfStay field if non-nil, zero value otherwise.

### GetMinLengthOfStayOk

`func (o *RateRestrictionsModel) GetMinLengthOfStayOk() (*int32, bool)`

GetMinLengthOfStayOk returns a tuple with the MinLengthOfStay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLengthOfStay

`func (o *RateRestrictionsModel) SetMinLengthOfStay(v int32)`

SetMinLengthOfStay sets MinLengthOfStay field to given value.

### HasMinLengthOfStay

`func (o *RateRestrictionsModel) HasMinLengthOfStay() bool`

HasMinLengthOfStay returns a boolean if a field has been set.

### GetMaxLengthOfStay

`func (o *RateRestrictionsModel) GetMaxLengthOfStay() int32`

GetMaxLengthOfStay returns the MaxLengthOfStay field if non-nil, zero value otherwise.

### GetMaxLengthOfStayOk

`func (o *RateRestrictionsModel) GetMaxLengthOfStayOk() (*int32, bool)`

GetMaxLengthOfStayOk returns a tuple with the MaxLengthOfStay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLengthOfStay

`func (o *RateRestrictionsModel) SetMaxLengthOfStay(v int32)`

SetMaxLengthOfStay sets MaxLengthOfStay field to given value.

### HasMaxLengthOfStay

`func (o *RateRestrictionsModel) HasMaxLengthOfStay() bool`

HasMaxLengthOfStay returns a boolean if a field has been set.

### GetClosed

`func (o *RateRestrictionsModel) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *RateRestrictionsModel) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *RateRestrictionsModel) SetClosed(v bool)`

SetClosed sets Closed field to given value.


### GetClosedOnArrival

`func (o *RateRestrictionsModel) GetClosedOnArrival() bool`

GetClosedOnArrival returns the ClosedOnArrival field if non-nil, zero value otherwise.

### GetClosedOnArrivalOk

`func (o *RateRestrictionsModel) GetClosedOnArrivalOk() (*bool, bool)`

GetClosedOnArrivalOk returns a tuple with the ClosedOnArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedOnArrival

`func (o *RateRestrictionsModel) SetClosedOnArrival(v bool)`

SetClosedOnArrival sets ClosedOnArrival field to given value.


### GetClosedOnDeparture

`func (o *RateRestrictionsModel) GetClosedOnDeparture() bool`

GetClosedOnDeparture returns the ClosedOnDeparture field if non-nil, zero value otherwise.

### GetClosedOnDepartureOk

`func (o *RateRestrictionsModel) GetClosedOnDepartureOk() (*bool, bool)`

GetClosedOnDepartureOk returns a tuple with the ClosedOnDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedOnDeparture

`func (o *RateRestrictionsModel) SetClosedOnDeparture(v bool)`

SetClosedOnDeparture sets ClosedOnDeparture field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


