# OfferModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arrival** | **time.Time** | The earliest arrival date and time for this offer&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Departure** | **time.Time** | The latest departure date and time for this offer&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**UnitGroup** | [**OfferUnitGroupModel**](OfferUnitGroupModel.md) |  | 
**MinGuaranteeType** | **string** | The minimum guarantee type for this offer | 
**AvailableUnits** | **int32** | The number of available units for that offer | 
**RatePlan** | [**EmbeddedRatePlanModel**](EmbeddedRatePlanModel.md) |  | 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**CancellationFee** | [**OfferCancellationFeeModel**](OfferCancellationFeeModel.md) |  | 
**NoShowFee** | [**OfferNoShowFeeModel**](OfferNoShowFeeModel.md) |  | 
**TimeSlices** | [**[]OfferTimeSliceModel**](OfferTimeSliceModel.md) | The breakdown for each time slice for this offer | 
**Services** | Pointer to [**[]ServiceOfferModel**](ServiceOfferModel.md) | The list of the mandatory services for this offer. Such services will be automatically booked when booking this offer | [optional] 
**Fees** | Pointer to [**[]OfferFeeModel**](OfferFeeModel.md) | The details of the fees that will be added on top of the Apaleo.Api.Modules.Booking.Models.Offer.StayOffer.OfferModel.TotalGrossAmount when creating the booking | [optional] 
**TaxDetails** | [**[]TaxDetailModel**](TaxDetailModel.md) | Tax breakdown, displaying net and tax amount for each VAT type | 
**ValidationMessages** | Pointer to [**[]OfferValidationMessageModel**](OfferValidationMessageModel.md) | Validation rules that were applied to the offer and show the reason why the offer is not bookable | [optional] 
**CompanyId** | Pointer to **string** | ID of the company the offer is created for | [optional] 
**CorporateCode** | Pointer to **string** | The corporate rate code the offer is created for | [optional] 
**IsCorporate** | **bool** | Whether the offer is for a corporate rate plan | 
**PrePaymentAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**CityTax** | Pointer to [**AmountModel**](AmountModel.md) |  | [optional] 

## Methods

### NewOfferModel

`func NewOfferModel(arrival time.Time, departure time.Time, unitGroup OfferUnitGroupModel, minGuaranteeType string, availableUnits int32, ratePlan EmbeddedRatePlanModel, totalGrossAmount MonetaryValueModel, cancellationFee OfferCancellationFeeModel, noShowFee OfferNoShowFeeModel, timeSlices []OfferTimeSliceModel, taxDetails []TaxDetailModel, isCorporate bool, prePaymentAmount MonetaryValueModel, ) *OfferModel`

NewOfferModel instantiates a new OfferModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferModelWithDefaults

`func NewOfferModelWithDefaults() *OfferModel`

NewOfferModelWithDefaults instantiates a new OfferModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrival

`func (o *OfferModel) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *OfferModel) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *OfferModel) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *OfferModel) GetDeparture() time.Time`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *OfferModel) GetDepartureOk() (*time.Time, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *OfferModel) SetDeparture(v time.Time)`

SetDeparture sets Departure field to given value.


### GetUnitGroup

`func (o *OfferModel) GetUnitGroup() OfferUnitGroupModel`

GetUnitGroup returns the UnitGroup field if non-nil, zero value otherwise.

### GetUnitGroupOk

`func (o *OfferModel) GetUnitGroupOk() (*OfferUnitGroupModel, bool)`

GetUnitGroupOk returns a tuple with the UnitGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroup

`func (o *OfferModel) SetUnitGroup(v OfferUnitGroupModel)`

SetUnitGroup sets UnitGroup field to given value.


### GetMinGuaranteeType

`func (o *OfferModel) GetMinGuaranteeType() string`

GetMinGuaranteeType returns the MinGuaranteeType field if non-nil, zero value otherwise.

### GetMinGuaranteeTypeOk

`func (o *OfferModel) GetMinGuaranteeTypeOk() (*string, bool)`

GetMinGuaranteeTypeOk returns a tuple with the MinGuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGuaranteeType

`func (o *OfferModel) SetMinGuaranteeType(v string)`

SetMinGuaranteeType sets MinGuaranteeType field to given value.


### GetAvailableUnits

`func (o *OfferModel) GetAvailableUnits() int32`

GetAvailableUnits returns the AvailableUnits field if non-nil, zero value otherwise.

### GetAvailableUnitsOk

`func (o *OfferModel) GetAvailableUnitsOk() (*int32, bool)`

GetAvailableUnitsOk returns a tuple with the AvailableUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUnits

`func (o *OfferModel) SetAvailableUnits(v int32)`

SetAvailableUnits sets AvailableUnits field to given value.


### GetRatePlan

`func (o *OfferModel) GetRatePlan() EmbeddedRatePlanModel`

GetRatePlan returns the RatePlan field if non-nil, zero value otherwise.

### GetRatePlanOk

`func (o *OfferModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool)`

GetRatePlanOk returns a tuple with the RatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlan

`func (o *OfferModel) SetRatePlan(v EmbeddedRatePlanModel)`

SetRatePlan sets RatePlan field to given value.


### GetTotalGrossAmount

`func (o *OfferModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *OfferModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *OfferModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetCancellationFee

`func (o *OfferModel) GetCancellationFee() OfferCancellationFeeModel`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *OfferModel) GetCancellationFeeOk() (*OfferCancellationFeeModel, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *OfferModel) SetCancellationFee(v OfferCancellationFeeModel)`

SetCancellationFee sets CancellationFee field to given value.


### GetNoShowFee

`func (o *OfferModel) GetNoShowFee() OfferNoShowFeeModel`

GetNoShowFee returns the NoShowFee field if non-nil, zero value otherwise.

### GetNoShowFeeOk

`func (o *OfferModel) GetNoShowFeeOk() (*OfferNoShowFeeModel, bool)`

GetNoShowFeeOk returns a tuple with the NoShowFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoShowFee

`func (o *OfferModel) SetNoShowFee(v OfferNoShowFeeModel)`

SetNoShowFee sets NoShowFee field to given value.


### GetTimeSlices

`func (o *OfferModel) GetTimeSlices() []OfferTimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *OfferModel) GetTimeSlicesOk() (*[]OfferTimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *OfferModel) SetTimeSlices(v []OfferTimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.


### GetServices

`func (o *OfferModel) GetServices() []ServiceOfferModel`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OfferModel) GetServicesOk() (*[]ServiceOfferModel, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OfferModel) SetServices(v []ServiceOfferModel)`

SetServices sets Services field to given value.

### HasServices

`func (o *OfferModel) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetFees

`func (o *OfferModel) GetFees() []OfferFeeModel`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *OfferModel) GetFeesOk() (*[]OfferFeeModel, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *OfferModel) SetFees(v []OfferFeeModel)`

SetFees sets Fees field to given value.

### HasFees

`func (o *OfferModel) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetTaxDetails

`func (o *OfferModel) GetTaxDetails() []TaxDetailModel`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *OfferModel) GetTaxDetailsOk() (*[]TaxDetailModel, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *OfferModel) SetTaxDetails(v []TaxDetailModel)`

SetTaxDetails sets TaxDetails field to given value.


### GetValidationMessages

`func (o *OfferModel) GetValidationMessages() []OfferValidationMessageModel`

GetValidationMessages returns the ValidationMessages field if non-nil, zero value otherwise.

### GetValidationMessagesOk

`func (o *OfferModel) GetValidationMessagesOk() (*[]OfferValidationMessageModel, bool)`

GetValidationMessagesOk returns a tuple with the ValidationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessages

`func (o *OfferModel) SetValidationMessages(v []OfferValidationMessageModel)`

SetValidationMessages sets ValidationMessages field to given value.

### HasValidationMessages

`func (o *OfferModel) HasValidationMessages() bool`

HasValidationMessages returns a boolean if a field has been set.

### GetCompanyId

`func (o *OfferModel) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *OfferModel) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *OfferModel) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *OfferModel) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCorporateCode

`func (o *OfferModel) GetCorporateCode() string`

GetCorporateCode returns the CorporateCode field if non-nil, zero value otherwise.

### GetCorporateCodeOk

`func (o *OfferModel) GetCorporateCodeOk() (*string, bool)`

GetCorporateCodeOk returns a tuple with the CorporateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCode

`func (o *OfferModel) SetCorporateCode(v string)`

SetCorporateCode sets CorporateCode field to given value.

### HasCorporateCode

`func (o *OfferModel) HasCorporateCode() bool`

HasCorporateCode returns a boolean if a field has been set.

### GetIsCorporate

`func (o *OfferModel) GetIsCorporate() bool`

GetIsCorporate returns the IsCorporate field if non-nil, zero value otherwise.

### GetIsCorporateOk

`func (o *OfferModel) GetIsCorporateOk() (*bool, bool)`

GetIsCorporateOk returns a tuple with the IsCorporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCorporate

`func (o *OfferModel) SetIsCorporate(v bool)`

SetIsCorporate sets IsCorporate field to given value.


### GetPrePaymentAmount

`func (o *OfferModel) GetPrePaymentAmount() MonetaryValueModel`

GetPrePaymentAmount returns the PrePaymentAmount field if non-nil, zero value otherwise.

### GetPrePaymentAmountOk

`func (o *OfferModel) GetPrePaymentAmountOk() (*MonetaryValueModel, bool)`

GetPrePaymentAmountOk returns a tuple with the PrePaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaymentAmount

`func (o *OfferModel) SetPrePaymentAmount(v MonetaryValueModel)`

SetPrePaymentAmount sets PrePaymentAmount field to given value.


### GetCityTax

`func (o *OfferModel) GetCityTax() AmountModel`

GetCityTax returns the CityTax field if non-nil, zero value otherwise.

### GetCityTaxOk

`func (o *OfferModel) GetCityTaxOk() (*AmountModel, bool)`

GetCityTaxOk returns a tuple with the CityTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTax

`func (o *OfferModel) SetCityTax(v AmountModel)`

SetCityTax sets CityTax field to given value.

### HasCityTax

`func (o *OfferModel) HasCityTax() bool`

HasCityTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


