# PerOccupancyPriceItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adults** | **int32** | Number of adults | 
**Price** | [**PriceModel**](PriceModel.md) |  | 

## Methods

### NewPerOccupancyPriceItemModel

`func NewPerOccupancyPriceItemModel(adults int32, price PriceModel, ) *PerOccupancyPriceItemModel`

NewPerOccupancyPriceItemModel instantiates a new PerOccupancyPriceItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerOccupancyPriceItemModelWithDefaults

`func NewPerOccupancyPriceItemModelWithDefaults() *PerOccupancyPriceItemModel`

NewPerOccupancyPriceItemModelWithDefaults instantiates a new PerOccupancyPriceItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdults

`func (o *PerOccupancyPriceItemModel) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *PerOccupancyPriceItemModel) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *PerOccupancyPriceItemModel) SetAdults(v int32)`

SetAdults sets Adults field to given value.


### GetPrice

`func (o *PerOccupancyPriceItemModel) GetPrice() PriceModel`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PerOccupancyPriceItemModel) GetPriceOk() (*PriceModel, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PerOccupancyPriceItemModel) SetPrice(v PriceModel)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


