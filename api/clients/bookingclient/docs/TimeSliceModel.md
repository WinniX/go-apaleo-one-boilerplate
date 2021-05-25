# TimeSliceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | The start date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | The end date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**ServiceDate** | **string** | The service date for this time slice | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**UnitGroup** | [**EmbeddedUnitGroupModel**](EmbeddedUnitGroupModel.md) |  | 
**Unit** | Pointer to [**EmbeddedUnitModel**](EmbeddedUnitModel.md) |  | [optional] 
**BaseAmount** | [**AmountModel**](AmountModel.md) |  | 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**IncludedServices** | Pointer to [**[]ReservationServiceModel**](ReservationServiceModel.md) | The list of services included in the rate plan (package elements) | [optional] 
**Actions** | Pointer to [**[]ActionModelReservationTimeSliceActionNotAllowedReservationTimeSliceActionReason**](ActionModelReservationTimeSliceActionNotAllowedReservationTimeSliceActionReason.md) | The list of actions allowed for this time slice | [optional] 

## Methods

### NewTimeSliceModel

`func NewTimeSliceModel(from time.Time, to time.Time, serviceDate string, ratePlan EmbeddedRatePlanModel, unitGroup EmbeddedUnitGroupModel, baseAmount AmountModel, totalGrossAmount MonetaryValueModel, ) *TimeSliceModel`

NewTimeSliceModel instantiates a new TimeSliceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSliceModelWithDefaults

`func NewTimeSliceModelWithDefaults() *TimeSliceModel`

NewTimeSliceModelWithDefaults instantiates a new TimeSliceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TimeSliceModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TimeSliceModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TimeSliceModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *TimeSliceModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TimeSliceModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TimeSliceModel) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetServiceDate

`func (o *TimeSliceModel) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *TimeSliceModel) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *TimeSliceModel) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetRatePlan

`func (o *TimeSliceModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *TimeSliceModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *TimeSliceModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetUnitGroup

`func (o *TimeSliceModel) GetUnitGroup() EmbeddedUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *TimeSliceModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *TimeSliceModel) SetUnitGroup(v EmbeddedUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetUnit

`func (o *TimeSliceModel) GetUnit() EmbeddedUnitModel`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeSliceModel) GetUnitOk() (*EmbeddedUnitModel, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeSliceModel) SetUnit(v EmbeddedUnitModel)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeSliceModel) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetBaseAmount

`func (o *TimeSliceModel) GetBaseAmount() AmountModel`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *TimeSliceModel) GetBaseAmountOk() (*AmountModel, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *TimeSliceModel) SetBaseAmount(v AmountModel)`

SetBaseAmount sets BaseAmount field to given value.


### GetTotalGrossAmount

`func (o *TimeSliceModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *TimeSliceModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *TimeSliceModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetIncludedServices

`func (o *TimeSliceModel) GetIncludedServices() []ReservationServiceModel`

GetIncludedServices returns the IncludedServices field if non-nil, zero value otherwise.

### GetIncludedServicesOk

`func (o *TimeSliceModel) GetIncludedServicesOk() (*[]ReservationServiceModel, bool)`

GetIncludedServicesOk returns a tuple with the IncludedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedServices

`func (o *TimeSliceModel) SetIncludedServices(v []ReservationServiceModel)`

SetIncludedServices sets IncludedServices field to given value.

### HasIncludedServices

`func (o *TimeSliceModel) HasIncludedServices() bool`

HasIncludedServices returns a boolean if a field has been set.

### GetActions

`func (o *TimeSliceModel) GetActions() []ActionModelReservationTimeSliceActionNotAllowedReservationTimeSliceActionReason`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TimeSliceModel) GetActionsOk() (*[]ActionModelReservationTimeSliceActionNotAllowedReservationTimeSliceActionReason, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TimeSliceModel) SetActions(v []ActionModelReservationTimeSliceActionNotAllowedReservationTimeSliceActionReason)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TimeSliceModel) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


