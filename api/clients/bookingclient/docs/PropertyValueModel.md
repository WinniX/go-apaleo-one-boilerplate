# PropertyValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**TotalGrossAmount** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**Balance** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewPropertyValueModel

`func NewPropertyValueModel(property EmbeddedPropertyModel, totalGrossAmount MonetaryValueModel, balance MonetaryValueModel, ) *PropertyValueModel`

NewPropertyValueModel instantiates a new PropertyValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyValueModelWithDefaults

`func NewPropertyValueModelWithDefaults() *PropertyValueModel`

NewPropertyValueModelWithDefaults instantiates a new PropertyValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *PropertyValueModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PropertyValueModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PropertyValueModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetTotalGrossAmount

`func (o *PropertyValueModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *PropertyValueModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *PropertyValueModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.


### GetBalance

`func (o *PropertyValueModel) GetBalance() MonetaryValueModel`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PropertyValueModel) GetBalanceOk() (*MonetaryValueModel, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PropertyValueModel) SetBalance(v MonetaryValueModel)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


