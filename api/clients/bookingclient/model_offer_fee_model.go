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

// OfferFeeModel struct for OfferFeeModel
type OfferFeeModel struct {
	// The fee id
	Id string `json:"id"`
	// The code for the fee
	Code string `json:"code"`
	// The name for the fee
	Name string `json:"name"`
	TotalAmount AmountModel `json:"totalAmount"`
}

// NewOfferFeeModel instantiates a new OfferFeeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferFeeModel(id string, code string, name string, totalAmount AmountModel) *OfferFeeModel {
	this := OfferFeeModel{}
	this.Id = id
	this.Code = code
	this.Name = name
	this.TotalAmount = totalAmount
	return &this
}

// NewOfferFeeModelWithDefaults instantiates a new OfferFeeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferFeeModelWithDefaults() *OfferFeeModel {
	this := OfferFeeModel{}
	return &this
}

// GetId returns the Id field value
func (o *OfferFeeModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OfferFeeModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OfferFeeModel) SetId(v string) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *OfferFeeModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OfferFeeModel) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OfferFeeModel) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *OfferFeeModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OfferFeeModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OfferFeeModel) SetName(v string) {
	o.Name = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *OfferFeeModel) GetTotalAmount() AmountModel {
	if o == nil {
		var ret AmountModel
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *OfferFeeModel) GetTotalAmountOk() (*AmountModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *OfferFeeModel) SetTotalAmount(v AmountModel) {
	o.TotalAmount = v
}

func (o OfferFeeModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	return json.Marshal(toSerialize)
}

type NullableOfferFeeModel struct {
	value *OfferFeeModel
	isSet bool
}

func (v NullableOfferFeeModel) Get() *OfferFeeModel {
	return v.value
}

func (v *NullableOfferFeeModel) Set(val *OfferFeeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferFeeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferFeeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferFeeModel(val *OfferFeeModel) *NullableOfferFeeModel {
	return &NullableOfferFeeModel{value: val, isSet: true}
}

func (v NullableOfferFeeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferFeeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


