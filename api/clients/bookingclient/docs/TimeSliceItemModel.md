# TimeSliceItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | Date and time the time slice begins&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**To** | **time.Time** | Date and time the time slice ends&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
**Offers** | Pointer to [**[]TimeSliceOfferItemModel**](TimeSliceOfferItemModel.md) | List of offers for this time slice | [optional] 

## Methods

### NewTimeSliceItemModel

`func NewTimeSliceItemModel(from time.Time, to time.Time, ) *TimeSliceItemModel`

NewTimeSliceItemModel instantiates a new TimeSliceItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSliceItemModelWithDefaults

`func NewTimeSliceItemModelWithDefaults() *TimeSliceItemModel`

NewTimeSliceItemModelWithDefaults instantiates a new TimeSliceItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TimeSliceItemModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TimeSliceItemModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TimeSliceItemModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *TimeSliceItemModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TimeSliceItemModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TimeSliceItemModel) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetOffers

`func (o *TimeSliceItemModel) GetOffers() []TimeSliceOfferItemModel`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *TimeSliceItemModel) GetOffersOk() (*[]TimeSliceOfferItemModel, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *TimeSliceItemModel) SetOffers(v []TimeSliceOfferItemModel)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *TimeSliceItemModel) HasOffers() bool`

HasOffers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


