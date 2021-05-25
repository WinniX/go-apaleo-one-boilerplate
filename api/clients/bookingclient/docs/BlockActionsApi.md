# \BlockActionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingBlockActionsByIdAmendPut**](BlockActionsApi.md#BookingBlockActionsByIdAmendPut) | **Put** /booking/v1/block-actions/{id}/amend | Allow to modify a block
[**BookingBlockActionsByIdCancelPut**](BlockActionsApi.md#BookingBlockActionsByIdCancelPut) | **Put** /booking/v1/block-actions/{id}/cancel | Cancel a block.
[**BookingBlockActionsByIdConfirmPut**](BlockActionsApi.md#BookingBlockActionsByIdConfirmPut) | **Put** /booking/v1/block-actions/{id}/confirm | Confirm a block.
[**BookingBlockActionsByIdReleasePut**](BlockActionsApi.md#BookingBlockActionsByIdReleasePut) | **Put** /booking/v1/block-actions/{id}/release | Release a block.
[**BookingBlockActionsByIdWashPut**](BlockActionsApi.md#BookingBlockActionsByIdWashPut) | **Put** /booking/v1/block-actions/{id}/wash | Wash a block.



## BookingBlockActionsByIdAmendPut

> BookingBlockActionsByIdAmendPut(ctx, id).Body(body).Execute()

Allow to modify a block



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
    body := *openapiclient.NewReplaceBlockModel("From_example", "To_example", *openapiclient.NewMonetaryValueModel(float64(123), "Currency_example"), []openapiclient.CreateBlockTimeSliceModel{*openapiclient.NewCreateBlockTimeSliceModel(int32(123))}) // ReplaceBlockModel | The definition of the block.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockActionsApi.BookingBlockActionsByIdAmendPut(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockActionsApi.BookingBlockActionsByIdAmendPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBookingBlockActionsByIdAmendPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplaceBlockModel**](ReplaceBlockModel.md) | The definition of the block. | 

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


## BookingBlockActionsByIdCancelPut

> BookingBlockActionsByIdCancelPut(ctx, id).Execute()

Cancel a block.



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
    id := "id_example" // string | Id of the block that should be processed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockActionsApi.BookingBlockActionsByIdCancelPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockActionsApi.BookingBlockActionsByIdCancelPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the block that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlockActionsByIdCancelPutRequest struct via the builder pattern


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


## BookingBlockActionsByIdConfirmPut

> BookingBlockActionsByIdConfirmPut(ctx, id).Execute()

Confirm a block.



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
    id := "id_example" // string | Id of the block that should be processed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockActionsApi.BookingBlockActionsByIdConfirmPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockActionsApi.BookingBlockActionsByIdConfirmPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the block that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlockActionsByIdConfirmPutRequest struct via the builder pattern


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


## BookingBlockActionsByIdReleasePut

> BookingBlockActionsByIdReleasePut(ctx, id).Execute()

Release a block.



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
    id := "id_example" // string | Id of the block that should be processed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockActionsApi.BookingBlockActionsByIdReleasePut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockActionsApi.BookingBlockActionsByIdReleasePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the block that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlockActionsByIdReleasePutRequest struct via the builder pattern


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


## BookingBlockActionsByIdWashPut

> BookingBlockActionsByIdWashPut(ctx, id).Execute()

Wash a block.



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
    id := "id_example" // string | Id of the block that should be processed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockActionsApi.BookingBlockActionsByIdWashPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockActionsApi.BookingBlockActionsByIdWashPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the block that should be processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingBlockActionsByIdWashPutRequest struct via the builder pattern


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

