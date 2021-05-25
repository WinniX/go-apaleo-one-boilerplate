# ReplaceBlockModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Start date and time from which the inventory will be blocked&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **string** | End date and time until which the inventory will be blocked. Cannot be more than 5 years after the start date.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**GrossDailyRate** | [**MonetaryValueModel**](MonetaryValueModel.md) |  | 
**TimeSlices** | [**[]CreateBlockTimeSliceModel**](CreateBlockTimeSliceModel.md) | The list of time slices | 

## Methods

### NewReplaceBlockModel

`func NewReplaceBlockModel(from string, to string, grossDailyRate MonetaryValueModel, timeSlices []CreateBlockTimeSliceModel, ) *ReplaceBlockModel`

NewReplaceBlockModel instantiates a new ReplaceBlockModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceBlockModelWithDefaults

`func NewReplaceBlockModelWithDefaults() *ReplaceBlockModel`

NewReplaceBlockModelWithDefaults instantiates a new ReplaceBlockModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *ReplaceBlockModel) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReplaceBlockModel) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReplaceBlockModel) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *ReplaceBlockModel) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReplaceBlockModel) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReplaceBlockModel) SetTo(v string)`

SetTo sets To field to given value.


### GetGrossDailyRate

`func (o *ReplaceBlockModel) GetGrossDailyRate() MonetaryValueModel`

GetGrossDailyRate returns the GrossDailyRate field if non-nil, zero value otherwise.

### GetGrossDailyRateOk

`func (o *ReplaceBlockModel) GetGrossDailyRateOk() (*MonetaryValueModel, bool)`

GetGrossDailyRateOk returns a tuple with the GrossDailyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossDailyRate

`func (o *ReplaceBlockModel) SetGrossDailyRate(v MonetaryValueModel)`

SetGrossDailyRate sets GrossDailyRate field to given value.


### GetTimeSlices

`func (o *ReplaceBlockModel) GetTimeSlices() []CreateBlockTimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *ReplaceBlockModel) GetTimeSlicesOk() (*[]CreateBlockTimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *ReplaceBlockModel) SetTimeSlices(v []CreateBlockTimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


