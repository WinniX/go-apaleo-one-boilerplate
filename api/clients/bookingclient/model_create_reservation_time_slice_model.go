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

// CreateReservationTimeSliceModel struct for CreateReservationTimeSliceModel
type CreateReservationTimeSliceModel struct {
	// The rate plan id for this time slice
	RatePlanId string `json:"ratePlanId"`
	TotalAmount *MonetaryValueModel `json:"totalAmount,omitempty"`
}

// NewCreateReservationTimeSliceModel instantiates a new CreateReservationTimeSliceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReservationTimeSliceModel(ratePlanId string) *CreateReservationTimeSliceModel {
	this := CreateReservationTimeSliceModel{}
	this.RatePlanId = ratePlanId
	return &this
}

// NewCreateReservationTimeSliceModelWithDefaults instantiates a new CreateReservationTimeSliceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReservationTimeSliceModelWithDefaults() *CreateReservationTimeSliceModel {
	this := CreateReservationTimeSliceModel{}
	return &this
}

// GetRatePlanId returns the RatePlanId field value
func (o *CreateReservationTimeSliceModel) GetRatePlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RatePlanId
}

// GetRatePlanIdOk returns a tuple with the RatePlanId field value
// and a boolean to check if the value has been set.
func (o *CreateReservationTimeSliceModel) GetRatePlanIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RatePlanId, true
}

// SetRatePlanId sets field value
func (o *CreateReservationTimeSliceModel) SetRatePlanId(v string) {
	o.RatePlanId = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *CreateReservationTimeSliceModel) GetTotalAmount() MonetaryValueModel {
	if o == nil || o.TotalAmount == nil {
		var ret MonetaryValueModel
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReservationTimeSliceModel) GetTotalAmountOk() (*MonetaryValueModel, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *CreateReservationTimeSliceModel) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given MonetaryValueModel and assigns it to the TotalAmount field.
func (o *CreateReservationTimeSliceModel) SetTotalAmount(v MonetaryValueModel) {
	o.TotalAmount = &v
}

func (o CreateReservationTimeSliceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ratePlanId"] = o.RatePlanId
	}
	if o.TotalAmount != nil {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReservationTimeSliceModel struct {
	value *CreateReservationTimeSliceModel
	isSet bool
}

func (v NullableCreateReservationTimeSliceModel) Get() *CreateReservationTimeSliceModel {
	return v.value
}

func (v *NullableCreateReservationTimeSliceModel) Set(val *CreateReservationTimeSliceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReservationTimeSliceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReservationTimeSliceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReservationTimeSliceModel(val *CreateReservationTimeSliceModel) *NullableCreateReservationTimeSliceModel {
	return &NullableCreateReservationTimeSliceModel{value: val, isSet: true}
}

func (v NullableCreateReservationTimeSliceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReservationTimeSliceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


