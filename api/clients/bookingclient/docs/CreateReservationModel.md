# CreateReservationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arrival** | **string** | Date and optional time of arrival&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Departure** | **string** | Date and optional time of departure. Cannot be more than 5 years after arrival.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Adults** | **int32** | Number of adults | 
**ChildrenAges** | Pointer to **[]int32** | Ages of the children | [optional] 
**Comment** | Pointer to **string** | Additional information and comments | [optional] 
**GuestComment** | Pointer to **string** | Additional information and comments by the guest | [optional] 
**ExternalCode** | Pointer to **string** | Code in some system | [optional] 
**ChannelCode** | **string** | Channel code | 
**Source** | Pointer to **string** | Source of the reservation | [optional] 
**PrimaryGuest** | Pointer to [**GuestModel**](GuestModel.md) |  | [optional] 
**AdditionalGuests** | Pointer to [**[]GuestModel**](GuestModel.md) | Additional guests of the reservation. | [optional] 
**GuaranteeType** | Pointer to **string** | The guarantee that has to be applied for this reservation. It has to be the same or stronger than  the minimum guarantee required by the selected rate plan | [optional] 
**TravelPurpose** | Pointer to **string** | Purpose of the trip, leisure or business | [optional] 
**TimeSlices** | [**[]CreateReservationTimeSliceModel**](CreateReservationTimeSliceModel.md) | Gross prices including services and taxes for each time slice. They will be applied to the reservation timeslices  in the order specified from arrival to departure | 
**Services** | Pointer to [**[]BookReservationServiceModel**](BookReservationServiceModel.md) | Additional services (extras, add-ons) that should be added to the reservation | [optional] 
**CompanyId** | Pointer to **string** | Set this if this reservation belongs to a company | [optional] 
**CorporateCode** | Pointer to **string** | Corporate code provided during creation. Used to find offers during amend. | [optional] 
**PrePaymentAmount** | Pointer to [**MonetaryValueModel**](MonetaryValueModel.md) |  | [optional] 
**Commission** | Pointer to [**CommissionModel**](CommissionModel.md) |  | [optional] 
**PromoCode** | Pointer to **string** | The promo code associated with a certain special offer | [optional] 

## Methods

### NewCreateReservationModel

`func NewCreateReservationModel(arrival string, departure string, adults int32, channelCode string, timeSlices []CreateReservationTimeSliceModel, ) *CreateReservationModel`

NewCreateReservationModel instantiates a new CreateReservationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReservationModelWithDefaults

`func NewCreateReservationModelWithDefaults() *CreateReservationModel`

NewCreateReservationModelWithDefaults instantiates a new CreateReservationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrival

`func (o *CreateReservationModel) GetArrival() string`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *CreateReservationModel) GetArrivalOk() (*string, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *CreateReservationModel) SetArrival(v string)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *CreateReservationModel) GetDeparture() string`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *CreateReservationModel) GetDepartureOk() (*string, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *CreateReservationModel) SetDeparture(v string)`

SetDeparture sets Departure field to given value.


### GetAdults

`func (o *CreateReservationModel) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *CreateReservationModel) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *CreateReservationModel) SetAdults(v int32)`

SetAdults sets Adults field to given value.


### GetChildrenAges

`func (o *CreateReservationModel) GetChildrenAges() []int32`

GetChildrenAges returns the ChildrenAges field if non-nil, zero value otherwise.

### GetChildrenAgesOk

`func (o *CreateReservationModel) GetChildrenAgesOk() (*[]int32, bool)`

GetChildrenAgesOk returns a tuple with the ChildrenAges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAges

`func (o *CreateReservationModel) SetChildrenAges(v []int32)`

SetChildrenAges sets ChildrenAges field to given value.

### HasChildrenAges

`func (o *CreateReservationModel) HasChildrenAges() bool`

HasChildrenAges returns a boolean if a field has been set.

### GetComment

`func (o *CreateReservationModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateReservationModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateReservationModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateReservationModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGuestComment

`func (o *CreateReservationModel) GetGuestComment() string`

GetGuestComment returns the GuestComment field if non-nil, zero value otherwise.

### GetGuestCommentOk

`func (o *CreateReservationModel) GetGuestCommentOk() (*string, bool)`

GetGuestCommentOk returns a tuple with the GuestComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestComment

`func (o *CreateReservationModel) SetGuestComment(v string)`

SetGuestComment sets GuestComment field to given value.

### HasGuestComment

`func (o *CreateReservationModel) HasGuestComment() bool`

HasGuestComment returns a boolean if a field has been set.

### GetExternalCode

`func (o *CreateReservationModel) GetExternalCode() string`

GetExternalCode returns the ExternalCode field if non-nil, zero value otherwise.

### GetExternalCodeOk

`func (o *CreateReservationModel) GetExternalCodeOk() (*string, bool)`

GetExternalCodeOk returns a tuple with the ExternalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCode

`func (o *CreateReservationModel) SetExternalCode(v string)`

SetExternalCode sets ExternalCode field to given value.

### HasExternalCode

`func (o *CreateReservationModel) HasExternalCode() bool`

HasExternalCode returns a boolean if a field has been set.

### GetChannelCode

`func (o *CreateReservationModel) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *CreateReservationModel) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *CreateReservationModel) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetSource

`func (o *CreateReservationModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateReservationModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateReservationModel) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateReservationModel) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrimaryGuest

`func (o *CreateReservationModel) GetPrimaryGuest() GuestModel`

GetPrimaryGuest returns the PrimaryGuest field if non-nil, zero value otherwise.

### GetPrimaryGuestOk

`func (o *CreateReservationModel) GetPrimaryGuestOk() (*GuestModel, bool)`

GetPrimaryGuestOk returns a tuple with the PrimaryGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGuest

`func (o *CreateReservationModel) SetPrimaryGuest(v GuestModel)`

SetPrimaryGuest sets PrimaryGuest field to given value.

### HasPrimaryGuest

`func (o *CreateReservationModel) HasPrimaryGuest() bool`

HasPrimaryGuest returns a boolean if a field has been set.

### GetAdditionalGuests

`func (o *CreateReservationModel) GetAdditionalGuests() []GuestModel`

GetAdditionalGuests returns the AdditionalGuests field if non-nil, zero value otherwise.

### GetAdditionalGuestsOk

`func (o *CreateReservationModel) GetAdditionalGuestsOk() (*[]GuestModel, bool)`

GetAdditionalGuestsOk returns a tuple with the AdditionalGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGuests

`func (o *CreateReservationModel) SetAdditionalGuests(v []GuestModel)`

SetAdditionalGuests sets AdditionalGuests field to given value.

### HasAdditionalGuests

`func (o *CreateReservationModel) HasAdditionalGuests() bool`

HasAdditionalGuests returns a boolean if a field has been set.

### GetGuaranteeType

`func (o *CreateReservationModel) GetGuaranteeType() string`

GetGuaranteeType returns the GuaranteeType field if non-nil, zero value otherwise.

### GetGuaranteeTypeOk

`func (o *CreateReservationModel) GetGuaranteeTypeOk() (*string, bool)`

GetGuaranteeTypeOk returns a tuple with the GuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeType

`func (o *CreateReservationModel) SetGuaranteeType(v string)`

SetGuaranteeType sets GuaranteeType field to given value.

### HasGuaranteeType

`func (o *CreateReservationModel) HasGuaranteeType() bool`

HasGuaranteeType returns a boolean if a field has been set.

### GetTravelPurpose

`func (o *CreateReservationModel) GetTravelPurpose() string`

GetTravelPurpose returns the TravelPurpose field if non-nil, zero value otherwise.

### GetTravelPurposeOk

`func (o *CreateReservationModel) GetTravelPurposeOk() (*string, bool)`

GetTravelPurposeOk returns a tuple with the TravelPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelPurpose

`func (o *CreateReservationModel) SetTravelPurpose(v string)`

SetTravelPurpose sets TravelPurpose field to given value.

### HasTravelPurpose

`func (o *CreateReservationModel) HasTravelPurpose() bool`

HasTravelPurpose returns a boolean if a field has been set.

### GetTimeSlices

`func (o *CreateReservationModel) GetTimeSlices() []CreateReservationTimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *CreateReservationModel) GetTimeSlicesOk() (*[]CreateReservationTimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *CreateReservationModel) SetTimeSlices(v []CreateReservationTimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.


### GetServices

`func (o *CreateReservationModel) GetServices() []BookReservationServiceModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CreateReservationModel) GetServicesOk() (*[]BookReservationServiceModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CreateReservationModel) SetServices(v []BookReservationServiceModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *CreateReservationModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetCompanyId

`func (o *CreateReservationModel) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateReservationModel) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateReservationModel) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CreateReservationModel) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCorporateCode

`func (o *CreateReservationModel) GetCorporateCode() string`

GetCorporateCode returns the CorporateCode field if non-nil, zero value otherwise.

### GetCorporateCodeOk

`func (o *CreateReservationModel) GetCorporateCodeOk() (*string, bool)`

GetCorporateCodeOk returns a tuple with the CorporateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCode

`func (o *CreateReservationModel) SetCorporateCode(v string)`

SetCorporateCode sets CorporateCode field to given value.

### HasCorporateCode

`func (o *CreateReservationModel) HasCorporateCode() bool`

HasCorporateCode returns a boolean if a field has been set.

### GetPrePaymentAmount

`func (o *CreateReservationModel) GetPrePaymentAmount() MonetaryValueModel`

GetPrePaymentAmount returns the PrePaymentAmount field if non-nil, zero value otherwise.

### GetPrePaymentAmountOk

`func (o *CreateReservationModel) GetPrePaymentAmountOk() (*MonetaryValueModel, bool)`

GetPrePaymentAmountOk returns a tuple with the PrePaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaymentAmount

`func (o *CreateReservationModel) SetPrePaymentAmount(v MonetaryValueModel)`

SetPrePaymentAmount sets PrePaymentAmount field to given value.

### HasPrePaymentAmount

`func (o *CreateReservationModel) HasPrePaymentAmount() bool`

HasPrePaymentAmount returns a boolean if a field has been set.

### GetCommission

`func (o *CreateReservationModel) GetCommission() CommissionModel`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *CreateReservationModel) GetCommissionOk() (*CommissionModel, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *CreateReservationModel) SetCommission(v CommissionModel)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *CreateReservationModel) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetPromoCode

`func (o *CreateReservationModel) GetPromoCode() string`

GetPromoCode returns the PromoCode field if non-nil, zero value otherwise.

### GetPromoCodeOk

`func (o *CreateReservationModel) GetPromoCodeOk() (*string, bool)`

GetPromoCodeOk returns a tuple with the PromoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCode

`func (o *CreateReservationModel) SetPromoCode(v string)`

SetPromoCode sets PromoCode field to given value.

### HasPromoCode

`func (o *CreateReservationModel) HasPromoCode() bool`

HasPromoCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


