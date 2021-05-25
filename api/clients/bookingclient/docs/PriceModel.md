# PriceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrossAmount** | **float64** | Price including all included services and VAT - &lt;b&gt;DEPRECATED: This field will be removed soon, use BeforeTax + Taxes.Tax instead&lt;/b&gt; | 
**BeforeTax** | **float64** | Price including all included services without VAT or any other taxes like city tax | 
**AfterTax** | **float64** | Price including all included services, VAT and any other taxes like city tax | 
**Taxes** | [**TaxesModel**](TaxesModel.md) |  | 
**Currency** | **string** | The currency for all prices and tax details | 

## Methods

### NewPriceModel

`func NewPriceModel(grossAmount float64, beforeTax float64, afterTax float64, taxes TaxesModel, currency string, ) *PriceModel`

NewPriceModel instantiates a new PriceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceModelWithDefaults

`func NewPriceModelWithDefaults() *PriceModel`

NewPriceModelWithDefaults instantiates a new PriceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrossAmount

`func (o *PriceModel) GetGrossAmount() float64`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *PriceModel) GetGrossAmountOk() (*float64, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *PriceModel) SetGrossAmount(v float64)`

SetGrossAmount sets GrossAmount field to given value.


### GetBeforeTax

`func (o *PriceModel) GetBeforeTax() float64`

GetBeforeTax returns the BeforeTax field if non-nil, zero value otherwise.

### GetBeforeTaxOk

`func (o *PriceModel) GetBeforeTaxOk() (*float64, bool)`

GetBeforeTaxOk returns a tuple with the BeforeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeTax

`func (o *PriceModel) SetBeforeTax(v float64)`

SetBeforeTax sets BeforeTax field to given value.


### GetAfterTax

`func (o *PriceModel) GetAfterTax() float64`

GetAfterTax returns the AfterTax field if non-nil, zero value otherwise.

### GetAfterTaxOk

`func (o *PriceModel) GetAfterTaxOk() (*float64, bool)`

GetAfterTaxOk returns a tuple with the AfterTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterTax

`func (o *PriceModel) SetAfterTax(v float64)`

SetAfterTax sets AfterTax field to given value.


### GetTaxes

`func (o *PriceModel) GetTaxes() TaxesModel`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *PriceModel) GetTaxesOk() (*TaxesModel, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *PriceModel) SetTaxes(v TaxesModel)`

SetTaxes sets Taxes field to given value.


### GetCurrency

`func (o *PriceModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


