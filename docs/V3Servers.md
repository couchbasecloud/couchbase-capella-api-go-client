# V3Servers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** | How many nodes, value &gt;&#x3D; 3, &lt;&#x3D; 27. Total nodes in the cluster &lt;&#x3D; 27 | 
**Compute** | **string** | Compute instance type | 
**Services** | [**[]V3CouchbaseServices**](V3CouchbaseServices.md) | Array of couchbase services to be installed on the nodes | 
**Storage** | [**V3ServersStorage**](V3ServersStorage.md) |  | 

## Methods

### NewV3Servers

`func NewV3Servers(size int32, compute string, services []V3CouchbaseServices, storage V3ServersStorage, ) *V3Servers`

NewV3Servers instantiates a new V3Servers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ServersWithDefaults

`func NewV3ServersWithDefaults() *V3Servers`

NewV3ServersWithDefaults instantiates a new V3Servers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *V3Servers) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V3Servers) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V3Servers) SetSize(v int32)`

SetSize sets Size field to given value.


### GetCompute

`func (o *V3Servers) GetCompute() string`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *V3Servers) GetComputeOk() (*string, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *V3Servers) SetCompute(v string)`

SetCompute sets Compute field to given value.


### GetServices

`func (o *V3Servers) GetServices() []V3CouchbaseServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *V3Servers) GetServicesOk() (*[]V3CouchbaseServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *V3Servers) SetServices(v []V3CouchbaseServices)`

SetServices sets Services field to given value.


### GetStorage

`func (o *V3Servers) GetStorage() V3ServersStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *V3Servers) GetStorageOk() (*V3ServersStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *V3Servers) SetStorage(v V3ServersStorage)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


