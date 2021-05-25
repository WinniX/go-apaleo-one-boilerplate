# ReservationNoShowFeeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the no-show policy applied | 
**Code** | **string** | The code of the no-show policy applied | 
**Name** | **string** | The name of the no-show policy applied | 
**Description** | **string** | The description of the no-show policy applied | 
**Fee** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewReservationNoShowFeeModel

`func NewReservationNoShowFeeModel(id string, code string, name string, description string, fee MonetaryValueModel, ) *ReservationNoShowFeeModel`

NewReservationNoShowFeeModel instantiates a new ReservationNoShowFeeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationNoShowFeeModelWithDefaults

`func NewReservationNoShowFeeModelWithDefaults() *ReservationNoShowFeeModel`

NewReservationNoShowFeeModelWithDefaults instantiates a new ReservationNoShowFeeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationNoShowFeeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationNoShowFeeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationNoShowFeeModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *ReservationNoShowFeeModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReservationNoShowFeeModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReservationNoShowFeeModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *ReservationNoShowFeeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReservationNoShowFeeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReservationNoShowFeeModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReservationNoShowFeeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReservationNoShowFeeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReservationNoShowFeeModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFee

`func (o *ReservationNoShowFeeModel) GetFee() MonetaryValueModel`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ReservationNoShowFeeModel) GetFeeOk() (*MonetaryValueModel, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ReservationNoShowFeeModel) SetFee(v MonetaryValueModel)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


