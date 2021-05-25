# ReservationListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservations** | [**[]ReservationItemModel**](ReservationItemModel.md) | List of reservations | 
**Count** | **int64** | Total count of items | 

## Methods

### NewReservationListModel

`func NewReservationListModel(reservations []ReservationItemModel, count int64, ) *ReservationListModel`

NewReservationListModel instantiates a new ReservationListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationListModelWithDefaults

`func NewReservationListModelWithDefaults() *ReservationListModel`

NewReservationListModelWithDefaults instantiates a new ReservationListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservations

`func (o *ReservationListModel) GetReservations() []ReservationItemModel`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *ReservationListModel) GetReservationsOk() (*[]ReservationItemModel, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *ReservationListModel) SetReservations(v []ReservationItemModel)`

SetReservations sets Reservations field to given value.


### GetCount

`func (o *ReservationListModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReservationListModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReservationListModel) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


