# PeriodModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hours** | Pointer to **int64** | The number of hours within the period | [optional] 
**Days** | Pointer to **int32** | The number of days within the period | [optional] 
**Months** | Pointer to **int32** | The number of months within the period | [optional] 

## Methods

### NewPeriodModel

`func NewPeriodModel() *PeriodModel`

NewPeriodModel instantiates a new PeriodModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodModelWithDefaults

`func NewPeriodModelWithDefaults() *PeriodModel`

NewPeriodModelWithDefaults instantiates a new PeriodModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHours

`func (o *PeriodModel) GetHours() int64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *PeriodModel) GetHoursOk() (*int64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *PeriodModel) SetHours(v int64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *PeriodModel) HasHours() bool`

HasHours returns a boolean if a field has been set.

### GetDays

`func (o *PeriodModel) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PeriodModel) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PeriodModel) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *PeriodModel) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetMonths

`func (o *PeriodModel) GetMonths() int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *PeriodModel) GetMonthsOk() (*int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *PeriodModel) SetMonths(v int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *PeriodModel) HasMonths() bool`

HasMonths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


