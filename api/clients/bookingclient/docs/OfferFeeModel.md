# OfferFeeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The fee id | 
**Code** | **string** | The code for the fee | 
**Name** | **string** | The name for the fee | 
**TotalAmount** | [**AmountModel**](AmountModel.md) |  | 

## Methods

### NewOfferFeeModel

`func NewOfferFeeModel(id string, code string, name string, totalAmount AmountModel, ) *OfferFeeModel`

NewOfferFeeModel instantiates a new OfferFeeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferFeeModelWithDefaults

`func NewOfferFeeModelWithDefaults() *OfferFeeModel`

NewOfferFeeModelWithDefaults instantiates a new OfferFeeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OfferFeeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfferFeeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfferFeeModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *OfferFeeModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OfferFeeModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OfferFeeModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *OfferFeeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfferFeeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfferFeeModel) SetName(v string)`

SetName sets Name field to given value.


### GetTotalAmount

`func (o *OfferFeeModel) GetTotalAmount() AmountModel`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *OfferFeeModel) GetTotalAmountOk() (*AmountModel, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *OfferFeeModel) SetTotalAmount(v AmountModel)`

SetTotalAmount sets TotalAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


