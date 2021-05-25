# OfferTimeSliceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | The start date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | The end date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**AvailableUnits** | **int32** | The number of available units for that time slice | 
**BaseAmount** | [**AmountModel**](AmountModel.md) |  | 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**IncludedServices** | Pointer to [**[]OfferServiceModel**](OfferServiceModel.md) | The breakdown for services included in the offer | [optional] 

## Methods

### NewOfferTimeSliceModel

`func NewOfferTimeSliceModel(from time.Time, to time.Time, availableUnits int32, baseAmount AmountModel, totalGrossAmount MonetaryValueModel, ) *OfferTimeSliceModel`

NewOfferTimeSliceModel instantiates a new OfferTimeSliceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferTimeSliceModelWithDefaults

`func NewOfferTimeSliceModelWithDefaults() *OfferTimeSliceModel`

NewOfferTimeSliceModelWithDefaults instantiates a new OfferTimeSliceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *OfferTimeSliceModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *OfferTimeSliceModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *OfferTimeSliceModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *OfferTimeSliceModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *OfferTimeSliceModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *OfferTimeSliceModel) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetAvailableUnits

`func (o *OfferTimeSliceModel) GetAvailableUnits() int32`

GetAvailableUnits returns the AvailableUnits field if non-nil, zero value otherwise.

### GetAvailableUnitsOk

`func (o *OfferTimeSliceModel) GetAvailableUnitsOk() (*int32, bool)`

GetAvailableUnitsOk returns a tuple with the AvailableUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUnits

`func (o *OfferTimeSliceModel) SetAvailableUnits(v int32)`

SetAvailableUnits sets AvailableUnits field to given value.


### GetBaseAmount

`func (o *OfferTimeSliceModel) GetBaseAmount() AmountModel`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *OfferTimeSliceModel) GetBaseAmountOk() (*AmountModel, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *OfferTimeSliceModel) SetBaseAmount(v AmountModel)`

SetBaseAmount sets BaseAmount field to given value.


### GetTotalGrossAmount

`func (o *OfferTimeSliceModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *OfferTimeSliceModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *OfferTimeSliceModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetIncludedServices

`func (o *OfferTimeSliceModel) GetIncludedServices() []OfferServiceModel`

GetIncludedServices returns the IncludedServices field if non-nil, zero value otherwise.

### GetIncludedServicesOk

`func (o *OfferTimeSliceModel) GetIncludedServicesOk() (*[]OfferServiceModel, bool)`

GetIncludedServicesOk returns a tuple with the IncludedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedServices

`func (o *OfferTimeSliceModel) SetIncludedServices(v []OfferServiceModel)`

SetIncludedServices sets IncludedServices field to given value.

### HasIncludedServices

`func (o *OfferTimeSliceModel) HasIncludedServices() bool`

HasIncludedServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


