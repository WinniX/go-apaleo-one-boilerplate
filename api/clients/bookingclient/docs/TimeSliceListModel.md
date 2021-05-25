# TimeSliceListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeSlices** | Pointer to [**[]TimeSliceItemModel**](TimeSliceItemModel.md) | List of time slices | [optional] 
**Count** | **int64** | Total count of items | 

## Methods

### NewTimeSliceListModel

`func NewTimeSliceListModel(count int64, ) *TimeSliceListModel`

NewTimeSliceListModel instantiates a new TimeSliceListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSliceListModelWithDefaults

`func NewTimeSliceListModelWithDefaults() *TimeSliceListModel`

NewTimeSliceListModelWithDefaults instantiates a new TimeSliceListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeSlices

`func (o *TimeSliceListModel) GetTimeSlices() []TimeSliceItemModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *TimeSliceListModel) GetTimeSlicesOk() (*[]TimeSliceItemModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *TimeSliceListModel) SetTimeSlices(v []TimeSliceItemModel)`

SetTimeSlices sets TimeSlices field to given value.

### HasTimeSlices

`func (o *TimeSliceListModel) HasTimeSlices() bool`

HasTimeSlices returns a boolean if a field has been set.

### GetCount

`func (o *TimeSliceListModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TimeSliceListModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TimeSliceListModel) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


