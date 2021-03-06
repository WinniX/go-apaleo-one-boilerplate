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

// StayOffersModel struct for StayOffersModel
type StayOffersModel struct {
	Property EmbeddedPropertyModel `json:"property"`
	// List of offered unit groups with rates
	Offers []OfferModel `json:"offers"`
}

// NewStayOffersModel instantiates a new StayOffersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStayOffersModel(property EmbeddedPropertyModel, offers []OfferModel) *StayOffersModel {
	this := StayOffersModel{}
	this.Property = property
	this.Offers = offers
	return &this
}

// NewStayOffersModelWithDefaults instantiates a new StayOffersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStayOffersModelWithDefaults() *StayOffersModel {
	this := StayOffersModel{}
	return &this
}

// GetProperty returns the Property field value
func (o *StayOffersModel) GetProperty() EmbeddedPropertyModel {
	if o == nil {
		var ret EmbeddedPropertyModel
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *StayOffersModel) GetPropertyOk() (*EmbeddedPropertyModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *StayOffersModel) SetProperty(v EmbeddedPropertyModel) {
	o.Property = v
}

// GetOffers returns the Offers field value
func (o *StayOffersModel) GetOffers() []OfferModel {
	if o == nil {
		var ret []OfferModel
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *StayOffersModel) GetOffersOk() (*[]OfferModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offers, true
}

// SetOffers sets field value
func (o *StayOffersModel) SetOffers(v []OfferModel) {
	o.Offers = v
}

func (o StayOffersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["property"] = o.Property
	}
	if true {
		toSerialize["offers"] = o.Offers
	}
	return json.Marshal(toSerialize)
}

type NullableStayOffersModel struct {
	value *StayOffersModel
	isSet bool
}

func (v NullableStayOffersModel) Get() *StayOffersModel {
	return v.value
}

func (v *NullableStayOffersModel) Set(val *StayOffersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStayOffersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStayOffersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStayOffersModel(val *StayOffersModel) *NullableStayOffersModel {
	return &NullableStayOffersModel{value: val, isSet: true}
}

func (v NullableStayOffersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStayOffersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


