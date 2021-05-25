# BlockTimeSliceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | Start date and time from which units will be blocked&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | End date and time until which units will be blocked&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**BlockedUnits** | **int32** | Number of units blocked for this time slice | 
**PickedUnits** | **int32** | Number of units which have picked reservations for this time slice | 
**BaseAmount** | [**AmountModel**](AmountModel.md) |  | 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewBlockTimeSliceModel

`func NewBlockTimeSliceModel(from time.Time, to time.Time, blockedUnits int32, pickedUnits int32, baseAmount AmountModel, totalGrossAmount MonetaryValueModel, ) *BlockTimeSliceModel`

NewBlockTimeSliceModel instantiates a new BlockTimeSliceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockTimeSliceModelWithDefaults

`func NewBlockTimeSliceModelWithDefaults() *BlockTimeSliceModel`

NewBlockTimeSliceModelWithDefaults instantiates a new BlockTimeSliceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BlockTimeSliceModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BlockTimeSliceModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BlockTimeSliceModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *BlockTimeSliceModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BlockTimeSliceModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BlockTimeSliceModel) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetBlockedUnits

`func (o *BlockTimeSliceModel) GetBlockedUnits() int32`

GetBlockedUnits returns the BlockedUnits field if non-nil, zero value otherwise.

### GetBlockedUnitsOk

`func (o *BlockTimeSliceModel) GetBlockedUnitsOk() (*int32, bool)`

GetBlockedUnitsOk returns a tuple with the BlockedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUnits

`func (o *BlockTimeSliceModel) SetBlockedUnits(v int32)`

SetBlockedUnits sets BlockedUnits field to given value.


### GetPickedUnits

`func (o *BlockTimeSliceModel) GetPickedUnits() int32`

GetPickedUnits returns the PickedUnits field if non-nil, zero value otherwise.

### GetPickedUnitsOk

`func (o *BlockTimeSliceModel) GetPickedUnitsOk() (*int32, bool)`

GetPickedUnitsOk returns a tuple with the PickedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickedUnits

`func (o *BlockTimeSliceModel) SetPickedUnits(v int32)`

SetPickedUnits sets PickedUnits field to given value.


### GetBaseAmount

`func (o *BlockTimeSliceModel) GetBaseAmount() AmountModel`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *BlockTimeSliceModel) GetBaseAmountOk() (*AmountModel, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *BlockTimeSliceModel) SetBaseAmount(v AmountModel)`

SetBaseAmount sets BaseAmount field to given value.


### GetTotalGrossAmount

`func (o *BlockTimeSliceModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *BlockTimeSliceModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *BlockTimeSliceModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


