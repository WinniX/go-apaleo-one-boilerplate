# PickUpReservationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservations** | [**[]PickUpReservationModel**](PickUpReservationModel.md) | List of reservations to pick up to the existing group booking | 

## Methods

### NewPickUpReservationsModel

`func NewPickUpReservationsModel(reservations []PickUpReservationModel, ) *PickUpReservationsModel`

NewPickUpReservationsModel instantiates a new PickUpReservationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickUpReservationsModelWithDefaults

`func NewPickUpReservationsModelWithDefaults() *PickUpReservationsModel`

NewPickUpReservationsModelWithDefaults instantiates a new PickUpReservationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservations

`func (o *PickUpReservationsModel) GetReservations() []PickUpReservationModel`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *PickUpReservationsModel) GetReservationsOk() (*[]PickUpReservationModel, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *PickUpReservationsModel) SetReservations(v []PickUpReservationModel)`

SetReservations sets Reservations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


