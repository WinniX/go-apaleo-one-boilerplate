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

// ReservationListModel struct for ReservationListModel
type ReservationListModel struct {
	// List of reservations
	Reservations []ReservationItemModel `json:"reservations"`
	// Total count of items
	Count int64 `json:"count"`
}

// NewReservationListModel instantiates a new ReservationListModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationListModel(reservations []ReservationItemModel, count int64) *ReservationListModel {
	this := ReservationListModel{}
	this.Reservations = reservations
	this.Count = count
	return &this
}

// NewReservationListModelWithDefaults instantiates a new ReservationListModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationListModelWithDefaults() *ReservationListModel {
	this := ReservationListModel{}
	return &this
}

// GetReservations returns the Reservations field value
func (o *ReservationListModel) GetReservations() []ReservationItemModel {
	if o == nil {
		var ret []ReservationItemModel
		return ret
	}

	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value
// and a boolean to check if the value has been set.
func (o *ReservationListModel) GetReservationsOk() (*[]ReservationItemModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reservations, true
}

// SetReservations sets field value
func (o *ReservationListModel) SetReservations(v []ReservationItemModel) {
	o.Reservations = v
}

// GetCount returns the Count field value
func (o *ReservationListModel) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ReservationListModel) GetCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ReservationListModel) SetCount(v int64) {
	o.Count = v
}

func (o ReservationListModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reservations"] = o.Reservations
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableReservationListModel struct {
	value *ReservationListModel
	isSet bool
}

func (v NullableReservationListModel) Get() *ReservationListModel {
	return v.value
}

func (v *NullableReservationListModel) Set(val *ReservationListModel) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationListModel) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationListModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationListModel(val *ReservationListModel) *NullableReservationListModel {
	return &NullableReservationListModel{value: val, isSet: true}
}

func (v NullableReservationListModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationListModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


