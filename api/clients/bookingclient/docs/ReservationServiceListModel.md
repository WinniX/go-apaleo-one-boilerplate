# ReservationServiceListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | [**[]ReservationServiceItemModel**](ReservationServiceItemModel.md) | The list of services booked for the reservation | 
**Count** | **int64** | Total count of items | 

## Methods

### NewReservationServiceListModel

`func NewReservationServiceListModel(services []ReservationServiceItemModel, count int64, ) *ReservationServiceListModel`

NewReservationServiceListModel instantiates a new ReservationServiceListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationServiceListModelWithDefaults

`func NewReservationServiceListModelWithDefaults() *ReservationServiceListModel`

NewReservationServiceListModelWithDefaults instantiates a new ReservationServiceListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ReservationServiceListModel) GetServices() []ReservationServiceItemModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ReservationServiceListModel) GetServicesOk() (*[]ReservationServiceItemModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ReservationServiceListModel) SetServices(v []ReservationServiceItemModel)`

SetServices sets Services field to given value.


### GetCount

`func (o *ReservationServiceListModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReservationServiceListModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReservationServiceListModel) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


