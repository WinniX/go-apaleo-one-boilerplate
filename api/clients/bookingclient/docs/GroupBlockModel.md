# GroupBlockModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Block id | 
**Status** | **string** | Status of the block | 
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**UnitGroup** | [**EmbeddedUnitGroupModel**](EmbeddedUnitGroupModel.md) |  | 
**GrossDailyRate** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**From** | **time.Time** | Start date and time from which the inventory will be blocked&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | End date and time until which the inventory will be blocked&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**BlockedUnits** | **int32** | Number of units blocked | 
**PickedReservations** | **int32** | Number of reservations already picked from this block | 
**Created** | **time.Time** | Date of creation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Modified** | **time.Time** | Date of last modification&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 

## Methods

### NewGroupBlockModel

`func NewGroupBlockModel(id string, status string, property EmbeddedPropertyModel, ratePlan EmbeddedRatePlanModel, unitGroup EmbeddedUnitGroupModel, grossDailyRate MonetaryValueModel, from time.Time, to time.Time, blockedUnits int32, pickedReservations int32, created time.Time, modified time.Time, ) *GroupBlockModel`

NewGroupBlockModel instantiates a new GroupBlockModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupBlockModelWithDefaults

`func NewGroupBlockModelWithDefaults() *GroupBlockModel`

NewGroupBlockModelWithDefaults instantiates a new GroupBlockModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupBlockModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupBlockModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupBlockModel) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *GroupBlockModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupBlockModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupBlockModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProperty

`func (o *GroupBlockModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *GroupBlockModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *GroupBlockModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetRatePlan

`func (o *GroupBlockModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *GroupBlockModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *GroupBlockModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetUnitGroup

`func (o *GroupBlockModel) GetUnitGroup() EmbeddedUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *GroupBlockModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *GroupBlockModel) SetUnitGroup(v EmbeddedUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetGrossDailyRate

`func (o *GroupBlockModel) GetGrossDailyRate() MonetaryValueModel`

GetGrossDailyRate returns the GrossDailyRate field if non-nil, zero value otherwise.

### GetGrossDailyRateOk

`func (o *GroupBlockModel) GetGrossDailyRateOk() (*MonetaryValueModel, bool)`

GetGrossDailyRateOk returns a tuple with the GrossDailyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossDailyRate

`func (o *GroupBlockModel) SetGrossDailyRate(v MonetaryValueModel)`

SetGrossDailyRate sets GrossDailyRate field to given value.


### GetFrom

`func (o *GroupBlockModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GroupBlockModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GroupBlockModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *GroupBlockModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GroupBlockModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GroupBlockModel) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetBlockedUnits

`func (o *GroupBlockModel) GetBlockedUnits() int32`

GetBlockedUnits returns the BlockedUnits field if non-nil, zero value otherwise.

### GetBlockedUnitsOk

`func (o *GroupBlockModel) GetBlockedUnitsOk() (*int32, bool)`

GetBlockedUnitsOk returns a tuple with the BlockedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUnits

`func (o *GroupBlockModel) SetBlockedUnits(v int32)`

SetBlockedUnits sets BlockedUnits field to given value.


### GetPickedReservations

`func (o *GroupBlockModel) GetPickedReservations() int32`

GetPickedReservations returns the PickedReservations field if non-nil, zero value otherwise.

### GetPickedReservationsOk

`func (o *GroupBlockModel) GetPickedReservationsOk() (*int32, bool)`

GetPickedReservationsOk returns a tuple with the PickedReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickedReservations

`func (o *GroupBlockModel) SetPickedReservations(v int32)`

SetPickedReservations sets PickedReservations field to given value.


### GetCreated

`func (o *GroupBlockModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupBlockModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupBlockModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *GroupBlockModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GroupBlockModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GroupBlockModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


