# PickUpReservationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockId** | **string** | ID of the block | 
**Services** | Pointer to [**[]BookReservationServiceModel**](BookReservationServiceModel.md) | Additional services (extras, add-ons) that should be added to the reservation | [optional] 
**Arrival** | **string** | Date of arrival and the optional time with UTC offset&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Departure** | **string** | Date of departure and the optional time with UTC offset&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Adults** | **int32** | Number of adults | 
**ChildrenAges** | Pointer to **[]int32** | The ages of the children | [optional] 
**Comment** | Pointer to **string** | Additional information and comments | [optional] 
**GuestComment** | Pointer to **string** | Additional information and comment by the guest | [optional] 
**PrimaryGuest** | Pointer to [**GuestModel**](GuestModel.md) |  | [optional] 
**AdditionalGuests** | Pointer to [**[]GuestModel**](GuestModel.md) | Additional guests of the reservation. | [optional] 
**TravelPurpose** | Pointer to **string** | The purpose of the trip, leisure or business | [optional] 

## Methods

### NewPickUpReservationModel

`func NewPickUpReservationModel(blockId string, arrival string, departure string, adults int32, ) *PickUpReservationModel`

NewPickUpReservationModel instantiates a new PickUpReservationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickUpReservationModelWithDefaults

`func NewPickUpReservationModelWithDefaults() *PickUpReservationModel`

NewPickUpReservationModelWithDefaults instantiates a new PickUpReservationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockId

`func (o *PickUpReservationModel) GetBlockId() string`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *PickUpReservationModel) GetBlockIdOk() (*string, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *PickUpReservationModel) SetBlockId(v string)`

SetBlockId sets BlockId field to given value.


### GetServices

`func (o *PickUpReservationModel) GetServices() []BookReservationServiceModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *PickUpReservationModel) GetServicesOk() (*[]BookReservationServiceModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *PickUpReservationModel) SetServices(v []BookReservationServiceModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *PickUpReservationModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetArrival

`func (o *PickUpReservationModel) GetArrival() string`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *PickUpReservationModel) GetArrivalOk() (*string, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *PickUpReservationModel) SetArrival(v string)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *PickUpReservationModel) GetDeparture() string`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *PickUpReservationModel) GetDepartureOk() (*string, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *PickUpReservationModel) SetDeparture(v string)`

SetDeparture sets Departure field to given value.


### GetAdults

`func (o *PickUpReservationModel) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *PickUpReservationModel) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *PickUpReservationModel) SetAdults(v int32)`

SetAdults sets Adults field to given value.


### GetChildrenAges

`func (o *PickUpReservationModel) GetChildrenAges() []int32`

GetChildrenAges returns the ChildrenAges field if non-nil, zero value otherwise.

### GetChildrenAgesOk

`func (o *PickUpReservationModel) GetChildrenAgesOk() (*[]int32, bool)`

GetChildrenAgesOk returns a tuple with the ChildrenAges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAges

`func (o *PickUpReservationModel) SetChildrenAges(v []int32)`

SetChildrenAges sets ChildrenAges field to given value.

### HasChildrenAges

`func (o *PickUpReservationModel) HasChildrenAges() bool`

HasChildrenAges returns a boolean if a field has been set.

### GetComment

`func (o *PickUpReservationModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PickUpReservationModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PickUpReservationModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PickUpReservationModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGuestComment

`func (o *PickUpReservationModel) GetGuestComment() string`

GetGuestComment returns the GuestComment field if non-nil, zero value otherwise.

### GetGuestCommentOk

`func (o *PickUpReservationModel) GetGuestCommentOk() (*string, bool)`

GetGuestCommentOk returns a tuple with the GuestComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestComment

`func (o *PickUpReservationModel) SetGuestComment(v string)`

SetGuestComment sets GuestComment field to given value.

### HasGuestComment

`func (o *PickUpReservationModel) HasGuestComment() bool`

HasGuestComment returns a boolean if a field has been set.

### GetPrimaryGuest

`func (o *PickUpReservationModel) GetPrimaryGuest() GuestModel`

GetPrimaryGuest returns the PrimaryGuest field if non-nil, zero value otherwise.

### GetPrimaryGuestOk

`func (o *PickUpReservationModel) GetPrimaryGuestOk() (*GuestModel, bool)`

GetPrimaryGuestOk returns a tuple with the PrimaryGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGuest

`func (o *PickUpReservationModel) SetPrimaryGuest(v GuestModel)`

SetPrimaryGuest sets PrimaryGuest field to given value.

### HasPrimaryGuest

`func (o *PickUpReservationModel) HasPrimaryGuest() bool`

HasPrimaryGuest returns a boolean if a field has been set.

### GetAdditionalGuests

`func (o *PickUpReservationModel) GetAdditionalGuests() []GuestModel`

GetAdditionalGuests returns the AdditionalGuests field if non-nil, zero value otherwise.

### GetAdditionalGuestsOk

`func (o *PickUpReservationModel) GetAdditionalGuestsOk() (*[]GuestModel, bool)`

GetAdditionalGuestsOk returns a tuple with the AdditionalGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGuests

`func (o *PickUpReservationModel) SetAdditionalGuests(v []GuestModel)`

SetAdditionalGuests sets AdditionalGuests field to given value.

### HasAdditionalGuests

`func (o *PickUpReservationModel) HasAdditionalGuests() bool`

HasAdditionalGuests returns a boolean if a field has been set.

### GetTravelPurpose

`func (o *PickUpReservationModel) GetTravelPurpose() string`

GetTravelPurpose returns the TravelPurpose field if non-nil, zero value otherwise.

### GetTravelPurposeOk

`func (o *PickUpReservationModel) GetTravelPurposeOk() (*string, bool)`

GetTravelPurposeOk returns a tuple with the TravelPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelPurpose

`func (o *PickUpReservationModel) SetTravelPurpose(v string)`

SetTravelPurpose sets TravelPurpose field to given value.

### HasTravelPurpose

`func (o *PickUpReservationModel) HasTravelPurpose() bool`

HasTravelPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


