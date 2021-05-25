# DesiredStayDetailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arrival** | **string** | Date and optional time of arrival&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Departure** | **string** | Date and optional time of departure. Cannot be more than 5 years after arrival.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Adults** | **int32** | Number of adults | 
**ChildrenAges** | Pointer to **[]int32** | Ages of the children | [optional] 
**Requote** | Pointer to **bool** | Whether the prices for time slices with no change to the rate plan should be re-quoted based on current prices, or if  only additions like change of number of adults should be calculated. Defaults to &#39;false&#39;. | [optional] 
**TimeSlices** | [**[]DesiredTimeSliceModel**](DesiredTimeSliceModel.md) | The list of time slices | 

## Methods

### NewDesiredStayDetailsModel

`func NewDesiredStayDetailsModel(arrival string, departure string, adults int32, timeSlices []DesiredTimeSliceModel, ) *DesiredStayDetailsModel`

NewDesiredStayDetailsModel instantiates a new DesiredStayDetailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesiredStayDetailsModelWithDefaults

`func NewDesiredStayDetailsModelWithDefaults() *DesiredStayDetailsModel`

NewDesiredStayDetailsModelWithDefaults instantiates a new DesiredStayDetailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrival

`func (o *DesiredStayDetailsModel) GetArrival() string`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *DesiredStayDetailsModel) GetArrivalOk() (*string, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *DesiredStayDetailsModel) SetArrival(v string)`

SetArrival sets Arrival field to given value.


### GetDeparture

`func (o *DesiredStayDetailsModel) GetDeparture() string`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *DesiredStayDetailsModel) GetDepartureOk() (*string, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *DesiredStayDetailsModel) SetDeparture(v string)`

SetDeparture sets Departure field to given value.


### GetAdults

`func (o *DesiredStayDetailsModel) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *DesiredStayDetailsModel) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *DesiredStayDetailsModel) SetAdults(v int32)`

SetAdults sets Adults field to given value.


### GetChildrenAges

`func (o *DesiredStayDetailsModel) GetChildrenAges() []int32`

GetChildrenAges returns the ChildrenAges field if non-nil, zero value otherwise.

### GetChildrenAgesOk

`func (o *DesiredStayDetailsModel) GetChildrenAgesOk() (*[]int32, bool)`

GetChildrenAgesOk returns a tuple with the ChildrenAges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAges

`func (o *DesiredStayDetailsModel) SetChildrenAges(v []int32)`

SetChildrenAges sets ChildrenAges field to given value.

### HasChildrenAges

`func (o *DesiredStayDetailsModel) HasChildrenAges() bool`

HasChildrenAges returns a boolean if a field has been set.

### GetRequote

`func (o *DesiredStayDetailsModel) GetRequote() bool`

GetRequote returns the Requote field if non-nil, zero value otherwise.

### GetRequoteOk

`func (o *DesiredStayDetailsModel) GetRequoteOk() (*bool, bool)`

GetRequoteOk returns a tuple with the Requote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequote

`func (o *DesiredStayDetailsModel) SetRequote(v bool)`

SetRequote sets Requote field to given value.

### HasRequote

`func (o *DesiredStayDetailsModel) HasRequote() bool`

HasRequote returns a boolean if a field has been set.

### GetTimeSlices

`func (o *DesiredStayDetailsModel) GetTimeSlices() []DesiredTimeSliceModel`

GetTimeSlices returns the TimeSlices field if non-nil, zero value otherwise.

### GetTimeSlicesOk

`func (o *DesiredStayDetailsModel) GetTimeSlicesOk() (*[]DesiredTimeSliceModel, bool)`

GetTimeSlicesOk returns a tuple with the TimeSlices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlices

`func (o *DesiredStayDetailsModel) SetTimeSlices(v []DesiredTimeSliceModel)`

SetTimeSlices sets TimeSlices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


