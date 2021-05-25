# OfferCancellationFeeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code of the cancellation policy applied | 
**Name** | **string** | The name of the cancellation policy applied | 
**Description** | **string** | The description of the cancellation policy applied | 
**DueDateTime** | **time.Time** | The date and time the cancellation fee will be due. After that time this fee will  be charged in case of cancellation&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Fee** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 

## Methods

### NewOfferCancellationFeeModel

`func NewOfferCancellationFeeModel(code string, name string, description string, dueDateTime time.Time, fee MonetaryValueModel, ) *OfferCancellationFeeModel`

NewOfferCancellationFeeModel instantiates a new OfferCancellationFeeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferCancellationFeeModelWithDefaults

`func NewOfferCancellationFeeModelWithDefaults() *OfferCancellationFeeModel`

NewOfferCancellationFeeModelWithDefaults instantiates a new OfferCancellationFeeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OfferCancellationFeeModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OfferCancellationFeeModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OfferCancellationFeeModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *OfferCancellationFeeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfferCancellationFeeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfferCancellationFeeModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OfferCancellationFeeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OfferCancellationFeeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OfferCancellationFeeModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDueDateTime

`func (o *OfferCancellationFeeModel) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *OfferCancellationFeeModel) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *OfferCancellationFeeModel) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.


### GetFee

`func (o *OfferCancellationFeeModel) GetFee() MonetaryValueModel`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *OfferCancellationFeeModel) GetFeeOk() (*MonetaryValueModel, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *OfferCancellationFeeModel) SetFee(v MonetaryValueModel)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


