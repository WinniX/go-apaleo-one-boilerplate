# ReservationCancellationFeeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the cancellation policy applied | 
**Code** | **string** | The code of the cancellation policy applied | 
**Name** | **string** | The name of the cancellation policy applied | 
**Description** | **string** | The description of the cancellation policy applied | 
**DueDateTime** | **time.Time** | The date and time the cancellation fee will be due. After that time this fee will  be charged in case of cancellation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Fee** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewReservationCancellationFeeModel

`func NewReservationCancellationFeeModel(id string, code string, name string, description string, dueDateTime time.Time, fee MonetaryValueModel, ) *ReservationCancellationFeeModel`

NewReservationCancellationFeeModel instantiates a new ReservationCancellationFeeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationCancellationFeeModelWithDefaults

`func NewReservationCancellationFeeModelWithDefaults() *ReservationCancellationFeeModel`

NewReservationCancellationFeeModelWithDefaults instantiates a new ReservationCancellationFeeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationCancellationFeeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationCancellationFeeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationCancellationFeeModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *ReservationCancellationFeeModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReservationCancellationFeeModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReservationCancellationFeeModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *ReservationCancellationFeeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReservationCancellationFeeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReservationCancellationFeeModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReservationCancellationFeeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReservationCancellationFeeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReservationCancellationFeeModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDueDateTime

`func (o *ReservationCancellationFeeModel) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *ReservationCancellationFeeModel) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *ReservationCancellationFeeModel) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.


### GetFee

`func (o *ReservationCancellationFeeModel) GetFee() MonetaryValueModel`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ReservationCancellationFeeModel) GetFeeOk() (*MonetaryValueModel, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ReservationCancellationFeeModel) SetFee(v MonetaryValueModel)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


