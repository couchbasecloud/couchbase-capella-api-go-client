# V3Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TenantId** | **string** |  | 
**ProjectId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Status** | **string** |  | 
**Version** | [**V3ClusterVersion**](V3ClusterVersion.md) |  | 
**EndpointsSrv** | Pointer to **string** |  | [optional] 
**Environment** | **string** |  | 
**Place** | [**V3ClusterPlace**](V3ClusterPlace.md) |  | 
**Servers** | [**[]V3ClusterServers**](V3ClusterServers.md) |  | 
**AvailabilityZones** | **[]string** |  | 
**Support** | **string** |  | 

## Methods

### NewV3Cluster

`func NewV3Cluster(id string, name string, tenantId string, projectId string, createdAt time.Time, updatedAt time.Time, status string, version V3ClusterVersion, environment string, place V3ClusterPlace, servers []V3ClusterServers, availabilityZones []string, support string, ) *V3Cluster`

NewV3Cluster instantiates a new V3Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ClusterWithDefaults

`func NewV3ClusterWithDefaults() *V3Cluster`

NewV3ClusterWithDefaults instantiates a new V3Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V3Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V3Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V3Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *V3Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetTenantId

`func (o *V3Cluster) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *V3Cluster) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *V3Cluster) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetProjectId

`func (o *V3Cluster) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *V3Cluster) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *V3Cluster) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *V3Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V3Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V3Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *V3Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V3Cluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V3Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *V3Cluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V3Cluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V3Cluster) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *V3Cluster) GetVersion() V3ClusterVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V3Cluster) GetVersionOk() (*V3ClusterVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V3Cluster) SetVersion(v V3ClusterVersion)`

SetVersion sets Version field to given value.


### GetEndpointsSrv

`func (o *V3Cluster) GetEndpointsSrv() string`

GetEndpointsSrv returns the EndpointsSrv field if non-nil, zero value otherwise.

### GetEndpointsSrvOk

`func (o *V3Cluster) GetEndpointsSrvOk() (*string, bool)`

GetEndpointsSrvOk returns a tuple with the EndpointsSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsSrv

`func (o *V3Cluster) SetEndpointsSrv(v string)`

SetEndpointsSrv sets EndpointsSrv field to given value.

### HasEndpointsSrv

`func (o *V3Cluster) HasEndpointsSrv() bool`

HasEndpointsSrv returns a boolean if a field has been set.

### GetEnvironment

`func (o *V3Cluster) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V3Cluster) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V3Cluster) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetPlace

`func (o *V3Cluster) GetPlace() V3ClusterPlace`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *V3Cluster) GetPlaceOk() (*V3ClusterPlace, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *V3Cluster) SetPlace(v V3ClusterPlace)`

SetPlace sets Place field to given value.


### GetServers

`func (o *V3Cluster) GetServers() []V3ClusterServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *V3Cluster) GetServersOk() (*[]V3ClusterServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *V3Cluster) SetServers(v []V3ClusterServers)`

SetServers sets Servers field to given value.


### GetAvailabilityZones

`func (o *V3Cluster) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *V3Cluster) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *V3Cluster) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetSupport

`func (o *V3Cluster) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *V3Cluster) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *V3Cluster) SetSupport(v string)`

SetSupport sets Support field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


