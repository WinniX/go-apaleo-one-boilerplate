# StayOffersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**Offers** | [**[]OfferModel**](OfferModel.md) | List of offered unit groups with rates | 

## Methods

### NewStayOffersModel

`func NewStayOffersModel(property EmbeddedPropertyModel, offers []OfferModel, ) *StayOffersModel`

NewStayOffersModel instantiates a new StayOffersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStayOffersModelWithDefaults

`func NewStayOffersModelWithDefaults() *StayOffersModel`

NewStayOffersModelWithDefaults instantiates a new StayOffersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *StayOffersModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StayOffersModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StayOffersModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetOffers

`func (o *StayOffersModel) GetOffers() []OfferModel`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *StayOffersModel) GetOffersOk() (*[]OfferModel, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *StayOffersModel) SetOffers(v []OfferModel)`

SetOffers sets Offers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


