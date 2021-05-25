# CreateReservationTimeSliceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatePlanId** | **string** | The rate plan id for this time slice | 
**TotalAmount** | Pointer to [**MonetaryValueModel**](MonetaryValueModel.md) |  | [optional] 

## Methods

### NewCreateReservationTimeSliceModel

`func NewCreateReservationTimeSliceModel(ratePlanId string, ) *CreateReservationTimeSliceModel`

NewCreateReservationTimeSliceModel instantiates a new CreateReservationTimeSliceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReservationTimeSliceModelWithDefaults

`func NewCreateReservationTimeSliceModelWithDefaults() *CreateReservationTimeSliceModel`

NewCreateReservationTimeSliceModelWithDefaults instantiates a new CreateReservationTimeSliceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatePlanId

`func (o *CreateReservationTimeSliceModel) GetRatePlanId() string`

GetRatePlanId returns the RatePlanId field if non-nil, zero value otherwise.

### GetRatePlanIdOk

`func (o *CreateReservationTimeSliceModel) GetRatePlanIdOk() (*string, bool)`

GetRatePlanIdOk returns a tuple with the RatePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlanId

`func (o *CreateReservationTimeSliceModel) SetRatePlanId(v string)`

SetRatePlanId sets RatePlanId field to given value.


### GetTotalAmount

`func (o *CreateReservationTimeSliceModel) GetTotalAmount() MonetaryValueModel`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CreateReservationTimeSliceModel) GetTotalAmountOk() (*MonetaryValueModel, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CreateReservationTimeSliceModel) SetTotalAmount(v MonetaryValueModel)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *CreateReservationTimeSliceModel) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


