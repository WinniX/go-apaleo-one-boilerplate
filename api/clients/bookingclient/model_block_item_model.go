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

// BlockItemModel struct for BlockItemModel
type BlockItemModel struct {
	// Block id
	Id string `json:"id"`
	Group EmbeddedGroupModel `json:"group"`
	// Status of the block
	Status string `json:"status"`
	Property EmbeddedPropertyModel `json:"property"`
	RatePlan EmbeddedRatePlanModel `json:"ratePlan"`
	UnitGroup EmbeddedUnitGroupModel `json:"unitGroup"`
	GrossDailyRate MonetaryValueModel `json:"grossDailyRate"`
	// Start date and time from which the inventory will be blocked<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	From time.Time `json:"from"`
	// End date and time until which the inventory will be blocked<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	To time.Time `json:"to"`
	// Number of reservations already picked from this block
	PickedReservations int32 `json:"pickedReservations"`
	// The promo code associated with a certain special offer used to create the block
	PromoCode *string `json:"promoCode,omitempty"`
	// The corporate code associated with a certain special offer used to create the block
	CorporateCode *string `json:"corporateCode,omitempty"`
	// Date of creation<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Created time.Time `json:"created"`
	// Date of last modification<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Modified time.Time `json:"modified"`
	// The list of blocked units for each time slice
	TimeSlices *[]BlockTimeSliceModel `json:"timeSlices,omitempty"`
	// The list of actions for this block
	Actions *[]ActionModelBlockActionNotAllowedBlockActionReason `json:"actions,omitempty"`
}

// NewBlockItemModel instantiates a new BlockItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockItemModel(id string, group EmbeddedGroupModel, status string, property EmbeddedPropertyModel, ratePlan EmbeddedRatePlanModel, unitGroup EmbeddedUnitGroupModel, grossDailyRate MonetaryValueModel, from time.Time, to time.Time, pickedReservations int32, created time.Time, modified time.Time) *BlockItemModel {
	this := BlockItemModel{}
	this.Id = id
	this.Group = group
	this.Status = status
	this.Property = property
	this.RatePlan = ratePlan
	this.UnitGroup = unitGroup
	this.GrossDailyRate = grossDailyRate
	this.From = from
	this.To = to
	this.PickedReservations = pickedReservations
	this.Created = created
	this.Modified = modified
	return &this
}

// NewBlockItemModelWithDefaults instantiates a new BlockItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockItemModelWithDefaults() *BlockItemModel {
	this := BlockItemModel{}
	return &this
}

// GetId returns the Id field value
func (o *BlockItemModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BlockItemModel) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *BlockItemModel) GetGroup() EmbeddedGroupModel {
	if o == nil {
		var ret EmbeddedGroupModel
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetGroupOk() (*EmbeddedGroupModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *BlockItemModel) SetGroup(v EmbeddedGroupModel) {
	o.Group = v
}

// GetStatus returns the Status field value
func (o *BlockItemModel) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BlockItemModel) SetStatus(v string) {
	o.Status = v
}

// GetProperty returns the Property field value
func (o *BlockItemModel) GetProperty() EmbeddedPropertyModel {
	if o == nil {
		var ret EmbeddedPropertyModel
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetPropertyOk() (*EmbeddedPropertyModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *BlockItemModel) SetProperty(v EmbeddedPropertyModel) {
	o.Property = v
}

// GetRatePlan returns the RatePlan field value
func (o *BlockItemModel) GetRatePlan() EmbeddedRatePlanModel {
	if o == nil {
		var ret EmbeddedRatePlanModel
		return ret
	}

	return o.RatePlan
}

// GetRatePlanOk returns a tuple with the RatePlan field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetRatePlanOk() (*EmbeddedRatePlanModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RatePlan, true
}

// SetRatePlan sets field value
func (o *BlockItemModel) SetRatePlan(v EmbeddedRatePlanModel) {
	o.RatePlan = v
}

// GetUnitGroup returns the UnitGroup field value
func (o *BlockItemModel) GetUnitGroup() EmbeddedUnitGroupModel {
	if o == nil {
		var ret EmbeddedUnitGroupModel
		return ret
	}

	return o.UnitGroup
}

// GetUnitGroupOk returns a tuple with the UnitGroup field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnitGroup, true
}

// SetUnitGroup sets field value
func (o *BlockItemModel) SetUnitGroup(v EmbeddedUnitGroupModel) {
	o.UnitGroup = v
}

// GetGrossDailyRate returns the GrossDailyRate field value
func (o *BlockItemModel) GetGrossDailyRate() MonetaryValueModel {
	if o == nil {
		var ret MonetaryValueModel
		return ret
	}

	return o.GrossDailyRate
}

// GetGrossDailyRateOk returns a tuple with the GrossDailyRate field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetGrossDailyRateOk() (*MonetaryValueModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GrossDailyRate, true
}

// SetGrossDailyRate sets field value
func (o *BlockItemModel) SetGrossDailyRate(v MonetaryValueModel) {
	o.GrossDailyRate = v
}

// GetFrom returns the From field value
func (o *BlockItemModel) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetFromOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *BlockItemModel) SetFrom(v time.Time) {
	o.From = v
}

// GetTo returns the To field value
func (o *BlockItemModel) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetToOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *BlockItemModel) SetTo(v time.Time) {
	o.To = v
}

// GetPickedReservations returns the PickedReservations field value
func (o *BlockItemModel) GetPickedReservations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PickedReservations
}

// GetPickedReservationsOk returns a tuple with the PickedReservations field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetPickedReservationsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PickedReservations, true
}

// SetPickedReservations sets field value
func (o *BlockItemModel) SetPickedReservations(v int32) {
	o.PickedReservations = v
}

// GetPromoCode returns the PromoCode field value if set, zero value otherwise.
func (o *BlockItemModel) GetPromoCode() string {
	if o == nil || o.PromoCode == nil {
		var ret string
		return ret
	}
	return *o.PromoCode
}

// GetPromoCodeOk returns a tuple with the PromoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetPromoCodeOk() (*string, bool) {
	if o == nil || o.PromoCode == nil {
		return nil, false
	}
	return o.PromoCode, true
}

// HasPromoCode returns a boolean if a field has been set.
func (o *BlockItemModel) HasPromoCode() bool {
	if o != nil && o.PromoCode != nil {
		return true
	}

	return false
}

// SetPromoCode gets a reference to the given string and assigns it to the PromoCode field.
func (o *BlockItemModel) SetPromoCode(v string) {
	o.PromoCode = &v
}

// GetCorporateCode returns the CorporateCode field value if set, zero value otherwise.
func (o *BlockItemModel) GetCorporateCode() string {
	if o == nil || o.CorporateCode == nil {
		var ret string
		return ret
	}
	return *o.CorporateCode
}

// GetCorporateCodeOk returns a tuple with the CorporateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetCorporateCodeOk() (*string, bool) {
	if o == nil || o.CorporateCode == nil {
		return nil, false
	}
	return o.CorporateCode, true
}

// HasCorporateCode returns a boolean if a field has been set.
func (o *BlockItemModel) HasCorporateCode() bool {
	if o != nil && o.CorporateCode != nil {
		return true
	}

	return false
}

// SetCorporateCode gets a reference to the given string and assigns it to the CorporateCode field.
func (o *BlockItemModel) SetCorporateCode(v string) {
	o.CorporateCode = &v
}

// GetCreated returns the Created field value
func (o *BlockItemModel) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *BlockItemModel) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *BlockItemModel) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetModifiedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *BlockItemModel) SetModified(v time.Time) {
	o.Modified = v
}

// GetTimeSlices returns the TimeSlices field value if set, zero value otherwise.
func (o *BlockItemModel) GetTimeSlices() []BlockTimeSliceModel {
	if o == nil || o.TimeSlices == nil {
		var ret []BlockTimeSliceModel
		return ret
	}
	return *o.TimeSlices
}

// GetTimeSlicesOk returns a tuple with the TimeSlices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetTimeSlicesOk() (*[]BlockTimeSliceModel, bool) {
	if o == nil || o.TimeSlices == nil {
		return nil, false
	}
	return o.TimeSlices, true
}

// HasTimeSlices returns a boolean if a field has been set.
func (o *BlockItemModel) HasTimeSlices() bool {
	if o != nil && o.TimeSlices != nil {
		return true
	}

	return false
}

// SetTimeSlices gets a reference to the given []BlockTimeSliceModel and assigns it to the TimeSlices field.
func (o *BlockItemModel) SetTimeSlices(v []BlockTimeSliceModel) {
	o.TimeSlices = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *BlockItemModel) GetActions() []ActionModelBlockActionNotAllowedBlockActionReason {
	if o == nil || o.Actions == nil {
		var ret []ActionModelBlockActionNotAllowedBlockActionReason
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockItemModel) GetActionsOk() (*[]ActionModelBlockActionNotAllowedBlockActionReason, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *BlockItemModel) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ActionModelBlockActionNotAllowedBlockActionReason and assigns it to the Actions field.
func (o *BlockItemModel) SetActions(v []ActionModelBlockActionNotAllowedBlockActionReason) {
	o.Actions = &v
}

func (o BlockItemModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["property"] = o.Property
	}
	if true {
		toSerialize["ratePlan"] = o.RatePlan
	}
	if true {
		toSerialize["unitGroup"] = o.UnitGroup
	}
	if true {
		toSerialize["grossDailyRate"] = o.GrossDailyRate
	}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["pickedReservations"] = o.PickedReservations
	}
	if o.PromoCode != nil {
		toSerialize["promoCode"] = o.PromoCode
	}
	if o.CorporateCode != nil {
		toSerialize["corporateCode"] = o.CorporateCode
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if o.TimeSlices != nil {
		toSerialize["timeSlices"] = o.TimeSlices
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableBlockItemModel struct {
	value *BlockItemModel
	isSet bool
}

func (v NullableBlockItemModel) Get() *BlockItemModel {
	return v.value
}

func (v *NullableBlockItemModel) Set(val *BlockItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockItemModel(val *BlockItemModel) *NullableBlockItemModel {
	return &NullableBlockItemModel{value: val, isSet: true}
}

func (v NullableBlockItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


