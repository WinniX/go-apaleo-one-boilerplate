# AutoAssignedUnitItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | [**EmbeddedUnitModel**](EmbeddedUnitModel.md) |  | 
**From** | **time.Time** | The start date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | The end date and time for this time slice&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 

## Methods

### NewAutoAssignedUnitItemModel

`func NewAutoAssignedUnitItemModel(unit EmbeddedUnitModel, from time.Time, to time.Time, ) *AutoAssignedUnitItemModel`

NewAutoAssignedUnitItemModel instantiates a new AutoAssignedUnitItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAssignedUnitItemModelWithDefaults

`func NewAutoAssignedUnitItemModelWithDefaults() *AutoAssignedUnitItemModel`

NewAutoAssignedUnitItemModelWithDefaults instantiates a new AutoAssignedUnitItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *AutoAssignedUnitItemModel) GetUnit() EmbeddedUnitModel`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AutoAssignedUnitItemModel) GetUnitOk() (*EmbeddedUnitModel, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AutoAssignedUnitItemModel) SetUnit(v EmbeddedUnitModel)`

SetUnit sets Unit field to given value.


### GetFrom

`func (o *AutoAssignedUnitItemModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AutoAssignedUnitItemModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AutoAssignedUnitItemModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *AutoAssignedUnitItemModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AutoAssignedUnitItemModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AutoAssignedUnitItemModel) SetTo(v time.Time)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


