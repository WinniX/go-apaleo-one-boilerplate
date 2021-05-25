# ReservationServiceItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | [**ServiceModel**](ServiceModel.md) |  | 
**TotalAmount** | [**AmountModel**](AmountModel.md) |  | 
**Dates** | [**[]ServiceDateItemModel**](ServiceDateItemModel.md) | The dates the service will be delivered with its price | 

## Methods

### NewReservationServiceItemModel

`func NewReservationServiceItemModel(service ServiceModel, totalAmount AmountModel, dates []ServiceDateItemModel, ) *ReservationServiceItemModel`

NewReservationServiceItemModel instantiates a new ReservationServiceItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationServiceItemModelWithDefaults

`func NewReservationServiceItemModelWithDefaults() *ReservationServiceItemModel`

NewReservationServiceItemModelWithDefaults instantiates a new ReservationServiceItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ReservationServiceItemModel) GetService() ServiceModel`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ReservationServiceItemModel) GetServiceOk() (*ServiceModel, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ReservationServiceItemModel) SetService(v ServiceModel)`

SetService sets Service field to given value.


### GetTotalAmount

`func (o *ReservationServiceItemModel) GetTotalAmount() AmountModel`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ReservationServiceItemModel) GetTotalAmountOk() (*AmountModel, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ReservationServiceItemModel) SetTotalAmount(v AmountModel)`

SetTotalAmount sets TotalAmount field to given value.


### GetDates

`func (o *ReservationServiceItemModel) GetDates() []ServiceDateItemModel`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *ReservationServiceItemModel) GetDatesOk() (*[]ServiceDateItemModel, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *ReservationServiceItemModel) SetDates(v []ServiceDateItemModel)`

SetDates sets Dates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


