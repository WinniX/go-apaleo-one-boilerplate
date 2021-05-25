# \BookingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingBookingsByIdGet**](BookingApi.md#BookingBookingsByIdGet) | **Get** /booking/v1/bookings/{id} | Returns a specific booking.
[**BookingBookingsByIdPatch**](BookingApi.md#BookingBookingsByIdPatch) | **Patch** /booking/v1/bookings/{id} | Allows to modify certain booking properties
[**BookingBookingsByIdReservationsPost**](BookingApi.md#BookingBookingsByIdReservationsPost) | **Post** /booking/v1/bookings/{id}/reservations | Add one or multiple reservations to an existing booking.
[**BookingBookingsByIdReservationsforcePost**](BookingApi.md#BookingBookingsByIdReservationsforcePost) | **Post** /booking/v1/bookings/{id}/reservations/$force | Add one or multiple reservations to an existing booking regardless of availability or restrictions.
[**BookingBookingsGet**](BookingApi.md#BookingBookingsGet) | **Get** /booking/v1/bookings | Returns a list of all bookings, filtered by the specified parameters.
[**BookingBookingsPost**](BookingApi.md#BookingBookingsPost) | **Post** /booking/v1/bookings | Creates a booking for one or more reservations.
[**BookingBookingsforcePost**](BookingApi.md#BookingBookingsforcePost) | **Post** /booking/v1/bookings/$force | Creates a booking for one or more reservations regardless of availability or restrictions.



## BookingBookingsByIdGet

> BookingModel BookingBookingsByIdGet(ctx, id).Expand(expand).Execute()

Returns a specific booking.



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
    id := "id_example" // string | Id of the booking to be retrieved.
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: property, unitGroup, ratePlan, services, reservations, propertyValues. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingApi.BookingBookingsByIdGet(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingApi.BookingBookingsByIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBookingsByIdGet`: BookingModel
    fmt.Fprintf(os.Stdout, "Response from `BookingApi.BookingBookingsByIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the booking to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBookingsByIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: property, unitGroup, ratePlan, services, reservations, propertyValues. All other values will be silently ignored. | 

### Return type

[**BookingModel**](BookingModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBookingsByIdPatch

> BookingBookingsByIdPatch(ctx, id).Body(body).Execute()

Allows to modify certain booking properties



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
    id := "id_example" // string | Id of the booking to be modified.
    body := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation | Define the list of operations to be applied to the resource. Learn more about JSON Patch here: http://jsonpatch.com/.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingApi.BookingBookingsByIdPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingApi.BookingBookingsByIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the booking to be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBookingsByIdPatchRequest struct via the builder pattern


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


## BookingBookingsByIdReservationsPost

> ReservationsCreatedModel BookingBookingsByIdReservationsPost(ctx, id).Body(body).IdempotencyKey(idempotencyKey).Execute()

Add one or multiple reservations to an existing booking.



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
    id := "id_example" // string | Id of the booking the reservations should be attached to.
    body := *openapiclient.NewAddReservationsModel([]openapiclient.CreateReservationModel{*openapiclient.NewCreateReservationModel("Arrival_example", "Departure_example", int32(123), "ChannelCode_example", []openapiclient.CreateReservationTimeSliceModel{*openapiclient.NewCreateReservationTimeSliceModel("RatePlanId_example")})}) // AddReservationsModel | The list of reservations you want to add.
    idempotencyKey := "idempotencyKey_example" // string | Unique key for safely retrying requests without accidentally performing the same operation twice.  We'll always send back the same response for requests made with the same key,  and keys can't be reused with different request parameters. Keys expire after 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingApi.BookingBookingsByIdReservationsPost(context.Background(), id).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingApi.BookingBookingsByIdReservationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBookingsByIdReservationsPost`: ReservationsCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `BookingApi.BookingBookingsByIdReservationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the booking the reservations should be attached to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBookingsByIdReservationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AddReservationsModel**](AddReservationsModel.md) | The list of reservations you want to add. | 
 **idempotencyKey** | **string** | Unique key for safely retrying requests without accidentally performing the same operation twice.  We&#39;ll always send back the same response for requests made with the same key,  and keys can&#39;t be reused with different request parameters. Keys expire after 24 hours. | 

### Return type

[**ReservationsCreatedModel**](ReservationsCreatedModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBookingsByIdReservationsforcePost

> ReservationsCreatedModel BookingBookingsByIdReservationsforcePost(ctx, id).Body(body).IdempotencyKey(idempotencyKey).Execute()

Add one or multiple reservations to an existing booking regardless of availability or restrictions.



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
    id := "id_example" // string | Id of the booking the reservations should be attached to.
    body := *openapiclient.NewAddReservationsModel([]openapiclient.CreateReservationModel{*openapiclient.NewCreateReservationModel("Arrival_example", "Departure_example", int32(123), "ChannelCode_example", []openapiclient.CreateReservationTimeSliceModel{*openapiclient.NewCreateReservationTimeSliceModel("RatePlanId_example")})}) // AddReservationsModel | The list of reservations you want to add.
    idempotencyKey := "idempotencyKey_example" // string | Unique key for safely retrying requests without accidentally performing the same operation twice.  We'll always send back the same response for requests made with the same key,  and keys can't be reused with different request parameters. Keys expire after 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingApi.BookingBookingsByIdReservationsforcePost(context.Background(), id).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingApi.BookingBookingsByIdReservationsforcePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBookingsByIdReservationsforcePost`: ReservationsCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `BookingApi.BookingBookingsByIdReservationsforcePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the booking the reservations should be attached to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBookingsByIdReservationsforcePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AddReservationsModel**](AddReservationsModel.md) | The list of reservations you want to add. | 
 **idempotencyKey** | **string** | Unique key for safely retrying requests without accidentally performing the same operation twice.  We&#39;ll always send back the same response for requests made with the same key,  and keys can&#39;t be reused with different request parameters. Keys expire after 24 hours. | 

### Return type

[**ReservationsCreatedModel**](ReservationsCreatedModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBookingsGet

> BookingListModel BookingBookingsGet(ctx).ReservationId(reservationId).GroupId(groupId).ChannelCode(channelCode).ExternalCode(externalCode).TextSearch(textSearch).PageNumber(pageNumber).PageSize(pageSize).Expand(expand).Execute()

Returns a list of all bookings, filtered by the specified parameters.



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
    reservationId := "reservationId_example" // string | Filter result by reservation id. The result set will contain all bookings having reservations with the specified id (optional)
    groupId := "groupId_example" // string | Filter result by group id. The result set will contain all bookings having groups with the specified id (optional)
    channelCode := []string{"ChannelCode_example"} // []string | Filter result by the channel code. The resul set will contain all bookings having reservations with the specified channel code (optional)
    externalCode := "externalCode_example" // string | Filter result by the external code. The result set will contain all bookings having reservations with external code starting with provided value (optional)
    textSearch := "textSearch_example" // string | This will filter all bookings for the provided free text. Currently it only looks up if either the lastname, firstname, email or company name of the booker  contains one of the provided values (optional)
    pageNumber := int32(56) // int32 | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. (optional) (default to 1)
    pageSize := int32(56) // int32 | Page size. If this is not set, the pageNumber will be ignored and all values returned. (optional) (default to 100)
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: property, unitGroup, ratePlan, services, reservations. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingApi.BookingBookingsGet(context.Background()).ReservationId(reservationId).GroupId(groupId).ChannelCode(channelCode).ExternalCode(externalCode).TextSearch(textSearch).PageNumber(pageNumber).PageSize(pageSize).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingApi.BookingBookingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBookingsGet`: BookingListModel
    fmt.Fprintf(os.Stdout, "Response from `BookingApi.BookingBookingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingBookingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reservationId** | **string** | Filter result by reservation id. The result set will contain all bookings having reservations with the specified id | 
 **groupId** | **string** | Filter result by group id. The result set will contain all bookings having groups with the specified id | 
 **channelCode** | **[]string** | Filter result by the channel code. The resul set will contain all bookings having reservations with the specified channel code | 
 **externalCode** | **string** | Filter result by the external code. The result set will contain all bookings having reservations with external code starting with provided value | 
 **textSearch** | **string** | This will filter all bookings for the provided free text. Currently it only looks up if either the lastname, firstname, email or company name of the booker  contains one of the provided values | 
 **pageNumber** | **int32** | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. | [default to 1]
 **pageSize** | **int32** | Page size. If this is not set, the pageNumber will be ignored and all values returned. | [default to 100]
 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: property, unitGroup, ratePlan, services, reservations. All other values will be silently ignored. | 

### Return type

[**BookingListModel**](BookingListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBookingsPost

> BookingCreatedModel BookingBookingsPost(ctx).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates a booking for one or more reservations.



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
    body := *openapiclient.NewCreateBookingModel(*openapiclient.NewBookerModel("LastName_example"), []openapiclient.CreateReservationModel{*openapiclient.NewCreateReservationModel("Arrival_example", "Departure_example", int32(123), "ChannelCode_example", []openapiclient.CreateReservationTimeSliceModel{*openapiclient.NewCreateReservationTimeSliceModel("RatePlanId_example")})}) // CreateBookingModel | The list of reservations you want to create.
    idempotencyKey := "idempotencyKey_example" // string | Unique key for safely retrying requests without accidentally performing the same operation twice.  We'll always send back the same response for requests made with the same key,  and keys can't be reused with different request parameters. Keys expire after 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingApi.BookingBookingsPost(context.Background()).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingApi.BookingBookingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBookingsPost`: BookingCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `BookingApi.BookingBookingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingBookingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBookingModel**](CreateBookingModel.md) | The list of reservations you want to create. | 
 **idempotencyKey** | **string** | Unique key for safely retrying requests without accidentally performing the same operation twice.  We&#39;ll always send back the same response for requests made with the same key,  and keys can&#39;t be reused with different request parameters. Keys expire after 24 hours. | 

### Return type

[**BookingCreatedModel**](BookingCreatedModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBookingsforcePost

> BookingCreatedModel BookingBookingsforcePost(ctx).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates a booking for one or more reservations regardless of availability or restrictions.



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
    body := *openapiclient.NewCreateBookingModel(*openapiclient.NewBookerModel("LastName_example"), []openapiclient.CreateReservationModel{*openapiclient.NewCreateReservationModel("Arrival_example", "Departure_example", int32(123), "ChannelCode_example", []openapiclient.CreateReservationTimeSliceModel{*openapiclient.NewCreateReservationTimeSliceModel("RatePlanId_example")})}) // CreateBookingModel | The list of reservations you want to create.
    idempotencyKey := "idempotencyKey_example" // string | Unique key for safely retrying requests without accidentally performing the same operation twice.  We'll always send back the same response for requests made with the same key,  and keys can't be reused with different request parameters. Keys expire after 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingApi.BookingBookingsforcePost(context.Background()).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingApi.BookingBookingsforcePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBookingsforcePost`: BookingCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `BookingApi.BookingBookingsforcePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingBookingsforcePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBookingModel**](CreateBookingModel.md) | The list of reservations you want to create. | 
 **idempotencyKey** | **string** | Unique key for safely retrying requests without accidentally performing the same operation twice.  We&#39;ll always send back the same response for requests made with the same key,  and keys can&#39;t be reused with different request parameters. Keys expire after 24 hours. | 

### Return type

[**BookingCreatedModel**](BookingCreatedModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

