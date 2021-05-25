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
)

// CreateBlockModel struct for CreateBlockModel
type CreateBlockModel struct {
	// ID of the group that reserved the block
	GroupId string `json:"groupId"`
	// The rate plan
	RatePlanId string `json:"ratePlanId"`
	// Start date and time from which the inventory will be blocked<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	From string `json:"from"`
	// End date and time until which the inventory will be blocked. Cannot be more than 5 years after the start date.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	To string `json:"to"`
	GrossDailyRate MonetaryValueModel `json:"grossDailyRate"`
	// The list of blocked units for each time slice
	TimeSlices *[]CreateBlockTimeSliceModel `json:"timeSlices,omitempty"`
	// Number of units to block for the defined time period
	BlockedUnits *int32 `json:"blockedUnits,omitempty"`
	// The promo code associated with a certain special offer
	PromoCode *string `json:"promoCode,omitempty"`
	// The corporate code associated with a certain special offer
	CorporateCode *string `json:"corporateCode,omitempty"`
}

// NewCreateBlockModel instantiates a new CreateBlockModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBlockModel(groupId string, ratePlanId string, from string, to string, grossDailyRate MonetaryValueModel) *CreateBlockModel {
	this := CreateBlockModel{}
	this.GroupId = groupId
	this.RatePlanId = ratePlanId
	this.From = from
	this.To = to
	this.GrossDailyRate = grossDailyRate
	return &this
}

// NewCreateBlockModelWithDefaults instantiates a new CreateBlockModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBlockModelWithDefaults() *CreateBlockModel {
	this := CreateBlockModel{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *CreateBlockModel) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *CreateBlockModel) SetGroupId(v string) {
	o.GroupId = v
}

// GetRatePlanId returns the RatePlanId field value
func (o *CreateBlockModel) GetRatePlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RatePlanId
}

// GetRatePlanIdOk returns a tuple with the RatePlanId field value
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetRatePlanIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RatePlanId, true
}

// SetRatePlanId sets field value
func (o *CreateBlockModel) SetRatePlanId(v string) {
	o.RatePlanId = v
}

// GetFrom returns the From field value
func (o *CreateBlockModel) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetFromOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CreateBlockModel) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *CreateBlockModel) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *CreateBlockModel) SetTo(v string) {
	o.To = v
}

// GetGrossDailyRate returns the GrossDailyRate field value
func (o *CreateBlockModel) GetGrossDailyRate() MonetaryValueModel {
	if o == nil {
		var ret MonetaryValueModel
		return ret
	}

	return o.GrossDailyRate
}

// GetGrossDailyRateOk returns a tuple with the GrossDailyRate field value
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetGrossDailyRateOk() (*MonetaryValueModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GrossDailyRate, true
}

// SetGrossDailyRate sets field value
func (o *CreateBlockModel) SetGrossDailyRate(v MonetaryValueModel) {
	o.GrossDailyRate = v
}

// GetTimeSlices returns the TimeSlices field value if set, zero value otherwise.
func (o *CreateBlockModel) GetTimeSlices() []CreateBlockTimeSliceModel {
	if o == nil || o.TimeSlices == nil {
		var ret []CreateBlockTimeSliceModel
		return ret
	}
	return *o.TimeSlices
}

// GetTimeSlicesOk returns a tuple with the TimeSlices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetTimeSlicesOk() (*[]CreateBlockTimeSliceModel, bool) {
	if o == nil || o.TimeSlices == nil {
		return nil, false
	}
	return o.TimeSlices, true
}

// HasTimeSlices returns a boolean if a field has been set.
func (o *CreateBlockModel) HasTimeSlices() bool {
	if o != nil && o.TimeSlices != nil {
		return true
	}

	return false
}

// SetTimeSlices gets a reference to the given []CreateBlockTimeSliceModel and assigns it to the TimeSlices field.
func (o *CreateBlockModel) SetTimeSlices(v []CreateBlockTimeSliceModel) {
	o.TimeSlices = &v
}

// GetBlockedUnits returns the BlockedUnits field value if set, zero value otherwise.
func (o *CreateBlockModel) GetBlockedUnits() int32 {
	if o == nil || o.BlockedUnits == nil {
		var ret int32
		return ret
	}
	return *o.BlockedUnits
}

// GetBlockedUnitsOk returns a tuple with the BlockedUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetBlockedUnitsOk() (*int32, bool) {
	if o == nil || o.BlockedUnits == nil {
		return nil, false
	}
	return o.BlockedUnits, true
}

// HasBlockedUnits returns a boolean if a field has been set.
func (o *CreateBlockModel) HasBlockedUnits() bool {
	if o != nil && o.BlockedUnits != nil {
		return true
	}

	return false
}

// SetBlockedUnits gets a reference to the given int32 and assigns it to the BlockedUnits field.
func (o *CreateBlockModel) SetBlockedUnits(v int32) {
	o.BlockedUnits = &v
}

// GetPromoCode returns the PromoCode field value if set, zero value otherwise.
func (o *CreateBlockModel) GetPromoCode() string {
	if o == nil || o.PromoCode == nil {
		var ret string
		return ret
	}
	return *o.PromoCode
}

// GetPromoCodeOk returns a tuple with the PromoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetPromoCodeOk() (*string, bool) {
	if o == nil || o.PromoCode == nil {
		return nil, false
	}
	return o.PromoCode, true
}

// HasPromoCode returns a boolean if a field has been set.
func (o *CreateBlockModel) HasPromoCode() bool {
	if o != nil && o.PromoCode != nil {
		return true
	}

	return false
}

// SetPromoCode gets a reference to the given string and assigns it to the PromoCode field.
func (o *CreateBlockModel) SetPromoCode(v string) {
	o.PromoCode = &v
}

// GetCorporateCode returns the CorporateCode field value if set, zero value otherwise.
func (o *CreateBlockModel) GetCorporateCode() string {
	if o == nil || o.CorporateCode == nil {
		var ret string
		return ret
	}
	return *o.CorporateCode
}

// GetCorporateCodeOk returns a tuple with the CorporateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBlockModel) GetCorporateCodeOk() (*string, bool) {
	if o == nil || o.CorporateCode == nil {
		return nil, false
	}
	return o.CorporateCode, true
}

// HasCorporateCode returns a boolean if a field has been set.
func (o *CreateBlockModel) HasCorporateCode() bool {
	if o != nil && o.CorporateCode != nil {
		return true
	}

	return false
}

// SetCorporateCode gets a reference to the given string and assigns it to the CorporateCode field.
func (o *CreateBlockModel) SetCorporateCode(v string) {
	o.CorporateCode = &v
}

func (o CreateBlockModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["ratePlanId"] = o.RatePlanId
	}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["grossDailyRate"] = o.GrossDailyRate
	}
	if o.TimeSlices != nil {
		toSerialize["timeSlices"] = o.TimeSlices
	}
	if o.BlockedUnits != nil {
		toSerialize["blockedUnits"] = o.BlockedUnits
	}
	if o.PromoCode != nil {
		toSerialize["promoCode"] = o.PromoCode
	}
	if o.CorporateCode != nil {
		toSerialize["corporateCode"] = o.CorporateCode
	}
	return json.Marshal(toSerialize)
}

type NullableCreateBlockModel struct {
	value *CreateBlockModel
	isSet bool
}

func (v NullableCreateBlockModel) Get() *CreateBlockModel {
	return v.value
}

func (v *NullableCreateBlockModel) Set(val *CreateBlockModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBlockModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBlockModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBlockModel(val *CreateBlockModel) *NullableCreateBlockModel {
	return &NullableCreateBlockModel{value: val, isSet: true}
}

func (v NullableCreateBlockModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBlockModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


