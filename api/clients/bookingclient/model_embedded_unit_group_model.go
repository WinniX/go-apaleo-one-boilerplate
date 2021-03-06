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

// EmbeddedUnitGroupModel struct for EmbeddedUnitGroupModel
type EmbeddedUnitGroupModel struct {
	// The unit group id
	Id string `json:"id"`
	// The code for the unit group that can be shown in reports and table views
	Code *string `json:"code,omitempty"`
	// The name for the unit group
	Name *string `json:"name,omitempty"`
	// The description for the unit group
	Description *string `json:"description,omitempty"`
}

// NewEmbeddedUnitGroupModel instantiates a new EmbeddedUnitGroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedUnitGroupModel(id string) *EmbeddedUnitGroupModel {
	this := EmbeddedUnitGroupModel{}
	this.Id = id
	return &this
}

// NewEmbeddedUnitGroupModelWithDefaults instantiates a new EmbeddedUnitGroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedUnitGroupModelWithDefaults() *EmbeddedUnitGroupModel {
	this := EmbeddedUnitGroupModel{}
	return &this
}

// GetId returns the Id field value
func (o *EmbeddedUnitGroupModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmbeddedUnitGroupModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmbeddedUnitGroupModel) SetId(v string) {
	o.Id = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EmbeddedUnitGroupModel) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedUnitGroupModel) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EmbeddedUnitGroupModel) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EmbeddedUnitGroupModel) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmbeddedUnitGroupModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedUnitGroupModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmbeddedUnitGroupModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmbeddedUnitGroupModel) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EmbeddedUnitGroupModel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedUnitGroupModel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EmbeddedUnitGroupModel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EmbeddedUnitGroupModel) SetDescription(v string) {
	o.Description = &v
}

func (o EmbeddedUnitGroupModel) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableEmbeddedUnitGroupModel struct {
	value *EmbeddedUnitGroupModel
	isSet bool
}

func (v NullableEmbeddedUnitGroupModel) Get() *EmbeddedUnitGroupModel {
	return v.value
}

func (v *NullableEmbeddedUnitGroupModel) Set(val *EmbeddedUnitGroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedUnitGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedUnitGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedUnitGroupModel(val *EmbeddedUnitGroupModel) *NullableEmbeddedUnitGroupModel {
	return &NullableEmbeddedUnitGroupModel{value: val, isSet: true}
}

func (v NullableEmbeddedUnitGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedUnitGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


