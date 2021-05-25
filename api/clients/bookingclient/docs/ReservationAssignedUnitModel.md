# ReservationAssignedUnitModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | [**EmbeddedUnitModel**](EmbeddedUnitModel.md) |  | 
**TimeRanges** | [**[]ReservationAssignedUnitTimeRangeModel**](ReservationAssignedUnitTimeRangeModel.md) | The time ranges for which the unit is assigned to the reservation | 

## Methods

### NewReservationAssignedUnitModel

`func NewReservationAssignedUnitModel(unit EmbeddedUnitModel, timeRanges []ReservationAssignedUnitTimeRangeModel, ) *ReservationAssignedUnitModel`

NewReservationAssignedUnitModel instantiates a new ReservationAssignedUnitModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationAssignedUnitModelWithDefaults

`func NewReservationAssignedUnitModelWithDefaults() *ReservationAssignedUnitModel`

NewReservationAssignedUnitModelWithDefaults instantiates a new ReservationAssignedUnitModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ReservationAssignedUnitModel) GetUnit() EmbeddedUnitModel`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ReservationAssignedUnitModel) GetUnitOk() (*EmbeddedUnitModel, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ReservationAssignedUnitModel) SetUnit(v EmbeddedUnitModel)`

SetUnit sets Unit field to given value.


### GetTimeRanges

`func (o *ReservationAssignedUnitModel) GetTimeRanges() []ReservationAssignedUnitTimeRangeModel`

GetTimeRanges returns the TimeRanges field if non-nil, zero value otherwise.

### GetTimeRangesOk

`func (o *ReservationAssignedUnitModel) GetTimeRangesOk() (*[]ReservationAssignedUnitTimeRangeModel, bool)`

GetTimeRangesOk returns a tuple with the TimeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRanges

`func (o *ReservationAssignedUnitModel) SetTimeRanges(v []ReservationAssignedUnitTimeRangeModel)`

SetTimeRanges sets TimeRanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


