# ReservationItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Reservation id | 
**BookingId** | **string** | Booking id | 
**BlockId** | Pointer to **string** | Block id | [optional] 
**GroupName** | Pointer to **string** | Name of the group | [optional] 
**Status** | **string** | Status of the reservation | 
**CheckInTime** | Pointer to **time.Time** | Time of check-in&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**CheckOutTime** | Pointer to **time.Time** | Time of check-out&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**CancellationTime** | Pointer to **time.Time** | Time of cancellation, if the reservation was canceled&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**NoShowTime** | Pointer to **time.Time** | Time of setting no-show reservation status&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**UnitGroup** | [**EmbeddedUnitGroupModel**](EmbeddedUnitGroupModel.md) |  | 
**Unit** | Pointer to [**EmbeddedUnitModel**](EmbeddedUnitModel.md) |  | [optional] 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**Arrival** | **time.Time** | Date of arrival&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Departure** | **time.Time** | Date of departure&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Created** | **time.Time** | Date of creation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Modified** | **time.Time** | Date of last modification&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Adults** | **int32** | Number of adults | 
**ChildrenAges** | Pointer to **[]int32** | The ages of the children | [optional] 
**Comment** | Pointer to **string** | Additional information and comments | [optional] 
**GuestComment** | Pointer to **string** | Additional information and comment by the guest | [optional] 
**ExternalCode** | Pointer to **string** | Code in external system | [optional] 
**ChannelCode** | **string** | Channel code | 
**Source** | Pointer to **string** | Source of the reservation (e.g Hotels.com, Orbitz, etc.) | [optional] 
**PrimaryGuest** | Pointer to [**GuestModel**](GuestModel.md) |  | [optional] 
**AdditionalGuests** | Pointer to [**[]GuestModel**](GuestModel.md) | Additional guests of the reservation. | [optional] 
**Booker** | Pointer to [**BookerModel**](BookerModel.md) |  | [optional] 
**PaymentAccount** | Pointer to [**PaymentAccountModel**](PaymentAccountModel.md) |  | [optional] 
**GuaranteeType** | **string** | The strongest guarantee for the rate plans booked in this reservation | 
**CancellationFee** | [**ReservationCancellationFeeModel**](ReservationCancellationFeeModel.md) |  | 
**NoShowFee** | [**ReservationNoShowFeeModel**](ReservationNoShowFeeModel.md) |  | 
**TravelPurpose** | Pointer to **string** | The purpose of the trip, leisure or business | [optional] 
**Balance** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**AssignedUnits** | Pointer to [**[]ReservationAssignedUnitModel**](ReservationAssignedUnitModel.md) | The list of units assigned to this reservation | [optional] 
**TimeSlices** | Pointer to [**[]TimeSliceModel**](TimeSliceModel.md) | The list of time slices with the reserved units / unit groups for the stay | [optional] 
**Services** | Pointer to [**[]ReservationServiceItemModel**](ReservationServiceItemModel.md) | The list of additional services (extras, add-ons) reserved for the stay | [optional] 
**ValidationMessages** | Pointer to [**[]ReservationValidationMessageModel**](ReservationValidationMessageModel.md) | Validation rules are applied to reservations during their lifetime.  For example a reservation that was created while the house or unit group is already fully booked.  Whenever a rule was or is currently violated, a validation message will be added to this list.  They can be deleted whenever the hotel staff worked them off. | [optional] 
**Actions** | Pointer to [**[]ActionModelReservationActionNotAllowedReservationActionReason**](ActionModelReservationActionNotAllowedReservationActionReason.md) | The list of actions for this reservation | [optional] 
**Company** | Pointer to [**EmbeddedCompanyModel**](EmbeddedCompanyModel.md) |  | [optional] 
**CorporateCode** | Pointer to **string** | Corporate code provided during creation. Used to find offers during amend. | [optional] 
**AllFoliosHaveInvoice** | Pointer to **bool** | Whether all folios of a reservation have an invoice | [optional] 
**HasCityTax** | **bool** | Whether the city tax has already been added to the reservation. Set to false, if the property does not have city tax configured | 
**Commission** | Pointer to [**CommissionModel**](CommissionModel.md) |  | [optional] 
**PromoCode** | Pointer to **string** | The promo code associated with a certain special offer used to create the reservation | [optional] 

## Methods

### NewReservationItemModel

`func NewReservationItemModel(id string, bookingId string, status string, property EmbeddedPropertyModel, ratePlan EmbeddedRatePlanModel, unitGroup EmbeddedUnitGroupModel, totalGrossAmount MonetaryValueModel, arrival time.Time, departure time.Time, created time.Time, modified time.Time, adults int32, channelCode string, guaranteeType string, cancellationFee ReservationCancellationFeeModel, noShowFee ReservationNoShowFeeModel, balance MonetaryValueModel, hasCityTax bool, ) *ReservationItemModel`

NewReservationItemModel instantiates a new ReservationItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationItemModelWithDefaults

`func NewReservationItemModelWithDefaults() *ReservationItemModel`

NewReservationItemModelWithDefaults instantiates a new ReservationItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationItemModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationItemModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationItemModel) SetId(v string)`

SetId sets Id field to given value.


### GetBookingId

`func (o *ReservationItemModel) GetBookingId() string`

GetBookingId returns the BookingId field if non-nil, zero value otherwise.

### GetBookingIdOk

`func (o *ReservationItemModel) GetBookingIdOk() (*string, bool)`

GetBookingIdOk returns a tuple with the BookingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingId

`func (o *ReservationItemModel) SetBookingId(v string)`

SetBookingId sets BookingId field to given value.


### GetBlockId

`func (o *ReservationItemModel) GetBlockId() string`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *ReservationItemModel) GetBlockIdOk() (*string, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *ReservationItemModel) SetBlockId(v string)`

SetBlockId sets BlockId field to given value.

### HasBlockId

`func (o *ReservationItemModel) HasBlockId() bool`

HasBlockId returns a boolean if a field has been set.

### GetGroupName

`func (o *ReservationItemModel) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ReservationItemModel) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ReservationItemModel) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ReservationItemModel) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStatus

`func (o *ReservationItemModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReservationItemModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReservationItemModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCheckInTime

`func (o *ReservationItemModel) GetCheckInTime() time.Time`

GetCheckInTime returns the CheckInTime field if non-nil, zero value otherwise.

### GetCheckInTimeOk

`func (o *ReservationItemModel) GetCheckInTimeOk() (*time.Time, bool)`

GetCheckInTimeOk returns a tuple with the CheckInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInTime

`func (o *ReservationItemModel) SetCheckInTime(v time.Time)`

SetCheckInTime sets CheckInTime field to given value.

### HasCheckInTime

`func (o *ReservationItemModel) HasCheckInTime() bool`

HasCheckInTime returns a boolean if a field has been set.

### GetCheckOutTime

`func (o *ReservationItemModel) GetCheckOutTime() time.Time`

GetCheckOutTime returns the CheckOutTime field if non-nil, zero value otherwise.

### GetCheckOutTimeOk

`func (o *ReservationItemModel) GetCheckOutTimeOk() (*time.Time, bool)`

GetCheckOutTimeOk returns a tuple with the CheckOutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOutTime

`func (o *ReservationItemModel) SetCheckOutTime(v time.Time)`

SetCheckOutTime sets CheckOutTime field to given value.

### HasCheckOutTime

`func (o *ReservationItemModel) HasCheckOutTime() bool`

HasCheckOutTime returns a boolean if a field has been set.

### GetCancellationTime

`func (o *ReservationItemModel) GetCancellationTime() time.Time`

GetCancellationTime returns the CancellationTime field if non-nil, zero value otherwise.

### GetCancellationTimeOk

`func (o *ReservationItemModel) GetCancellationTimeOk() (*time.Time, bool)`

GetCancellationTimeOk returns a tuple with the CancellationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationTime

`func (o *ReservationItemModel) SetCancellationTime(v time.Time)`

SetCancellationTime sets CancellationTime field to given value.

### HasCancellationTime

`func (o *ReservationItemModel) HasCancellationTime() bool`

HasCancellationTime returns a boolean if a field has been set.

### GetNoShowTime

`func (o *ReservationItemModel) GetNoShowTime() time.Time`

GetNoShowTime returns the NoShowTime field if non-nil, zero value otherwise.

### GetNoShowTimeOk

`func (o *ReservationItemModel) GetNoShowTimeOk() (*time.Time, bool)`

GetNoShowTimeOk returns a tuple with the NoShowTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShowTime

`func (o *ReservationItemModel) SetNoShowTime(v time.Time)`

SetNoShowTime sets NoShowTime field to given value.

### HasNoShowTime

`func (o *ReservationItemModel) HasNoShowTime() bool`

HasNoShowTime returns a boolean if a field has been set.

### GetProperty

`func (o *ReservationItemModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ReservationItemModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ReservationItemModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetRatePlan

`func (o *ReservationItemModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *ReservationItemModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *ReservationItemModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetUnitGroup

`func (o *ReservationItemModel) GetUnitGroup() EmbeddedUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *ReservationItemModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *ReservationItemModel) SetUnitGroup(v EmbeddedUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetUnit

`func (o *ReservationItemModel) GetUnit() EmbeddedUnitModel`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ReservationItemModel) GetUnitOk() (*EmbeddedUnitModel, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ReservationItemModel) SetUnit(v EmbeddedUnitModel)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ReservationItemModel) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTotalGrossAmount

`func (o *ReservationItemModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *ReservationItemModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *ReservationItemModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetArrival

`func (o *ReservationItemModel) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *ReservationItemModel) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *ReservationItemModel) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *ReservationItemModel) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *ReservationItemModel) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *ReservationItemModel) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.


### GetCreated

`func (o *ReservationItemModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReservationItemModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReservationItemModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *ReservationItemModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ReservationItemModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ReservationItemModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAdults

`func (o *ReservationItemModel) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *ReservationItemModel) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *ReservationItemModel) SetAdults(v int32)`

SetAdults sets Adults field to given value.


### GetChildrenAges

`func (o *ReservationItemModel) GetChildrenAges() []int32`

GetChildrenAges returns the ChildrenAges field if non-nil, zero value otherwise.

### GetChildrenAgesOk

`func (o *ReservationItemModel) GetChildrenAgesOk() (*[]int32, bool)`

GetChildrenAgesOk returns a tuple with the ChildrenAges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAges

`func (o *ReservationItemModel) SetChildrenAges(v []int32)`

SetChildrenAges sets ChildrenAges field to given value.

### HasChildrenAges

`func (o *ReservationItemModel) HasChildrenAges() bool`

HasChildrenAges returns a boolean if a field has been set.

### GetComment

`func (o *ReservationItemModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReservationItemModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReservationItemModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReservationItemModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGuestComment

`func (o *ReservationItemModel) GetGuestComment() string`

GetGuestComment returns the GuestComment field if non-nil, zero value otherwise.

### GetGuestCommentOk

`func (o *ReservationItemModel) GetGuestCommentOk() (*string, bool)`

GetGuestCommentOk returns a tuple with the GuestComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestComment

`func (o *ReservationItemModel) SetGuestComment(v string)`

SetGuestComment sets GuestComment field to given value.

### HasGuestComment

`func (o *ReservationItemModel) HasGuestComment() bool`

HasGuestComment returns a boolean if a field has been set.

### GetExternalCode

`func (o *ReservationItemModel) GetExternalCode() string`

GetExternalCode returns the ExternalCode field if non-nil, zero value otherwise.

### GetExternalCodeOk

`func (o *ReservationItemModel) GetExternalCodeOk() (*string, bool)`

GetExternalCodeOk returns a tuple with the ExternalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCode

`func (o *ReservationItemModel) SetExternalCode(v string)`

SetExternalCode sets ExternalCode field to given value.

### HasExternalCode

`func (o *ReservationItemModel) HasExternalCode() bool`

HasExternalCode returns a boolean if a field has been set.

### GetChannelCode

`func (o *ReservationItemModel) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *ReservationItemModel) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *ReservationItemModel) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetSource

`func (o *ReservationItemModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReservationItemModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReservationItemModel) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReservationItemModel) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrimaryGuest

`func (o *ReservationItemModel) GetPrimaryGuest() GuestModel`

GetPrimaryGuest returns the PrimaryGuest field if non-nil, zero value otherwise.

### GetPrimaryGuestOk

`func (o *ReservationItemModel) GetPrimaryGuestOk() (*GuestModel, bool)`

GetPrimaryGuestOk returns a tuple with the PrimaryGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGuest

`func (o *ReservationItemModel) SetPrimaryGuest(v GuestModel)`

SetPrimaryGuest sets PrimaryGuest field to given value.

### HasPrimaryGuest

`func (o *ReservationItemModel) HasPrimaryGuest() bool`

HasPrimaryGuest returns a boolean if a field has been set.

### GetAdditionalGuests

`func (o *ReservationItemModel) GetAdditionalGuests() []GuestModel`

GetAdditionalGuests returns the AdditionalGuests field if non-nil, zero value otherwise.

### GetAdditionalGuestsOk

`func (o *ReservationItemModel) GetAdditionalGuestsOk() (*[]GuestModel, bool)`

GetAdditionalGuestsOk returns a tuple with the AdditionalGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGuests

`func (o *ReservationItemModel) SetAdditionalGuests(v []GuestModel)`

SetAdditionalGuests sets AdditionalGuests field to given value.

### HasAdditionalGuests

`func (o *ReservationItemModel) HasAdditionalGuests() bool`

HasAdditionalGuests returns a boolean if a field has been set.

### GetBooker

`func (o *ReservationItemModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *ReservationItemModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *ReservationItemModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.

### HasBooker

`func (o *ReservationItemModel) HasBooker() bool`

HasBooker returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *ReservationItemModel) GetPaymentAccount() PaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *ReservationItemModel) GetPaymentAccountOk() (*PaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *ReservationItemModel) SetPaymentAccount(v PaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *ReservationItemModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetGuaranteeType

`func (o *ReservationItemModel) GetGuaranteeType() string`

GetGuaranteeType returns the GuaranteeType field if non-nil, zero value otherwise.

### GetGuaranteeTypeOk

`func (o *ReservationItemModel) GetGuaranteeTypeOk() (*string, bool)`

GetGuaranteeTypeOk returns a tuple with the GuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeType

`func (o *ReservationItemModel) SetGuaranteeType(v string)`

SetGuaranteeType sets GuaranteeType field to given value.


### GetCancellationFee

`func (o *ReservationItemModel) GetCancellationFee() ReservationCancellationFeeModel`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *ReservationItemModel) GetCancellationFeeOk() (*ReservationCancellationFeeModel, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *ReservationItemModel) SetCancellationFee(v ReservationCancellationFeeModel)`

SetCancellationFee sets CancellationFee field to given value.


### GetNoShowFee

`func (o *ReservationItemModel) GetNoShowFee() ReservationNoShowFeeModel`

GetNoShowFee returns the NoShowFee field if non-nil, zero value otherwise.

### GetNoShowFeeOk

`func (o *ReservationItemModel) GetNoShowFeeOk() (*ReservationNoShowFeeModel, bool)`

GetNoShowFeeOk returns a tuple with the NoShowFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShowFee

`func (o *ReservationItemModel) SetNoShowFee(v ReservationNoShowFeeModel)`

SetNoShowFee sets NoShowFee field to given value.


### GetTravelPurpose

`func (o *ReservationItemModel) GetTravelPurpose() string`

GetTravelPurpose returns the TravelPurpose field if non-nil, zero value otherwise.

### GetTravelPurposeOk

`func (o *ReservationItemModel) GetTravelPurposeOk() (*string, bool)`

GetTravelPurposeOk returns a tuple with the TravelPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelPurpose

`func (o *ReservationItemModel) SetTravelPurpose(v string)`

SetTravelPurpose sets TravelPurpose field to given value.

### HasTravelPurpose

`func (o *ReservationItemModel) HasTravelPurpose() bool`

HasTravelPurpose returns a boolean if a field has been set.

### GetBalance

`func (o *ReservationItemModel) GetBalance() MonetaryValueModel`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ReservationItemModel) GetBalanceOk() (*MonetaryValueModel, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ReservationItemModel) SetBalance(v MonetaryValueModel)`

SetBalance sets Balance field to given value.


### GetAssignedUnits

`func (o *ReservationItemModel) GetAssignedUnits() []ReservationAssignedUnitModel`

GetAssignedUnits returns the AssignedUnits field if non-nil, zero value otherwise.

### GetAssignedUnitsOk

`func (o *ReservationItemModel) GetAssignedUnitsOk() (*[]ReservationAssignedUnitModel, bool)`

GetAssignedUnitsOk returns a tuple with the AssignedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUnits

`func (o *ReservationItemModel) SetAssignedUnits(v []ReservationAssignedUnitModel)`

SetAssignedUnits sets AssignedUnits field to given value.

### HasAssignedUnits

`func (o *ReservationItemModel) HasAssignedUnits() bool`

HasAssignedUnits returns a boolean if a field has been set.

### GetTimeSlices

`func (o *ReservationItemModel) GetTimeSlices() []TimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *ReservationItemModel) GetTimeSlicesOk() (*[]TimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *ReservationItemModel) SetTimeSlices(v []TimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.

### HasTimeSlices

`func (o *ReservationItemModel) HasTimeSlices() bool`

HasTimeSlices returns a boolean if a field has been set.

### GetServices

`func (o *ReservationItemModel) GetServices() []ReservationServiceItemModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ReservationItemModel) GetServicesOk() (*[]ReservationServiceItemModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ReservationItemModel) SetServices(v []ReservationServiceItemModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *ReservationItemModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetValidationMessages

`func (o *ReservationItemModel) GetValidationMessages() []ReservationValidationMessageModel`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *ReservationItemModel) GetValidationMessagesOk() (*[]ReservationValidationMessageModel, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *ReservationItemModel) SetValidationMessages(v []ReservationValidationMessageModel)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *ReservationItemModel) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.

### GetActions

`func (o *ReservationItemModel) GetActions() []ActionModelReservationActionNotAllowedReservationActionReason`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ReservationItemModel) GetActionsOk() (*[]ActionModelReservationActionNotAllowedReservationActionReason, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ReservationItemModel) SetActions(v []ActionModelReservationActionNotAllowedReservationActionReason)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ReservationItemModel) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCompany

`func (o *ReservationItemModel) GetCompany() EmbeddedCompanyModel`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ReservationItemModel) GetCompanyOk() (*EmbeddedCompanyModel, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ReservationItemModel) SetCompany(v EmbeddedCompanyModel)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ReservationItemModel) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCorporateCode

`func (o *ReservationItemModel) GetCorporateCode() string`

GetCorporateCode returns the CorporateCode field if non-nil, zero value otherwise.

### GetCorporateCodeOk

`func (o *ReservationItemModel) GetCorporateCodeOk() (*string, bool)`

GetCorporateCodeOk returns a tuple with the CorporateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCode

`func (o *ReservationItemModel) SetCorporateCode(v string)`

SetCorporateCode sets CorporateCode field to given value.

### HasCorporateCode

`func (o *ReservationItemModel) HasCorporateCode() bool`

HasCorporateCode returns a boolean if a field has been set.

### GetAllFoliosHaveInvoice

`func (o *ReservationItemModel) GetAllFoliosHaveInvoice() bool`

GetAllFoliosHaveInvoice returns the AllFoliosHaveInvoice field if non-nil, zero value otherwise.

### GetAllFoliosHaveInvoiceOk

`func (o *ReservationItemModel) GetAllFoliosHaveInvoiceOk() (*bool, bool)`

GetAllFoliosHaveInvoiceOk returns a tuple with the AllFoliosHaveInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFoliosHaveInvoice

`func (o *ReservationItemModel) SetAllFoliosHaveInvoice(v bool)`

SetAllFoliosHaveInvoice sets AllFoliosHaveInvoice field to given value.

### HasAllFoliosHaveInvoice

`func (o *ReservationItemModel) HasAllFoliosHaveInvoice() bool`

HasAllFoliosHaveInvoice returns a boolean if a field has been set.

### GetHasCityTax

`func (o *ReservationItemModel) GetHasCityTax() bool`

GetHasCityTax returns the HasCityTax field if non-nil, zero value otherwise.

### GetHasCityTaxOk

`func (o *ReservationItemModel) GetHasCityTaxOk() (*bool, bool)`

GetHasCityTaxOk returns a tuple with the HasCityTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCityTax

`func (o *ReservationItemModel) SetHasCityTax(v bool)`

SetHasCityTax sets HasCityTax field to given value.


### GetCommission

`func (o *ReservationItemModel) GetCommission() CommissionModel`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *ReservationItemModel) GetCommissionOk() (*CommissionModel, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *ReservationItemModel) SetCommission(v CommissionModel)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *ReservationItemModel) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetPromoCode

`func (o *ReservationItemModel) GetPromoCode() string`

GetPromoCode returns the PromoCode field if non-nil, zero value otherwise.

### GetPromoCodeOk

`func (o *ReservationItemModel) GetPromoCodeOk() (*string, bool)`

GetPromoCodeOk returns a tuple with the PromoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCode

`func (o *ReservationItemModel) SetPromoCode(v string)`

SetPromoCode sets PromoCode field to given value.

### HasPromoCode

`func (o *ReservationItemModel) HasPromoCode() bool`

HasPromoCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


