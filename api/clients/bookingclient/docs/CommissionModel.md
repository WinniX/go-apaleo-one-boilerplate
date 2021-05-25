# CommissionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommissionAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**BeforeCommissionAmount** | Pointer to [**MonetaryValueModel**](MonetaryValueModel.md) |  | [optional] 

## Methods

### NewCommissionModel

`func NewCommissionModel(commissionAmount MonetaryValueModel, ) *CommissionModel`

NewCommissionModel instantiates a new CommissionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionModelWithDefaults

`func NewCommissionModelWithDefaults() *CommissionModel`

NewCommissionModelWithDefaults instantiates a new CommissionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommissionAmount

`func (o *CommissionModel) GetCommissionAmount() MonetaryValueModel`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *CommissionModel) GetCommissionAmountOk() (*MonetaryValueModel, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *CommissionModel) SetCommissionAmount(v MonetaryValueModel)`

SetCommissionAmount sets CommissionAmount field to given value.


### GetBeforeCommissionAmount

`func (o *CommissionModel) GetBeforeCommissionAmount() MonetaryValueModel`

GetBeforeCommissionAmount returns the BeforeCommissionAmount field if non-nil, zero value otherwise.

### GetBeforeCommissionAmountOk

`func (o *CommissionModel) GetBeforeCommissionAmountOk() (*MonetaryValueModel, bool)`

GetBeforeCommissionAmountOk returns a tuple with the BeforeCommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeCommissionAmount

`func (o *CommissionModel) SetBeforeCommissionAmount(v MonetaryValueModel)`

SetBeforeCommissionAmount sets BeforeCommissionAmount field to given value.

### HasBeforeCommissionAmount

`func (o *CommissionModel) HasBeforeCommissionAmount() bool`

HasBeforeCommissionAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


