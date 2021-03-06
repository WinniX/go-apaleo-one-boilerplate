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

// EmbeddedCompanyModel struct for EmbeddedCompanyModel
type EmbeddedCompanyModel struct {
	// The company ID
	Id string `json:"id"`
	// The code of the company
	Code *string `json:"code,omitempty"`
	// The name of the company
	Name *string `json:"name,omitempty"`
	// Whether or not the company can check out on AR
	CanCheckOutOnAr *bool `json:"canCheckOutOnAr,omitempty"`
}

// NewEmbeddedCompanyModel instantiates a new EmbeddedCompanyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedCompanyModel(id string) *EmbeddedCompanyModel {
	this := EmbeddedCompanyModel{}
	this.Id = id
	return &this
}

// NewEmbeddedCompanyModelWithDefaults instantiates a new EmbeddedCompanyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedCompanyModelWithDefaults() *EmbeddedCompanyModel {
	this := EmbeddedCompanyModel{}
	return &this
}

// GetId returns the Id field value
func (o *EmbeddedCompanyModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmbeddedCompanyModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmbeddedCompanyModel) SetId(v string) {
	o.Id = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EmbeddedCompanyModel) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedCompanyModel) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EmbeddedCompanyModel) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EmbeddedCompanyModel) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmbeddedCompanyModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedCompanyModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmbeddedCompanyModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmbeddedCompanyModel) SetName(v string) {
	o.Name = &v
}

// GetCanCheckOutOnAr returns the CanCheckOutOnAr field value if set, zero value otherwise.
func (o *EmbeddedCompanyModel) GetCanCheckOutOnAr() bool {
	if o == nil || o.CanCheckOutOnAr == nil {
		var ret bool
		return ret
	}
	return *o.CanCheckOutOnAr
}

// GetCanCheckOutOnArOk returns a tuple with the CanCheckOutOnAr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedCompanyModel) GetCanCheckOutOnArOk() (*bool, bool) {
	if o == nil || o.CanCheckOutOnAr == nil {
		return nil, false
	}
	return o.CanCheckOutOnAr, true
}

// HasCanCheckOutOnAr returns a boolean if a field has been set.
func (o *EmbeddedCompanyModel) HasCanCheckOutOnAr() bool {
	if o != nil && o.CanCheckOutOnAr != nil {
		return true
	}

	return false
}

// SetCanCheckOutOnAr gets a reference to the given bool and assigns it to the CanCheckOutOnAr field.
func (o *EmbeddedCompanyModel) SetCanCheckOutOnAr(v bool) {
	o.CanCheckOutOnAr = &v
}

func (o EmbeddedCompanyModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CanCheckOutOnAr != nil {
		toSerialize["canCheckOutOnAr"] = o.CanCheckOutOnAr
	}
	return json.Marshal(toSerialize)
}

type NullableEmbeddedCompanyModel struct {
	value *EmbeddedCompanyModel
	isSet bool
}

func (v NullableEmbeddedCompanyModel) Get() *EmbeddedCompanyModel {
	return v.value
}

func (v *NullableEmbeddedCompanyModel) Set(val *EmbeddedCompanyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedCompanyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedCompanyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedCompanyModel(val *EmbeddedCompanyModel) *NullableEmbeddedCompanyModel {
	return &NullableEmbeddedCompanyModel{value: val, isSet: true}
}

func (v NullableEmbeddedCompanyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedCompanyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


