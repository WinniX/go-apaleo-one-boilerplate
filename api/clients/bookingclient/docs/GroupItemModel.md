# GroupItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Group id | 
**From** | Pointer to **time.Time** | Start date and time of the earliest block for this group&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**To** | Pointer to **time.Time** | End date and time of the latest block for this group&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | [optional] 
**Name** | **string** | Name of the group | 
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

### NewGroupItemModel

`func NewGroupItemModel(id string, name string, created time.Time, modified time.Time, propertyIds []string, ) *GroupItemModel`

NewGroupItemModel instantiates a new GroupItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupItemModelWithDefaults

`func NewGroupItemModelWithDefaults() *GroupItemModel`

NewGroupItemModelWithDefaults instantiates a new GroupItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupItemModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupItemModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupItemModel) SetId(v string)`

SetId sets Id field to given value.


### GetFrom

`func (o *GroupItemModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GroupItemModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GroupItemModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GroupItemModel) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GroupItemModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GroupItemModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GroupItemModel) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *GroupItemModel) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetName

`func (o *GroupItemModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupItemModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupItemModel) SetName(v string)`

SetName sets Name field to given value.


### GetBooker

`func (o *GroupItemModel) GetBooker() BookerModel`

GetBooker returns the Booker field if non-nil, zero value otherwise.

### GetBookerOk

`func (o *GroupItemModel) GetBookerOk() (*BookerModel, bool)`

GetBookerOk returns a tuple with the Booker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooker

`func (o *GroupItemModel) SetBooker(v BookerModel)`

SetBooker sets Booker field to given value.

### HasBooker

`func (o *GroupItemModel) HasBooker() bool`

HasBooker returns a boolean if a field has been set.

### GetComment

`func (o *GroupItemModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GroupItemModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GroupItemModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GroupItemModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetBookerComment

`func (o *GroupItemModel) GetBookerComment() string`

GetBookerComment returns the BookerComment field if non-nil, zero value otherwise.

### GetBookerCommentOk

`func (o *GroupItemModel) GetBookerCommentOk() (*string, bool)`

GetBookerCommentOk returns a tuple with the BookerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookerComment

`func (o *GroupItemModel) SetBookerComment(v string)`

SetBookerComment sets BookerComment field to given value.

### HasBookerComment

`func (o *GroupItemModel) HasBookerComment() bool`

HasBookerComment returns a boolean if a field has been set.

### GetPaymentAccount

`func (o *GroupItemModel) GetPaymentAccount() PaymentAccountModel`

GetPaymentAccount returns the PaymentAccount field if non-nil, zero value otherwise.

### GetPaymentAccountOk

`func (o *GroupItemModel) GetPaymentAccountOk() (*PaymentAccountModel, bool)`

GetPaymentAccountOk returns a tuple with the PaymentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccount

`func (o *GroupItemModel) SetPaymentAccount(v PaymentAccountModel)`

SetPaymentAccount sets PaymentAccount field to given value.

### HasPaymentAccount

`func (o *GroupItemModel) HasPaymentAccount() bool`

HasPaymentAccount returns a boolean if a field has been set.

### GetCreated

`func (o *GroupItemModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupItemModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupItemModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *GroupItemModel) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GroupItemModel) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GroupItemModel) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetBlocks

`func (o *GroupItemModel) GetBlocks() []GroupBlockModel`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *GroupItemModel) GetBlocksOk() (*[]GroupBlockModel, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *GroupItemModel) SetBlocks(v []GroupBlockModel)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *GroupItemModel) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetActions

`func (o *GroupItemModel) GetActions() []ActionModelGroupActionNotAllowedGroupActionReason`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GroupItemModel) GetActionsOk() (*[]ActionModelGroupActionNotAllowedGroupActionReason, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GroupItemModel) SetActions(v []ActionModelGroupActionNotAllowedGroupActionReason)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GroupItemModel) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPropertyIds

`func (o *GroupItemModel) GetPropertyIds() []string`

GetPropertyIds returns the PropertyIds field if non-nil, zero value otherwise.

### GetPropertyIdsOk

`func (o *GroupItemModel) GetPropertyIdsOk() (*[]string, bool)`

GetPropertyIdsOk returns a tuple with the PropertyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIds

`func (o *GroupItemModel) SetPropertyIds(v []string)`

SetPropertyIds sets PropertyIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


