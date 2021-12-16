# V3ClusterServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** |  | 
**Compute** | **string** | Compute instance type | 
**Services** | **[]string** |  | 
**Storage** | [**V3ClusterStorage**](V3ClusterStorage.md) |  | 

## Methods

### NewV3ClusterServers

`func NewV3ClusterServers(size int32, compute string, services []string, storage V3ClusterStorage, ) *V3ClusterServers`

NewV3ClusterServers instantiates a new V3ClusterServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ClusterServersWithDefaults

`func NewV3ClusterServersWithDefaults() *V3ClusterServers`

NewV3ClusterServersWithDefaults instantiates a new V3ClusterServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *V3ClusterServers) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V3ClusterServers) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V3ClusterServers) SetSize(v int32)`

SetSize sets Size field to given value.


### GetCompute

`func (o *V3ClusterServers) GetCompute() string`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *V3ClusterServers) GetComputeOk() (*string, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *V3ClusterServers) SetCompute(v string)`

SetCompute sets Compute field to given value.


### GetServices

`func (o *V3ClusterServers) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *V3ClusterServers) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *V3ClusterServers) SetServices(v []string)`

SetServices sets Services field to given value.


### GetStorage

`func (o *V3ClusterServers) GetStorage() V3ClusterStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *V3ClusterServers) GetStorageOk() (*V3ClusterStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *V3ClusterServers) SetStorage(v V3ClusterStorage)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


