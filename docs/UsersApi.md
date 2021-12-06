# \UsersApi

All URIs are relative to *https://cloudapi.cloud.couchbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersList**](UsersApi.md#UsersList) | **Get** /v2/users | List Users



## UsersList

> ListUsersResponse UsersList(ctx).Page(page).PerPage(perPage).SortBy(sortBy).Execute()

List Users



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
    page := int32(56) // int32 | Sets how many results you would like to have on each page (optional)
    perPage := int32(56) // int32 | Sets what page you would like to view (optional)
    sortBy := "name.desc" // string | Sets order of how you would like to sort results and also the key you would like to order by (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersList(context.Background()).Page(page).PerPage(perPage).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersList`: ListUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Sets how many results you would like to have on each page | 
 **perPage** | **int32** | Sets what page you would like to view | 
 **sortBy** | **string** | Sets order of how you would like to sort results and also the key you would like to order by | 

### Return type

[**ListUsersResponse**](ListUsersResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

