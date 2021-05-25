# Date

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceDate** | **string** | The date the service is delivered | 
**Count** | Pointer to **int32** | The number of services to book for this date. It defaults to the service offer count when not specified. | [optional] 
**Amount** | Pointer to [**MonetaryValueModel**](MonetaryValueModel.md) |  | [optional] 

## Methods

### NewDate

`func NewDate(serviceDate string, ) *Date`

NewDate instantiates a new Date object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateWithDefaults

`func NewDateWithDefaults() *Date`

NewDateWithDefaults instantiates a new Date object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceDate

`func (o *Date) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *Date) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *Date) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetCount

`func (o *Date) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Date) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Date) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Date) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAmount

`func (o *Date) GetAmount() MonetaryValueModel`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Date) GetAmountOk() (*MonetaryValueModel, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Date) SetAmount(v MonetaryValueModel)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Date) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


