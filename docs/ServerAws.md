# ServerAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSize** | [**AwsInstances**](AwsInstances.md) |  | 
**EbsSizeGib** | **int32** |  | 

## Methods

### NewServerAws

`func NewServerAws(instanceSize AwsInstances, ebsSizeGib int32, ) *ServerAws`

NewServerAws instantiates a new ServerAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerAwsWithDefaults

`func NewServerAwsWithDefaults() *ServerAws`

NewServerAwsWithDefaults instantiates a new ServerAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSize

`func (o *ServerAws) GetInstanceSize() AwsInstances`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ServerAws) GetInstanceSizeOk() (*AwsInstances, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ServerAws) SetInstanceSize(v AwsInstances)`

SetInstanceSize sets InstanceSize field to given value.


### GetEbsSizeGib

`func (o *ServerAws) GetEbsSizeGib() int32`

GetEbsSizeGib returns the EbsSizeGib field if non-nil, zero value otherwise.

### GetEbsSizeGibOk

`func (o *ServerAws) GetEbsSizeGibOk() (*int32, bool)`

GetEbsSizeGibOk returns a tuple with the EbsSizeGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsSizeGib

`func (o *ServerAws) SetEbsSizeGib(v int32)`

SetEbsSizeGib sets EbsSizeGib field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


