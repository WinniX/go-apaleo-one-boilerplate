# CreateGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the group | 
**Booker** | [**BookerModel**](BookerModel.md) |  | 
**Comment** | Pointer to **string** | Additional information and comments | [optional] 
**BookerComment** | Pointer to **string** | Additional information and comment by the booker | [optional] 
**PaymentAccount** | Pointer to [**CreatePaymentAccountModel**](CreatePaymentAccountModel.md) |  | [optional] 
**PropertyIds** | **[]string** | List of property ids the group booking belongs to | 

## Methods

### NewCreateGroupModel

`func NewCreateGroupModel(name string, booker BookerModel, propertyIds []string, ) *CreateGroupModel`

NewCreateGroupModel instantiates a new CreateGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupModelWithDefaults

`func NewCreateGroupModelWithDefaults() *CreateGroupModel`

NewCreateGroupModelWithDefaults instantiates a new CreateGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetBooker

`func (o *CreateGroupModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *CreateGroupModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *CreateGroupModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.


### GetComment

`func (o *CreateGroupModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGroupModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGroupModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGroupModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetBookerComment

`func (o *CreateGroupModel) GetBookerComment() string`

GetBookerComment returns the BookerComment field if non-nil, zero value otherwise.

### GetBookerCommentOk

`func (o *CreateGroupModel) GetBookerCommentOk() (*string, bool)`

GetBookerCommentOk returns a tuple with the BookerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookerComment

`func (o *CreateGroupModel) SetBookerComment(v string)`

SetBookerComment sets BookerComment field to given value.

### HasBookerComment

`func (o *CreateGroupModel) HasBookerComment() bool`

HasBookerComment returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *CreateGroupModel) GetPaymentAccount() CreatePaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *CreateGroupModel) GetPaymentAccountOk() (*CreatePaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *CreateGroupModel) SetPaymentAccount(v CreatePaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *CreateGroupModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetPropertyIds

`func (o *CreateGroupModel) GetPropertyIds() []string`

GetPropertyIds returns the PropertyIds field if non-nil, zero value otherwise.

### GetPropertyIdsOk

`func (o *CreateGroupModel) GetPropertyIdsOk() (*[]string, bool)`

GetPropertyIdsOk returns a tuple with the PropertyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIds

`func (o *CreateGroupModel) SetPropertyIds(v []string)`

SetPropertyIds sets PropertyIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


