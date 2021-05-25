# ServiceOfferItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceDate** | **string** | The date this service is delivered | 
**Amount** | [**AmountModel**](AmountModel.md) |  | 
**IsDefaultDate** | **bool** | Depending on the postNextDay setting of the service it will by default be posted before or after midnight.  Breakfast is usually delivered on the next morning, so all the dates from the day after arrival to the departure day  are default dates and will have this flag set to true. Those are also the dates the service will be booked for if  you do not specify dates in the book-service call. Still, you can override this and also book the dates set to IsDefaultDate &#x3D; false. | 
**IsMandatory** | **bool** | Rate plans can have additional services. When booking an offer for such rate plans, those services are automatically booked.  They are marked as mandatory and they cannot be removed. | 

## Methods

### NewServiceOfferItemModel

`func NewServiceOfferItemModel(serviceDate string, amount AmountModel, isDefaultDate bool, isMandatory bool, ) *ServiceOfferItemModel`

NewServiceOfferItemModel instantiates a new ServiceOfferItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOfferItemModelWithDefaults

`func NewServiceOfferItemModelWithDefaults() *ServiceOfferItemModel`

NewServiceOfferItemModelWithDefaults instantiates a new ServiceOfferItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceDate

`func (o *ServiceOfferItemModel) GetServiceDate() string`

GetServiceDate returns the ServiceDate field if non-nil, zero value otherwise.

### GetServiceDateOk

`func (o *ServiceOfferItemModel) GetServiceDateOk() (*string, bool)`

GetServiceDateOk returns a tuple with the ServiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDate

`func (o *ServiceOfferItemModel) SetServiceDate(v string)`

SetServiceDate sets ServiceDate field to given value.


### GetAmount

`func (o *ServiceOfferItemModel) GetAmount() AmountModel`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServiceOfferItemModel) GetAmountOk() (*AmountModel, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServiceOfferItemModel) SetAmount(v AmountModel)`

SetAmount sets Amount field to given value.


### GetIsDefaultDate

`func (o *ServiceOfferItemModel) GetIsDefaultDate() bool`

GetIsDefaultDate returns the IsDefaultDate field if non-nil, zero value otherwise.

### GetIsDefaultDateOk

`func (o *ServiceOfferItemModel) GetIsDefaultDateOk() (*bool, bool)`

GetIsDefaultDateOk returns a tuple with the IsDefaultDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultDate

`func (o *ServiceOfferItemModel) SetIsDefaultDate(v bool)`

SetIsDefaultDate sets IsDefaultDate field to given value.


### GetIsMandatory

`func (o *ServiceOfferItemModel) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *ServiceOfferItemModel) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *ServiceOfferItemModel) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


