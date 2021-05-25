# TaxesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tax** | **float64** | The amount of taxes, which are VAT or Sales Taxes depending on the country of the property | 
**CityTax** | **float64** | The amount of City Tax including VAT | 

## Methods

### NewTaxesModel

`func NewTaxesModel(tax float64, cityTax float64, ) *TaxesModel`

NewTaxesModel instantiates a new TaxesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxesModelWithDefaults

`func NewTaxesModelWithDefaults() *TaxesModel`

NewTaxesModelWithDefaults instantiates a new TaxesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTax

`func (o *TaxesModel) GetTax() float64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *TaxesModel) GetTaxOk() (*float64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *TaxesModel) SetTax(v float64)`

SetTax sets Tax field to given value.


### GetCityTax

`func (o *TaxesModel) GetCityTax() float64`

GetCityTax returns the CityTax field if non-nil, zero value otherwise.

### GetCityTaxOk

`func (o *TaxesModel) GetCityTaxOk() (*float64, bool)`

GetCityTaxOk returns a tuple with the CityTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTax

`func (o *TaxesModel) SetCityTax(v float64)`

SetCityTax sets CityTax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


