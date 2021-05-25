# GroupListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]GroupItemModel**](GroupItemModel.md) |  | 
**Count** | **int64** | Total count of items | 

## Methods

### NewGroupListModel

`func NewGroupListModel(groups []GroupItemModel, count int64, ) *GroupListModel`

NewGroupListModel instantiates a new GroupListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupListModelWithDefaults

`func NewGroupListModelWithDefaults() *GroupListModel`

NewGroupListModelWithDefaults instantiates a new GroupListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GroupListModel) GetGroups() []GroupItemModel`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupListModel) GetGroupsOk() (*[]GroupItemModel, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupListModel) SetGroups(v []GroupItemModel)`

SetGroups sets Groups field to given value.


### GetCount

`func (o *GroupListModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupListModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupListModel) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


