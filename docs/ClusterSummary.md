# ClusterSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TenantId** | **string** |  | 
**CloudId** | **string** |  | 
**ProjectId** | **string** |  | 
**Services** | [**[]CouchbaseServices**](CouchbaseServices.md) |  | 
**Nodes** | **int32** |  | 

## Methods

### NewClusterSummary

`func NewClusterSummary(id string, name string, tenantId string, cloudId string, projectId string, services []CouchbaseServices, nodes int32, ) *ClusterSummary`

NewClusterSummary instantiates a new ClusterSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSummaryWithDefaults

`func NewClusterSummaryWithDefaults() *ClusterSummary`

NewClusterSummaryWithDefaults instantiates a new ClusterSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ClusterSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterSummary) SetName(v string)`

SetName sets Name field to given value.


### GetTenantId

`func (o *ClusterSummary) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ClusterSummary) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ClusterSummary) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCloudId

`func (o *ClusterSummary) GetCloudId() string`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *ClusterSummary) GetCloudIdOk() (*string, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *ClusterSummary) SetCloudId(v string)`

SetCloudId sets CloudId field to given value.


### GetProjectId

`func (o *ClusterSummary) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ClusterSummary) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ClusterSummary) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetServices

`func (o *ClusterSummary) GetServices() []CouchbaseServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ClusterSummary) GetServicesOk() (*[]CouchbaseServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ClusterSummary) SetServices(v []CouchbaseServices)`

SetServices sets Services field to given value.


### GetNodes

`func (o *ClusterSummary) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterSummary) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterSummary) SetNodes(v int32)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


