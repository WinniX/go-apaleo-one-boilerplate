# DesiredTimeSliceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatePlanId** | **string** | The rate plan id for this time slice | 
**TotalGrossAmount** | Pointer to [**MonetaryValueModel**](MonetaryValueModel.md) |  | [optional] 

## Methods

### NewDesiredTimeSliceModel

`func NewDesiredTimeSliceModel(ratePlanId string, ) *DesiredTimeSliceModel`

NewDesiredTimeSliceModel instantiates a new DesiredTimeSliceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesiredTimeSliceModelWithDefaults

`func NewDesiredTimeSliceModelWithDefaults() *DesiredTimeSliceModel`

NewDesiredTimeSliceModelWithDefaults instantiates a new DesiredTimeSliceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatePlanId

`func (o *DesiredTimeSliceModel) GetRatePlanId() string`

GetRatePlanId returns the RatePlanId field if non-nil, zero value otherwise.

### GetRatePlanIdOk

`func (o *DesiredTimeSliceModel) GetRatePlanIdOk() (*string, bool)`

GetRatePlanIdOk returns a tuple with the RatePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlanId

`func (o *DesiredTimeSliceModel) SetRatePlanId(v string)`

SetRatePlanId sets RatePlanId field to given value.


### GetTotalGrossAmount

`func (o *DesiredTimeSliceModel) GetTotalGrossAmount() MonetaryValueModel`

GetTotalGrossAmount returns the TotalGrossAmount field if non-nil, zero value otherwise.

### GetTotalGrossAmountOk

`func (o *DesiredTimeSliceModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool)`

GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossAmount

`func (o *DesiredTimeSliceModel) SetTotalGrossAmount(v MonetaryValueModel)`

SetTotalGrossAmount sets TotalGrossAmount field to given value.

### HasTotalGrossAmount

`func (o *DesiredTimeSliceModel) HasTotalGrossAmount() bool`

HasTotalGrossAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


