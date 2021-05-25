# ActionModelBlockActionNotAllowedBlockActionReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**IsAllowed** | **bool** |  | 
**Reasons** | Pointer to [**[]ActionReasonModelNotAllowedBlockActionReason**](ActionReasonModelNotAllowedBlockActionReason.md) |  | [optional] 

## Methods

### NewActionModelBlockActionNotAllowedBlockActionReason

`func NewActionModelBlockActionNotAllowedBlockActionReason(action string, isAllowed bool, ) *ActionModelBlockActionNotAllowedBlockActionReason`

NewActionModelBlockActionNotAllowedBlockActionReason instantiates a new ActionModelBlockActionNotAllowedBlockActionReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionModelBlockActionNotAllowedBlockActionReasonWithDefaults

`func NewActionModelBlockActionNotAllowedBlockActionReasonWithDefaults() *ActionModelBlockActionNotAllowedBlockActionReason`

NewActionModelBlockActionNotAllowedBlockActionReasonWithDefaults instantiates a new ActionModelBlockActionNotAllowedBlockActionReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) SetAction(v string)`

SetAction sets Action field to given value.


### GetIsAllowed

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) GetIsAllowed() bool`

GetIsAllowed returns the IsAllowed field if non-nil, zero value otherwise.

### GetIsAllowedOk

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) GetIsAllowedOk() (*bool, bool)`

GetIsAllowedOk returns a tuple with the IsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowed

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) SetIsAllowed(v bool)`

SetIsAllowed sets IsAllowed field to given value.


### GetReasons

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) GetReasons() []ActionReasonModelNotAllowedBlockActionReason`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) GetReasonsOk() (*[]ActionReasonModelNotAllowedBlockActionReason, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) SetReasons(v []ActionReasonModelNotAllowedBlockActionReason)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ActionModelBlockActionNotAllowedBlockActionReason) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


