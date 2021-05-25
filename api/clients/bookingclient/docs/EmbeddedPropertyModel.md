# EmbeddedPropertyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The property id | 
**Code** | Pointer to **string** | The code for the property that can be shown in reports and table views | [optional] 
**Name** | Pointer to **string** | The name for the property | [optional] 
**Description** | Pointer to **string** | The description for the property | [optional] [readonly] 

## Methods

### NewEmbeddedPropertyModel

`func NewEmbeddedPropertyModel(id string, ) *EmbeddedPropertyModel`

NewEmbeddedPropertyModel instantiates a new EmbeddedPropertyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedPropertyModelWithDefaults

`func NewEmbeddedPropertyModelWithDefaults() *EmbeddedPropertyModel`

NewEmbeddedPropertyModelWithDefaults instantiates a new EmbeddedPropertyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmbeddedPropertyModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmbeddedPropertyModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmbeddedPropertyModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *EmbeddedPropertyModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EmbeddedPropertyModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EmbeddedPropertyModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EmbeddedPropertyModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *EmbeddedPropertyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmbeddedPropertyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmbeddedPropertyModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmbeddedPropertyModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EmbeddedPropertyModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmbeddedPropertyModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmbeddedPropertyModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmbeddedPropertyModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


