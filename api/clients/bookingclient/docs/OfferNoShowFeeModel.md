# OfferNoShowFeeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code of the no-show policy applied | 
**Name** | **string** | The name of the no-show policy applied | 
**Description** | **string** | The description of the no-show policy applied | 
**Fee** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewOfferNoShowFeeModel

`func NewOfferNoShowFeeModel(code string, name string, description string, fee MonetaryValueModel, ) *OfferNoShowFeeModel`

NewOfferNoShowFeeModel instantiates a new OfferNoShowFeeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferNoShowFeeModelWithDefaults

`func NewOfferNoShowFeeModelWithDefaults() *OfferNoShowFeeModel`

NewOfferNoShowFeeModelWithDefaults instantiates a new OfferNoShowFeeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OfferNoShowFeeModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OfferNoShowFeeModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OfferNoShowFeeModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *OfferNoShowFeeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfferNoShowFeeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfferNoShowFeeModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OfferNoShowFeeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OfferNoShowFeeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OfferNoShowFeeModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFee

`func (o *OfferNoShowFeeModel) GetFee() MonetaryValueModel`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *OfferNoShowFeeModel) GetFeeOk() (*MonetaryValueModel, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *OfferNoShowFeeModel) SetFee(v MonetaryValueModel)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


