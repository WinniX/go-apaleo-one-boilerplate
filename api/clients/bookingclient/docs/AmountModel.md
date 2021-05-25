# AmountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrossAmount** | **float64** |  | 
**NetAmount** | **float64** |  | 
**VatType** | **string** |  | 
**VatPercent** | **float64** |  | 
**Currency** | **string** |  | 

## Methods

### NewAmountModel

`func NewAmountModel(grossAmount float64, netAmount float64, vatType string, vatPercent float64, currency string, ) *AmountModel`

NewAmountModel instantiates a new AmountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountModelWithDefaults

`func NewAmountModelWithDefaults() *AmountModel`

NewAmountModelWithDefaults instantiates a new AmountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrossAmount

`func (o *AmountModel) GetGrossAmount() float64`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *AmountModel) GetGrossAmountOk() (*float64, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *AmountModel) SetGrossAmount(v float64)`

SetGrossAmount sets GrossAmount field to given value.


### GetNetAmount

`func (o *AmountModel) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *AmountModel) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *AmountModel) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.


### GetVatType

`func (o *AmountModel) GetVatType() string`

GetVatType returns the VatType field if non-nil, zero value otherwise.

### GetVatTypeOk

`func (o *AmountModel) GetVatTypeOk() (*string, bool)`

GetVatTypeOk returns a tuple with the VatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatType

`func (o *AmountModel) SetVatType(v string)`

SetVatType sets VatType field to given value.


### GetVatPercent

`func (o *AmountModel) GetVatPercent() float64`

GetVatPercent returns the VatPercent field if non-nil, zero value otherwise.

### GetVatPercentOk

`func (o *AmountModel) GetVatPercentOk() (*float64, bool)`

GetVatPercentOk returns a tuple with the VatPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercent

`func (o *AmountModel) SetVatPercent(v float64)`

SetVatPercent sets VatPercent field to given value.


### GetCurrency

`func (o *AmountModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AmountModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AmountModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


