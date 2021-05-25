# ServiceDateItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceDate** | **string** | The date this service is delivered | 
**Count** | **int32** | The count of booked services | 
**Amount** | [**AmountModel**](AmountModel.md) |  | 
**IsMandatory** | **bool** | Rate plans can have additional services. When booking an offer for such rate plans, those services are automatically booked.  They are marked as mandatory and they cannot be removed. | 

## Methods

### NewServiceDateItemModel

`func NewServiceDateItemModel(serviceDate string, count int32, amount AmountModel, isMandatory bool, ) *ServiceDateItemModel`

NewServiceDateItemModel instantiates a new ServiceDateItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDateItemModelWithDefaults

`func NewServiceDateItemModelWithDefaults() *ServiceDateItemModel`

NewServiceDateItemModelWithDefaults instantiates a new ServiceDateItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceDate

`func (o *ServiceDateItemModel) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *ServiceDateItemModel) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *ServiceDateItemModel) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetCount

`func (o *ServiceDateItemModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ServiceDateItemModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ServiceDateItemModel) SetCount(v int32)`

SetCount sets Count field to given value.


### GetAmount

`func (o *ServiceDateItemModel) GetAmount() AmountModel`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServiceDateItemModel) GetAmountOk() (*AmountModel, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServiceDateItemModel) SetAmount(v AmountModel)`

SetAmount sets Amount field to given value.


### GetIsMandatory

`func (o *ServiceDateItemModel) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *ServiceDateItemModel) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *ServiceDateItemModel) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


