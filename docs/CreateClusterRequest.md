# CreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CloudId** | **string** |  | 
**ProjectId** | **string** |  | 
**Servers** | Pointer to [**[]Server**](Server.md) |  | [optional] 
**SupportPackage** | Pointer to [**CreateClusterRequestSupportPackage**](CreateClusterRequestSupportPackage.md) |  | [optional] 
**Version** | Pointer to [**ClusterVersions**](ClusterVersions.md) |  | [optional] 

## Methods

### NewCreateClusterRequest

`func NewCreateClusterRequest(name string, cloudId string, projectId string, ) *CreateClusterRequest`

NewCreateClusterRequest instantiates a new CreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterRequestWithDefaults

`func NewCreateClusterRequestWithDefaults() *CreateClusterRequest`

NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCloudId

`func (o *CreateClusterRequest) GetCloudId() string`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *CreateClusterRequest) GetCloudIdOk() (*string, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *CreateClusterRequest) SetCloudId(v string)`

SetCloudId sets CloudId field to given value.


### GetProjectId

`func (o *CreateClusterRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateClusterRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateClusterRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetServers

`func (o *CreateClusterRequest) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *CreateClusterRequest) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *CreateClusterRequest) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *CreateClusterRequest) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSupportPackage

`func (o *CreateClusterRequest) GetSupportPackage() CreateClusterRequestSupportPackage`

GetSupportPackage returns the SupportPackage field if non-nil, zero value otherwise.

### GetSupportPackageOk

`func (o *CreateClusterRequest) GetSupportPackageOk() (*CreateClusterRequestSupportPackage, bool)`

GetSupportPackageOk returns a tuple with the SupportPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPackage

`func (o *CreateClusterRequest) SetSupportPackage(v CreateClusterRequestSupportPackage)`

SetSupportPackage sets SupportPackage field to given value.

### HasSupportPackage

`func (o *CreateClusterRequest) HasSupportPackage() bool`

HasSupportPackage returns a boolean if a field has been set.

### GetVersion

`func (o *CreateClusterRequest) GetVersion() ClusterVersions`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateClusterRequest) GetVersionOk() (*ClusterVersions, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateClusterRequest) SetVersion(v ClusterVersions)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateClusterRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


