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

// CommissionModel struct for CommissionModel
type CommissionModel struct {
	CommissionAmount MonetaryValueModel `json:"commissionAmount"`
	BeforeCommissionAmount *MonetaryValueModel `json:"beforeCommissionAmount,omitempty"`
}

// NewCommissionModel instantiates a new CommissionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommissionModel(commissionAmount MonetaryValueModel) *CommissionModel {
	this := CommissionModel{}
	this.CommissionAmount = commissionAmount
	return &this
}

// NewCommissionModelWithDefaults instantiates a new CommissionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommissionModelWithDefaults() *CommissionModel {
	this := CommissionModel{}
	return &this
}

// GetCommissionAmount returns the CommissionAmount field value
func (o *CommissionModel) GetCommissionAmount() MonetaryValueModel {
	if o == nil {
		var ret MonetaryValueModel
		return ret
	}

	return o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value
// and a boolean to check if the value has been set.
func (o *CommissionModel) GetCommissionAmountOk() (*MonetaryValueModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommissionAmount, true
}

// SetCommissionAmount sets field value
func (o *CommissionModel) SetCommissionAmount(v MonetaryValueModel) {
	o.CommissionAmount = v
}

// GetBeforeCommissionAmount returns the BeforeCommissionAmount field value if set, zero value otherwise.
func (o *CommissionModel) GetBeforeCommissionAmount() MonetaryValueModel {
	if o == nil || o.BeforeCommissionAmount == nil {
		var ret MonetaryValueModel
		return ret
	}
	return *o.BeforeCommissionAmount
}

// GetBeforeCommissionAmountOk returns a tuple with the BeforeCommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionModel) GetBeforeCommissionAmountOk() (*MonetaryValueModel, bool) {
	if o == nil || o.BeforeCommissionAmount == nil {
		return nil, false
	}
	return o.BeforeCommissionAmount, true
}

// HasBeforeCommissionAmount returns a boolean if a field has been set.
func (o *CommissionModel) HasBeforeCommissionAmount() bool {
	if o != nil && o.BeforeCommissionAmount != nil {
		return true
	}

	return false
}

// SetBeforeCommissionAmount gets a reference to the given MonetaryValueModel and assigns it to the BeforeCommissionAmount field.
func (o *CommissionModel) SetBeforeCommissionAmount(v MonetaryValueModel) {
	o.BeforeCommissionAmount = &v
}

func (o CommissionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if o.BeforeCommissionAmount != nil {
		toSerialize["beforeCommissionAmount"] = o.BeforeCommissionAmount
	}
	return json.Marshal(toSerialize)
}

type NullableCommissionModel struct {
	value *CommissionModel
	isSet bool
}

func (v NullableCommissionModel) Get() *CommissionModel {
	return v.value
}

func (v *NullableCommissionModel) Set(val *CommissionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCommissionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCommissionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommissionModel(val *CommissionModel) *NullableCommissionModel {
	return &NullableCommissionModel{value: val, isSet: true}
}

func (v NullableCommissionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommissionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


