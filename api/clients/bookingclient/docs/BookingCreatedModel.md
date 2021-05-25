# BookingCreatedModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Booking id | 
**ReservationIds** | [**[]ReservationCreatedModel**](ReservationCreatedModel.md) | List of ids for newly created reservations | 

## Methods

### NewBookingCreatedModel

`func NewBookingCreatedModel(id string, reservationIds []ReservationCreatedModel, ) *BookingCreatedModel`

NewBookingCreatedModel instantiates a new BookingCreatedModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingCreatedModelWithDefaults

`func NewBookingCreatedModelWithDefaults() *BookingCreatedModel`

NewBookingCreatedModelWithDefaults instantiates a new BookingCreatedModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookingCreatedModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookingCreatedModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookingCreatedModel) SetId(v string)`

SetId sets Id field to given value.


### GetReservationIds

`func (o *BookingCreatedModel) GetReservationIds() []ReservationCreatedModel`

GetReservationIds returns the ReservationIds field if non-nil, zero value otherwise.

### GetReservationIdsOk

`func (o *BookingCreatedModel) GetReservationIdsOk() (*[]ReservationCreatedModel, bool)`

GetReservationIdsOk returns a tuple with the ReservationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationIds

`func (o *BookingCreatedModel) SetReservationIds(v []ReservationCreatedModel)`

SetReservationIds sets ReservationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


