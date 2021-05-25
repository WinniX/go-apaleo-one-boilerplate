# TaxDetailModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VatType** | **string** |  | 
**VatPercent** | **float64** |  | 
**Net** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**Tax** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewTaxDetailModel

`func NewTaxDetailModel(vatType string, vatPercent float64, net MonetaryValueModel, tax MonetaryValueModel, ) *TaxDetailModel`

NewTaxDetailModel instantiates a new TaxDetailModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxDetailModelWithDefaults

`func NewTaxDetailModelWithDefaults() *TaxDetailModel`

NewTaxDetailModelWithDefaults instantiates a new TaxDetailModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVatType

`func (o *TaxDetailModel) GetVatType() string`

GetVatType returns the VatType field if non-nil, zero value otherwise.

### GetVatTypeOk

`func (o *TaxDetailModel) GetVatTypeOk() (*string, bool)`

GetVatTypeOk returns a tuple with the VatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatType

`func (o *TaxDetailModel) SetVatType(v string)`

SetVatType sets VatType field to given value.


### GetVatPercent

`func (o *TaxDetailModel) GetVatPercent() float64`

GetVatPercent returns the VatPercent field if non-nil, zero value otherwise.

### GetVatPercentOk

`func (o *TaxDetailModel) GetVatPercentOk() (*float64, bool)`

GetVatPercentOk returns a tuple with the VatPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercent

`func (o *TaxDetailModel) SetVatPercent(v float64)`

SetVatPercent sets VatPercent field to given value.


### GetNet

`func (o *TaxDetailModel) GetNet() MonetaryValueModel`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *TaxDetailModel) GetNetOk() (*MonetaryValueModel, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *TaxDetailModel) SetNet(v MonetaryValueModel)`

SetNet sets Net field to given value.


### GetTax

`func (o *TaxDetailModel) GetTax() MonetaryValueModel`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *TaxDetailModel) GetTaxOk() (*MonetaryValueModel, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *TaxDetailModel) SetTax(v MonetaryValueModel)`

SetTax sets Tax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


