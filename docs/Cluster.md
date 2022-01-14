# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TenantId** | **string** |  | 
**CloudId** | **string** |  | 
**ProjectId** | **string** |  | 
**Status** | [**ClusterStatus**](ClusterStatus.md) |  | 
**ResourceIdentifier** | **string** | Resource identifier name linked with the Cloud Provider  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeployedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to [**ClusterVersion**](ClusterVersion.md) |  | [optional] 
**EndpointsURL** | Pointer to **[]string** |  | [optional] 
**PrivateEndpointURL** | Pointer to **[]string** |  | [optional] 
**EndpointsSrv** | Pointer to **string** |  | [optional] 
**PrivateEndpointsSrv** | Pointer to **string** |  | [optional] 

## Methods

### NewCluster

`func NewCluster(id string, name string, tenantId string, cloudId string, projectId string, status ClusterStatus, resourceIdentifier string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetTenantId

`func (o *Cluster) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Cluster) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Cluster) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCloudId

`func (o *Cluster) GetCloudId() string`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *Cluster) GetCloudIdOk() (*string, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *Cluster) SetCloudId(v string)`

SetCloudId sets CloudId field to given value.


### GetProjectId

`func (o *Cluster) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Cluster) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Cluster) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStatus

`func (o *Cluster) GetStatus() ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v ClusterStatus)`

SetStatus sets Status field to given value.


### GetResourceIdentifier

`func (o *Cluster) GetResourceIdentifier() string`

GetResourceIdentifier returns the ResourceIdentifier field if non-nil, zero value otherwise.

### GetResourceIdentifierOk

`func (o *Cluster) GetResourceIdentifierOk() (*string, bool)`

GetResourceIdentifierOk returns a tuple with the ResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIdentifier

`func (o *Cluster) SetResourceIdentifier(v string)`

SetResourceIdentifier sets ResourceIdentifier field to given value.


### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Cluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Cluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *Cluster) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *Cluster) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *Cluster) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *Cluster) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Cluster) GetVersion() ClusterVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*ClusterVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v ClusterVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEndpointsURL

`func (o *Cluster) GetEndpointsURL() []string`

GetEndpointsURL returns the EndpointsURL field if non-nil, zero value otherwise.

### GetEndpointsURLOk

`func (o *Cluster) GetEndpointsURLOk() (*[]string, bool)`

GetEndpointsURLOk returns a tuple with the EndpointsURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsURL

`func (o *Cluster) SetEndpointsURL(v []string)`

SetEndpointsURL sets EndpointsURL field to given value.

### HasEndpointsURL

`func (o *Cluster) HasEndpointsURL() bool`

HasEndpointsURL returns a boolean if a field has been set.

### GetPrivateEndpointURL

`func (o *Cluster) GetPrivateEndpointURL() []string`

GetPrivateEndpointURL returns the PrivateEndpointURL field if non-nil, zero value otherwise.

### GetPrivateEndpointURLOk

`func (o *Cluster) GetPrivateEndpointURLOk() (*[]string, bool)`

GetPrivateEndpointURLOk returns a tuple with the PrivateEndpointURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointURL

`func (o *Cluster) SetPrivateEndpointURL(v []string)`

SetPrivateEndpointURL sets PrivateEndpointURL field to given value.

### HasPrivateEndpointURL

`func (o *Cluster) HasPrivateEndpointURL() bool`

HasPrivateEndpointURL returns a boolean if a field has been set.

### GetEndpointsSrv

`func (o *Cluster) GetEndpointsSrv() string`

GetEndpointsSrv returns the EndpointsSrv field if non-nil, zero value otherwise.

### GetEndpointsSrvOk

`func (o *Cluster) GetEndpointsSrvOk() (*string, bool)`

GetEndpointsSrvOk returns a tuple with the EndpointsSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsSrv

`func (o *Cluster) SetEndpointsSrv(v string)`

SetEndpointsSrv sets EndpointsSrv field to given value.

### HasEndpointsSrv

`func (o *Cluster) HasEndpointsSrv() bool`

HasEndpointsSrv returns a boolean if a field has been set.

### GetPrivateEndpointsSrv

`func (o *Cluster) GetPrivateEndpointsSrv() string`

GetPrivateEndpointsSrv returns the PrivateEndpointsSrv field if non-nil, zero value otherwise.

### GetPrivateEndpointsSrvOk

`func (o *Cluster) GetPrivateEndpointsSrvOk() (*string, bool)`

GetPrivateEndpointsSrvOk returns a tuple with the PrivateEndpointsSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointsSrv

`func (o *Cluster) SetPrivateEndpointsSrv(v string)`

SetPrivateEndpointsSrv sets PrivateEndpointsSrv field to given value.

### HasPrivateEndpointsSrv

`func (o *Cluster) HasPrivateEndpointsSrv() bool`

HasPrivateEndpointsSrv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


