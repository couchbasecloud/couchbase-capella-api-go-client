# ServerAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSize** | [**AzureInstances**](AzureInstances.md) |  | 
**VolumeType** | [**AzureVolumeTypes**](AzureVolumeTypes.md) |  | 

## Methods

### NewServerAzure

`func NewServerAzure(instanceSize AzureInstances, volumeType AzureVolumeTypes, ) *ServerAzure`

NewServerAzure instantiates a new ServerAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerAzureWithDefaults

`func NewServerAzureWithDefaults() *ServerAzure`

NewServerAzureWithDefaults instantiates a new ServerAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSize

`func (o *ServerAzure) GetInstanceSize() AzureInstances`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ServerAzure) GetInstanceSizeOk() (*AzureInstances, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ServerAzure) SetInstanceSize(v AzureInstances)`

SetInstanceSize sets InstanceSize field to given value.


### GetVolumeType

`func (o *ServerAzure) GetVolumeType() AzureVolumeTypes`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ServerAzure) GetVolumeTypeOk() (*AzureVolumeTypes, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ServerAzure) SetVolumeType(v AzureVolumeTypes)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


