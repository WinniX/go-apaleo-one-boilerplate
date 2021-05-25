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

// ServiceDateItemModel struct for ServiceDateItemModel
type ServiceDateItemModel struct {
	// The date this service is delivered
	ServiceDate string `json:"serviceDate"`
	// The count of booked services
	Count int32 `json:"count"`
	Amount AmountModel `json:"amount"`
	// Rate plans can have additional services. When booking an offer for such rate plans, those services are automatically booked.  They are marked as mandatory and they cannot be removed.
	IsMandatory bool `json:"isMandatory"`
}

// NewServiceDateItemModel instantiates a new ServiceDateItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDateItemModel(serviceDate string, count int32, amount AmountModel, isMandatory bool) *ServiceDateItemModel {
	this := ServiceDateItemModel{}
	this.ServiceDate = serviceDate
	this.Count = count
	this.Amount = amount
	this.IsMandatory = isMandatory
	return &this
}

// NewServiceDateItemModelWithDefaults instantiates a new ServiceDateItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDateItemModelWithDefaults() *ServiceDateItemModel {
	this := ServiceDateItemModel{}
	return &this
}

// GetServiceDate returns the ServiceDate field value
func (o *ServiceDateItemModel) GetServiceDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceDate
}

// GetServiceDateOk returns a tuple with the ServiceDate field value
// and a boolean to check if the value has been set.
func (o *ServiceDateItemModel) GetServiceDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceDate, true
}

// SetServiceDate sets field value
func (o *ServiceDateItemModel) SetServiceDate(v string) {
	o.ServiceDate = v
}

// GetCount returns the Count field value
func (o *ServiceDateItemModel) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ServiceDateItemModel) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ServiceDateItemModel) SetCount(v int32) {
	o.Count = v
}

// GetAmount returns the Amount field value
func (o *ServiceDateItemModel) GetAmount() AmountModel {
	if o == nil {
		var ret AmountModel
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ServiceDateItemModel) GetAmountOk() (*AmountModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ServiceDateItemModel) SetAmount(v AmountModel) {
	o.Amount = v
}

// GetIsMandatory returns the IsMandatory field value
func (o *ServiceDateItemModel) GetIsMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value
// and a boolean to check if the value has been set.
func (o *ServiceDateItemModel) GetIsMandatoryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsMandatory, true
}

// SetIsMandatory sets field value
func (o *ServiceDateItemModel) SetIsMandatory(v bool) {
	o.IsMandatory = v
}

func (o ServiceDateItemModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceDate"] = o.ServiceDate
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["isMandatory"] = o.IsMandatory
	}
	return json.Marshal(toSerialize)
}

type NullableServiceDateItemModel struct {
	value *ServiceDateItemModel
	isSet bool
}

func (v NullableServiceDateItemModel) Get() *ServiceDateItemModel {
	return v.value
}

func (v *NullableServiceDateItemModel) Set(val *ServiceDateItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDateItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDateItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDateItemModel(val *ServiceDateItemModel) *NullableServiceDateItemModel {
	return &NullableServiceDateItemModel{value: val, isSet: true}
}

func (v NullableServiceDateItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDateItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


