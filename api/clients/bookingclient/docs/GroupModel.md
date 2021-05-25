# GroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Group id | 
**Name** | **string** | Name of the group | 
**From** | Pointer to **time.Time** | Start date and time of the earliest block for this group&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**To** | Pointer to **time.Time** | End date and time of the latest block for this group&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**Booker** | Pointer to [**BookerModel**](BookerModel.md) |  | [optional] 
**Comment** | Pointer to **string** | Additional information and comments | [optional] 
**BookerComment** | Pointer to **string** | Additional information and comment by the booker | [optional] 
**PaymentAccount** | Pointer to [**PaymentAccountModel**](PaymentAccountModel.md) |  | [optional] 
**Created** | **time.Time** | Date of creation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Modified** | **time.Time** | Date of last modification&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Blocks** | Pointer to [**[]GroupBlockModel**](GroupBlockModel.md) | Blocks within this group | [optional] 
**Actions** | Pointer to [**[]ActionModelGroupActionNotAllowedGroupActionReason**](ActionModelGroupActionNotAllowedGroupActionReason.md) | The list of actions for this group | [optional] 
**PropertyIds** | **[]string** | The list of property ids this group belongs to | 

## Methods

### NewGroupModel

`func NewGroupModel(id string, name string, created time.Time, modified time.Time, propertyIds []string, ) *GroupModel`

NewGroupModel instantiates a new GroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupModelWithDefaults

`func NewGroupModelWithDefaults() *GroupModel`

NewGroupModelWithDefaults instantiates a new GroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetFrom

`func (o *GroupModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GroupModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GroupModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GroupModel) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GroupModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GroupModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GroupModel) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *GroupModel) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetBooker

`func (o *GroupModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *GroupModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *GroupModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.

### HasBooker

`func (o *GroupModel) HasBooker() bool`

HasBooker returns a boolean if a field has been set.

### GetComment

`func (o *GroupModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GroupModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GroupModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GroupModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetBookerComment

`func (o *GroupModel) GetBookerComment() string`

GetBookerComment returns the BookerComment field if non-nil, zero value otherwise.

### GetBookerCommentOk

`func (o *GroupModel) GetBookerCommentOk() (*string, bool)`

GetBookerCommentOk returns a tuple with the BookerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookerComment

`func (o *GroupModel) SetBookerComment(v string)`

SetBookerComment sets BookerComment field to given value.

### HasBookerComment

`func (o *GroupModel) HasBookerComment() bool`

HasBookerComment returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *GroupModel) GetPaymentAccount() PaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *GroupModel) GetPaymentAccountOk() (*PaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *GroupModel) SetPaymentAccount(v PaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *GroupModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetCreated

`func (o *GroupModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *GroupModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GroupModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GroupModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetBlocks

`func (o *GroupModel) GetBlocks() []GroupBlockModel`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *GroupModel) GetBlocksOk() (*[]GroupBlockModel, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *GroupModel) SetBlocks(v []GroupBlockModel)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *GroupModel) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetActions

`func (o *GroupModel) GetActions() []ActionModelGroupActionNotAllowedGroupActionReason`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GroupModel) GetActionsOk() (*[]ActionModelGroupActionNotAllowedGroupActionReason, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GroupModel) SetActions(v []ActionModelGroupActionNotAllowedGroupActionReason)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GroupModel) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPropertyIds

`func (o *GroupModel) GetPropertyIds() []string`

GetPropertyIds returns the PropertyIds field if non-nil, zero value otherwise.

### GetPropertyIdsOk

`func (o *GroupModel) GetPropertyIdsOk() (*[]string, bool)`

GetPropertyIdsOk returns a tuple with the PropertyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIds

`func (o *GroupModel) SetPropertyIds(v []string)`

SetPropertyIds sets PropertyIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


