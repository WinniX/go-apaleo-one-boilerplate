# OfferUnitGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unit group id | 
**Code** | **string** | The code for the unit group that can be shown in reports and table views | 
**Name** | **string** | The name for the unit group | 
**Description** | **string** | The description for the unit group | 
**MaxPersons** | **int32** | Maximum number of persons for the unit group | 
**Rank** | Pointer to **int32** | The unit group rank | [optional] 

## Methods

### NewOfferUnitGroupModel

`func NewOfferUnitGroupModel(id string, code string, name string, description string, maxPersons int32, ) *OfferUnitGroupModel`

NewOfferUnitGroupModel instantiates a new OfferUnitGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferUnitGroupModelWithDefaults

`func NewOfferUnitGroupModelWithDefaults() *OfferUnitGroupModel`

NewOfferUnitGroupModelWithDefaults instantiates a new OfferUnitGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OfferUnitGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfferUnitGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfferUnitGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *OfferUnitGroupModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OfferUnitGroupModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OfferUnitGroupModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *OfferUnitGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfferUnitGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfferUnitGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OfferUnitGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OfferUnitGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OfferUnitGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMaxPersons

`func (o *OfferUnitGroupModel) GetMaxPersons() int32`

GetMaxPersons returns the MaxPersons field if non-nil, zero value otherwise.

### GetMaxPersonsOk

`func (o *OfferUnitGroupModel) GetMaxPersonsOk() (*int32, bool)`

GetMaxPersonsOk returns a tuple with the MaxPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPersons

`func (o *OfferUnitGroupModel) SetMaxPersons(v int32)`

SetMaxPersons sets MaxPersons field to given value.


### GetRank

`func (o *OfferUnitGroupModel) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *OfferUnitGroupModel) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *OfferUnitGroupModel) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *OfferUnitGroupModel) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


