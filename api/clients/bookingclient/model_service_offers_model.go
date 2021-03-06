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

// ServiceOffersModel struct for ServiceOffersModel
type ServiceOffersModel struct {
	Services []ServiceOfferModel `json:"services"`
}

// NewServiceOffersModel instantiates a new ServiceOffersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOffersModel(services []ServiceOfferModel) *ServiceOffersModel {
	this := ServiceOffersModel{}
	this.Services = services
	return &this
}

// NewServiceOffersModelWithDefaults instantiates a new ServiceOffersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOffersModelWithDefaults() *ServiceOffersModel {
	this := ServiceOffersModel{}
	return &this
}

// GetServices returns the Services field value
func (o *ServiceOffersModel) GetServices() []ServiceOfferModel {
	if o == nil {
		var ret []ServiceOfferModel
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *ServiceOffersModel) GetServicesOk() (*[]ServiceOfferModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *ServiceOffersModel) SetServices(v []ServiceOfferModel) {
	o.Services = v
}

func (o ServiceOffersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableServiceOffersModel struct {
	value *ServiceOffersModel
	isSet bool
}

func (v NullableServiceOffersModel) Get() *ServiceOffersModel {
	return v.value
}

func (v *NullableServiceOffersModel) Set(val *ServiceOffersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOffersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOffersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOffersModel(val *ServiceOffersModel) *NullableServiceOffersModel {
	return &NullableServiceOffersModel{value: val, isSet: true}
}

func (v NullableServiceOffersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOffersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


