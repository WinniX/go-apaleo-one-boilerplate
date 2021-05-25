# BookingModel

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
**PropertyValues** | Pointer to [**[]PropertyValueModel**](PropertyValueModel.md) | Property specific values like total amount and balance | [optional] 
**Reservations** | Pointer to [**[]BookingReservationModel**](BookingReservationModel.md) | Reservations within this booking | [optional] 

## Methods

### NewBookingModel

`func NewBookingModel(id string, created time.Time, modified time.Time, ) *BookingModel`

NewBookingModel instantiates a new BookingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingModelWithDefaults

`func NewBookingModelWithDefaults() *BookingModel`

NewBookingModelWithDefaults instantiates a new BookingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookingModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookingModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookingModel) SetId(v string)`

SetId sets Id field to given value.


### GetGroupId

`func (o *BookingModel) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BookingModel) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BookingModel) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BookingModel) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetBooker

`func (o *BookingModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *BookingModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *BookingModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.

### HasBooker

`func (o *BookingModel) HasBooker() bool`

HasBooker returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *BookingModel) GetPaymentAccount() PaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *BookingModel) GetPaymentAccountOk() (*PaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *BookingModel) SetPaymentAccount(v PaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *BookingModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetComment

`func (o *BookingModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *BookingModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *BookingModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *BookingModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetBookerComment

`func (o *BookingModel) GetBookerComment() string`

GetBookerComment returns the BookerComment field if non-nil, zero value otherwise.

### GetBookerCommentOk

`func (o *BookingModel) GetBookerCommentOk() (*string, bool)`

GetBookerCommentOk returns a tuple with the BookerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookerComment

`func (o *BookingModel) SetBookerComment(v string)`

SetBookerComment sets BookerComment field to given value.

### HasBookerComment

`func (o *BookingModel) HasBookerComment() bool`

HasBookerComment returns a boolean if a field has been set.

### GetCreated

`func (o *BookingModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BookingModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BookingModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *BookingModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BookingModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BookingModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetPropertyValues

`func (o *BookingModel) GetPropertyValues() []PropertyValueModel`

GetPropertyValues returns the PropertyValues field if non-nil, zero value otherwise.

### GetPropertyValuesOk

`func (o *BookingModel) GetPropertyValuesOk() (*[]PropertyValueModel, bool)`

GetPropertyValuesOk returns a tuple with the PropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyValues

`func (o *BookingModel) SetPropertyValues(v []PropertyValueModel)`

SetPropertyValues sets PropertyValues field to given value.

### HasPropertyValues

`func (o *BookingModel) HasPropertyValues() bool`

HasPropertyValues returns a boolean if a field has been set.

### GetReservations

`func (o *BookingModel) GetReservations() []BookingReservationModel`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *BookingModel) GetReservationsOk() (*[]BookingReservationModel, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *BookingModel) SetReservations(v []BookingReservationModel)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *BookingModel) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


