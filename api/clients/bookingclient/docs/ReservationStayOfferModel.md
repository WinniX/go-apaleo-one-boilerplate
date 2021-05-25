# ReservationStayOfferModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arrival** | **time.Time** | The earliest arrival date and time for this offer&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Departure** | **time.Time** | The latest departure date and time for this offer&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**MinGuaranteeType** | **string** | The minimum guarantee type for this offer | 
**AvailableUnits** | **int32** | The number of available units for that offer | 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**CancellationFee** | [**OfferCancellationFeeModel**](OfferCancellationFeeModel.md) |  | 
**NoShowFee** | [**OfferNoShowFeeModel**](OfferNoShowFeeModel.md) |  | 
**TimeSlices** | [**[]ReservationStayOfferTimeSliceModel**](ReservationStayOfferTimeSliceModel.md) | The breakdown for each time slice for this offer | 
**Services** | Pointer to [**[]ServiceOfferModel**](ServiceOfferModel.md) | The breakdown for extra services reserved for this offer | [optional] 
**TaxDetails** | [**[]TaxDetailModel**](TaxDetailModel.md) | Tax breakdown, displaying net and tax amount for each VAT type | 
**ValidationMessages** | Pointer to [**[]OfferValidationMessageModel**](OfferValidationMessageModel.md) | Validation rules that were applied to the offer and show the reason why the offer is not bookable | [optional] 
**CompanyId** | Pointer to **string** | ID of the company the offer is created for | [optional] 
**CorporateCode** | Pointer to **string** | The corporate rate code the offer is created for | [optional] 
**IsCorporate** | **bool** | Whether the offer is for a corporate rate plan | 
**CityTax** | Pointer to [**AmountModel**](AmountModel.md) |  | [optional] 

## Methods

### NewReservationStayOfferModel

`func NewReservationStayOfferModel(arrival time.Time, departure time.Time, minGuaranteeType string, availableUnits int32, totalGrossAmount MonetaryValueModel, cancellationFee OfferCancellationFeeModel, noShowFee OfferNoShowFeeModel, timeSlices []ReservationStayOfferTimeSliceModel, taxDetails []TaxDetailModel, isCorporate bool, ) *ReservationStayOfferModel`

NewReservationStayOfferModel instantiates a new ReservationStayOfferModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationStayOfferModelWithDefaults

`func NewReservationStayOfferModelWithDefaults() *ReservationStayOfferModel`

NewReservationStayOfferModelWithDefaults instantiates a new ReservationStayOfferModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrival

`func (o *ReservationStayOfferModel) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *ReservationStayOfferModel) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *ReservationStayOfferModel) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *ReservationStayOfferModel) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *ReservationStayOfferModel) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *ReservationStayOfferModel) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.


### GetMinGuaranteeType

`func (o *ReservationStayOfferModel) GetMinGuaranteeType() string`

GetMinGuaranteeType returns the MinGuaranteeType field if non-nil, zero value otherwise.

### GetMinGuaranteeTypeOk

`func (o *ReservationStayOfferModel) GetMinGuaranteeTypeOk() (*string, bool)`

GetMinGuaranteeTypeOk returns a tuple with the MinGuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGuaranteeType

`func (o *ReservationStayOfferModel) SetMinGuaranteeType(v string)`

SetMinGuaranteeType sets MinGuaranteeType field to given value.


### GetAvailableUnits

`func (o *ReservationStayOfferModel) GetAvailableUnits() int32`

GetAvailableUnits returns the AvailableUnits field if non-nil, zero value otherwise.

### GetAvailableUnitsOk

`func (o *ReservationStayOfferModel) GetAvailableUnitsOk() (*int32, bool)`

GetAvailableUnitsOk returns a tuple with the AvailableUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUnits

`func (o *ReservationStayOfferModel) SetAvailableUnits(v int32)`

SetAvailableUnits sets AvailableUnits field to given value.


### GetTotalGrossAmount

`func (o *ReservationStayOfferModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *ReservationStayOfferModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *ReservationStayOfferModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetCancellationFee

`func (o *ReservationStayOfferModel) GetCancellationFee() OfferCancellationFeeModel`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *ReservationStayOfferModel) GetCancellationFeeOk() (*OfferCancellationFeeModel, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *ReservationStayOfferModel) SetCancellationFee(v OfferCancellationFeeModel)`

SetCancellationFee sets CancellationFee field to given value.


### GetNoShowFee

`func (o *ReservationStayOfferModel) GetNoShowFee() OfferNoShowFeeModel`

GetNoShowFee returns the NoShowFee field if non-nil, zero value otherwise.

### GetNoShowFeeOk

`func (o *ReservationStayOfferModel) GetNoShowFeeOk() (*OfferNoShowFeeModel, bool)`

GetNoShowFeeOk returns a tuple with the NoShowFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShowFee

`func (o *ReservationStayOfferModel) SetNoShowFee(v OfferNoShowFeeModel)`

SetNoShowFee sets NoShowFee field to given value.


### GetTimeSlices

`func (o *ReservationStayOfferModel) GetTimeSlices() []ReservationStayOfferTimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *ReservationStayOfferModel) GetTimeSlicesOk() (*[]ReservationStayOfferTimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *ReservationStayOfferModel) SetTimeSlices(v []ReservationStayOfferTimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.


### GetServices

`func (o *ReservationStayOfferModel) GetServices() []ServiceOfferModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ReservationStayOfferModel) GetServicesOk() (*[]ServiceOfferModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ReservationStayOfferModel) SetServices(v []ServiceOfferModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *ReservationStayOfferModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTaxDetails

`func (o *ReservationStayOfferModel) GetTaxDetails() []TaxDetailModel`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *ReservationStayOfferModel) GetTaxDetailsOk() (*[]TaxDetailModel, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *ReservationStayOfferModel) SetTaxDetails(v []TaxDetailModel)`

SetTaxDetails sets TaxDetails field to given value.


### GetValidationMessages

`func (o *ReservationStayOfferModel) GetValidationMessages() []OfferValidationMessageModel`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *ReservationStayOfferModel) GetValidationMessagesOk() (*[]OfferValidationMessageModel, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *ReservationStayOfferModel) SetValidationMessages(v []OfferValidationMessageModel)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *ReservationStayOfferModel) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.

### GetCompanyId

`func (o *ReservationStayOfferModel) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ReservationStayOfferModel) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ReservationStayOfferModel) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ReservationStayOfferModel) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCorporateCode

`func (o *ReservationStayOfferModel) GetCorporateCode() string`

GetCorporateCode returns the CorporateCode field if non-nil, zero value otherwise.

### GetCorporateCodeOk

`func (o *ReservationStayOfferModel) GetCorporateCodeOk() (*string, bool)`

GetCorporateCodeOk returns a tuple with the CorporateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCode

`func (o *ReservationStayOfferModel) SetCorporateCode(v string)`

SetCorporateCode sets CorporateCode field to given value.

### HasCorporateCode

`func (o *ReservationStayOfferModel) HasCorporateCode() bool`

HasCorporateCode returns a boolean if a field has been set.

### GetIsCorporate

`func (o *ReservationStayOfferModel) GetIsCorporate() bool`

GetIsCorporate returns the IsCorporate field if non-nil, zero value otherwise.

### GetIsCorporateOk

`func (o *ReservationStayOfferModel) GetIsCorporateOk() (*bool, bool)`

GetIsCorporateOk returns a tuple with the IsCorporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCorporate

`func (o *ReservationStayOfferModel) SetIsCorporate(v bool)`

SetIsCorporate sets IsCorporate field to given value.


### GetCityTax

`func (o *ReservationStayOfferModel) GetCityTax() AmountModel`

GetCityTax returns the CityTax field if non-nil, zero value otherwise.

### GetCityTaxOk

`func (o *ReservationStayOfferModel) GetCityTaxOk() (*AmountModel, bool)`

GetCityTaxOk returns a tuple with the CityTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTax

`func (o *ReservationStayOfferModel) SetCityTax(v AmountModel)`

SetCityTax sets CityTax field to given value.

### HasCityTax

`func (o *ReservationStayOfferModel) HasCityTax() bool`

HasCityTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


