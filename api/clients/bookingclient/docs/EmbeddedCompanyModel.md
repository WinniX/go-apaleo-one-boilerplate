# EmbeddedCompanyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The company ID | 
**Code** | Pointer to **string** | The code of the company | [optional] 
**Name** | Pointer to **string** | The name of the company | [optional] 
**CanCheckOutOnAr** | Pointer to **bool** | Whether or not the company can check out on AR | [optional] 

## Methods

### NewEmbeddedCompanyModel

`func NewEmbeddedCompanyModel(id string, ) *EmbeddedCompanyModel`

NewEmbeddedCompanyModel instantiates a new EmbeddedCompanyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedCompanyModelWithDefaults

`func NewEmbeddedCompanyModelWithDefaults() *EmbeddedCompanyModel`

NewEmbeddedCompanyModelWithDefaults instantiates a new EmbeddedCompanyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmbeddedCompanyModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmbeddedCompanyModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmbeddedCompanyModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *EmbeddedCompanyModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EmbeddedCompanyModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EmbeddedCompanyModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EmbeddedCompanyModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *EmbeddedCompanyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmbeddedCompanyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmbeddedCompanyModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmbeddedCompanyModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCanCheckOutOnAr

`func (o *EmbeddedCompanyModel) GetCanCheckOutOnAr() bool`

GetCanCheckOutOnAr returns the CanCheckOutOnAr field if non-nil, zero value otherwise.

### GetCanCheckOutOnArOk

`func (o *EmbeddedCompanyModel) GetCanCheckOutOnArOk() (*bool, bool)`

GetCanCheckOutOnArOk returns a tuple with the CanCheckOutOnAr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCheckOutOnAr

`func (o *EmbeddedCompanyModel) SetCanCheckOutOnAr(v bool)`

SetCanCheckOutOnAr sets CanCheckOutOnAr field to given value.

### HasCanCheckOutOnAr

`func (o *EmbeddedCompanyModel) HasCanCheckOutOnAr() bool`

HasCanCheckOutOnAr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


