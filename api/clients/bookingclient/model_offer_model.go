/*
 * apaleo Booking API
 *
 * Resources and methods to manage guest journeys.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bookingclient

import (
	"encoding/json"
	"time"
)

// OfferModel struct for OfferModel
type OfferModel struct {
	// The earliest arrival date and time for this offer<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Arrival time.Time `json:"arrival"`
	// The latest departure date and time for this offer<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Departure time.Time `json:"departure"`
	UnitGroup OfferUnitGroupModel `json:"unitGroup"`
	// The minimum guarantee type for this offer
	MinGuaranteeType string `json:"minGuaranteeType"`
	// The number of available units for that offer
	AvailableUnits int32 `json:"availableUnits"`
	RatePlan EmbeddedRatePlanModel `json:"ratePlan"`
	TotalGrossAmount MonetaryValueModel `json:"totalGrossAmount"`
	CancellationFee OfferCancellationFeeModel `json:"cancellationFee"`
	NoShowFee OfferNoShowFeeModel `json:"noShowFee"`
	// The breakdown for each time slice for this offer
	TimeSlices []OfferTimeSliceModel `json:"timeSlices"`
	// The list of the mandatory services for this offer. Such services will be automatically booked when booking this offer
	Services *[]ServiceOfferModel `json:"services,omitempty"`
	// The details of the fees that will be added on top of the Apaleo.Api.Modules.Booking.Models.Offer.StayOffer.OfferModel.TotalGrossAmount when creating the booking
	Fees *[]OfferFeeModel `json:"fees,omitempty"`
	// Tax breakdown, displaying net and tax amount for each VAT type
	TaxDetails []TaxDetailModel `json:"taxDetails"`
	// Validation rules that were applied to the offer and show the reason why the offer is not bookable
	ValidationMessages *[]OfferValidationMessageModel `json:"validationMessages,omitempty"`
	// ID of the company the offer is created for
	CompanyId *string `json:"companyId,omitempty"`
	// The corporate rate code the offer is created for
	CorporateCode *string `json:"corporateCode,omitempty"`
	// Whether the offer is for a corporate rate plan
	IsCorporate bool `json:"isCorporate"`
	PrePaymentAmount MonetaryValueModel `json:"prePaymentAmount"`
	CityTax *AmountModel `json:"cityTax,omitempty"`
}

// NewOfferModel instantiates a new OfferModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferModel(arrival time.Time, departure time.Time, unitGroup OfferUnitGroupModel, minGuaranteeType string, availableUnits int32, ratePlan EmbeddedRatePlanModel, totalGrossAmount MonetaryValueModel, cancellationFee OfferCancellationFeeModel, noShowFee OfferNoShowFeeModel, timeSlices []OfferTimeSliceModel, taxDetails []TaxDetailModel, isCorporate bool, prePaymentAmount MonetaryValueModel) *OfferModel {
	this := OfferModel{}
	this.Arrival = arrival
	this.Departure = departure
	this.UnitGroup = unitGroup
	this.MinGuaranteeType = minGuaranteeType
	this.AvailableUnits = availableUnits
	this.RatePlan = ratePlan
	this.TotalGrossAmount = totalGrossAmount
	this.CancellationFee = cancellationFee
	this.NoShowFee = noShowFee
	this.TimeSlices = timeSlices
	this.TaxDetails = taxDetails
	this.IsCorporate = isCorporate
	this.PrePaymentAmount = prePaymentAmount
	return &this
}

// NewOfferModelWithDefaults instantiates a new OfferModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferModelWithDefaults() *OfferModel {
	this := OfferModel{}
	return &this
}

// GetArrival returns the Arrival field value
func (o *OfferModel) GetArrival() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Arrival
}

// GetArrivalOk returns a tuple with the Arrival field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetArrivalOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Arrival, true
}

// SetArrival sets field value
func (o *OfferModel) SetArrival(v time.Time) {
	o.Arrival = v
}

// GetDeparture returns the Departure field value
func (o *OfferModel) GetDeparture() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Departure
}

// GetDepartureOk returns a tuple with the Departure field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetDepartureOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Departure, true
}

// SetDeparture sets field value
func (o *OfferModel) SetDeparture(v time.Time) {
	o.Departure = v
}

// GetUnitGroup returns the UnitGroup field value
func (o *OfferModel) GetUnitGroup() OfferUnitGroupModel {
	if o == nil {
		var ret OfferUnitGroupModel
		return ret
	}

	return o.UnitGroup
}

// GetUnitGroupOk returns a tuple with the UnitGroup field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetUnitGroupOk() (*OfferUnitGroupModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnitGroup, true
}

// SetUnitGroup sets field value
func (o *OfferModel) SetUnitGroup(v OfferUnitGroupModel) {
	o.UnitGroup = v
}

// GetMinGuaranteeType returns the MinGuaranteeType field value
func (o *OfferModel) GetMinGuaranteeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinGuaranteeType
}

// GetMinGuaranteeTypeOk returns a tuple with the MinGuaranteeType field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetMinGuaranteeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinGuaranteeType, true
}

// SetMinGuaranteeType sets field value
func (o *OfferModel) SetMinGuaranteeType(v string) {
	o.MinGuaranteeType = v
}

// GetAvailableUnits returns the AvailableUnits field value
func (o *OfferModel) GetAvailableUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailableUnits
}

// GetAvailableUnitsOk returns a tuple with the AvailableUnits field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetAvailableUnitsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableUnits, true
}

// SetAvailableUnits sets field value
func (o *OfferModel) SetAvailableUnits(v int32) {
	o.AvailableUnits = v
}

// GetRatePlan returns the RatePlan field value
func (o *OfferModel) GetRatePlan() EmbeddedRatePlanModel {
	if o == nil {
		var ret EmbeddedRatePlanModel
		return ret
	}

	return o.RatePlan
}

// GetRatePlanOk returns a tuple with the RatePlan field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RatePlan, true
}

// SetRatePlan sets field value
func (o *OfferModel) SetRatePlan(v EmbeddedRatePlanModel) {
	o.RatePlan = v
}

// GetTotalGrossAmount returns the TotalGrossAmount field value
func (o *OfferModel) GetTotalGrossAmount() MonetaryValueModel {
	if o == nil {
		var ret MonetaryValueModel
		return ret
	}

	return o.TotalGrossAmount
}

// GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalGrossAmount, true
}

// SetTotalGrossAmount sets field value
func (o *OfferModel) SetTotalGrossAmount(v MonetaryValueModel) {
	o.TotalGrossAmount = v
}

// GetCancellationFee returns the CancellationFee field value
func (o *OfferModel) GetCancellationFee() OfferCancellationFeeModel {
	if o == nil {
		var ret OfferCancellationFeeModel
		return ret
	}

	return o.CancellationFee
}

// GetCancellationFeeOk returns a tuple with the CancellationFee field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetCancellationFeeOk() (*OfferCancellationFeeModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CancellationFee, true
}

// SetCancellationFee sets field value
func (o *OfferModel) SetCancellationFee(v OfferCancellationFeeModel) {
	o.CancellationFee = v
}

// GetNoShowFee returns the NoShowFee field value
func (o *OfferModel) GetNoShowFee() OfferNoShowFeeModel {
	if o == nil {
		var ret OfferNoShowFeeModel
		return ret
	}

	return o.NoShowFee
}

// GetNoShowFeeOk returns a tuple with the NoShowFee field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetNoShowFeeOk() (*OfferNoShowFeeModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NoShowFee, true
}

// SetNoShowFee sets field value
func (o *OfferModel) SetNoShowFee(v OfferNoShowFeeModel) {
	o.NoShowFee = v
}

// GetTimeSlices returns the TimeSlices field value
func (o *OfferModel) GetTimeSlices() []OfferTimeSliceModel {
	if o == nil {
		var ret []OfferTimeSliceModel
		return ret
	}

	return o.TimeSlices
}

// GetTimeSlicesOk returns a tuple with the TimeSlices field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetTimeSlicesOk() (*[]OfferTimeSliceModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeSlices, true
}

// SetTimeSlices sets field value
func (o *OfferModel) SetTimeSlices(v []OfferTimeSliceModel) {
	o.TimeSlices = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *OfferModel) GetServices() []ServiceOfferModel {
	if o == nil || o.Services == nil {
		var ret []ServiceOfferModel
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferModel) GetServicesOk() (*[]ServiceOfferModel, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *OfferModel) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []ServiceOfferModel and assigns it to the Services field.
func (o *OfferModel) SetServices(v []ServiceOfferModel) {
	o.Services = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *OfferModel) GetFees() []OfferFeeModel {
	if o == nil || o.Fees == nil {
		var ret []OfferFeeModel
		return ret
	}
	return *o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferModel) GetFeesOk() (*[]OfferFeeModel, bool) {
	if o == nil || o.Fees == nil {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *OfferModel) HasFees() bool {
	if o != nil && o.Fees != nil {
		return true
	}

	return false
}

// SetFees gets a reference to the given []OfferFeeModel and assigns it to the Fees field.
func (o *OfferModel) SetFees(v []OfferFeeModel) {
	o.Fees = &v
}

// GetTaxDetails returns the TaxDetails field value
func (o *OfferModel) GetTaxDetails() []TaxDetailModel {
	if o == nil {
		var ret []TaxDetailModel
		return ret
	}

	return o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetTaxDetailsOk() (*[]TaxDetailModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TaxDetails, true
}

// SetTaxDetails sets field value
func (o *OfferModel) SetTaxDetails(v []TaxDetailModel) {
	o.TaxDetails = v
}

// GetValidationMessages returns the ValidationMessages field value if set, zero value otherwise.
func (o *OfferModel) GetValidationMessages() []OfferValidationMessageModel {
	if o == nil || o.ValidationMessages == nil {
		var ret []OfferValidationMessageModel
		return ret
	}
	return *o.ValidationMessages
}

// GetValidationMessagesOk returns a tuple with the ValidationMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferModel) GetValidationMessagesOk() (*[]OfferValidationMessageModel, bool) {
	if o == nil || o.ValidationMessages == nil {
		return nil, false
	}
	return o.ValidationMessages, true
}

// HasValidationMessages returns a boolean if a field has been set.
func (o *OfferModel) HasValidationMessages() bool {
	if o != nil && o.ValidationMessages != nil {
		return true
	}

	return false
}

// SetValidationMessages gets a reference to the given []OfferValidationMessageModel and assigns it to the ValidationMessages field.
func (o *OfferModel) SetValidationMessages(v []OfferValidationMessageModel) {
	o.ValidationMessages = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *OfferModel) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferModel) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *OfferModel) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *OfferModel) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCorporateCode returns the CorporateCode field value if set, zero value otherwise.
func (o *OfferModel) GetCorporateCode() string {
	if o == nil || o.CorporateCode == nil {
		var ret string
		return ret
	}
	return *o.CorporateCode
}

// GetCorporateCodeOk returns a tuple with the CorporateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferModel) GetCorporateCodeOk() (*string, bool) {
	if o == nil || o.CorporateCode == nil {
		return nil, false
	}
	return o.CorporateCode, true
}

// HasCorporateCode returns a boolean if a field has been set.
func (o *OfferModel) HasCorporateCode() bool {
	if o != nil && o.CorporateCode != nil {
		return true
	}

	return false
}

// SetCorporateCode gets a reference to the given string and assigns it to the CorporateCode field.
func (o *OfferModel) SetCorporateCode(v string) {
	o.CorporateCode = &v
}

// GetIsCorporate returns the IsCorporate field value
func (o *OfferModel) GetIsCorporate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCorporate
}

// GetIsCorporateOk returns a tuple with the IsCorporate field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetIsCorporateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsCorporate, true
}

// SetIsCorporate sets field value
func (o *OfferModel) SetIsCorporate(v bool) {
	o.IsCorporate = v
}

// GetPrePaymentAmount returns the PrePaymentAmount field value
func (o *OfferModel) GetPrePaymentAmount() MonetaryValueModel {
	if o == nil {
		var ret MonetaryValueModel
		return ret
	}

	return o.PrePaymentAmount
}

// GetPrePaymentAmountOk returns a tuple with the PrePaymentAmount field value
// and a boolean to check if the value has been set.
func (o *OfferModel) GetPrePaymentAmountOk() (*MonetaryValueModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrePaymentAmount, true
}

// SetPrePaymentAmount sets field value
func (o *OfferModel) SetPrePaymentAmount(v MonetaryValueModel) {
	o.PrePaymentAmount = v
}

// GetCityTax returns the CityTax field value if set, zero value otherwise.
func (o *OfferModel) GetCityTax() AmountModel {
	if o == nil || o.CityTax == nil {
		var ret AmountModel
		return ret
	}
	return *o.CityTax
}

// GetCityTaxOk returns a tuple with the CityTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferModel) GetCityTaxOk() (*AmountModel, bool) {
	if o == nil || o.CityTax == nil {
		return nil, false
	}
	return o.CityTax, true
}

// HasCityTax returns a boolean if a field has been set.
func (o *OfferModel) HasCityTax() bool {
	if o != nil && o.CityTax != nil {
		return true
	}

	return false
}

// SetCityTax gets a reference to the given AmountModel and assigns it to the CityTax field.
func (o *OfferModel) SetCityTax(v AmountModel) {
	o.CityTax = &v
}

func (o OfferModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["arrival"] = o.Arrival
	}
	if true {
		toSerialize["departure"] = o.Departure
	}
	if true {
		toSerialize["unitGroup"] = o.UnitGroup
	}
	if true {
		toSerialize["minGuaranteeType"] = o.MinGuaranteeType
	}
	if true {
		toSerialize["availableUnits"] = o.AvailableUnits
	}
	if true {
		toSerialize["ratePlan"] = o.RatePlan
	}
	if true {
		toSerialize["totalGrossAmount"] = o.TotalGrossAmount
	}
	if true {
		toSerialize["cancellationFee"] = o.CancellationFee
	}
	if true {
		toSerialize["noShowFee"] = o.NoShowFee
	}
	if true {
		toSerialize["timeSlices"] = o.TimeSlices
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Fees != nil {
		toSerialize["fees"] = o.Fees
	}
	if true {
		toSerialize["taxDetails"] = o.TaxDetails
	}
	if o.ValidationMessages != nil {
		toSerialize["validationMessages"] = o.ValidationMessages
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.CorporateCode != nil {
		toSerialize["corporateCode"] = o.CorporateCode
	}
	if true {
		toSerialize["isCorporate"] = o.IsCorporate
	}
	if true {
		toSerialize["prePaymentAmount"] = o.PrePaymentAmount
	}
	if o.CityTax != nil {
		toSerialize["cityTax"] = o.CityTax
	}
	return json.Marshal(toSerialize)
}

type NullableOfferModel struct {
	value *OfferModel
	isSet bool
}

func (v NullableOfferModel) Get() *OfferModel {
	return v.value
}

func (v *NullableOfferModel) Set(val *OfferModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferModel(val *OfferModel) *NullableOfferModel {
	return &NullableOfferModel{value: val, isSet: true}
}

func (v NullableOfferModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


