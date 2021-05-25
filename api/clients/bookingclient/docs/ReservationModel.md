# ReservationModel

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
**Unit** | Pointer to [**EmbeddedUnitModel**](EmbeddedUnitModel.md) |  | [optional] 
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**UnitGroup** | [**EmbeddedUnitGroupModel**](EmbeddedUnitGroupModel.md) |  | 
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
**TimeSlices** | Pointer to [**[]TimeSliceModel**](TimeSliceModel.md) | The list of time slices with the reserved units / unit groups for the stay | [optional] 
**Services** | Pointer to [**[]ReservationServiceItemModel**](ReservationServiceItemModel.md) | The list of additional services (extras, add-ons) reserved for the stay | [optional] 
**GuaranteeType** | **string** | The strongest guarantee for the rate plans booked in this reservation | 
**CancellationFee** | [**ReservationCancellationFeeModel**](ReservationCancellationFeeModel.md) |  | 
**NoShowFee** | [**ReservationNoShowFeeModel**](ReservationNoShowFeeModel.md) |  | 
**TravelPurpose** | Pointer to **string** | The purpose of the trip, leisure or business | [optional] 
**Balance** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**AssignedUnits** | Pointer to [**[]ReservationAssignedUnitModel**](ReservationAssignedUnitModel.md) | The list of units assigned to this reservation | [optional] 
**ValidationMessages** | Pointer to [**[]ReservationValidationMessageModel**](ReservationValidationMessageModel.md) | Validation rules are applied to reservations during their lifetime.  For example a reservation that was created while the house or unit group is already fully booked.  Whenever a rule was or is currently violated, a validation message will be added to this list.  They can be deleted whenever the hotel staff worked them off. | [optional] 
**Actions** | Pointer to [**[]ActionModelReservationActionNotAllowedReservationActionReason**](ActionModelReservationActionNotAllowedReservationActionReason.md) | The list of actions for this reservation | [optional] 
**Company** | Pointer to [**EmbeddedCompanyModel**](EmbeddedCompanyModel.md) |  | [optional] 
**CorporateCode** | Pointer to **string** | Corporate code provided during creation. Used to find offers during amend. | [optional] 
**AllFoliosHaveInvoice** | Pointer to **bool** | Whether all folios of a reservation have an invoice | [optional] 
**TaxDetails** | [**[]TaxDetailModel**](TaxDetailModel.md) | Tax breakdown, displaying net and tax amount for each VAT type | 
**HasCityTax** | **bool** | Whether the city tax has already been added to the reservation. Set to false, if the property does not have city tax configured | 
**Commission** | Pointer to [**CommissionModel**](CommissionModel.md) |  | [optional] 
**PromoCode** | Pointer to **string** | The promo code associated with a certain special offer used to create the reservation | [optional] 
**PayableAmount** | [**PayableAmountModel**](PayableAmountModel.md) |  | 

## Methods

### NewReservationModel

`func NewReservationModel(id string, bookingId string, status string, property EmbeddedPropertyModel, ratePlan EmbeddedRatePlanModel, unitGroup EmbeddedUnitGroupModel, totalGrossAmount MonetaryValueModel, arrival time.Time, departure time.Time, created time.Time, modified time.Time, adults int32, channelCode string, guaranteeType string, cancellationFee ReservationCancellationFeeModel, noShowFee ReservationNoShowFeeModel, balance MonetaryValueModel, taxDetails []TaxDetailModel, hasCityTax bool, payableAmount PayableAmountModel, ) *ReservationModel`

NewReservationModel instantiates a new ReservationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationModelWithDefaults

`func NewReservationModelWithDefaults() *ReservationModel`

NewReservationModelWithDefaults instantiates a new ReservationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationModel) SetId(v string)`

SetId sets Id field to given value.


### GetBookingId

`func (o *ReservationModel) GetBookingId() string`

GetBookingId returns the BookingId field if non-nil, zero value otherwise.

### GetBookingIdOk

`func (o *ReservationModel) GetBookingIdOk() (*string, bool)`

GetBookingIdOk returns a tuple with the BookingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingId

`func (o *ReservationModel) SetBookingId(v string)`

SetBookingId sets BookingId field to given value.


### GetBlockId

`func (o *ReservationModel) GetBlockId() string`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *ReservationModel) GetBlockIdOk() (*string, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *ReservationModel) SetBlockId(v string)`

SetBlockId sets BlockId field to given value.

### HasBlockId

`func (o *ReservationModel) HasBlockId() bool`

HasBlockId returns a boolean if a field has been set.

### GetGroupName

`func (o *ReservationModel) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ReservationModel) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ReservationModel) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ReservationModel) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStatus

`func (o *ReservationModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReservationModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReservationModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCheckInTime

`func (o *ReservationModel) GetCheckInTime() time.Time`

GetCheckInTime returns the CheckInTime field if non-nil, zero value otherwise.

### GetCheckInTimeOk

`func (o *ReservationModel) GetCheckInTimeOk() (*time.Time, bool)`

GetCheckInTimeOk returns a tuple with the CheckInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInTime

`func (o *ReservationModel) SetCheckInTime(v time.Time)`

SetCheckInTime sets CheckInTime field to given value.

### HasCheckInTime

`func (o *ReservationModel) HasCheckInTime() bool`

HasCheckInTime returns a boolean if a field has been set.

### GetCheckOutTime

`func (o *ReservationModel) GetCheckOutTime() time.Time`

GetCheckOutTime returns the CheckOutTime field if non-nil, zero value otherwise.

### GetCheckOutTimeOk

`func (o *ReservationModel) GetCheckOutTimeOk() (*time.Time, bool)`

GetCheckOutTimeOk returns a tuple with the CheckOutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOutTime

`func (o *ReservationModel) SetCheckOutTime(v time.Time)`

SetCheckOutTime sets CheckOutTime field to given value.

### HasCheckOutTime

`func (o *ReservationModel) HasCheckOutTime() bool`

HasCheckOutTime returns a boolean if a field has been set.

### GetCancellationTime

`func (o *ReservationModel) GetCancellationTime() time.Time`

GetCancellationTime returns the CancellationTime field if non-nil, zero value otherwise.

### GetCancellationTimeOk

`func (o *ReservationModel) GetCancellationTimeOk() (*time.Time, bool)`

GetCancellationTimeOk returns a tuple with the CancellationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationTime

`func (o *ReservationModel) SetCancellationTime(v time.Time)`

SetCancellationTime sets CancellationTime field to given value.

### HasCancellationTime

`func (o *ReservationModel) HasCancellationTime() bool`

HasCancellationTime returns a boolean if a field has been set.

### GetNoShowTime

`func (o *ReservationModel) GetNoShowTime() time.Time`

GetNoShowTime returns the NoShowTime field if non-nil, zero value otherwise.

### GetNoShowTimeOk

`func (o *ReservationModel) GetNoShowTimeOk() (*time.Time, bool)`

GetNoShowTimeOk returns a tuple with the NoShowTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShowTime

`func (o *ReservationModel) SetNoShowTime(v time.Time)`

SetNoShowTime sets NoShowTime field to given value.

### HasNoShowTime

`func (o *ReservationModel) HasNoShowTime() bool`

HasNoShowTime returns a boolean if a field has been set.

### GetUnit

`func (o *ReservationModel) GetUnit() EmbeddedUnitModel`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ReservationModel) GetUnitOk() (*EmbeddedUnitModel, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ReservationModel) SetUnit(v EmbeddedUnitModel)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ReservationModel) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetProperty

`func (o *ReservationModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ReservationModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ReservationModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetRatePlan

`func (o *ReservationModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *ReservationModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *ReservationModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetUnitGroup

`func (o *ReservationModel) GetUnitGroup() EmbeddedUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *ReservationModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *ReservationModel) SetUnitGroup(v EmbeddedUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetTotalGrossAmount

`func (o *ReservationModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *ReservationModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *ReservationModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetArrival

`func (o *ReservationModel) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *ReservationModel) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *ReservationModel) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *ReservationModel) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *ReservationModel) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *ReservationModel) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.


### GetCreated

`func (o *ReservationModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReservationModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReservationModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *ReservationModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ReservationModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ReservationModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAdults

`func (o *ReservationModel) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *ReservationModel) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *ReservationModel) SetAdults(v int32)`

SetAdults sets Adults field to given value.


### GetChildrenAges

`func (o *ReservationModel) GetChildrenAges() []int32`

GetChildrenAges returns the ChildrenAges field if non-nil, zero value otherwise.

### GetChildrenAgesOk

`func (o *ReservationModel) GetChildrenAgesOk() (*[]int32, bool)`

GetChildrenAgesOk returns a tuple with the ChildrenAges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAges

`func (o *ReservationModel) SetChildrenAges(v []int32)`

SetChildrenAges sets ChildrenAges field to given value.

### HasChildrenAges

`func (o *ReservationModel) HasChildrenAges() bool`

HasChildrenAges returns a boolean if a field has been set.

### GetComment

`func (o *ReservationModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReservationModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReservationModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReservationModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGuestComment

`func (o *ReservationModel) GetGuestComment() string`

GetGuestComment returns the GuestComment field if non-nil, zero value otherwise.

### GetGuestCommentOk

`func (o *ReservationModel) GetGuestCommentOk() (*string, bool)`

GetGuestCommentOk returns a tuple with the GuestComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestComment

`func (o *ReservationModel) SetGuestComment(v string)`

SetGuestComment sets GuestComment field to given value.

### HasGuestComment

`func (o *ReservationModel) HasGuestComment() bool`

HasGuestComment returns a boolean if a field has been set.

### GetExternalCode

`func (o *ReservationModel) GetExternalCode() string`

GetExternalCode returns the ExternalCode field if non-nil, zero value otherwise.

### GetExternalCodeOk

`func (o *ReservationModel) GetExternalCodeOk() (*string, bool)`

GetExternalCodeOk returns a tuple with the ExternalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCode

`func (o *ReservationModel) SetExternalCode(v string)`

SetExternalCode sets ExternalCode field to given value.

### HasExternalCode

`func (o *ReservationModel) HasExternalCode() bool`

HasExternalCode returns a boolean if a field has been set.

### GetChannelCode

`func (o *ReservationModel) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *ReservationModel) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *ReservationModel) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetSource

`func (o *ReservationModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReservationModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReservationModel) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReservationModel) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrimaryGuest

`func (o *ReservationModel) GetPrimaryGuest() GuestModel`

GetPrimaryGuest returns the PrimaryGuest field if non-nil, zero value otherwise.

### GetPrimaryGuestOk

`func (o *ReservationModel) GetPrimaryGuestOk() (*GuestModel, bool)`

GetPrimaryGuestOk returns a tuple with the PrimaryGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGuest

`func (o *ReservationModel) SetPrimaryGuest(v GuestModel)`

SetPrimaryGuest sets PrimaryGuest field to given value.

### HasPrimaryGuest

`func (o *ReservationModel) HasPrimaryGuest() bool`

HasPrimaryGuest returns a boolean if a field has been set.

### GetAdditionalGuests

`func (o *ReservationModel) GetAdditionalGuests() []GuestModel`

GetAdditionalGuests returns the AdditionalGuests field if non-nil, zero value otherwise.

### GetAdditionalGuestsOk

`func (o *ReservationModel) GetAdditionalGuestsOk() (*[]GuestModel, bool)`

GetAdditionalGuestsOk returns a tuple with the AdditionalGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGuests

`func (o *ReservationModel) SetAdditionalGuests(v []GuestModel)`

SetAdditionalGuests sets AdditionalGuests field to given value.

### HasAdditionalGuests

`func (o *ReservationModel) HasAdditionalGuests() bool`

HasAdditionalGuests returns a boolean if a field has been set.

### GetBooker

`func (o *ReservationModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *ReservationModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *ReservationModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.

### HasBooker

`func (o *ReservationModel) HasBooker() bool`

HasBooker returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *ReservationModel) GetPaymentAccount() PaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *ReservationModel) GetPaymentAccountOk() (*PaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *ReservationModel) SetPaymentAccount(v PaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *ReservationModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetTimeSlices

`func (o *ReservationModel) GetTimeSlices() []TimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *ReservationModel) GetTimeSlicesOk() (*[]TimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *ReservationModel) SetTimeSlices(v []TimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.

### HasTimeSlices

`func (o *ReservationModel) HasTimeSlices() bool`

HasTimeSlices returns a boolean if a field has been set.

### GetServices

`func (o *ReservationModel) GetServices() []ReservationServiceItemModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ReservationModel) GetServicesOk() (*[]ReservationServiceItemModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ReservationModel) SetServices(v []ReservationServiceItemModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *ReservationModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetGuaranteeType

`func (o *ReservationModel) GetGuaranteeType() string`

GetGuaranteeType returns the GuaranteeType field if non-nil, zero value otherwise.

### GetGuaranteeTypeOk

`func (o *ReservationModel) GetGuaranteeTypeOk() (*string, bool)`

GetGuaranteeTypeOk returns a tuple with the GuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeType

`func (o *ReservationModel) SetGuaranteeType(v string)`

SetGuaranteeType sets GuaranteeType field to given value.


### GetCancellationFee

`func (o *ReservationModel) GetCancellationFee() ReservationCancellationFeeModel`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *ReservationModel) GetCancellationFeeOk() (*ReservationCancellationFeeModel, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *ReservationModel) SetCancellationFee(v ReservationCancellationFeeModel)`

SetCancellationFee sets CancellationFee field to given value.


### GetNoShowFee

`func (o *ReservationModel) GetNoShowFee() ReservationNoShowFeeModel`

GetNoShowFee returns the NoShowFee field if non-nil, zero value otherwise.

### GetNoShowFeeOk

`func (o *ReservationModel) GetNoShowFeeOk() (*ReservationNoShowFeeModel, bool)`

GetNoShowFeeOk returns a tuple with the NoShowFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShowFee

`func (o *ReservationModel) SetNoShowFee(v ReservationNoShowFeeModel)`

SetNoShowFee sets NoShowFee field to given value.


### GetTravelPurpose

`func (o *ReservationModel) GetTravelPurpose() string`

GetTravelPurpose returns the TravelPurpose field if non-nil, zero value otherwise.

### GetTravelPurposeOk

`func (o *ReservationModel) GetTravelPurposeOk() (*string, bool)`

GetTravelPurposeOk returns a tuple with the TravelPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelPurpose

`func (o *ReservationModel) SetTravelPurpose(v string)`

SetTravelPurpose sets TravelPurpose field to given value.

### HasTravelPurpose

`func (o *ReservationModel) HasTravelPurpose() bool`

HasTravelPurpose returns a boolean if a field has been set.

### GetBalance

`func (o *ReservationModel) GetBalance() MonetaryValueModel`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ReservationModel) GetBalanceOk() (*MonetaryValueModel, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ReservationModel) SetBalance(v MonetaryValueModel)`

SetBalance sets Balance field to given value.


### GetAssignedUnits

`func (o *ReservationModel) GetAssignedUnits() []ReservationAssignedUnitModel`

GetAssignedUnits returns the AssignedUnits field if non-nil, zero value otherwise.

### GetAssignedUnitsOk

`func (o *ReservationModel) GetAssignedUnitsOk() (*[]ReservationAssignedUnitModel, bool)`

GetAssignedUnitsOk returns a tuple with the AssignedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUnits

`func (o *ReservationModel) SetAssignedUnits(v []ReservationAssignedUnitModel)`

SetAssignedUnits sets AssignedUnits field to given value.

### HasAssignedUnits

`func (o *ReservationModel) HasAssignedUnits() bool`

HasAssignedUnits returns a boolean if a field has been set.

### GetValidationMessages

`func (o *ReservationModel) GetValidationMessages() []ReservationValidationMessageModel`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *ReservationModel) GetValidationMessagesOk() (*[]ReservationValidationMessageModel, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *ReservationModel) SetValidationMessages(v []ReservationValidationMessageModel)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *ReservationModel) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.

### GetActions

`func (o *ReservationModel) GetActions() []ActionModelReservationActionNotAllowedReservationActionReason`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ReservationModel) GetActionsOk() (*[]ActionModelReservationActionNotAllowedReservationActionReason, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ReservationModel) SetActions(v []ActionModelReservationActionNotAllowedReservationActionReason)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ReservationModel) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCompany

`func (o *ReservationModel) GetCompany() EmbeddedCompanyModel`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ReservationModel) GetCompanyOk() (*EmbeddedCompanyModel, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ReservationModel) SetCompany(v EmbeddedCompanyModel)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ReservationModel) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCorporateCode

`func (o *ReservationModel) GetCorporateCode() string`

GetCorporateCode returns the CorporateCode field if non-nil, zero value otherwise.

### GetCorporateCodeOk

`func (o *ReservationModel) GetCorporateCodeOk() (*string, bool)`

GetCorporateCodeOk returns a tuple with the CorporateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCode

`func (o *ReservationModel) SetCorporateCode(v string)`

SetCorporateCode sets CorporateCode field to given value.

### HasCorporateCode

`func (o *ReservationModel) HasCorporateCode() bool`

HasCorporateCode returns a boolean if a field has been set.

### GetAllFoliosHaveInvoice

`func (o *ReservationModel) GetAllFoliosHaveInvoice() bool`

GetAllFoliosHaveInvoice returns the AllFoliosHaveInvoice field if non-nil, zero value otherwise.

### GetAllFoliosHaveInvoiceOk

`func (o *ReservationModel) GetAllFoliosHaveInvoiceOk() (*bool, bool)`

GetAllFoliosHaveInvoiceOk returns a tuple with the AllFoliosHaveInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFoliosHaveInvoice

`func (o *ReservationModel) SetAllFoliosHaveInvoice(v bool)`

SetAllFoliosHaveInvoice sets AllFoliosHaveInvoice field to given value.

### HasAllFoliosHaveInvoice

`func (o *ReservationModel) HasAllFoliosHaveInvoice() bool`

HasAllFoliosHaveInvoice returns a boolean if a field has been set.

### GetTaxDetails

`func (o *ReservationModel) GetTaxDetails() []TaxDetailModel`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *ReservationModel) GetTaxDetailsOk() (*[]TaxDetailModel, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *ReservationModel) SetTaxDetails(v []TaxDetailModel)`

SetTaxDetails sets TaxDetails field to given value.


### GetHasCityTax

`func (o *ReservationModel) GetHasCityTax() bool`

GetHasCityTax returns the HasCityTax field if non-nil, zero value otherwise.

### GetHasCityTaxOk

`func (o *ReservationModel) GetHasCityTaxOk() (*bool, bool)`

GetHasCityTaxOk returns a tuple with the HasCityTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCityTax

`func (o *ReservationModel) SetHasCityTax(v bool)`

SetHasCityTax sets HasCityTax field to given value.


### GetCommission

`func (o *ReservationModel) GetCommission() CommissionModel`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *ReservationModel) GetCommissionOk() (*CommissionModel, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *ReservationModel) SetCommission(v CommissionModel)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *ReservationModel) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetPromoCode

`func (o *ReservationModel) GetPromoCode() string`

GetPromoCode returns the PromoCode field if non-nil, zero value otherwise.

### GetPromoCodeOk

`func (o *ReservationModel) GetPromoCodeOk() (*string, bool)`

GetPromoCodeOk returns a tuple with the PromoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCode

`func (o *ReservationModel) SetPromoCode(v string)`

SetPromoCode sets PromoCode field to given value.

### HasPromoCode

`func (o *ReservationModel) HasPromoCode() bool`

HasPromoCode returns a boolean if a field has been set.

### GetPayableAmount

`func (o *ReservationModel) GetPayableAmount() PayableAmountModel`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *ReservationModel) GetPayableAmountOk() (*PayableAmountModel, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *ReservationModel) SetPayableAmount(v PayableAmountModel)`

SetPayableAmount sets PayableAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


