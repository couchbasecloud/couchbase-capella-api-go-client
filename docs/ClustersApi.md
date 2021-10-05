# \ClustersApi

All URIs are relative to *https://cloudapi.cloud.couchbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersCreate**](ClustersApi.md#ClustersCreate) | **Post** /v2/clusters | Create cluster
[**ClustersCreateAllowlistEntry**](ClustersApi.md#ClustersCreateAllowlistEntry) | **Post** /v2/clusters/{id}/allowlist | Adds entry to allowlist
[**ClustersCreateBucket**](ClustersApi.md#ClustersCreateBucket) | **Post** /v2/clusters/{id}/buckets | Create bucket in cluster
[**ClustersCreateUser**](ClustersApi.md#ClustersCreateUser) | **Post** /v2/clusters/{id}/users | Create Database User
[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /v2/clusters/{id} | Delete cluster
[**ClustersDeleteAllowlistEntry**](ClustersApi.md#ClustersDeleteAllowlistEntry) | **Delete** /v2/clusters/{id}/allowlist | Delete entry from allowlist
[**ClustersDeleteBucket**](ClustersApi.md#ClustersDeleteBucket) | **Delete** /v2/clusters/{id}/buckets | Delete bucket in cluster
[**ClustersDeleteUser**](ClustersApi.md#ClustersDeleteUser) | **Delete** /v2/clusters/{id}/users/{username} | Delete Database User
[**ClustersGetAllowlist**](ClustersApi.md#ClustersGetAllowlist) | **Get** /v2/clusters/{id}/allowlist | Get current allowlist
[**ClustersGetCertificate**](ClustersApi.md#ClustersGetCertificate) | **Get** /v2/clusters/{id}/certificate | Get Cluster Certificate
[**ClustersList**](ClustersApi.md#ClustersList) | **Get** /v2/clusters | List Clusters
[**ClustersListBuckets**](ClustersApi.md#ClustersListBuckets) | **Get** /v2/clusters/{id}/buckets | List cluster buckets
[**ClustersListUsers**](ClustersApi.md#ClustersListUsers) | **Get** /v2/clusters/{id}/users | List Database Users
[**ClustersShow**](ClustersApi.md#ClustersShow) | **Get** /v2/clusters/{id} | Get Cluster



## ClustersCreate

> Cluster ClustersCreate(ctx).CreateClusterRequest(createClusterRequest).Execute()

Create cluster



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
    createClusterRequest := *openapiclient.NewCreateClusterRequest("Name_example", "CloudId_example", "ProjectId_example") // CreateClusterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersCreate(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersCreate`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersCreateAllowlistEntry

> ClustersCreateAllowlistEntry(ctx, id).AppendAllowListRequest(appendAllowListRequest).Execute()

Adds entry to allowlist



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
    appendAllowListRequest := *openapiclient.NewAppendAllowListRequest("CidrBlock_example", openapiclient.allowListRules("temporary")) // AppendAllowListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersCreateAllowlistEntry(context.Background(), id).AppendAllowListRequest(appendAllowListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersCreateAllowlistEntry``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClustersCreateAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appendAllowListRequest** | [**AppendAllowListRequest**](AppendAllowListRequest.md) |  | 

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


## ClustersCreateBucket

> CouchbaseBucketSpec ClustersCreateBucket(ctx, id).CouchbaseBucketSpec(couchbaseBucketSpec).Execute()

Create bucket in cluster



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
    couchbaseBucketSpec := *openapiclient.NewCouchbaseBucketSpec("Name_example", int32(123)) // CouchbaseBucketSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersCreateBucket(context.Background(), id).CouchbaseBucketSpec(couchbaseBucketSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersCreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersCreateBucket`: CouchbaseBucketSpec
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersCreateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **couchbaseBucketSpec** | [**CouchbaseBucketSpec**](CouchbaseBucketSpec.md) |  | 

### Return type

[**CouchbaseBucketSpec**](CouchbaseBucketSpec.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersCreateUser

> ClustersCreateUser(ctx, id).CreateDatabaseUserRequest(createDatabaseUserRequest).Execute()

Create Database User



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
    createDatabaseUserRequest := *openapiclient.NewCreateDatabaseUserRequest("Username_example", "Password_example", []openapiclient.BucketRole{*openapiclient.NewBucketRole("Name_example", []openapiclient.BucketRoleTypes{openapiclient.bucketRoleTypes("data_writer")})}) // CreateDatabaseUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersCreateUser(context.Background(), id).CreateDatabaseUserRequest(createDatabaseUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersCreateUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClustersCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDatabaseUserRequest** | [**CreateDatabaseUserRequest**](CreateDatabaseUserRequest.md) |  | 

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


## ClustersDelete

> ClustersDelete(ctx, id).Execute()

Delete cluster



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
    resp, r, err := api_client.ClustersApi.ClustersDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClustersDeleteRequest struct via the builder pattern


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


## ClustersDeleteAllowlistEntry

> ClustersDeleteAllowlistEntry(ctx, id).DeleteAllowListEntryRequest(deleteAllowListEntryRequest).Execute()

Delete entry from allowlist



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
    deleteAllowListEntryRequest := *openapiclient.NewDeleteAllowListEntryRequest("CidrBlock_example") // DeleteAllowListEntryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersDeleteAllowlistEntry(context.Background(), id).DeleteAllowListEntryRequest(deleteAllowListEntryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDeleteAllowlistEntry``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClustersDeleteAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteAllowListEntryRequest** | [**DeleteAllowListEntryRequest**](DeleteAllowListEntryRequest.md) |  | 

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


## ClustersDeleteBucket

> ClustersDeleteBucket(ctx, id).DeleteBucketRequest(deleteBucketRequest).Execute()

Delete bucket in cluster



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
    deleteBucketRequest := *openapiclient.NewDeleteBucketRequest("Name_example") // DeleteBucketRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersDeleteBucket(context.Background(), id).DeleteBucketRequest(deleteBucketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDeleteBucket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClustersDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteBucketRequest** | [**DeleteBucketRequest**](DeleteBucketRequest.md) |  | 

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


## ClustersDeleteUser

> ClustersDeleteUser(ctx, id, username).Execute()

Delete Database User



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
    username := "foobar" // string | Database User username

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersDeleteUser(context.Background(), id, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 
**username** | **string** | Database User username | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersDeleteUserRequest struct via the builder pattern


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


## ClustersGetAllowlist

> []AllowListEntry ClustersGetAllowlist(ctx, id).Execute()

Get current allowlist



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
    resp, r, err := api_client.ClustersApi.ClustersGetAllowlist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGetAllowlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersGetAllowlist`: []AllowListEntry
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGetAllowlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersGetAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AllowListEntry**](AllowListEntry.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersGetCertificate

> GetCertificateResponse ClustersGetCertificate(ctx, id).Execute()

Get Cluster Certificate



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
    resp, r, err := api_client.ClustersApi.ClustersGetCertificate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGetCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersGetCertificate`: GetCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCertificateResponse**](GetCertificateResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersList

> ListClustersResponse ClustersList(ctx).Page(page).PerPage(perPage).SortBy(sortBy).CloudId(cloudId).ProjectId(projectId).Execute()

List Clusters



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
    cloudId := TODO // string | Cloud ID for filtering cloud clusters.  (optional)
    projectId := TODO // string | Project ID for filtering project clusters.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ClustersList(context.Background()).Page(page).PerPage(perPage).SortBy(sortBy).CloudId(cloudId).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersList`: ListClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Sets how many results you would like to have on each page | 
 **perPage** | **int32** | Sets what page you would like to view | 
 **sortBy** | **string** | Sets order of how you would like to sort results and also the key you would like to order by | 
 **cloudId** | [**string**](string.md) | Cloud ID for filtering cloud clusters.  | 
 **projectId** | [**string**](string.md) | Project ID for filtering project clusters.  | 

### Return type

[**ListClustersResponse**](ListClustersResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersListBuckets

> []ListBucketItem ClustersListBuckets(ctx, id).Execute()

List cluster buckets



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
    resp, r, err := api_client.ClustersApi.ClustersListBuckets(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersListBuckets`: []ListBucketItem
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersListBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListBucketItem**](ListBucketItem.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersListUsers

> []ListDatabaseUsersResponseItem ClustersListUsers(ctx, id).Execute()

List Database Users



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
    resp, r, err := api_client.ClustersApi.ClustersListUsers(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersListUsers`: []ListDatabaseUsersResponseItem
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListDatabaseUsersResponseItem**](ListDatabaseUsersResponseItem.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersShow

> Cluster ClustersShow(ctx, id).Execute()

Get Cluster



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
    resp, r, err := api_client.ClustersApi.ClustersShow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersShow`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cluster**](Cluster.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

