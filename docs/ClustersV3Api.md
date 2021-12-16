# \ClustersV3Api

All URIs are relative to *https://cloudapi.cloud.couchbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersV3create**](ClustersV3Api.md#ClustersV3create) | **Post** /v3/clusters | Create Cluster V3
[**ClustersV3createUser**](ClustersV3Api.md#ClustersV3createUser) | **Post** /v3/clusters/{id}/users | Create cluster user
[**ClustersV3delete**](ClustersV3Api.md#ClustersV3delete) | **Delete** /v3/clusters/{id} | Delete Cluster V3
[**ClustersV3list**](ClustersV3Api.md#ClustersV3list) | **Get** /v3/clusters | List all clusters v3
[**ClustersV3show**](ClustersV3Api.md#ClustersV3show) | **Get** /v3/clusters/{id} | Get Cluster Info V3
[**ClustersV3status**](ClustersV3Api.md#ClustersV3status) | **Get** /v3/clusters/{id}/status | Get Cluster Status V3
[**ClustersV3updateMeta**](ClustersV3Api.md#ClustersV3updateMeta) | **Put** /v3/clusters/{id}/meta | Update Cluster Metadata V3
[**ClustersV3updateServers**](ClustersV3Api.md#ClustersV3updateServers) | **Put** /v3/clusters/{id}/servers | Update Cluster Servers V3
[**ClustersV3updateSupport**](ClustersV3Api.md#ClustersV3updateSupport) | **Put** /v3/clusters/{id}/support | Update Cluster Support Package V3



## ClustersV3create

> ClustersV3create(ctx).V3CreateClusterRequest(v3CreateClusterRequest).Execute()

Create Cluster V3



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
    v3CreateClusterRequest := *openapiclient.NewV3CreateClusterRequest(openapiclient.v3Environment("hosted"), "Demo G2 Cluster", "99d9487a-235b-4b5f-b610-577a60edb911", *openapiclient.NewV3Place(true), []openapiclient.V3Servers{*openapiclient.NewV3Servers(int32(4), "m5.xlarge", []openapiclient.V3CouchbaseServices{openapiclient.v3CouchbaseServices("data")}, *openapiclient.NewV3ServersStorage(openapiclient.v3StorageType("GP3"), int32(3000), int32(50)))}, *openapiclient.NewV3SupportPackage(openapiclient.v3SupportPackageTimezones("ET"), openapiclient.v3SupportPackageType("Basic"))) // V3CreateClusterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3create(context.Background()).V3CreateClusterRequest(v3CreateClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3createRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3CreateClusterRequest** | [**V3CreateClusterRequest**](V3CreateClusterRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersV3createUser

> ClustersV3createUser(ctx, id).V3CreateClusterUserRequest(v3CreateClusterUserRequest).Execute()

Create cluster user



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
    id := TODO // string | Cluster ID
    v3CreateClusterUserRequest := *openapiclient.NewV3CreateClusterUserRequest("user1", "Password123!") // V3CreateClusterUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3createUser(context.Background(), id).V3CreateClusterUserRequest(v3CreateClusterUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3createUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3createUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3CreateClusterUserRequest** | [**V3CreateClusterUserRequest**](V3CreateClusterUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersV3delete

> ClustersV3delete(ctx, id).Execute()

Delete Cluster V3



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
    id := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3delete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3deleteRequest struct via the builder pattern


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


## ClustersV3list

> V3ClusterList ClustersV3list(ctx).Page(page).PerPage(perPage).CloudId(cloudId).ProjectId(projectId).Execute()

List all clusters v3



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
    cloudId := TODO // string | Cloud ID for filtering cloud clusters. (optional)
    projectId := TODO // string | Project ID for filtering project clusters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3list(context.Background()).Page(page).PerPage(perPage).CloudId(cloudId).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3list``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersV3list`: V3ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersV3Api.ClustersV3list`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3listRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Sets how many results you would like to have on each page | 
 **perPage** | **int32** | Sets what page you would like to view | 
 **cloudId** | [**string**](string.md) | Cloud ID for filtering cloud clusters. | 
 **projectId** | [**string**](string.md) | Project ID for filtering project clusters. | 

### Return type

[**V3ClusterList**](V3ClusterList.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersV3show

> V3Cluster ClustersV3show(ctx, id).Execute()

Get Cluster Info V3



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
    id := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3show(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3show``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersV3show`: V3Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersV3Api.ClustersV3show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3showRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3Cluster**](V3Cluster.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersV3status

> V3ClusterStatusResponse ClustersV3status(ctx, id).Execute()

Get Cluster Status V3



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
    id := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3status(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3status``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersV3status`: V3ClusterStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersV3Api.ClustersV3status`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3statusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3ClusterStatusResponse**](V3ClusterStatusResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersV3updateMeta

> ClustersV3updateMeta(ctx, id).V3UpdateClusterMetaRequest(v3UpdateClusterMetaRequest).Execute()

Update Cluster Metadata V3



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
    id := TODO // string | Cluster ID
    v3UpdateClusterMetaRequest := *openapiclient.NewV3UpdateClusterMetaRequest() // V3UpdateClusterMetaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3updateMeta(context.Background(), id).V3UpdateClusterMetaRequest(v3UpdateClusterMetaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3updateMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3updateMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3UpdateClusterMetaRequest** | [**V3UpdateClusterMetaRequest**](V3UpdateClusterMetaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersV3updateServers

> ClustersV3updateServers(ctx, id).V3UpdateClusterServersRequest(v3UpdateClusterServersRequest).Execute()

Update Cluster Servers V3



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
    id := TODO // string | Cluster ID
    v3UpdateClusterServersRequest := *openapiclient.NewV3UpdateClusterServersRequest([]openapiclient.V3Servers{*openapiclient.NewV3Servers(int32(4), "m5.xlarge", []openapiclient.V3CouchbaseServices{openapiclient.v3CouchbaseServices("data")}, *openapiclient.NewV3ServersStorage(openapiclient.v3StorageType("GP3"), int32(3000), int32(50)))}) // V3UpdateClusterServersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3updateServers(context.Background(), id).V3UpdateClusterServersRequest(v3UpdateClusterServersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3updateServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3updateServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3UpdateClusterServersRequest** | [**V3UpdateClusterServersRequest**](V3UpdateClusterServersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersV3updateSupport

> ClustersV3updateSupport(ctx, id).V3UpdateClusterSupportRequest(v3UpdateClusterSupportRequest).Execute()

Update Cluster Support Package V3



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
    id := TODO // string | Cluster ID
    v3UpdateClusterSupportRequest := *openapiclient.NewV3UpdateClusterSupportRequest(*openapiclient.NewV3UpdateClusterSupportRequestSupportPackage(openapiclient.v3SupportPackageType("Basic"))) // V3UpdateClusterSupportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersV3Api.ClustersV3updateSupport(context.Background(), id).V3UpdateClusterSupportRequest(v3UpdateClusterSupportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersV3Api.ClustersV3updateSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersV3updateSupportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3UpdateClusterSupportRequest** | [**V3UpdateClusterSupportRequest**](V3UpdateClusterSupportRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

