# \TypesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingTypesSourcesGet**](TypesApi.md#BookingTypesSourcesGet) | **Get** /booking/v1/types/sources | Returns a list of supported sources.



## BookingTypesSourcesGet

> SourceListModel BookingTypesSourcesGet(ctx).Execute()

Returns a list of supported sources.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesApi.BookingTypesSourcesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.BookingTypesSourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingTypesSourcesGet`: SourceListModel
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.BookingTypesSourcesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBookingTypesSourcesGetRequest struct via the builder pattern


### Return type

[**SourceListModel**](SourceListModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

