# \ProjectsApi

All URIs are relative to *https://cloudapi.cloud.couchbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsCreate**](ProjectsApi.md#ProjectsCreate) | **Post** /v2/projects | Create project
[**ProjectsDelete**](ProjectsApi.md#ProjectsDelete) | **Delete** /v2/projects/{id} | Delete project
[**ProjectsList**](ProjectsApi.md#ProjectsList) | **Get** /v2/projects | List Projects
[**ProjectsShow**](ProjectsApi.md#ProjectsShow) | **Get** /v2/projects/{id} | Get project by ID



## ProjectsCreate

> Project ProjectsCreate(ctx).CreateProjectRequest(createProjectRequest).Execute()

Create project



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
    createProjectRequest := *openapiclient.NewCreateProjectRequest("Name_example") // CreateProjectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.ProjectsCreate(context.Background()).CreateProjectRequest(createProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreate`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectRequest** | [**CreateProjectRequest**](CreateProjectRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDelete

> ProjectsDelete(ctx, id).Execute()

Delete project



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
    id := TODO // string | Project ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.ProjectsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> ListProjectsResponse ProjectsList(ctx).Page(page).PerPage(perPage).SortBy(sortBy).Execute()

List Projects



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
    resp, r, err := api_client.ProjectsApi.ProjectsList(context.Background()).Page(page).PerPage(perPage).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsList`: ListProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Sets how many results you would like to have on each page | 
 **perPage** | **int32** | Sets what page you would like to view | 
 **sortBy** | **string** | Sets order of how you would like to sort results and also the key you would like to order by | 

### Return type

[**ListProjectsResponse**](ListProjectsResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsShow

> Project ProjectsShow(ctx, id).Execute()

Get project by ID



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
    id := TODO // string | Project ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.ProjectsShow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsShow`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

