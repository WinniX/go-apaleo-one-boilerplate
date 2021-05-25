# \GroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingGroupsByIdDelete**](GroupApi.md#BookingGroupsByIdDelete) | **Delete** /booking/v1/groups/{id} | Delete a certain group booking
[**BookingGroupsByIdGet**](GroupApi.md#BookingGroupsByIdGet) | **Get** /booking/v1/groups/{id} | Returns a specific group booking.
[**BookingGroupsByIdHead**](GroupApi.md#BookingGroupsByIdHead) | **Head** /booking/v1/groups/{id} | Check if a certain group booking exists
[**BookingGroupsByIdPatch**](GroupApi.md#BookingGroupsByIdPatch) | **Patch** /booking/v1/groups/{id} | Allows to modify certain group booking properties
[**BookingGroupsByIdReservationsPost**](GroupApi.md#BookingGroupsByIdReservationsPost) | **Post** /booking/v1/groups/{id}/reservations | Add one or multiple reservations to an existing group booking using blocked inventory.
[**BookingGroupsGet**](GroupApi.md#BookingGroupsGet) | **Get** /booking/v1/groups | Returns a list of all group bookings, filtered by the specified parameters.
[**BookingGroupsPost**](GroupApi.md#BookingGroupsPost) | **Post** /booking/v1/groups | Creates a group booking.
[**BookingGroupscountGet**](GroupApi.md#BookingGroupscountGet) | **Get** /booking/v1/groups/$count | Returns number of group bookings



## BookingGroupsByIdDelete

> BookingGroupsByIdDelete(ctx, id).Execute()

Delete a certain group booking



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
    id := "id_example" // string | Id of the group booking to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupsByIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupsByIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the group booking to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupsByIdDeleteRequest struct via the builder pattern


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


## BookingGroupsByIdGet

> GroupModel BookingGroupsByIdGet(ctx, id).Expand(expand).Execute()

Returns a specific group booking.



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
    id := "id_example" // string | Id of the group booking to be retrieved.
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: blocks, actions. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupsByIdGet(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupsByIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingGroupsByIdGet`: GroupModel
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.BookingGroupsByIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the group booking to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupsByIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: blocks, actions. All other values will be silently ignored. | 

### Return type

[**GroupModel**](GroupModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingGroupsByIdHead

> BookingGroupsByIdHead(ctx, id).Execute()

Check if a certain group booking exists



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
    id := "id_example" // string | Id of the group booking to be checked for existence.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupsByIdHead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupsByIdHead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the group booking to be checked for existence. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupsByIdHeadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingGroupsByIdPatch

> BookingGroupsByIdPatch(ctx, id).Body(body).Execute()

Allows to modify certain group booking properties



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
    id := "id_example" // string | Id of the group booking to be modified.
    body := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation | Define the list of operations to be applied to the resource. Learn more about JSON Patch here: http://jsonpatch.com/.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupsByIdPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupsByIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the group booking to be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupsByIdPatchRequest struct via the builder pattern


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


## BookingGroupsByIdReservationsPost

> ReservationsCreatedModel BookingGroupsByIdReservationsPost(ctx, id).Body(body).IdempotencyKey(idempotencyKey).Execute()

Add one or multiple reservations to an existing group booking using blocked inventory.



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
    id := "id_example" // string | Id of the group booking the reservations should be attached to.
    body := *openapiclient.NewPickUpReservationsModel([]openapiclient.PickUpReservationModel{*openapiclient.NewPickUpReservationModel("BlockId_example", "Arrival_example", "Departure_example", int32(123))}) // PickUpReservationsModel | The list of reservations you want to create.
    idempotencyKey := "idempotencyKey_example" // string | Unique key for safely retrying requests without accidentally performing the same operation twice.  We'll always send back the same response for requests made with the same key,  and keys can't be reused with different request parameters. Keys expire after 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupsByIdReservationsPost(context.Background(), id).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupsByIdReservationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingGroupsByIdReservationsPost`: ReservationsCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.BookingGroupsByIdReservationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the group booking the reservations should be attached to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupsByIdReservationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PickUpReservationsModel**](PickUpReservationsModel.md) | The list of reservations you want to create. | 
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


## BookingGroupsGet

> GroupListModel BookingGroupsGet(ctx).TextSearch(textSearch).PropertyIds(propertyIds).From(from).To(to).PageNumber(pageNumber).PageSize(pageSize).Expand(expand).Execute()

Returns a list of all group bookings, filtered by the specified parameters.



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
    textSearch := "textSearch_example" // string | This will filter all group bookings for the provided free text. Currently it only looks up if either the group name, lastname,  firstname, email or company name of the booker contains one of the provided values (optional)
    propertyIds := []string{"Inner_example"} // []string | Filter result by requested properties (optional)
    from := time.Now() // time.Time | The start of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    to := time.Now() // time.Time | The end of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    pageNumber := int32(56) // int32 | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. (optional) (default to 1)
    pageSize := int32(56) // int32 | Page size. If this is not set, the pageNumber will be ignored and all values returned. (optional) (default to 100)
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: blocks, actions. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupsGet(context.Background()).TextSearch(textSearch).PropertyIds(propertyIds).From(from).To(to).PageNumber(pageNumber).PageSize(pageSize).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingGroupsGet`: GroupListModel
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.BookingGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textSearch** | **string** | This will filter all group bookings for the provided free text. Currently it only looks up if either the group name, lastname,  firstname, email or company name of the booker contains one of the provided values | 
 **propertyIds** | **[]string** | Filter result by requested properties | 
 **from** | **time.Time** | The start of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **time.Time** | The end of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **pageNumber** | **int32** | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. | [default to 1]
 **pageSize** | **int32** | Page size. If this is not set, the pageNumber will be ignored and all values returned. | [default to 100]
 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: blocks, actions. All other values will be silently ignored. | 

### Return type

[**GroupListModel**](GroupListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingGroupsPost

> GroupCreatedModel BookingGroupsPost(ctx).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates a group booking.



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
    body := *openapiclient.NewCreateGroupModel("Name_example", *openapiclient.NewBookerModel("LastName_example"), []string{"PropertyIds_example"}) // CreateGroupModel | The details of the group that should be created.
    idempotencyKey := "idempotencyKey_example" // string | Unique key for safely retrying requests without accidentally performing the same operation twice.  We'll always send back the same response for requests made with the same key,  and keys can't be reused with different request parameters. Keys expire after 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupsPost(context.Background()).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingGroupsPost`: GroupCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.BookingGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateGroupModel**](CreateGroupModel.md) | The details of the group that should be created. | 
 **idempotencyKey** | **string** | Unique key for safely retrying requests without accidentally performing the same operation twice.  We&#39;ll always send back the same response for requests made with the same key,  and keys can&#39;t be reused with different request parameters. Keys expire after 24 hours. | 

### Return type

[**GroupCreatedModel**](GroupCreatedModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingGroupscountGet

> CountModel BookingGroupscountGet(ctx).TextSearch(textSearch).PropertyIds(propertyIds).From(from).To(to).Execute()

Returns number of group bookings



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
    textSearch := "textSearch_example" // string | This will filter all group bookings for the provided free text. Currently it only looks up if either the group name, lastname,  firstname, email or company name of the booker contains one of the provided values (optional)
    propertyIds := []string{"Inner_example"} // []string | Filter result by requested properties (optional)
    from := time.Now() // time.Time | The start of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    to := time.Now() // time.Time | The end of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.BookingGroupscountGet(context.Background()).TextSearch(textSearch).PropertyIds(propertyIds).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.BookingGroupscountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingGroupscountGet`: CountModel
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.BookingGroupscountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingGroupscountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textSearch** | **string** | This will filter all group bookings for the provided free text. Currently it only looks up if either the group name, lastname,  firstname, email or company name of the booker contains one of the provided values | 
 **propertyIds** | **[]string** | Filter result by requested properties | 
 **from** | **time.Time** | The start of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **time.Time** | The end of the time range. All groups that have blocks overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 

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

