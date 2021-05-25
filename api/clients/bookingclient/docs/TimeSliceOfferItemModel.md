# TimeSliceOfferItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitGroup** | [**EmbeddedUnitGroupModel**](EmbeddedUnitGroupModel.md) |  | 
**MinGuaranteeType** | Pointer to **string** | The minimum guarantee type for the offer | [optional] 
**MinAdvance** | Pointer to [**PeriodModel**](PeriodModel.md) |  | [optional] 
**MaxAdvance** | Pointer to [**PeriodModel**](PeriodModel.md) |  | [optional] 
**Available** | **int32** | The number of available units for the offer | 
**Restrictions** | Pointer to [**RateRestrictionsModel**](RateRestrictionsModel.md) |  | [optional] 
**Prices** | Pointer to [**[]PerOccupancyPriceItemModel**](PerOccupancyPriceItemModel.md) | The prices for this offer | [optional] 

## Methods

### NewTimeSliceOfferItemModel

`func NewTimeSliceOfferItemModel(unitGroup EmbeddedUnitGroupModel, available int32, ) *TimeSliceOfferItemModel`

NewTimeSliceOfferItemModel instantiates a new TimeSliceOfferItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSliceOfferItemModelWithDefaults

`func NewTimeSliceOfferItemModelWithDefaults() *TimeSliceOfferItemModel`

NewTimeSliceOfferItemModelWithDefaults instantiates a new TimeSliceOfferItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitGroup

`func (o *TimeSliceOfferItemModel) GetUnitGroup() EmbeddedUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *TimeSliceOfferItemModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *TimeSliceOfferItemModel) SetUnitGroup(v EmbeddedUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetMinGuaranteeType

`func (o *TimeSliceOfferItemModel) GetMinGuaranteeType() string`

GetMinGuaranteeType returns the MinGuaranteeType field if non-nil, zero value otherwise.

### GetMinGuaranteeTypeOk

`func (o *TimeSliceOfferItemModel) GetMinGuaranteeTypeOk() (*string, bool)`

GetMinGuaranteeTypeOk returns a tuple with the MinGuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGuaranteeType

`func (o *TimeSliceOfferItemModel) SetMinGuaranteeType(v string)`

SetMinGuaranteeType sets MinGuaranteeType field to given value.

### HasMinGuaranteeType

`func (o *TimeSliceOfferItemModel) HasMinGuaranteeType() bool`

HasMinGuaranteeType returns a boolean if a field has been set.

### GetMinAdvance

`func (o *TimeSliceOfferItemModel) GetMinAdvance() PeriodModel`

GetMinAdvance returns the MinAdvance field if non-nil, zero value otherwise.

### GetMinAdvanceOk

`func (o *TimeSliceOfferItemModel) GetMinAdvanceOk() (*PeriodModel, bool)`

GetMinAdvanceOk returns a tuple with the MinAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAdvance

`func (o *TimeSliceOfferItemModel) SetMinAdvance(v PeriodModel)`

SetMinAdvance sets MinAdvance field to given value.

### HasMinAdvance

`func (o *TimeSliceOfferItemModel) HasMinAdvance() bool`

HasMinAdvance returns a boolean if a field has been set.

### GetMaxAdvance

`func (o *TimeSliceOfferItemModel) GetMaxAdvance() PeriodModel`

GetMaxAdvance returns the MaxAdvance field if non-nil, zero value otherwise.

### GetMaxAdvanceOk

`func (o *TimeSliceOfferItemModel) GetMaxAdvanceOk() (*PeriodModel, bool)`

GetMaxAdvanceOk returns a tuple with the MaxAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAdvance

`func (o *TimeSliceOfferItemModel) SetMaxAdvance(v PeriodModel)`

SetMaxAdvance sets MaxAdvance field to given value.

### HasMaxAdvance

`func (o *TimeSliceOfferItemModel) HasMaxAdvance() bool`

HasMaxAdvance returns a boolean if a field has been set.

### GetAvailable

`func (o *TimeSliceOfferItemModel) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TimeSliceOfferItemModel) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TimeSliceOfferItemModel) SetAvailable(v int32)`

SetAvailable sets Available field to given value.


### GetRestrictions

`func (o *TimeSliceOfferItemModel) GetRestrictions() RateRestrictionsModel`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *TimeSliceOfferItemModel) GetRestrictionsOk() (*RateRestrictionsModel, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *TimeSliceOfferItemModel) SetRestrictions(v RateRestrictionsModel)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *TimeSliceOfferItemModel) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetPrices

`func (o *TimeSliceOfferItemModel) GetPrices() []PerOccupancyPriceItemModel`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *TimeSliceOfferItemModel) GetPricesOk() (*[]PerOccupancyPriceItemModel, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *TimeSliceOfferItemModel) SetPrices(v []PerOccupancyPriceItemModel)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *TimeSliceOfferItemModel) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


