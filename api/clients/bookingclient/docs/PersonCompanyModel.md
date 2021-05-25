# PersonCompanyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the company | [optional] 
**TaxId** | Pointer to **string** | Tax or Vat ID of the company | [optional] 

## Methods

### NewPersonCompanyModel

`func NewPersonCompanyModel() *PersonCompanyModel`

NewPersonCompanyModel instantiates a new PersonCompanyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonCompanyModelWithDefaults

`func NewPersonCompanyModelWithDefaults() *PersonCompanyModel`

NewPersonCompanyModelWithDefaults instantiates a new PersonCompanyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PersonCompanyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonCompanyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonCompanyModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersonCompanyModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxId

`func (o *PersonCompanyModel) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *PersonCompanyModel) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *PersonCompanyModel) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *PersonCompanyModel) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


