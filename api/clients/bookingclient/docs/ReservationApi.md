# \ReservationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingReservationsByIdGet**](ReservationApi.md#BookingReservationsByIdGet) | **Get** /booking/v1/reservations/{id} | Returns a specific reservation.
[**BookingReservationsByIdOffersGet**](ReservationApi.md#BookingReservationsByIdOffersGet) | **Get** /booking/v1/reservations/{id}/offers | Returns offers for one specific reservation.
[**BookingReservationsByIdPatch**](ReservationApi.md#BookingReservationsByIdPatch) | **Patch** /booking/v1/reservations/{id} | Allows to modify certain reservation properties
[**BookingReservationsByIdServiceOffersGet**](ReservationApi.md#BookingReservationsByIdServiceOffersGet) | **Get** /booking/v1/reservations/{id}/service-offers | Returns service offers for one specific reservation.
[**BookingReservationsByIdServicesDelete**](ReservationApi.md#BookingReservationsByIdServicesDelete) | **Delete** /booking/v1/reservations/{id}/services | Removes a service from a reservation.
[**BookingReservationsByIdServicesGet**](ReservationApi.md#BookingReservationsByIdServicesGet) | **Get** /booking/v1/reservations/{id}/services | Returns the services booked for a specific reservation.
[**BookingReservationsGet**](ReservationApi.md#BookingReservationsGet) | **Get** /booking/v1/reservations | Returns a list of all reservations, filtered by the specified parameters.
[**BookingReservationscountGet**](ReservationApi.md#BookingReservationscountGet) | **Get** /booking/v1/reservations/$count | Returns the number of reservations fulfilling the criteria specified in the parameters.



## BookingReservationsByIdGet

> ReservationModel BookingReservationsByIdGet(ctx, id).Expand(expand).Execute()

Returns a specific reservation.



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
    id := "id_example" // string | Id of the reservation to be retrieved.
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: timeSlices, services, booker, actions, company, assignedUnits. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationsByIdGet(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationsByIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationsByIdGet`: ReservationModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationApi.BookingReservationsByIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationsByIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: timeSlices, services, booker, actions, company, assignedUnits. All other values will be silently ignored. | 

### Return type

[**ReservationModel**](ReservationModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationsByIdOffersGet

> ReservationStayOffersModel BookingReservationsByIdOffersGet(ctx, id).Arrival(arrival).Departure(departure).Adults(adults).ChildrenAges(childrenAges).ChannelCode(channelCode).PromoCode(promoCode).Requote(requote).IncludeUnavailable(includeUnavailable).Execute()

Returns offers for one specific reservation.



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
    id := "id_example" // string | Id of the reservation to be amended.
    arrival := "arrival_example" // string | Date and optional time of arrival<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    departure := "departure_example" // string | Date and optional time of departure. Cannot be more than 5 years after arrival.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    adults := int32(56) // int32 | Number of adults (optional)
    childrenAges := []int32{int32(123)} // []int32 | Ages of children (optional)
    channelCode := "channelCode_example" // string | The channel code used to filter the rate plans (optional)
    promoCode := "promoCode_example" // string | The promo code associated with a certain special offer, like corporate rate (optional)
    requote := true // bool | Whether the offers should be re-quoted based on current prices, or only additions like change of number of adults should be calculated.  Defaults to 'false' (optional)
    includeUnavailable := true // bool | Return also offers that are currently not publicly bookable as restrictions are violated. By default only available offers are returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationsByIdOffersGet(context.Background(), id).Arrival(arrival).Departure(departure).Adults(adults).ChildrenAges(childrenAges).ChannelCode(channelCode).PromoCode(promoCode).Requote(requote).IncludeUnavailable(includeUnavailable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationsByIdOffersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationsByIdOffersGet`: ReservationStayOffersModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationApi.BookingReservationsByIdOffersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation to be amended. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationsByIdOffersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arrival** | **string** | Date and optional time of arrival&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **departure** | **string** | Date and optional time of departure. Cannot be more than 5 years after arrival.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **adults** | **int32** | Number of adults | 
 **childrenAges** | **[]int32** | Ages of children | 
 **channelCode** | **string** | The channel code used to filter the rate plans | 
 **promoCode** | **string** | The promo code associated with a certain special offer, like corporate rate | 
 **requote** | **bool** | Whether the offers should be re-quoted based on current prices, or only additions like change of number of adults should be calculated.  Defaults to &#39;false&#39; | 
 **includeUnavailable** | **bool** | Return also offers that are currently not publicly bookable as restrictions are violated. By default only available offers are returned | 

### Return type

[**ReservationStayOffersModel**](ReservationStayOffersModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationsByIdPatch

> BookingReservationsByIdPatch(ctx, id).Body(body).Execute()

Allows to modify certain reservation properties



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
    id := "id_example" // string | Id of the reservation to be modified.
    body := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation | Define the list of operations to be applied to the resource. Learn more about JSON Patch here: http://jsonpatch.com/.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationsByIdPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationsByIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation to be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationsByIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]Operation**](Operation.md) | Define the list of operations to be applied to the resource. Learn more about JSON Patch here: http://jsonpatch.com/. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationsByIdServiceOffersGet

> ServiceOffersModel BookingReservationsByIdServiceOffersGet(ctx, id).ChannelCode(channelCode).OnlyDefaultDates(onlyDefaultDates).Execute()

Returns service offers for one specific reservation.



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
    id := "id_example" // string | Id of the reservation.
    channelCode := "channelCode_example" // string | The channel code used to filter the services (optional)
    onlyDefaultDates := true // bool | Depending on the postNextDay setting of a service it will by default be posted before or after midnight.  Breakfast is usually delivered on the next morning, so all the dates from the day after arrival to the departure day  are default dates and will have this flag set to true. For services like a dinner it is the other way around.  With this query parameter, you can also ask for the dates, that usually the service will not be booked. It defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationsByIdServiceOffersGet(context.Background(), id).ChannelCode(channelCode).OnlyDefaultDates(onlyDefaultDates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationsByIdServiceOffersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationsByIdServiceOffersGet`: ServiceOffersModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationApi.BookingReservationsByIdServiceOffersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationsByIdServiceOffersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelCode** | **string** | The channel code used to filter the services | 
 **onlyDefaultDates** | **bool** | Depending on the postNextDay setting of a service it will by default be posted before or after midnight.  Breakfast is usually delivered on the next morning, so all the dates from the day after arrival to the departure day  are default dates and will have this flag set to true. For services like a dinner it is the other way around.  With this query parameter, you can also ask for the dates, that usually the service will not be booked. It defaults to true. | 

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


## BookingReservationsByIdServicesDelete

> BookingReservationsByIdServicesDelete(ctx, id).ServiceId(serviceId).Execute()

Removes a service from a reservation.



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
    id := "id_example" // string | Id of the reservation.
    serviceId := "serviceId_example" // string | The id of the service to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationsByIdServicesDelete(context.Background(), id).ServiceId(serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationsByIdServicesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationsByIdServicesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceId** | **string** | The id of the service to delete | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationsByIdServicesGet

> ReservationServiceListModel BookingReservationsByIdServicesGet(ctx, id).Execute()

Returns the services booked for a specific reservation.



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
    id := "id_example" // string | Id of the reservation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationsByIdServicesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationsByIdServicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationsByIdServicesGet`: ReservationServiceListModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationApi.BookingReservationsByIdServicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationsByIdServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReservationServiceListModel**](ReservationServiceListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationsGet

> ReservationListModel BookingReservationsGet(ctx).BookingId(bookingId).PropertyIds(propertyIds).RatePlanIds(ratePlanIds).CompanyIds(companyIds).UnitIds(unitIds).UnitGroupIds(unitGroupIds).UnitGroupTypes(unitGroupTypes).BlockIds(blockIds).Status(status).DateFilter(dateFilter).From(from).To(to).ChannelCode(channelCode).Sources(sources).ValidationMessageCategory(validationMessageCategory).ExternalCode(externalCode).TextSearch(textSearch).BalanceFilter(balanceFilter).AllFoliosHaveInvoice(allFoliosHaveInvoice).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Expand(expand).Execute()

Returns a list of all reservations, filtered by the specified parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    bookingId := "bookingId_example" // string | Filter result by booking id (optional)
    propertyIds := []string{"Inner_example"} // []string | Filter result by requested properties (optional)
    ratePlanIds := []string{"Inner_example"} // []string | Filter result by requested rate plans (optional)
    companyIds := []string{"Inner_example"} // []string | Filter result by requested companies (optional)
    unitIds := []string{"Inner_example"} // []string | Filter result by assigned units (optional)
    unitGroupIds := []string{"Inner_example"} // []string | Filter result by requested unit groups (optional)
    unitGroupTypes := []string{"UnitGroupTypes_example"} // []string | Filter result by requested unit group types (optional)
    blockIds := []string{"Inner_example"} // []string | Filter result by requested blocks (optional)
    status := []string{"Status_example"} // []string | Filter result by reservation status (optional)
    dateFilter := "dateFilter_example" // string | Filter by date and time attributes of reservation. Use in combination with the 'To' and 'From' attributes.  All filters will check if the date specified by the filter type is between from (included) and to (excluded).  The exception being filtering for 'stay', which will return all reservations that are overlapping with the interval specified by from and to. (optional)
    from := time.Now() // time.Time | The start of the time interval. When filtering by date, at least one of 'from' and 'to' has to be specified<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    to := time.Now() // time.Time | The end of the time interval, must be larger than 'from'. When filtering by date, at least one of 'from' and 'to' has to be specified<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    channelCode := []string{"ChannelCode_example"} // []string | Filter result by the channel code (optional)
    sources := []string{"Inner_example"} // []string | Filter result by source (optional)
    validationMessageCategory := []string{"ValidationMessageCategory_example"} // []string | Filter result by validation message category (optional)
    externalCode := "externalCode_example" // string | Filter result by the external code. The result set will contain all reservations that have an external code starting with the  provided value (optional)
    textSearch := "textSearch_example" // string | This will filter all reservations where the provided text is contained in: booker first name or last name or email or company name,  primary guest first name or last name or email or company name, external code, reservation id, unit name. The search is case insensitive. (optional)
    balanceFilter := []string{"Inner_example"} // []string | This will filter reservations based on their balance.<br />You can provide an array of string expressions which all need to apply.<br />Each expression has the form of 'OPERATION_VALUE' where VALUE needs to be of the valid format of the property type and OPERATION can be:<br />'eq' for equals<br />'neq' for not equals<br />'lt' for less than<br />'gt' for greater than<br />'lte' for less than or equals<br />'gte' for greater than or equals<br />For instance<br />'eq_5' would mean the value should equal 5<br />'lte_7' would mean the value should be less than or equal to 7 (optional)
    allFoliosHaveInvoice := true // bool | If set to {true}, returns only reservations, in which all folios are closed and have an invoice.  If set to {false}, returns only reservations, in which some of the folios are open or don't have an invoice (optional)
    pageNumber := int32(56) // int32 | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. (optional) (default to 1)
    pageSize := int32(56) // int32 | Page size. If this is not set, the pageNumber will be ignored and all values returned. (optional) (default to 100)
    sort := []string{"Sort_example"} // []string | List of all fields that can be used to sort the results. Possible values are: arrival:asc, arrival:desc, departure:asc, departure:desc, created:asc, created:desc, updated:asc, updated:desc, id:asc, id:desc, firstname:asc, firstname:desc, lastname:asc, lastname:desc, unitname:asc, unitname:desc. All other values will be silently ignored. (optional)
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: booker, actions, timeSlices, services, assignedUnits, company. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationsGet(context.Background()).BookingId(bookingId).PropertyIds(propertyIds).RatePlanIds(ratePlanIds).CompanyIds(companyIds).UnitIds(unitIds).UnitGroupIds(unitGroupIds).UnitGroupTypes(unitGroupTypes).BlockIds(blockIds).Status(status).DateFilter(dateFilter).From(from).To(to).ChannelCode(channelCode).Sources(sources).ValidationMessageCategory(validationMessageCategory).ExternalCode(externalCode).TextSearch(textSearch).BalanceFilter(balanceFilter).AllFoliosHaveInvoice(allFoliosHaveInvoice).PageNumber(pageNumber).PageSize(pageSize).Sort(sort).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationsGet`: ReservationListModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationApi.BookingReservationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookingId** | **string** | Filter result by booking id | 
 **propertyIds** | **[]string** | Filter result by requested properties | 
 **ratePlanIds** | **[]string** | Filter result by requested rate plans | 
 **companyIds** | **[]string** | Filter result by requested companies | 
 **unitIds** | **[]string** | Filter result by assigned units | 
 **unitGroupIds** | **[]string** | Filter result by requested unit groups | 
 **unitGroupTypes** | **[]string** | Filter result by requested unit group types | 
 **blockIds** | **[]string** | Filter result by requested blocks | 
 **status** | **[]string** | Filter result by reservation status | 
 **dateFilter** | **string** | Filter by date and time attributes of reservation. Use in combination with the &#39;To&#39; and &#39;From&#39; attributes.  All filters will check if the date specified by the filter type is between from (included) and to (excluded).  The exception being filtering for &#39;stay&#39;, which will return all reservations that are overlapping with the interval specified by from and to. | 
 **from** | **time.Time** | The start of the time interval. When filtering by date, at least one of &#39;from&#39; and &#39;to&#39; has to be specified&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **time.Time** | The end of the time interval, must be larger than &#39;from&#39;. When filtering by date, at least one of &#39;from&#39; and &#39;to&#39; has to be specified&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **channelCode** | **[]string** | Filter result by the channel code | 
 **sources** | **[]string** | Filter result by source | 
 **validationMessageCategory** | **[]string** | Filter result by validation message category | 
 **externalCode** | **string** | Filter result by the external code. The result set will contain all reservations that have an external code starting with the  provided value | 
 **textSearch** | **string** | This will filter all reservations where the provided text is contained in: booker first name or last name or email or company name,  primary guest first name or last name or email or company name, external code, reservation id, unit name. The search is case insensitive. | 
 **balanceFilter** | **[]string** | This will filter reservations based on their balance.&lt;br /&gt;You can provide an array of string expressions which all need to apply.&lt;br /&gt;Each expression has the form of &#39;OPERATION_VALUE&#39; where VALUE needs to be of the valid format of the property type and OPERATION can be:&lt;br /&gt;&#39;eq&#39; for equals&lt;br /&gt;&#39;neq&#39; for not equals&lt;br /&gt;&#39;lt&#39; for less than&lt;br /&gt;&#39;gt&#39; for greater than&lt;br /&gt;&#39;lte&#39; for less than or equals&lt;br /&gt;&#39;gte&#39; for greater than or equals&lt;br /&gt;For instance&lt;br /&gt;&#39;eq_5&#39; would mean the value should equal 5&lt;br /&gt;&#39;lte_7&#39; would mean the value should be less than or equal to 7 | 
 **allFoliosHaveInvoice** | **bool** | If set to {true}, returns only reservations, in which all folios are closed and have an invoice.  If set to {false}, returns only reservations, in which some of the folios are open or don&#39;t have an invoice | 
 **pageNumber** | **int32** | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. | [default to 1]
 **pageSize** | **int32** | Page size. If this is not set, the pageNumber will be ignored and all values returned. | [default to 100]
 **sort** | **[]string** | List of all fields that can be used to sort the results. Possible values are: arrival:asc, arrival:desc, departure:asc, departure:desc, created:asc, created:desc, updated:asc, updated:desc, id:asc, id:desc, firstname:asc, firstname:desc, lastname:asc, lastname:desc, unitname:asc, unitname:desc. All other values will be silently ignored. | 
 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: booker, actions, timeSlices, services, assignedUnits, company. All other values will be silently ignored. | 

### Return type

[**ReservationListModel**](ReservationListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationscountGet

> CountModel BookingReservationscountGet(ctx).BookingId(bookingId).PropertyIds(propertyIds).RatePlanIds(ratePlanIds).CompanyIds(companyIds).UnitIds(unitIds).UnitGroupIds(unitGroupIds).UnitGroupTypes(unitGroupTypes).BlockIds(blockIds).Status(status).DateFilter(dateFilter).From(from).To(to).ChannelCode(channelCode).Sources(sources).ValidationMessageCategory(validationMessageCategory).ExternalCode(externalCode).TextSearch(textSearch).BalanceFilter(balanceFilter).AllFoliosHaveInvoice(allFoliosHaveInvoice).Execute()

Returns the number of reservations fulfilling the criteria specified in the parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    bookingId := "bookingId_example" // string | Filter result by booking id (optional)
    propertyIds := []string{"Inner_example"} // []string | Filter result by requested properties (optional)
    ratePlanIds := []string{"Inner_example"} // []string | Filter result by requested rate plans (optional)
    companyIds := []string{"Inner_example"} // []string | Filter result by requested companies (optional)
    unitIds := []string{"Inner_example"} // []string | Filter result by assigned units (optional)
    unitGroupIds := []string{"Inner_example"} // []string | Filter result by requested unit groups (optional)
    unitGroupTypes := []string{"UnitGroupTypes_example"} // []string | Filter result by requested unit group types (optional)
    blockIds := []string{"Inner_example"} // []string | Filter result by requested blocks (optional)
    status := []string{"Status_example"} // []string | Filter result by reservation status (optional)
    dateFilter := "dateFilter_example" // string | Filter by date and time attributes of reservation. Use in combination with the 'To' and 'From' attributes.  All filters will check if the date specified by the filter type is between from (included) and to (excluded).  The exception being filtering for 'stay', which will return all reservations that are overlapping with the interval specified by from and to. (optional)
    from := time.Now() // time.Time | The start of the time interval. When filtering by date, at least one of 'from' and 'to' has to be specified<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    to := time.Now() // time.Time | The end of the time interval, must be larger than 'from'. When filtering by date, at least one of 'from' and 'to' has to be specified<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    channelCode := []string{"ChannelCode_example"} // []string | Filter result by the channel code (optional)
    sources := []string{"Inner_example"} // []string | Filter result by source (optional)
    validationMessageCategory := []string{"ValidationMessageCategory_example"} // []string | Filter result by validation message category (optional)
    externalCode := "externalCode_example" // string | Filter result by the external code. The result set will contain all reservations that have an external code starting with the  provided value (optional)
    textSearch := "textSearch_example" // string | This will filter all reservations where the provided text is contained in: booker first name or last name or email or company name,  primary guest first name or last name or email or company name, external code, reservation id, unit name. The search is case insensitive. (optional)
    balanceFilter := []string{"Inner_example"} // []string | This will filter reservations based on their balance.<br />You can provide an array of string expressions which all need to apply.<br />Each expression has the form of 'OPERATION_VALUE' where VALUE needs to be of the valid format of the property type and OPERATION can be:<br />'eq' for equals<br />'neq' for not equals<br />'lt' for less than<br />'gt' for greater than<br />'lte' for less than or equals<br />'gte' for greater than or equals<br />For instance<br />'eq_5' would mean the value should equal 5<br />'lte_7' would mean the value should be less than or equal to 7 (optional)
    allFoliosHaveInvoice := true // bool | If set to {true}, returns only reservations, in which all folios are closed and have an invoice.  If set to {false}, returns only reservations, in which some of the folios are open or don't have an invoice (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationApi.BookingReservationscountGet(context.Background()).BookingId(bookingId).PropertyIds(propertyIds).RatePlanIds(ratePlanIds).CompanyIds(companyIds).UnitIds(unitIds).UnitGroupIds(unitGroupIds).UnitGroupTypes(unitGroupTypes).BlockIds(blockIds).Status(status).DateFilter(dateFilter).From(from).To(to).ChannelCode(channelCode).Sources(sources).ValidationMessageCategory(validationMessageCategory).ExternalCode(externalCode).TextSearch(textSearch).BalanceFilter(balanceFilter).AllFoliosHaveInvoice(allFoliosHaveInvoice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationApi.BookingReservationscountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationscountGet`: CountModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationApi.BookingReservationscountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationscountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookingId** | **string** | Filter result by booking id | 
 **propertyIds** | **[]string** | Filter result by requested properties | 
 **ratePlanIds** | **[]string** | Filter result by requested rate plans | 
 **companyIds** | **[]string** | Filter result by requested companies | 
 **unitIds** | **[]string** | Filter result by assigned units | 
 **unitGroupIds** | **[]string** | Filter result by requested unit groups | 
 **unitGroupTypes** | **[]string** | Filter result by requested unit group types | 
 **blockIds** | **[]string** | Filter result by requested blocks | 
 **status** | **[]string** | Filter result by reservation status | 
 **dateFilter** | **string** | Filter by date and time attributes of reservation. Use in combination with the &#39;To&#39; and &#39;From&#39; attributes.  All filters will check if the date specified by the filter type is between from (included) and to (excluded).  The exception being filtering for &#39;stay&#39;, which will return all reservations that are overlapping with the interval specified by from and to. | 
 **from** | **time.Time** | The start of the time interval. When filtering by date, at least one of &#39;from&#39; and &#39;to&#39; has to be specified&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **time.Time** | The end of the time interval, must be larger than &#39;from&#39;. When filtering by date, at least one of &#39;from&#39; and &#39;to&#39; has to be specified&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **channelCode** | **[]string** | Filter result by the channel code | 
 **sources** | **[]string** | Filter result by source | 
 **validationMessageCategory** | **[]string** | Filter result by validation message category | 
 **externalCode** | **string** | Filter result by the external code. The result set will contain all reservations that have an external code starting with the  provided value | 
 **textSearch** | **string** | This will filter all reservations where the provided text is contained in: booker first name or last name or email or company name,  primary guest first name or last name or email or company name, external code, reservation id, unit name. The search is case insensitive. | 
 **balanceFilter** | **[]string** | This will filter reservations based on their balance.&lt;br /&gt;You can provide an array of string expressions which all need to apply.&lt;br /&gt;Each expression has the form of &#39;OPERATION_VALUE&#39; where VALUE needs to be of the valid format of the property type and OPERATION can be:&lt;br /&gt;&#39;eq&#39; for equals&lt;br /&gt;&#39;neq&#39; for not equals&lt;br /&gt;&#39;lt&#39; for less than&lt;br /&gt;&#39;gt&#39; for greater than&lt;br /&gt;&#39;lte&#39; for less than or equals&lt;br /&gt;&#39;gte&#39; for greater than or equals&lt;br /&gt;For instance&lt;br /&gt;&#39;eq_5&#39; would mean the value should equal 5&lt;br /&gt;&#39;lte_7&#39; would mean the value should be less than or equal to 7 | 
 **allFoliosHaveInvoice** | **bool** | If set to {true}, returns only reservations, in which all folios are closed and have an invoice.  If set to {false}, returns only reservations, in which some of the folios are open or don&#39;t have an invoice | 

### Return type

[**CountModel**](CountModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

