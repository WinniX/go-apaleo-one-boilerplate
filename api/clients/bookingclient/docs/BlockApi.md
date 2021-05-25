# \BlockApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingBlocksByIdDelete**](BlockApi.md#BookingBlocksByIdDelete) | **Delete** /booking/v1/blocks/{id} | Delete a specific block
[**BookingBlocksByIdGet**](BlockApi.md#BookingBlocksByIdGet) | **Get** /booking/v1/blocks/{id} | Returns a specific block.
[**BookingBlocksByIdHead**](BlockApi.md#BookingBlocksByIdHead) | **Head** /booking/v1/blocks/{id} | Check if a block exists
[**BookingBlocksByIdPatch**](BlockApi.md#BookingBlocksByIdPatch) | **Patch** /booking/v1/blocks/{id} | Allows to modify the block
[**BookingBlocksGet**](BlockApi.md#BookingBlocksGet) | **Get** /booking/v1/blocks | Returns a list of blocks
[**BookingBlocksPost**](BlockApi.md#BookingBlocksPost) | **Post** /booking/v1/blocks | Creates a block
[**BookingBlockscountGet**](BlockApi.md#BookingBlockscountGet) | **Get** /booking/v1/blocks/$count | Returns number of blocks



## BookingBlocksByIdDelete

> BookingBlocksByIdDelete(ctx, id).Execute()

Delete a specific block



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
    id := "id_example" // string | The id of the block.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BookingBlocksByIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BookingBlocksByIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the block. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlocksByIdDeleteRequest struct via the builder pattern


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


## BookingBlocksByIdGet

> BlockModel BookingBlocksByIdGet(ctx, id).Expand(expand).Execute()

Returns a specific block.



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
    id := "id_example" // string | Id of the block to be retrieved.
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: actions, timeSlices. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BookingBlocksByIdGet(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BookingBlocksByIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBlocksByIdGet`: BlockModel
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.BookingBlocksByIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the block to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlocksByIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: actions, timeSlices. All other values will be silently ignored. | 

### Return type

[**BlockModel**](BlockModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBlocksByIdHead

> BookingBlocksByIdHead(ctx, id).Execute()

Check if a block exists



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
    id := "id_example" // string | The id of the block.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BookingBlocksByIdHead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BookingBlocksByIdHead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the block. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlocksByIdHeadRequest struct via the builder pattern


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


## BookingBlocksByIdPatch

> BookingBlocksByIdPatch(ctx, id).Body(body).Execute()

Allows to modify the block



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
    id := "id_example" // string | Id of the block to be modified.
    body := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation | Define the list of operations to be applied to the resource. Learn more about JSON Patch here: http://jsonpatch.com/.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BookingBlocksByIdPatch(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BookingBlocksByIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the block to be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlocksByIdPatchRequest struct via the builder pattern


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


## BookingBlocksGet

> BlockListModel BookingBlocksGet(ctx).GroupId(groupId).PropertyIds(propertyIds).Status(status).UnitGroupIds(unitGroupIds).RatePlanIds(ratePlanIds).TimeSliceDefinitionIds(timeSliceDefinitionIds).UnitGroupTypes(unitGroupTypes).TimeSliceTemplate(timeSliceTemplate).From(from).To(to).PageNumber(pageNumber).PageSize(pageSize).Expand(expand).Execute()

Returns a list of blocks



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
    groupId := "groupId_example" // string | Return blocks for the specific group (optional)
    propertyIds := []string{"Inner_example"} // []string | Return blocks filtered by properties (optional)
    status := []string{"Status_example"} // []string | Return blocks filtered by statuses (optional)
    unitGroupIds := []string{"Inner_example"} // []string | Return blocks with any of the specified unit groups (optional)
    ratePlanIds := []string{"Inner_example"} // []string | Return blocks with any of the specified rate plans (optional)
    timeSliceDefinitionIds := []string{"Inner_example"} // []string | Return blocks with any of the specified time slice definitions (optional)
    unitGroupTypes := []string{"UnitGroupTypes_example"} // []string | Return blocks with any of the specified unit group types (optional)
    timeSliceTemplate := "timeSliceTemplate_example" // string | The time slice template, defaults to 'over night' (optional)
    from := time.Now() // time.Time | The start of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    to := time.Now() // time.Time | The end of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    pageNumber := int32(56) // int32 | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. (optional) (default to 1)
    pageSize := int32(56) // int32 | Page size. If this is not set, the pageNumber will be ignored and all values returned. (optional) (default to 100)
    expand := []string{"Expand_example"} // []string | List of all embedded resources that should be expanded in the response. Possible values are: actions, timeSlices. All other values will be silently ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BookingBlocksGet(context.Background()).GroupId(groupId).PropertyIds(propertyIds).Status(status).UnitGroupIds(unitGroupIds).RatePlanIds(ratePlanIds).TimeSliceDefinitionIds(timeSliceDefinitionIds).UnitGroupTypes(unitGroupTypes).TimeSliceTemplate(timeSliceTemplate).From(from).To(to).PageNumber(pageNumber).PageSize(pageSize).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BookingBlocksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBlocksGet`: BlockListModel
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.BookingBlocksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlocksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Return blocks for the specific group | 
 **propertyIds** | **[]string** | Return blocks filtered by properties | 
 **status** | **[]string** | Return blocks filtered by statuses | 
 **unitGroupIds** | **[]string** | Return blocks with any of the specified unit groups | 
 **ratePlanIds** | **[]string** | Return blocks with any of the specified rate plans | 
 **timeSliceDefinitionIds** | **[]string** | Return blocks with any of the specified time slice definitions | 
 **unitGroupTypes** | **[]string** | Return blocks with any of the specified unit group types | 
 **timeSliceTemplate** | **string** | The time slice template, defaults to &#39;over night&#39; | 
 **from** | **time.Time** | The start of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **time.Time** | The end of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **pageNumber** | **int32** | Page number, starting from 1 and defaulting to 1. Results in 204 if there are no items on that page. | [default to 1]
 **pageSize** | **int32** | Page size. If this is not set, the pageNumber will be ignored and all values returned. | [default to 100]
 **expand** | **[]string** | List of all embedded resources that should be expanded in the response. Possible values are: actions, timeSlices. All other values will be silently ignored. | 

### Return type

[**BlockListModel**](BlockListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBlocksPost

> BlockCreatedModel BookingBlocksPost(ctx).Body(body).IdempotencyKey(idempotencyKey).Execute()

Creates a block



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
    body := *openapiclient.NewCreateBlockModel("GroupId_example", "RatePlanId_example", "From_example", "To_example", *openapiclient.NewMonetaryValueModel(float64(123), "Currency_example")) // CreateBlockModel | The details for the block you want to create.
    idempotencyKey := "idempotencyKey_example" // string | Unique key for safely retrying requests without accidentally performing the same operation twice.  We'll always send back the same response for requests made with the same key,  and keys can't be reused with different request parameters. Keys expire after 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BookingBlocksPost(context.Background()).Body(body).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BookingBlocksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBlocksPost`: BlockCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.BookingBlocksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlocksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBlockModel**](CreateBlockModel.md) | The details for the block you want to create. | 
 **idempotencyKey** | **string** | Unique key for safely retrying requests without accidentally performing the same operation twice.  We&#39;ll always send back the same response for requests made with the same key,  and keys can&#39;t be reused with different request parameters. Keys expire after 24 hours. | 

### Return type

[**BlockCreatedModel**](BlockCreatedModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookingBlockscountGet

> CountModel BookingBlockscountGet(ctx).GroupId(groupId).PropertyIds(propertyIds).Status(status).UnitGroupIds(unitGroupIds).RatePlanIds(ratePlanIds).TimeSliceDefinitionIds(timeSliceDefinitionIds).UnitGroupTypes(unitGroupTypes).TimeSliceTemplate(timeSliceTemplate).From(from).To(to).Execute()

Returns number of blocks



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
    groupId := "groupId_example" // string | Return blocks for the specific group (optional)
    propertyIds := []string{"Inner_example"} // []string | Return blocks filtered by properties (optional)
    status := []string{"Status_example"} // []string | Return blocks filtered by statuses (optional)
    unitGroupIds := []string{"Inner_example"} // []string | Return blocks with any of the specified unit groups (optional)
    ratePlanIds := []string{"Inner_example"} // []string | Return blocks with any of the specified rate plans (optional)
    timeSliceDefinitionIds := []string{"Inner_example"} // []string | Return blocks with any of the specified time slice definitions (optional)
    unitGroupTypes := []string{"UnitGroupTypes_example"} // []string | Return blocks with any of the specified unit group types (optional)
    timeSliceTemplate := "timeSliceTemplate_example" // string | The time slice template, defaults to 'over night' (optional)
    from := time.Now() // time.Time | The start of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)
    to := time.Now() // time.Time | The end of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a> (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BookingBlockscountGet(context.Background()).GroupId(groupId).PropertyIds(propertyIds).Status(status).UnitGroupIds(unitGroupIds).RatePlanIds(ratePlanIds).TimeSliceDefinitionIds(timeSliceDefinitionIds).UnitGroupTypes(unitGroupTypes).TimeSliceTemplate(timeSliceTemplate).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BookingBlockscountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingBlockscountGet`: CountModel
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.BookingBlockscountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlockscountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Return blocks for the specific group | 
 **propertyIds** | **[]string** | Return blocks filtered by properties | 
 **status** | **[]string** | Return blocks filtered by statuses | 
 **unitGroupIds** | **[]string** | Return blocks with any of the specified unit groups | 
 **ratePlanIds** | **[]string** | Return blocks with any of the specified rate plans | 
 **timeSliceDefinitionIds** | **[]string** | Return blocks with any of the specified time slice definitions | 
 **unitGroupTypes** | **[]string** | Return blocks with any of the specified unit group types | 
 **timeSliceTemplate** | **string** | The time slice template, defaults to &#39;over night&#39; | 
 **from** | **time.Time** | The start of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 
 **to** | **time.Time** | The end of the time range. All blocks that are overlapping with the interval specified by from and to  will be returned&lt;br /&gt;A date and time (without fractional second part) in UTC or with UTC offset as defined in &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO8601:2004&lt;/a&gt; | 

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

