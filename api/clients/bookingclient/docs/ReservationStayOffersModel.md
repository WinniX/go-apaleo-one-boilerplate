# ReservationStayOffersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**EmbeddedPropertyModel**](EmbeddedPropertyModel.md) |  | 
**Offers** | [**[]ReservationStayOfferModel**](ReservationStayOfferModel.md) | List of offered unit groups with rates | 

## Methods

### NewReservationStayOffersModel

`func NewReservationStayOffersModel(property EmbeddedPropertyModel, offers []ReservationStayOfferModel, ) *ReservationStayOffersModel`

NewReservationStayOffersModel instantiates a new ReservationStayOffersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationStayOffersModelWithDefaults

`func NewReservationStayOffersModelWithDefaults() *ReservationStayOffersModel`

NewReservationStayOffersModelWithDefaults instantiates a new ReservationStayOffersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ReservationStayOffersModel) GetProperty() EmbeddedPropertyModel`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ReservationStayOffersModel) GetPropertyOk() (*EmbeddedPropertyModel, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ReservationStayOffersModel) SetProperty(v EmbeddedPropertyModel)`

SetProperty sets Property field to given value.


### GetOffers

`func (o *ReservationStayOffersModel) GetOffers() []ReservationStayOfferModel`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *ReservationStayOffersModel) GetOffersOk() (*[]ReservationStayOfferModel, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *ReservationStayOffersModel) SetOffers(v []ReservationStayOfferModel)`

SetOffers sets Offers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


