# \ReservationActionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingReservationActionsByIdAddCityTaxPut**](ReservationActionsApi.md#BookingReservationActionsByIdAddCityTaxPut) | **Put** /booking/v1/reservation-actions/{id}/add-city-tax | Adds the city tax to a reservation.
[**BookingReservationActionsByIdAmendPut**](ReservationActionsApi.md#BookingReservationActionsByIdAmendPut) | **Put** /booking/v1/reservation-actions/{id}/amend | Allows you to amend the stay details of a reservation
[**BookingReservationActionsByIdAmendforcePut**](ReservationActionsApi.md#BookingReservationActionsByIdAmendforcePut) | **Put** /booking/v1/reservation-actions/{id}/amend/$force | Allows you to amend the stay details of a reservation regardless of availability or restrictions.
[**BookingReservationActionsByIdAssignUnitByUnitIdPut**](ReservationActionsApi.md#BookingReservationActionsByIdAssignUnitByUnitIdPut) | **Put** /booking/v1/reservation-actions/{id}/assign-unit/{unitId} | Assign a specific unit to a reservation.
[**BookingReservationActionsByIdAssignUnitPut**](ReservationActionsApi.md#BookingReservationActionsByIdAssignUnitPut) | **Put** /booking/v1/reservation-actions/{id}/assign-unit | Assign a unit to a reservation.
[**BookingReservationActionsByIdBookServicePut**](ReservationActionsApi.md#BookingReservationActionsByIdBookServicePut) | **Put** /booking/v1/reservation-actions/{id}/book-service | Book the service for a specific reservation.
[**BookingReservationActionsByIdCancelPut**](ReservationActionsApi.md#BookingReservationActionsByIdCancelPut) | **Put** /booking/v1/reservation-actions/{id}/cancel | Cancel a reservation.
[**BookingReservationActionsByIdCheckinPut**](ReservationActionsApi.md#BookingReservationActionsByIdCheckinPut) | **Put** /booking/v1/reservation-actions/{id}/checkin | Check-in of a reservation.
[**BookingReservationActionsByIdCheckoutPut**](ReservationActionsApi.md#BookingReservationActionsByIdCheckoutPut) | **Put** /booking/v1/reservation-actions/{id}/checkout | Check-out of a reservation.
[**BookingReservationActionsByIdNoshowPut**](ReservationActionsApi.md#BookingReservationActionsByIdNoshowPut) | **Put** /booking/v1/reservation-actions/{id}/noshow | Set a reservation to No-show.
[**BookingReservationActionsByIdRemoveCityTaxPut**](ReservationActionsApi.md#BookingReservationActionsByIdRemoveCityTaxPut) | **Put** /booking/v1/reservation-actions/{id}/remove-city-tax | Removes the city tax from a reservation.
[**BookingReservationActionsByIdUnassignUnitsPut**](ReservationActionsApi.md#BookingReservationActionsByIdUnassignUnitsPut) | **Put** /booking/v1/reservation-actions/{id}/unassign-units | Unassign units from a reservation.



## BookingReservationActionsByIdAddCityTaxPut

> BookingReservationActionsByIdAddCityTaxPut(ctx, id).Execute()

Adds the city tax to a reservation.



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
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdAddCityTaxPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdAddCityTaxPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdAddCityTaxPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BookingReservationActionsByIdAmendPut

> BookingReservationActionsByIdAmendPut(ctx, id).Body(body).Execute()

Allows you to amend the stay details of a reservation



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
    id := "id_example" // string | Id of the reservation that should be modified
    body := *openapiclient.NewDesiredStayDetailsModel("Arrival_example", "Departure_example", int32(123), []openapiclient.DesiredTimeSliceModel{*openapiclient.NewDesiredTimeSliceModel("RatePlanId_example")}) // DesiredStayDetailsModel | The new stay details that should be applied to the reservation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdAmendPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdAmendPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation that should be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdAmendPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DesiredStayDetailsModel**](DesiredStayDetailsModel.md) | The new stay details that should be applied to the reservation. | 

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


## BookingReservationActionsByIdAmendforcePut

> BookingReservationActionsByIdAmendforcePut(ctx, id).Body(body).Execute()

Allows you to amend the stay details of a reservation regardless of availability or restrictions.



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
    id := "id_example" // string | Id of the reservation that should be modified
    body := *openapiclient.NewDesiredStayDetailsModel("Arrival_example", "Departure_example", int32(123), []openapiclient.DesiredTimeSliceModel{*openapiclient.NewDesiredTimeSliceModel("RatePlanId_example")}) // DesiredStayDetailsModel | The new stay details that should be applied to the reservation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdAmendforcePut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdAmendforcePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation that should be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdAmendforcePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DesiredStayDetailsModel**](DesiredStayDetailsModel.md) | The new stay details that should be applied to the reservation. | 

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


## BookingReservationActionsByIdAssignUnitByUnitIdPut

> AssignedUnitModel BookingReservationActionsByIdAssignUnitByUnitIdPut(ctx, id, unitId).From(from).To(to).Execute()

Assign a specific unit to a reservation.



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
    id := "id_example" // string | Id of the reservation the unit should be assigned to.
    unitId := "unitId_example" // string | The id of the unit to be assigned.
    from := "from_example" // string | The start date and optional time for the unit assignment. If not specified, the reservation's arrival will be used.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    to := "to_example" // string | The end date and optional time for the unit assignment. If not specified, the reservation's departure will be used.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdAssignUnitByUnitIdPut(context.Background(), id, unitId).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdAssignUnitByUnitIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationActionsByIdAssignUnitByUnitIdPut`: AssignedUnitModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationActionsApi.BookingReservationActionsByIdAssignUnitByUnitIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation the unit should be assigned to. | 
**unitId** | **string** | The id of the unit to be assigned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdAssignUnitByUnitIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **string** | The start date and optional time for the unit assignment. If not specified, the reservation&#39;s arrival will be used.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **string** | The end date and optional time for the unit assignment. If not specified, the reservation&#39;s departure will be used.&lt;br /&gt;Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 

### Return type

[**AssignedUnitModel**](AssignedUnitModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationActionsByIdAssignUnitPut

> AutoAssignedUnitListModel BookingReservationActionsByIdAssignUnitPut(ctx, id).Execute()

Assign a unit to a reservation.



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
    id := "id_example" // string | Id of the reservation a unit should be assigned to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdAssignUnitPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdAssignUnitPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingReservationActionsByIdAssignUnitPut`: AutoAssignedUnitListModel
    fmt.Fprintf(os.Stdout, "Response from `ReservationActionsApi.BookingReservationActionsByIdAssignUnitPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation a unit should be assigned to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdAssignUnitPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoAssignedUnitListModel**](AutoAssignedUnitListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingReservationActionsByIdBookServicePut

> BookingReservationActionsByIdBookServicePut(ctx, id).Body(body).Execute()

Book the service for a specific reservation.



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
    body := *openapiclient.NewBookReservationServiceModel("ServiceId_example") // BookReservationServiceModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdBookServicePut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdBookServicePut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdBookServicePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BookReservationServiceModel**](BookReservationServiceModel.md) |  | 

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


## BookingReservationActionsByIdCancelPut

> BookingReservationActionsByIdCancelPut(ctx, id).Execute()

Cancel a reservation.



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
    id := "id_example" // string | Id of the reservation that should be processed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdCancelPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdCancelPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdCancelPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BookingReservationActionsByIdCheckinPut

> BookingReservationActionsByIdCheckinPut(ctx, id).WithCityTax(withCityTax).Execute()

Check-in of a reservation.



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
    id := "id_example" // string | Id of the reservation that should be processed.
    withCityTax := true // bool | Define if city tax should be added for this reservation or not. The default is \"true\". (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdCheckinPut(context.Background(), id).WithCityTax(withCityTax).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdCheckinPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdCheckinPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withCityTax** | **bool** | Define if city tax should be added for this reservation or not. The default is \&quot;true\&quot;. | 

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


## BookingReservationActionsByIdCheckoutPut

> BookingReservationActionsByIdCheckoutPut(ctx, id).Execute()

Check-out of a reservation.



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
    id := "id_example" // string | Id of the reservation that should be processed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdCheckoutPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdCheckoutPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdCheckoutPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BookingReservationActionsByIdNoshowPut

> BookingReservationActionsByIdNoshowPut(ctx, id).Execute()

Set a reservation to No-show.



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
    id := "id_example" // string | Id of the reservation that should be processed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdNoshowPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdNoshowPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdNoshowPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BookingReservationActionsByIdRemoveCityTaxPut

> BookingReservationActionsByIdRemoveCityTaxPut(ctx, id).Execute()

Removes the city tax from a reservation.



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
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdRemoveCityTaxPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdRemoveCityTaxPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdRemoveCityTaxPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BookingReservationActionsByIdUnassignUnitsPut

> BookingReservationActionsByIdUnassignUnitsPut(ctx, id).Execute()

Unassign units from a reservation.



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
    id := "id_example" // string | Id of the reservation the unit should be unassigned for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReservationActionsApi.BookingReservationActionsByIdUnassignUnitsPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationActionsApi.BookingReservationActionsByIdUnassignUnitsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the reservation the unit should be unassigned for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingReservationActionsByIdUnassignUnitsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

