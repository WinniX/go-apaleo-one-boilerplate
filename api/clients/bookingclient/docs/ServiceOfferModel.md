# ServiceOfferModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | [**ServiceModel**](ServiceModel.md) |  | 
**Count** | **int32** | The default count of offered services. For services whose pricing unit is &#39;Person&#39; it will be based on the adults and children specified, otherwise 1. | 
**TotalAmount** | [**AmountModel**](AmountModel.md) |  | 
**PrePaymentAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**Fees** | Pointer to [**[]OfferFeeModel**](OfferFeeModel.md) | The details of the fees that will be added on top of the Apaleo.Api.Modules.Booking.Models.Offer.ServiceOffer.ServiceOfferModel.TotalAmount when booking the service | [optional] 
**Dates** | [**[]ServiceOfferItemModel**](ServiceOfferItemModel.md) | The dates the service will be delivered with its price | 

## Methods

### NewServiceOfferModel

`func NewServiceOfferModel(service ServiceModel, count int32, totalAmount AmountModel, prePaymentAmount MonetaryValueModel, dates []ServiceOfferItemModel, ) *ServiceOfferModel`

NewServiceOfferModel instantiates a new ServiceOfferModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOfferModelWithDefaults

`func NewServiceOfferModelWithDefaults() *ServiceOfferModel`

NewServiceOfferModelWithDefaults instantiates a new ServiceOfferModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceOfferModel) GetService() ServiceModel`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceOfferModel) GetServiceOk() (*ServiceModel, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceOfferModel) SetService(v ServiceModel)`

SetService sets Service field to given value.


### GetCount

`func (o *ServiceOfferModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ServiceOfferModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ServiceOfferModel) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotalAmount

`func (o *ServiceOfferModel) GetTotalAmount() AmountModel`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ServiceOfferModel) GetTotalAmountOk() (*AmountModel, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ServiceOfferModel) SetTotalAmount(v AmountModel)`

SetTotalAmount sets TotalAmount field to given value.


### GetPrePaymentAmount

`func (o *ServiceOfferModel) GetPrePaymentAmount() MonetaryValueModel`

GetPrePaymentAmount returns the PrePaymentAmount field if non-nil, zero value otherwise.

### GetPrePaymentAmountOk

`func (o *ServiceOfferModel) GetPrePaymentAmountOk() (*MonetaryValueModel, bool)`

GetPrePaymentAmountOk returns a tuple with the PrePaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaymentAmount

`func (o *ServiceOfferModel) SetPrePaymentAmount(v MonetaryValueModel)`

SetPrePaymentAmount sets PrePaymentAmount field to given value.


### GetFees

`func (o *ServiceOfferModel) GetFees() []OfferFeeModel`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ServiceOfferModel) GetFeesOk() (*[]OfferFeeModel, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ServiceOfferModel) SetFees(v []OfferFeeModel)`

SetFees sets Fees field to given value.

### HasFees

`func (o *ServiceOfferModel) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetDates

`func (o *ServiceOfferModel) GetDates() []ServiceOfferItemModel`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *ServiceOfferModel) GetDatesOk() (*[]ServiceOfferItemModel, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *ServiceOfferModel) SetDates(v []ServiceOfferItemModel)`

SetDates sets Dates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


