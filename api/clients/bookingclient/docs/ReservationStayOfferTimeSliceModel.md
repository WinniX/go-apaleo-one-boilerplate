# ReservationStayOfferTimeSliceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | The start date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | The end date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**UnitGroup** | [**OfferUnitGroupModel**](OfferUnitGroupModel.md) |  | 
**BaseAmount** | [**AmountModel**](AmountModel.md) |  | 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**IncludedServices** | Pointer to [**[]ReservationStayOfferServiceModel**](ReservationStayOfferServiceModel.md) | The breakdown for services included in the offer | [optional] 

## Methods

### NewReservationStayOfferTimeSliceModel

`func NewReservationStayOfferTimeSliceModel(from time.Time, to time.Time, ratePlan EmbeddedRatePlanModel, unitGroup OfferUnitGroupModel, baseAmount AmountModel, totalGrossAmount MonetaryValueModel, ) *ReservationStayOfferTimeSliceModel`

NewReservationStayOfferTimeSliceModel instantiates a new ReservationStayOfferTimeSliceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationStayOfferTimeSliceModelWithDefaults

`func NewReservationStayOfferTimeSliceModelWithDefaults() *ReservationStayOfferTimeSliceModel`

NewReservationStayOfferTimeSliceModelWithDefaults instantiates a new ReservationStayOfferTimeSliceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *ReservationStayOfferTimeSliceModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReservationStayOfferTimeSliceModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReservationStayOfferTimeSliceModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *ReservationStayOfferTimeSliceModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReservationStayOfferTimeSliceModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReservationStayOfferTimeSliceModel) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetRatePlan

`func (o *ReservationStayOfferTimeSliceModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *ReservationStayOfferTimeSliceModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *ReservationStayOfferTimeSliceModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetUnitGroup

`func (o *ReservationStayOfferTimeSliceModel) GetUnitGroup() OfferUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *ReservationStayOfferTimeSliceModel) GetUnitGroupOk() (*OfferUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *ReservationStayOfferTimeSliceModel) SetUnitGroup(v OfferUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetBaseAmount

`func (o *ReservationStayOfferTimeSliceModel) GetBaseAmount() AmountModel`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *ReservationStayOfferTimeSliceModel) GetBaseAmountOk() (*AmountModel, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *ReservationStayOfferTimeSliceModel) SetBaseAmount(v AmountModel)`

SetBaseAmount sets BaseAmount field to given value.


### GetTotalGrossAmount

`func (o *ReservationStayOfferTimeSliceModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *ReservationStayOfferTimeSliceModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *ReservationStayOfferTimeSliceModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetIncludedServices

`func (o *ReservationStayOfferTimeSliceModel) GetIncludedServices() []ReservationStayOfferServiceModel`

GetIncludedServices returns the IncludedServices field if non-nil, zero value otherwise.

### GetIncludedServicesOk

`func (o *ReservationStayOfferTimeSliceModel) GetIncludedServicesOk() (*[]ReservationStayOfferServiceModel, bool)`

GetIncludedServicesOk returns a tuple with the IncludedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedServices

`func (o *ReservationStayOfferTimeSliceModel) SetIncludedServices(v []ReservationStayOfferServiceModel)`

SetIncludedServices sets IncludedServices field to given value.

### HasIncludedServices

`func (o *ReservationStayOfferTimeSliceModel) HasIncludedServices() bool`

HasIncludedServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


