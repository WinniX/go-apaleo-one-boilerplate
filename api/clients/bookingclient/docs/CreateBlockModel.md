# CreateBlockModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | ID of the group that reserved the block | 
**RatePlanId** | **string** | The rate plan | 
**From** | **string** | Start date and time from which the inventory will be blocked&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **string** | End date and time until which the inventory will be blocked. Cannot be more than 5 years after the start date.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**GrossDailyRate** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**TimeSlices** | Pointer to [**[]CreateBlockTimeSliceModel**](CreateBlockTimeSliceModel.md) | The list of blocked units for each time slice | [optional] 
**BlockedUnits** | Pointer to **int32** | Number of units to block for the defined time period | [optional] 
**PromoCode** | Pointer to **string** | The promo code associated with a certain special offer | [optional] 
**CorporateCode** | Pointer to **string** | The corporate code associated with a certain special offer | [optional] 

## Methods

### NewCreateBlockModel

`func NewCreateBlockModel(groupId string, ratePlanId string, from string, to string, grossDailyRate MonetaryValueModel, ) *CreateBlockModel`

NewCreateBlockModel instantiates a new CreateBlockModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBlockModelWithDefaults

`func NewCreateBlockModelWithDefaults() *CreateBlockModel`

NewCreateBlockModelWithDefaults instantiates a new CreateBlockModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CreateBlockModel) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateBlockModel) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateBlockModel) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetRatePlanId

`func (o *CreateBlockModel) GetRatePlanId() string`

GetRatePlanId returns the RatePlanId field if non-nil, zero value otherwise.

### GetRatePlanIdOk

`func (o *CreateBlockModel) GetRatePlanIdOk() (*string, bool)`

GetRatePlanIdOk returns a tuple with the RatePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlanId

`func (o *CreateBlockModel) SetRatePlanId(v string)`

SetRatePlanId sets RatePlanId field to given value.


### GetFrom

`func (o *CreateBlockModel) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateBlockModel) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateBlockModel) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CreateBlockModel) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateBlockModel) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateBlockModel) SetTo(v string)`

SetTo sets To field to given value.


### GetGrossDailyRate

`func (o *CreateBlockModel) GetGrossDailyRate() MonetaryValueModel`

GetGrossDailyRate returns the GrossDailyRate field if non-nil, zero value otherwise.

### GetGrossDailyRateOk

`func (o *CreateBlockModel) GetGrossDailyRateOk() (*MonetaryValueModel, bool)`

GetGrossDailyRateOk returns a tuple with the GrossDailyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossDailyRate

`func (o *CreateBlockModel) SetGrossDailyRate(v MonetaryValueModel)`

SetGrossDailyRate sets GrossDailyRate field to given value.


### GetTimeSlices

`func (o *CreateBlockModel) GetTimeSlices() []CreateBlockTimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *CreateBlockModel) GetTimeSlicesOk() (*[]CreateBlockTimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *CreateBlockModel) SetTimeSlices(v []CreateBlockTimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.

### HasTimeSlices

`func (o *CreateBlockModel) HasTimeSlices() bool`

HasTimeSlices returns a boolean if a field has been set.

### GetBlockedUnits

`func (o *CreateBlockModel) GetBlockedUnits() int32`

GetBlockedUnits returns the BlockedUnits field if non-nil, zero value otherwise.

### GetBlockedUnitsOk

`func (o *CreateBlockModel) GetBlockedUnitsOk() (*int32, bool)`

GetBlockedUnitsOk returns a tuple with the BlockedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUnits

`func (o *CreateBlockModel) SetBlockedUnits(v int32)`

SetBlockedUnits sets BlockedUnits field to given value.

### HasBlockedUnits

`func (o *CreateBlockModel) HasBlockedUnits() bool`

HasBlockedUnits returns a boolean if a field has been set.

### GetPromoCode

`func (o *CreateBlockModel) GetPromoCode() string`

GetPromoCode returns the PromoCode field if non-nil, zero value otherwise.

### GetPromoCodeOk

`func (o *CreateBlockModel) GetPromoCodeOk() (*string, bool)`

GetPromoCodeOk returns a tuple with the PromoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCode

`func (o *CreateBlockModel) SetPromoCode(v string)`

SetPromoCode sets PromoCode field to given value.

### HasPromoCode

`func (o *CreateBlockModel) HasPromoCode() bool`

HasPromoCode returns a boolean if a field has been set.

### GetCorporateCode

`func (o *CreateBlockModel) GetCorporateCode() string`

GetCorporateCode returns the CorporateCode field if non-nil, zero value otherwise.

### GetCorporateCodeOk

`func (o *CreateBlockModel) GetCorporateCodeOk() (*string, bool)`

GetCorporateCodeOk returns a tuple with the CorporateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCode

`func (o *CreateBlockModel) SetCorporateCode(v string)`

SetCorporateCode sets CorporateCode field to given value.

### HasCorporateCode

`func (o *CreateBlockModel) HasCorporateCode() bool`

HasCorporateCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


