# ReservationAssignedUnitTimeRangeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | The start date and time of the period for which the unit is assigned to the reservation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | The end date and time of the period for which the unit is assigned to the reservation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 

## Methods

### NewReservationAssignedUnitTimeRangeModel

`func NewReservationAssignedUnitTimeRangeModel(from time.Time, to time.Time, ) *ReservationAssignedUnitTimeRangeModel`

NewReservationAssignedUnitTimeRangeModel instantiates a new ReservationAssignedUnitTimeRangeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationAssignedUnitTimeRangeModelWithDefaults

`func NewReservationAssignedUnitTimeRangeModelWithDefaults() *ReservationAssignedUnitTimeRangeModel`

NewReservationAssignedUnitTimeRangeModelWithDefaults instantiates a new ReservationAssignedUnitTimeRangeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *ReservationAssignedUnitTimeRangeModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReservationAssignedUnitTimeRangeModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReservationAssignedUnitTimeRangeModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *ReservationAssignedUnitTimeRangeModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReservationAssignedUnitTimeRangeModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReservationAssignedUnitTimeRangeModel) SetTo(v time.Time)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


