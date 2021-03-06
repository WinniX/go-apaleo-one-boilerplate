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

// ActionReasonModelNotAllowedReservationActionReason struct for ActionReasonModelNotAllowedReservationActionReason
type ActionReasonModelNotAllowedReservationActionReason struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

// NewActionReasonModelNotAllowedReservationActionReason instantiates a new ActionReasonModelNotAllowedReservationActionReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionReasonModelNotAllowedReservationActionReason(code string, message string) *ActionReasonModelNotAllowedReservationActionReason {
	this := ActionReasonModelNotAllowedReservationActionReason{}
	this.Code = code
	this.Message = message
	return &this
}

// NewActionReasonModelNotAllowedReservationActionReasonWithDefaults instantiates a new ActionReasonModelNotAllowedReservationActionReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionReasonModelNotAllowedReservationActionReasonWithDefaults() *ActionReasonModelNotAllowedReservationActionReason {
	this := ActionReasonModelNotAllowedReservationActionReason{}
	return &this
}

// GetCode returns the Code field value
func (o *ActionReasonModelNotAllowedReservationActionReason) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ActionReasonModelNotAllowedReservationActionReason) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ActionReasonModelNotAllowedReservationActionReason) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ActionReasonModelNotAllowedReservationActionReason) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ActionReasonModelNotAllowedReservationActionReason) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ActionReasonModelNotAllowedReservationActionReason) SetMessage(v string) {
	o.Message = v
}

func (o ActionReasonModelNotAllowedReservationActionReason) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableActionReasonModelNotAllowedReservationActionReason struct {
	value *ActionReasonModelNotAllowedReservationActionReason
	isSet bool
}

func (v NullableActionReasonModelNotAllowedReservationActionReason) Get() *ActionReasonModelNotAllowedReservationActionReason {
	return v.value
}

func (v *NullableActionReasonModelNotAllowedReservationActionReason) Set(val *ActionReasonModelNotAllowedReservationActionReason) {
	v.value = val
	v.isSet = true
}

func (v NullableActionReasonModelNotAllowedReservationActionReason) IsSet() bool {
	return v.isSet
}

func (v *NullableActionReasonModelNotAllowedReservationActionReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionReasonModelNotAllowedReservationActionReason(val *ActionReasonModelNotAllowedReservationActionReason) *NullableActionReasonModelNotAllowedReservationActionReason {
	return &NullableActionReasonModelNotAllowedReservationActionReason{value: val, isSet: true}
}

func (v NullableActionReasonModelNotAllowedReservationActionReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionReasonModelNotAllowedReservationActionReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


