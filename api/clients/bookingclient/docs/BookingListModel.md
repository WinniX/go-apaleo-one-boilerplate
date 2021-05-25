# BookingListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bookings** | [**[]BookingItemModel**](BookingItemModel.md) |  | 
**Count** | **int64** | Total count of items | 

## Methods

### NewBookingListModel

`func NewBookingListModel(bookings []BookingItemModel, count int64, ) *BookingListModel`

NewBookingListModel instantiates a new BookingListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingListModelWithDefaults

`func NewBookingListModelWithDefaults() *BookingListModel`

NewBookingListModelWithDefaults instantiates a new BookingListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookings

`func (o *BookingListModel) GetBookings() []BookingItemModel`

GetBookings returns the Bookings field if non-nil, zero value otherwise.

### GetBookingsOk

`func (o *BookingListModel) GetBookingsOk() (*[]BookingItemModel, bool)`

GetBookingsOk returns a tuple with the Bookings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookings

`func (o *BookingListModel) SetBookings(v []BookingItemModel)`

SetBookings sets Bookings field to given value.


### GetCount

`func (o *BookingListModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BookingListModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BookingListModel) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


