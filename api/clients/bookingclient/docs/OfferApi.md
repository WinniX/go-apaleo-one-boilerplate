# \OfferApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingOfferIndexGet**](OfferApi.md#BookingOfferIndexGet) | **Get** /booking/v1/offer-index | Returns offers with rates and availabilities for the specified range.
[**BookingOffersGet**](OfferApi.md#BookingOffersGet) | **Get** /booking/v1/offers | Returns offers for one specific stay.
[**BookingRatePlanOffersGet**](OfferApi.md#BookingRatePlanOffersGet) | **Get** /booking/v1/rate-plan-offers | Returns offers for a specific rate plan.
[**BookingServiceOffersGet**](OfferApi.md#BookingServiceOffersGet) | **Get** /booking/v1/service-offers | Returns service offers for one specific stay.



## BookingOfferIndexGet

> TimeSliceListModel BookingOfferIndexGet(ctx).RatePlanId(ratePlanId).From(from).To(to).ChannelCode(channelCode).PageNumber(pageNumber).PageSize(pageSize).Execute()

Returns offers with rates and availabilities for the specified range.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ratePlanId := "ratePlanId_example" // string | 
    from := "from_example" // string | <br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    to := "to_example" // string | <br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    channelCode := "channelCode_example" // string | 
    pageNumber := int32(56) // int32 | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. (optional) (default to 1)
    pageSize := int32(56) // int32 | Page size. If this is not set, the pageNumber will be ignored and all values returned. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfferApi.BookingOfferIndexGet(context.Background()).RatePlanId(ratePlanId).From(from).To(to).ChannelCode(channelCode).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferApi.BookingOfferIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingOfferIndexGet`: TimeSliceListModel
    fmt.Fprintf(os.Stdout, "Response from `OfferApi.BookingOfferIndexGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingOfferIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratePlanId** | **string** |  | 
 **from** | **string** | &lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **string** | &lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **channelCode** | **string** |  | 
 **pageNumber** | **int32** | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. | [default to 1]
 **pageSize** | **int32** | Page size. If this is not set, the pageNumber will be ignored and all values returned. | [default to 100]

### Return type

[**TimeSliceListModel**](TimeSliceListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingOffersGet

> StayOffersModel BookingOffersGet(ctx).PropertyId(propertyId).Arrival(arrival).Departure(departure).Adults(adults).TimeSliceTemplate(timeSliceTemplate).TimeSliceDefinitionIds(timeSliceDefinitionIds).UnitGroupIds(unitGroupIds).UnitGroupTypes(unitGroupTypes).ChannelCode(channelCode).PromoCode(promoCode).CorporateCode(corporateCode).ChildrenAges(childrenAges).IncludeUnavailable(includeUnavailable).Execute()

Returns offers for one specific stay.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    propertyId := "propertyId_example" // string | The property ID
    arrival := "arrival_example" // string | Date and optional time of arrival<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    departure := "departure_example" // string | Date and optional time of departure. Cannot be more than 5 years after arrival.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    adults := int32(56) // int32 | The number of adults you want offers for
    timeSliceTemplate := "timeSliceTemplate_example" // string | The time slice template used to filter the rate plans, defaults to 'over night' (optional)
    timeSliceDefinitionIds := []string{"Inner_example"} // []string | Time slice definition IDs, used to filter rate plans (optional)
    unitGroupIds := []string{"Inner_example"} // []string | Unit group IDs, used to filter rate plans (optional)
    unitGroupTypes := []string{"UnitGroupTypes_example"} // []string | Unit group types, used to filter rate plans (optional)
    channelCode := "channelCode_example" // string | Channel code, used to filter the rate plans (optional)
    promoCode := "promoCode_example" // string | The promo code associated with a certain special offer (optional)
    corporateCode := "corporateCode_example" // string | The code associated with a corporate rate (optional)
    childrenAges := []int32{int32(123)} // []int32 | The ages of the children you want offers for (optional)
    includeUnavailable := true // bool | Return also offers that are currently not publicly bookable as restrictions are violated. By default only available offers are returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfferApi.BookingOffersGet(context.Background()).PropertyId(propertyId).Arrival(arrival).Departure(departure).Adults(adults).TimeSliceTemplate(timeSliceTemplate).TimeSliceDefinitionIds(timeSliceDefinitionIds).UnitGroupIds(unitGroupIds).UnitGroupTypes(unitGroupTypes).ChannelCode(channelCode).PromoCode(promoCode).CorporateCode(corporateCode).ChildrenAges(childrenAges).IncludeUnavailable(includeUnavailable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferApi.BookingOffersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingOffersGet`: StayOffersModel
    fmt.Fprintf(os.Stdout, "Response from `OfferApi.BookingOffersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingOffersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **propertyId** | **string** | The property ID | 
 **arrival** | **string** | Date and optional time of arrival&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **departure** | **string** | Date and optional time of departure. Cannot be more than 5 years after arrival.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **adults** | **int32** | The number of adults you want offers for | 
 **timeSliceTemplate** | **string** | The time slice template used to filter the rate plans, defaults to &#39;over night&#39; | 
 **timeSliceDefinitionIds** | **[]string** | Time slice definition IDs, used to filter rate plans | 
 **unitGroupIds** | **[]string** | Unit group IDs, used to filter rate plans | 
 **unitGroupTypes** | **[]string** | Unit group types, used to filter rate plans | 
 **channelCode** | **string** | Channel code, used to filter the rate plans | 
 **promoCode** | **string** | The promo code associated with a certain special offer | 
 **corporateCode** | **string** | The code associated with a corporate rate | 
 **childrenAges** | **[]int32** | The ages of the children you want offers for | 
 **includeUnavailable** | **bool** | Return also offers that are currently not publicly bookable as restrictions are violated. By default only available offers are returned | 

### Return type

[**StayOffersModel**](StayOffersModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingRatePlanOffersGet

> StayOffersModel BookingRatePlanOffersGet(ctx).RatePlanId(ratePlanId).Arrival(arrival).Departure(departure).Adults(adults).ChannelCode(channelCode).ChildrenAges(childrenAges).IncludeUnavailable(includeUnavailable).Execute()

Returns offers for a specific rate plan.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ratePlanId := "ratePlanId_example" // string | The rate plan ID
    arrival := "arrival_example" // string | Date and optional time of arrival<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    departure := "departure_example" // string | Date and optional time of departure. Cannot be more than 5 years after arrival.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    adults := int32(56) // int32 | The number of adults you want offers for
    channelCode := "channelCode_example" // string | The channel code (optional)
    childrenAges := []int32{int32(123)} // []int32 | The ages of the children you want offers for (optional)
    includeUnavailable := true // bool | Return also offers that are currently not publicly bookable as restrictions are violated. By default only available offers are returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfferApi.BookingRatePlanOffersGet(context.Background()).RatePlanId(ratePlanId).Arrival(arrival).Departure(departure).Adults(adults).ChannelCode(channelCode).ChildrenAges(childrenAges).IncludeUnavailable(includeUnavailable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferApi.BookingRatePlanOffersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingRatePlanOffersGet`: StayOffersModel
    fmt.Fprintf(os.Stdout, "Response from `OfferApi.BookingRatePlanOffersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingRatePlanOffersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratePlanId** | **string** | The rate plan ID | 
 **arrival** | **string** | Date and optional time of arrival&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **departure** | **string** | Date and optional time of departure. Cannot be more than 5 years after arrival.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **adults** | **int32** | The number of adults you want offers for | 
 **channelCode** | **string** | The channel code | 
 **childrenAges** | **[]int32** | The ages of the children you want offers for | 
 **includeUnavailable** | **bool** | Return also offers that are currently not publicly bookable as restrictions are violated. By default only available offers are returned | 

### Return type

[**StayOffersModel**](StayOffersModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingServiceOffersGet

> ServiceOffersModel BookingServiceOffersGet(ctx).RatePlanId(ratePlanId).Arrival(arrival).Departure(departure).Adults(adults).ChannelCode(channelCode).ChildrenAges(childrenAges).OnlyDefaultDates(onlyDefaultDates).Execute()

Returns service offers for one specific stay.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ratePlanId := "ratePlanId_example" // string | The rate plan ID
    arrival := "arrival_example" // string | Date and optional time of arrival<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    departure := "departure_example" // string | Date and optional time of departure. Cannot be more than 5 years after arrival.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
    adults := int32(56) // int32 | The number of adults you want offers for
    channelCode := "channelCode_example" // string | The channel code used to filter the services (optional)
    childrenAges := []int32{int32(123)} // []int32 | The ages of the children you want offers for (optional)
    onlyDefaultDates := true // bool | Depending on the postNextDay setting of a service it will be posted before or after midnight.  Breakfast is usually delivered on the next morning, having 'postNextDay' set to true. Its 'default dates' are from the day after  arrival until the departure day. For services like dinner 'postNextDay' is false, and default dates are day of arrival until one  day before departure.  With this query parameter set to 'false', you can also ask for dates outside of those default dates. It defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfferApi.BookingServiceOffersGet(context.Background()).RatePlanId(ratePlanId).Arrival(arrival).Departure(departure).Adults(adults).ChannelCode(channelCode).ChildrenAges(childrenAges).OnlyDefaultDates(onlyDefaultDates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfferApi.BookingServiceOffersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingServiceOffersGet`: ServiceOffersModel
    fmt.Fprintf(os.Stdout, "Response from `OfferApi.BookingServiceOffersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingServiceOffersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratePlanId** | **string** | The rate plan ID | 
 **arrival** | **string** | Date and optional time of arrival&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **departure** | **string** | Date and optional time of departure. Cannot be more than 5 years after arrival.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **adults** | **int32** | The number of adults you want offers for | 
 **channelCode** | **string** | The channel code used to filter the services | 
 **childrenAges** | **[]int32** | The ages of the children you want offers for | 
 **onlyDefaultDates** | **bool** | Depending on the postNextDay setting of a service it will be posted before or after midnight.  Breakfast is usually delivered on the next morning, having &#39;postNextDay&#39; set to true. Its &#39;default dates&#39; are from the day after  arrival until the departure day. For services like dinner &#39;postNextDay&#39; is false, and default dates are day of arrival until one  day before departure.  With this query parameter set to &#39;false&#39;, you can also ask for dates outside of those default dates. It defaults to true. | 

### Return type

[**ServiceOffersModel**](ServiceOffersModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

