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

// OfferValidationMessageModel struct for OfferValidationMessageModel
type OfferValidationMessageModel struct {
	// The message Code
	Code string `json:"code"`
	// The message description
	Message string `json:"message"`
}

// NewOfferValidationMessageModel instantiates a new OfferValidationMessageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferValidationMessageModel(code string, message string) *OfferValidationMessageModel {
	this := OfferValidationMessageModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewOfferValidationMessageModelWithDefaults instantiates a new OfferValidationMessageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferValidationMessageModelWithDefaults() *OfferValidationMessageModel {
	this := OfferValidationMessageModel{}
	return &this
}

// GetCode returns the Code field value
func (o *OfferValidationMessageModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OfferValidationMessageModel) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OfferValidationMessageModel) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *OfferValidationMessageModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OfferValidationMessageModel) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OfferValidationMessageModel) SetMessage(v string) {
	o.Message = v
}

func (o OfferValidationMessageModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableOfferValidationMessageModel struct {
	value *OfferValidationMessageModel
	isSet bool
}

func (v NullableOfferValidationMessageModel) Get() *OfferValidationMessageModel {
	return v.value
}

func (v *NullableOfferValidationMessageModel) Set(val *OfferValidationMessageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferValidationMessageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferValidationMessageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferValidationMessageModel(val *OfferValidationMessageModel) *NullableOfferValidationMessageModel {
	return &NullableOfferValidationMessageModel{value: val, isSet: true}
}

func (v NullableOfferValidationMessageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferValidationMessageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

