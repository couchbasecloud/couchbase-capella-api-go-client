# V3ServersStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**V3StorageType**](V3StorageType.md) |  | 
**IOPS** | Pointer to **int32** | Required for GP3 and IO2. Min 3000 for GP3, 1000 if IO2. If storageType&#x3D;\&quot;GP3\&quot;, max &#x3D; 16000. If storageType&#x3D; \&quot;IO2\&quot;, max&#x3D; 64000 | [optional] 
**Size** | **int32** | Min 50GB, max 16TB | 

## Methods

### NewV3ServersStorage

`func NewV3ServersStorage(type_ V3StorageType, size int32, ) *V3ServersStorage`

NewV3ServersStorage instantiates a new V3ServersStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ServersStorageWithDefaults

`func NewV3ServersStorageWithDefaults() *V3ServersStorage`

NewV3ServersStorageWithDefaults instantiates a new V3ServersStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V3ServersStorage) GetType() V3StorageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V3ServersStorage) GetTypeOk() (*V3StorageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V3ServersStorage) SetType(v V3StorageType)`

SetType sets Type field to given value.


### GetIOPS

`func (o *V3ServersStorage) GetIOPS() int32`

GetIOPS returns the IOPS field if non-nil, zero value otherwise.

### GetIOPSOk

`func (o *V3ServersStorage) GetIOPSOk() (*int32, bool)`

GetIOPSOk returns a tuple with the IOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOPS

`func (o *V3ServersStorage) SetIOPS(v int32)`

SetIOPS sets IOPS field to given value.

### HasIOPS

`func (o *V3ServersStorage) HasIOPS() bool`

HasIOPS returns a boolean if a field has been set.

### GetSize

`func (o *V3ServersStorage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V3ServersStorage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V3ServersStorage) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


