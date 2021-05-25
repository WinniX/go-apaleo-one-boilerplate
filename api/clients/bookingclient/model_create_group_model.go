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

// CreateGroupModel With this request you can create a group booking request
type CreateGroupModel struct {
	// Name of the group
	Name string `json:"name"`
	Booker BookerModel `json:"booker"`
	// Additional information and comments
	Comment *string `json:"comment,omitempty"`
	// Additional information and comment by the booker
	BookerComment *string `json:"bookerComment,omitempty"`
	PaymentAccount *CreatePaymentAccountModel `json:"paymentAccount,omitempty"`
	// List of property ids the group booking belongs to
	PropertyIds []string `json:"propertyIds"`
}

// NewCreateGroupModel instantiates a new CreateGroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupModel(name string, booker BookerModel, propertyIds []string) *CreateGroupModel {
	this := CreateGroupModel{}
	this.Name = name
	this.Booker = booker
	this.PropertyIds = propertyIds
	return &this
}

// NewCreateGroupModelWithDefaults instantiates a new CreateGroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupModelWithDefaults() *CreateGroupModel {
	this := CreateGroupModel{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGroupModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGroupModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGroupModel) SetName(v string) {
	o.Name = v
}

// GetBooker returns the Booker field value
func (o *CreateGroupModel) GetBooker() BookerModel {
	if o == nil {
		var ret BookerModel
		return ret
	}

	return o.Booker
}

// GetBookerOk returns a tuple with the Booker field value
// and a boolean to check if the value has been set.
func (o *CreateGroupModel) GetBookerOk() (*BookerModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Booker, true
}

// SetBooker sets field value
func (o *CreateGroupModel) SetBooker(v BookerModel) {
	o.Booker = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateGroupModel) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupModel) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateGroupModel) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateGroupModel) SetComment(v string) {
	o.Comment = &v
}

// GetBookerComment returns the BookerComment field value if set, zero value otherwise.
func (o *CreateGroupModel) GetBookerComment() string {
	if o == nil || o.BookerComment == nil {
		var ret string
		return ret
	}
	return *o.BookerComment
}

// GetBookerCommentOk returns a tuple with the BookerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupModel) GetBookerCommentOk() (*string, bool) {
	if o == nil || o.BookerComment == nil {
		return nil, false
	}
	return o.BookerComment, true
}

// HasBookerComment returns a boolean if a field has been set.
func (o *CreateGroupModel) HasBookerComment() bool {
	if o != nil && o.BookerComment != nil {
		return true
	}

	return false
}

// SetBookerComment gets a reference to the given string and assigns it to the BookerComment field.
func (o *CreateGroupModel) SetBookerComment(v string) {
	o.BookerComment = &v
}

// GetPaymentAccount returns the PaymentAccount field value if set, zero value otherwise.
func (o *CreateGroupModel) GetPaymentAccount() CreatePaymentAccountModel {
	if o == nil || o.PaymentAccount == nil {
		var ret CreatePaymentAccountModel
		return ret
	}
	return *o.PaymentAccount
}

// GetPaymentAccountOk returns a tuple with the PaymentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupModel) GetPaymentAccountOk() (*CreatePaymentAccountModel, bool) {
	if o == nil || o.PaymentAccount == nil {
		return nil, false
	}
	return o.PaymentAccount, true
}

// HasPaymentAccount returns a boolean if a field has been set.
func (o *CreateGroupModel) HasPaymentAccount() bool {
	if o != nil && o.PaymentAccount != nil {
		return true
	}

	return false
}

// SetPaymentAccount gets a reference to the given CreatePaymentAccountModel and assigns it to the PaymentAccount field.
func (o *CreateGroupModel) SetPaymentAccount(v CreatePaymentAccountModel) {
	o.PaymentAccount = &v
}

// GetPropertyIds returns the PropertyIds field value
func (o *CreateGroupModel) GetPropertyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertyIds
}

// GetPropertyIdsOk returns a tuple with the PropertyIds field value
// and a boolean to check if the value has been set.
func (o *CreateGroupModel) GetPropertyIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyIds, true
}

// SetPropertyIds sets field value
func (o *CreateGroupModel) SetPropertyIds(v []string) {
	o.PropertyIds = v
}

func (o CreateGroupModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["booker"] = o.Booker
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.BookerComment != nil {
		toSerialize["bookerComment"] = o.BookerComment
	}
	if o.PaymentAccount != nil {
		toSerialize["paymentAccount"] = o.PaymentAccount
	}
	if true {
		toSerialize["propertyIds"] = o.PropertyIds
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGroupModel struct {
	value *CreateGroupModel
	isSet bool
}

func (v NullableCreateGroupModel) Get() *CreateGroupModel {
	return v.value
}

func (v *NullableCreateGroupModel) Set(val *CreateGroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupModel(val *CreateGroupModel) *NullableCreateGroupModel {
	return &NullableCreateGroupModel{value: val, isSet: true}
}

func (v NullableCreateGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

