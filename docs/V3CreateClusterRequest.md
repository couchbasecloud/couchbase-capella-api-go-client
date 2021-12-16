# V3CreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**V3Environment**](V3Environment.md) |  | 
**ClusterName** | **string** | Name of the cluster | 
**ProjectId** | **string** | Uniquely identifiable projectId | 
**Description** | Pointer to **string** | A short description that describes the cluster | [optional] 
**Place** | [**V3Place**](V3Place.md) |  | 
**Servers** | [**[]V3Servers**](V3Servers.md) | Server specifications for creating the cluster | 
**SupportPackage** | [**V3SupportPackage**](V3SupportPackage.md) |  | 

## Methods

### NewV3CreateClusterRequest

`func NewV3CreateClusterRequest(environment V3Environment, clusterName string, projectId string, place V3Place, servers []V3Servers, supportPackage V3SupportPackage, ) *V3CreateClusterRequest`

NewV3CreateClusterRequest instantiates a new V3CreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3CreateClusterRequestWithDefaults

`func NewV3CreateClusterRequestWithDefaults() *V3CreateClusterRequest`

NewV3CreateClusterRequestWithDefaults instantiates a new V3CreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *V3CreateClusterRequest) GetEnvironment() V3Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V3CreateClusterRequest) GetEnvironmentOk() (*V3Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V3CreateClusterRequest) SetEnvironment(v V3Environment)`

SetEnvironment sets Environment field to given value.


### GetClusterName

`func (o *V3CreateClusterRequest) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *V3CreateClusterRequest) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *V3CreateClusterRequest) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetProjectId

`func (o *V3CreateClusterRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *V3CreateClusterRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *V3CreateClusterRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDescription

`func (o *V3CreateClusterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V3CreateClusterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V3CreateClusterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V3CreateClusterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlace

`func (o *V3CreateClusterRequest) GetPlace() V3Place`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *V3CreateClusterRequest) GetPlaceOk() (*V3Place, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *V3CreateClusterRequest) SetPlace(v V3Place)`

SetPlace sets Place field to given value.


### GetServers

`func (o *V3CreateClusterRequest) GetServers() []V3Servers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *V3CreateClusterRequest) GetServersOk() (*[]V3Servers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *V3CreateClusterRequest) SetServers(v []V3Servers)`

SetServers sets Servers field to given value.


### GetSupportPackage

`func (o *V3CreateClusterRequest) GetSupportPackage() V3SupportPackage`

GetSupportPackage returns the SupportPackage field if non-nil, zero value otherwise.

### GetSupportPackageOk

`func (o *V3CreateClusterRequest) GetSupportPackageOk() (*V3SupportPackage, bool)`

GetSupportPackageOk returns a tuple with the SupportPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPackage

`func (o *V3CreateClusterRequest) SetSupportPackage(v V3SupportPackage)`

SetSupportPackage sets SupportPackage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


