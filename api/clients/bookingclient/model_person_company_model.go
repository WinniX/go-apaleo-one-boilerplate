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

// PersonCompanyModel struct for PersonCompanyModel
type PersonCompanyModel struct {
	// Name of the company
	Name *string `json:"name,omitempty"`
	// Tax or Vat ID of the company
	TaxId *string `json:"taxId,omitempty"`
}

// NewPersonCompanyModel instantiates a new PersonCompanyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonCompanyModel() *PersonCompanyModel {
	this := PersonCompanyModel{}
	return &this
}

// NewPersonCompanyModelWithDefaults instantiates a new PersonCompanyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonCompanyModelWithDefaults() *PersonCompanyModel {
	this := PersonCompanyModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PersonCompanyModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonCompanyModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PersonCompanyModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PersonCompanyModel) SetName(v string) {
	o.Name = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *PersonCompanyModel) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonCompanyModel) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *PersonCompanyModel) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *PersonCompanyModel) SetTaxId(v string) {
	o.TaxId = &v
}

func (o PersonCompanyModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TaxId != nil {
		toSerialize["taxId"] = o.TaxId
	}
	return json.Marshal(toSerialize)
}

type NullablePersonCompanyModel struct {
	value *PersonCompanyModel
	isSet bool
}

func (v NullablePersonCompanyModel) Get() *PersonCompanyModel {
	return v.value
}

func (v *NullablePersonCompanyModel) Set(val *PersonCompanyModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonCompanyModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonCompanyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonCompanyModel(val *PersonCompanyModel) *NullablePersonCompanyModel {
	return &NullablePersonCompanyModel{value: val, isSet: true}
}

func (v NullablePersonCompanyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonCompanyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


