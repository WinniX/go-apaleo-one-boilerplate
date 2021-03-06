# EmbeddedRatePlanModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The rate plan id | 
**Code** | Pointer to **string** | The code for the rate plan that can be shown in reports and table views | [optional] 
**Name** | Pointer to **string** | The name for the rate plan | [optional] 
**Description** | Pointer to **string** | The description for the rate plan | [optional] 
**IsSubjectToCityTax** | **bool** | Whether the rate plan is subject to city tax or not | 

## Methods

### NewEmbeddedRatePlanModel

`func NewEmbeddedRatePlanModel(id string, isSubjectToCityTax bool, ) *EmbeddedRatePlanModel`

NewEmbeddedRatePlanModel instantiates a new EmbeddedRatePlanModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedRatePlanModelWithDefaults

`func NewEmbeddedRatePlanModelWithDefaults() *EmbeddedRatePlanModel`

NewEmbeddedRatePlanModelWithDefaults instantiates a new EmbeddedRatePlanModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmbeddedRatePlanModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmbeddedRatePlanModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmbeddedRatePlanModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *EmbeddedRatePlanModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EmbeddedRatePlanModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EmbeddedRatePlanModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EmbeddedRatePlanModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *EmbeddedRatePlanModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmbeddedRatePlanModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmbeddedRatePlanModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmbeddedRatePlanModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EmbeddedRatePlanModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmbeddedRatePlanModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmbeddedRatePlanModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmbeddedRatePlanModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsSubjectToCityTax

`func (o *EmbeddedRatePlanModel) GetIsSubjectToCityTax() bool`

GetIsSubjectToCityTax returns the IsSubjectToCityTax field if non-nil, zero value otherwise.

### GetIsSubjectToCityTaxOk

`func (o *EmbeddedRatePlanModel) GetIsSubjectToCityTaxOk() (*bool, bool)`

GetIsSubjectToCityTaxOk returns a tuple with the IsSubjectToCityTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubjectToCityTax

`func (o *EmbeddedRatePlanModel) SetIsSubjectToCityTax(v bool)`

SetIsSubjectToCityTax sets IsSubjectToCityTax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


