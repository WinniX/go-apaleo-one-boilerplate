# BookingItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Booking id | 
**GroupId** | Pointer to **string** | Group id | [optional] 
**Booker** | Pointer to [**BookerModel**](BookerModel.md) |  | [optional] 
**PaymentAccount** | Pointer to [**PaymentAccountModel**](PaymentAccountModel.md) |  | [optional] 
**Comment** | Pointer to **string** | Additional information and comments | [optional] 
**BookerComment** | Pointer to **string** | Additional information and comment by the booker | [optional] 
**Created** | **time.Time** | Date of creation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Modified** | **time.Time** | Date of last modification&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Reservations** | Pointer to [**[]BookingReservationModel**](BookingReservationModel.md) | Reservations within this booking | [optional] 

## Methods

### NewBookingItemModel

`func NewBookingItemModel(id string, created time.Time, modified time.Time, ) *BookingItemModel`

NewBookingItemModel instantiates a new BookingItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingItemModelWithDefaults

`func NewBookingItemModelWithDefaults() *BookingItemModel`

NewBookingItemModelWithDefaults instantiates a new BookingItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookingItemModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookingItemModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookingItemModel) SetId(v string)`

SetId sets Id field to given value.


### GetGroupId

`func (o *BookingItemModel) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BookingItemModel) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BookingItemModel) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BookingItemModel) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetBooker

`func (o *BookingItemModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *BookingItemModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *BookingItemModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.

### HasBooker

`func (o *BookingItemModel) HasBooker() bool`

HasBooker returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *BookingItemModel) GetPaymentAccount() PaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *BookingItemModel) GetPaymentAccountOk() (*PaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *BookingItemModel) SetPaymentAccount(v PaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *BookingItemModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetComment

`func (o *BookingItemModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *BookingItemModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *BookingItemModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *BookingItemModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetBookerComment

`func (o *BookingItemModel) GetBookerComment() string`

GetBookerComment returns the BookerComment field if non-nil, zero value otherwise.

### GetBookerCommentOk

`func (o *BookingItemModel) GetBookerCommentOk() (*string, bool)`

GetBookerCommentOk returns a tuple with the BookerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookerComment

`func (o *BookingItemModel) SetBookerComment(v string)`

SetBookerComment sets BookerComment field to given value.

### HasBookerComment

`func (o *BookingItemModel) HasBookerComment() bool`

HasBookerComment returns a boolean if a field has been set.

### GetCreated

`func (o *BookingItemModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BookingItemModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BookingItemModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *BookingItemModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BookingItemModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BookingItemModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetReservations

`func (o *BookingItemModel) GetReservations() []BookingReservationModel`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *BookingItemModel) GetReservationsOk() (*[]BookingReservationModel, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *BookingItemModel) SetReservations(v []BookingReservationModel)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *BookingItemModel) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


