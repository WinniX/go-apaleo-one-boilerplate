# ServiceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The service id | 
**Code** | **string** | The code for the service | 
**Name** | **string** | The name for the service | 
**Description** | **string** | The description for the service | 
**PricingUnit** | **string** | Defines the granularity (room, person) for which this item is offered and priced | 
**DefaultGrossPrice** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewServiceModel

`func NewServiceModel(id string, code string, name string, description string, pricingUnit string, defaultGrossPrice MonetaryValueModel, ) *ServiceModel`

NewServiceModel instantiates a new ServiceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceModelWithDefaults

`func NewServiceModelWithDefaults() *ServiceModel`

NewServiceModelWithDefaults instantiates a new ServiceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *ServiceModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *ServiceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPricingUnit

`func (o *ServiceModel) GetPricingUnit() string`

GetPricingUnit returns the PricingUnit field if non-nil, zero value otherwise.

### GetPricingUnitOk

`func (o *ServiceModel) GetPricingUnitOk() (*string, bool)`

GetPricingUnitOk returns a tuple with the PricingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingUnit

`func (o *ServiceModel) SetPricingUnit(v string)`

SetPricingUnit sets PricingUnit field to given value.


### GetDefaultGrossPrice

`func (o *ServiceModel) GetDefaultGrossPrice() MonetaryValueModel`

GetDefaultGrossPrice returns the DefaultGrossPrice field if non-nil, zero value otherwise.

### GetDefaultGrossPriceOk

`func (o *ServiceModel) GetDefaultGrossPriceOk() (*MonetaryValueModel, bool)`

GetDefaultGrossPriceOk returns a tuple with the DefaultGrossPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGrossPrice

`func (o *ServiceModel) SetDefaultGrossPrice(v MonetaryValueModel)`

SetDefaultGrossPrice sets DefaultGrossPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


