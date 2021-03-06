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

// BlockCreatedModel struct for BlockCreatedModel
type BlockCreatedModel struct {
	// Block id
	Id string `json:"id"`
}

// NewBlockCreatedModel instantiates a new BlockCreatedModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockCreatedModel(id string) *BlockCreatedModel {
	this := BlockCreatedModel{}
	this.Id = id
	return &this
}

// NewBlockCreatedModelWithDefaults instantiates a new BlockCreatedModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockCreatedModelWithDefaults() *BlockCreatedModel {
	this := BlockCreatedModel{}
	return &this
}

// GetId returns the Id field value
func (o *BlockCreatedModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BlockCreatedModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BlockCreatedModel) SetId(v string) {
	o.Id = v
}

func (o BlockCreatedModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBlockCreatedModel struct {
	value *BlockCreatedModel
	isSet bool
}

func (v NullableBlockCreatedModel) Get() *BlockCreatedModel {
	return v.value
}

func (v *NullableBlockCreatedModel) Set(val *BlockCreatedModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockCreatedModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockCreatedModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockCreatedModel(val *BlockCreatedModel) *NullableBlockCreatedModel {
	return &NullableBlockCreatedModel{value: val, isSet: true}
}

func (v NullableBlockCreatedModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockCreatedModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


