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

// GroupCreatedModel struct for GroupCreatedModel
type GroupCreatedModel struct {
	// Id of newly created group booking
	Id string `json:"id"`
}

// NewGroupCreatedModel instantiates a new GroupCreatedModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCreatedModel(id string) *GroupCreatedModel {
	this := GroupCreatedModel{}
	this.Id = id
	return &this
}

// NewGroupCreatedModelWithDefaults instantiates a new GroupCreatedModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCreatedModelWithDefaults() *GroupCreatedModel {
	this := GroupCreatedModel{}
	return &this
}

// GetId returns the Id field value
func (o *GroupCreatedModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupCreatedModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupCreatedModel) SetId(v string) {
	o.Id = v
}

func (o GroupCreatedModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGroupCreatedModel struct {
	value *GroupCreatedModel
	isSet bool
}

func (v NullableGroupCreatedModel) Get() *GroupCreatedModel {
	return v.value
}

func (v *NullableGroupCreatedModel) Set(val *GroupCreatedModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCreatedModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCreatedModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCreatedModel(val *GroupCreatedModel) *NullableGroupCreatedModel {
	return &NullableGroupCreatedModel{value: val, isSet: true}
}

func (v NullableGroupCreatedModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCreatedModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


