# BlockItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Block id | 
**Group** | [**EmbeddedGroupModel**](EmbeddedGroupModel.md) |  | 
**Status** | **string** | Status of the block | 
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**UnitGroup** | [**EmbeddedUnitGroupModel**](EmbeddedUnitGroupModel.md) |  | 
**GrossDailyRate** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**From** | **time.Time** | Start date and time from which the inventory will be blocked&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | End date and time until which the inventory will be blocked&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**PickedReservations** | **int32** | Number of reservations already picked from this block | 
**PromoCode** | Pointer to **string** | The promo code associated with a certain special offer used to create the block | [optional] 
**CorporateCode** | Pointer to **string** | The corporate code associated with a certain special offer used to create the block | [optional] 
**Created** | **time.Time** | Date of creation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Modified** | **time.Time** | Date of last modification&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**TimeSlices** | Pointer to [**[]BlockTimeSliceModel**](BlockTimeSliceModel.md) | The list of blocked units for each time slice | [optional] 
**Actions** | Pointer to [**[]ActionModelBlockActionNotAllowedBlockActionReason**](ActionModelBlockActionNotAllowedBlockActionReason.md) | The list of actions for this block | [optional] 

## Methods

### NewBlockItemModel

`func NewBlockItemModel(id string, group EmbeddedGroupModel, status string, property EmbeddedPropertyModel, ratePlan EmbeddedRatePlanModel, unitGroup EmbeddedUnitGroupModel, grossDailyRate MonetaryValueModel, from time.Time, to time.Time, pickedReservations int32, created time.Time, modified time.Time, ) *BlockItemModel`

NewBlockItemModel instantiates a new BlockItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockItemModelWithDefaults

`func NewBlockItemModelWithDefaults() *BlockItemModel`

NewBlockItemModelWithDefaults instantiates a new BlockItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlockItemModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockItemModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockItemModel) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *BlockItemModel) GetGroup() EmbeddedGroupModel`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BlockItemModel) GetGroupOk() (*EmbeddedGroupModel, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BlockItemModel) SetGroup(v EmbeddedGroupModel)`

SetGroup sets Group field to given value.


### GetStatus

`func (o *BlockItemModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlockItemModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlockItemModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProperty

`func (o *BlockItemModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *BlockItemModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *BlockItemModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetRatePlan

`func (o *BlockItemModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *BlockItemModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *BlockItemModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetUnitGroup

`func (o *BlockItemModel) GetUnitGroup() EmbeddedUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *BlockItemModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *BlockItemModel) SetUnitGroup(v EmbeddedUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetGrossDailyRate

`func (o *BlockItemModel) GetGrossDailyRate() MonetaryValueModel`

GetGrossDailyRate returns the GrossDailyRate field if non-nil, zero value otherwise.

### GetGrossDailyRateOk

`func (o *BlockItemModel) GetGrossDailyRateOk() (*MonetaryValueModel, bool)`

GetGrossDailyRateOk returns a tuple with the GrossDailyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossDailyRate

`func (o *BlockItemModel) SetGrossDailyRate(v MonetaryValueModel)`

SetGrossDailyRate sets GrossDailyRate field to given value.


### GetFrom

`func (o *BlockItemModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BlockItemModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BlockItemModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *BlockItemModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BlockItemModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BlockItemModel) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetPickedReservations

`func (o *BlockItemModel) GetPickedReservations() int32`

GetPickedReservations returns the PickedReservations field if non-nil, zero value otherwise.

### GetPickedReservationsOk

`func (o *BlockItemModel) GetPickedReservationsOk() (*int32, bool)`

GetPickedReservationsOk returns a tuple with the PickedReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickedReservations

`func (o *BlockItemModel) SetPickedReservations(v int32)`

SetPickedReservations sets PickedReservations field to given value.


### GetPromoCode

`func (o *BlockItemModel) GetPromoCode() string`

GetPromoCode returns the PromoCode field if non-nil, zero value otherwise.

### GetPromoCodeOk

`func (o *BlockItemModel) GetPromoCodeOk() (*string, bool)`

GetPromoCodeOk returns a tuple with the PromoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCode

`func (o *BlockItemModel) SetPromoCode(v string)`

SetPromoCode sets PromoCode field to given value.

### HasPromoCode

`func (o *BlockItemModel) HasPromoCode() bool`

HasPromoCode returns a boolean if a field has been set.

### GetCorporateCode

`func (o *BlockItemModel) GetCorporateCode() string`

GetCorporateCode returns the CorporateCode field if non-nil, zero value otherwise.

### GetCorporateCodeOk

`func (o *BlockItemModel) GetCorporateCodeOk() (*string, bool)`

GetCorporateCodeOk returns a tuple with the CorporateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCode

`func (o *BlockItemModel) SetCorporateCode(v string)`

SetCorporateCode sets CorporateCode field to given value.

### HasCorporateCode

`func (o *BlockItemModel) HasCorporateCode() bool`

HasCorporateCode returns a boolean if a field has been set.

### GetCreated

`func (o *BlockItemModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BlockItemModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BlockItemModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *BlockItemModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BlockItemModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BlockItemModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetTimeSlices

`func (o *BlockItemModel) GetTimeSlices() []BlockTimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *BlockItemModel) GetTimeSlicesOk() (*[]BlockTimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *BlockItemModel) SetTimeSlices(v []BlockTimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.

### HasTimeSlices

`func (o *BlockItemModel) HasTimeSlices() bool`

HasTimeSlices returns a boolean if a field has been set.

### GetActions

`func (o *BlockItemModel) GetActions() []ActionModelBlockActionNotAllowedBlockActionReason`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BlockItemModel) GetActionsOk() (*[]ActionModelBlockActionNotAllowedBlockActionReason, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BlockItemModel) SetActions(v []ActionModelBlockActionNotAllowedBlockActionReason)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BlockItemModel) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


