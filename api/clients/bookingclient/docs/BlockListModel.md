# BlockListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | [**[]BlockItemModel**](BlockItemModel.md) | List of blocks | 
**Count** | **int64** | Total count of items | 

## Methods

### NewBlockListModel

`func NewBlockListModel(blocks []BlockItemModel, count int64, ) *BlockListModel`

NewBlockListModel instantiates a new BlockListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockListModelWithDefaults

`func NewBlockListModelWithDefaults() *BlockListModel`

NewBlockListModelWithDefaults instantiates a new BlockListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *BlockListModel) GetBlocks() []BlockItemModel`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BlockListModel) GetBlocksOk() (*[]BlockItemModel, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BlockListModel) SetBlocks(v []BlockItemModel)`

SetBlocks sets Blocks field to given value.


### GetCount

`func (o *BlockListModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BlockListModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BlockListModel) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


