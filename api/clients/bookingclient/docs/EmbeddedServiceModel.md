# EmbeddedServiceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The service id | 
**Code** | Pointer to **string** | The code for the service | [optional] 
**Name** | Pointer to **string** | The name for the service | [optional] 
**Description** | Pointer to **string** | The description for the service | [optional] 

## Methods

### NewEmbeddedServiceModel

`func NewEmbeddedServiceModel(id string, ) *EmbeddedServiceModel`

NewEmbeddedServiceModel instantiates a new EmbeddedServiceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedServiceModelWithDefaults

`func NewEmbeddedServiceModelWithDefaults() *EmbeddedServiceModel`

NewEmbeddedServiceModelWithDefaults instantiates a new EmbeddedServiceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmbeddedServiceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmbeddedServiceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmbeddedServiceModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *EmbeddedServiceModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EmbeddedServiceModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EmbeddedServiceModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EmbeddedServiceModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *EmbeddedServiceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmbeddedServiceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmbeddedServiceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmbeddedServiceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EmbeddedServiceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmbeddedServiceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmbeddedServiceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmbeddedServiceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


