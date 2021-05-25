# BookingReservationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Reservation id | 
**Status** | **string** | Status of the reservation | 
**ExternalCode** | Pointer to **string** | Code in external system | [optional] 
**ChannelCode** | **string** | Channel code | 
**Source** | Pointer to **string** | Source of the reservation (e.g Hotels.com, Orbitz, etc.) | [optional] 
**PaymentAccount** | Pointer to [**PaymentAccountModel**](PaymentAccountModel.md) |  | [optional] 
**Arrival** | **time.Time** | Date of arrival&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Departure** | **time.Time** | Date of departure&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Adults** | **int32** | Number of adults | 
**ChildrenAges** | Pointer to **[]int32** | The ages of the children | [optional] 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**UnitGroup** | [**EmbeddedUnitGroupModel**](EmbeddedUnitGroupModel.md) |  | 
**Services** | Pointer to [**[]ReservationServiceItemModel**](ReservationServiceItemModel.md) | The list of additional services (extras, add-ons) reserved for the stay | [optional] 
**GuestComment** | Pointer to **string** | Additional information and comment by the guest | [optional] 
**CancellationFee** | [**ReservationCancellationFeeModel**](ReservationCancellationFeeModel.md) |  | 
**NoShowFee** | [**ReservationNoShowFeeModel**](ReservationNoShowFeeModel.md) |  | 
**Company** | Pointer to [**EmbeddedCompanyModel**](EmbeddedCompanyModel.md) |  | [optional] 

## Methods

### NewBookingReservationModel

`func NewBookingReservationModel(id string, status string, channelCode string, arrival time.Time, departure time.Time, adults int32, totalGrossAmount MonetaryValueModel, property EmbeddedPropertyModel, ratePlan EmbeddedRatePlanModel, unitGroup EmbeddedUnitGroupModel, cancellationFee ReservationCancellationFeeModel, noShowFee ReservationNoShowFeeModel, ) *BookingReservationModel`

NewBookingReservationModel instantiates a new BookingReservationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingReservationModelWithDefaults

`func NewBookingReservationModelWithDefaults() *BookingReservationModel`

NewBookingReservationModelWithDefaults instantiates a new BookingReservationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookingReservationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookingReservationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookingReservationModel) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BookingReservationModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingReservationModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingReservationModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetExternalCode

`func (o *BookingReservationModel) GetExternalCode() string`

GetExternalCode returns the ExternalCode field if non-nil, zero value otherwise.

### GetExternalCodeOk

`func (o *BookingReservationModel) GetExternalCodeOk() (*string, bool)`

GetExternalCodeOk returns a tuple with the ExternalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCode

`func (o *BookingReservationModel) SetExternalCode(v string)`

SetExternalCode sets ExternalCode field to given value.

### HasExternalCode

`func (o *BookingReservationModel) HasExternalCode() bool`

HasExternalCode returns a boolean if a field has been set.

### GetChannelCode

`func (o *BookingReservationModel) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *BookingReservationModel) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *BookingReservationModel) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetSource

`func (o *BookingReservationModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BookingReservationModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BookingReservationModel) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BookingReservationModel) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *BookingReservationModel) GetPaymentAccount() PaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *BookingReservationModel) GetPaymentAccountOk() (*PaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *BookingReservationModel) SetPaymentAccount(v PaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *BookingReservationModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetArrival

`func (o *BookingReservationModel) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *BookingReservationModel) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *BookingReservationModel) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *BookingReservationModel) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *BookingReservationModel) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *BookingReservationModel) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.


### GetAdults

`func (o *BookingReservationModel) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *BookingReservationModel) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *BookingReservationModel) SetAdults(v int32)`

SetAdults sets Adults field to given value.


### GetChildrenAges

`func (o *BookingReservationModel) GetChildrenAges() []int32`

GetChildrenAges returns the ChildrenAges field if non-nil, zero value otherwise.

### GetChildrenAgesOk

`func (o *BookingReservationModel) GetChildrenAgesOk() (*[]int32, bool)`

GetChildrenAgesOk returns a tuple with the ChildrenAges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAges

`func (o *BookingReservationModel) SetChildrenAges(v []int32)`

SetChildrenAges sets ChildrenAges field to given value.

### HasChildrenAges

`func (o *BookingReservationModel) HasChildrenAges() bool`

HasChildrenAges returns a boolean if a field has been set.

### GetTotalGrossAmount

`func (o *BookingReservationModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *BookingReservationModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *BookingReservationModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetProperty

`func (o *BookingReservationModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *BookingReservationModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *BookingReservationModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetRatePlan

`func (o *BookingReservationModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *BookingReservationModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *BookingReservationModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetUnitGroup

`func (o *BookingReservationModel) GetUnitGroup() EmbeddedUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *BookingReservationModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *BookingReservationModel) SetUnitGroup(v EmbeddedUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetServices

`func (o *BookingReservationModel) GetServices() []ReservationServiceItemModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BookingReservationModel) GetServicesOk() (*[]ReservationServiceItemModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BookingReservationModel) SetServices(v []ReservationServiceItemModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *BookingReservationModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetGuestComment

`func (o *BookingReservationModel) GetGuestComment() string`

GetGuestComment returns the GuestComment field if non-nil, zero value otherwise.

### GetGuestCommentOk

`func (o *BookingReservationModel) GetGuestCommentOk() (*string, bool)`

GetGuestCommentOk returns a tuple with the GuestComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestComment

`func (o *BookingReservationModel) SetGuestComment(v string)`

SetGuestComment sets GuestComment field to given value.

### HasGuestComment

`func (o *BookingReservationModel) HasGuestComment() bool`

HasGuestComment returns a boolean if a field has been set.

### GetCancellationFee

`func (o *BookingReservationModel) GetCancellationFee() ReservationCancellationFeeModel`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *BookingReservationModel) GetCancellationFeeOk() (*ReservationCancellationFeeModel, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *BookingReservationModel) SetCancellationFee(v ReservationCancellationFeeModel)`

SetCancellationFee sets CancellationFee field to given value.


### GetNoShowFee

`func (o *BookingReservationModel) GetNoShowFee() ReservationNoShowFeeModel`

GetNoShowFee returns the NoShowFee field if non-nil, zero value otherwise.

### GetNoShowFeeOk

`func (o *BookingReservationModel) GetNoShowFeeOk() (*ReservationNoShowFeeModel, bool)`

GetNoShowFeeOk returns a tuple with the NoShowFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShowFee

`func (o *BookingReservationModel) SetNoShowFee(v ReservationNoShowFeeModel)`

SetNoShowFee sets NoShowFee field to given value.


### GetCompany

`func (o *BookingReservationModel) GetCompany() EmbeddedCompanyModel`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BookingReservationModel) GetCompanyOk() (*EmbeddedCompanyModel, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BookingReservationModel) SetCompany(v EmbeddedCompanyModel)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BookingReservationModel) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


