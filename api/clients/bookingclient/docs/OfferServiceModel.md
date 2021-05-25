# OfferServiceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | [**EmbeddedServiceModel**](EmbeddedServiceModel.md) |  | 
**ServiceDate** | **string** | The date this service is delivered | 
**Count** | **int32** | The default count of offered services. For services whose pricing unit is &#39;Person&#39; it will be based on the adults and children specified, otherwise 1. | 
**Amount** | [**AmountModel**](AmountModel.md) |  | 
**PricingMode** | **string** | Whether the service price is included in or added to the base rate | 

## Methods

### NewOfferServiceModel

`func NewOfferServiceModel(service EmbeddedServiceModel, serviceDate string, count int32, amount AmountModel, pricingMode string, ) *OfferServiceModel`

NewOfferServiceModel instantiates a new OfferServiceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferServiceModelWithDefaults

`func NewOfferServiceModelWithDefaults() *OfferServiceModel`

NewOfferServiceModelWithDefaults instantiates a new OfferServiceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *OfferServiceModel) GetService() EmbeddedServiceModel`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *OfferServiceModel) GetServiceOk() (*EmbeddedServiceModel, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *OfferServiceModel) SetService(v EmbeddedServiceModel)`

SetService sets Service field to given value.


### GetServiceDate

`func (o *OfferServiceModel) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *OfferServiceModel) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *OfferServiceModel) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetCount

`func (o *OfferServiceModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OfferServiceModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OfferServiceModel) SetCount(v int32)`

SetCount sets Count field to given value.


### GetAmount

`func (o *OfferServiceModel) GetAmount() AmountModel`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OfferServiceModel) GetAmountOk() (*AmountModel, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OfferServiceModel) SetAmount(v AmountModel)`

SetAmount sets Amount field to given value.


### GetPricingMode

`func (o *OfferServiceModel) GetPricingMode() string`

GetPricingMode returns the PricingMode field if non-nil, zero value otherwise.

### GetPricingModeOk

`func (o *OfferServiceModel) GetPricingModeOk() (*string, bool)`

GetPricingModeOk returns a tuple with the PricingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingMode

`func (o *OfferServiceModel) SetPricingMode(v string)`

SetPricingMode sets PricingMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


