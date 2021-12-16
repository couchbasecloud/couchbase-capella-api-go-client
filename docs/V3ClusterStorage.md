# V3ClusterStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IOPS** | **int32** |  | 
**Size** | **int32** |  | 

## Methods

### NewV3ClusterStorage

`func NewV3ClusterStorage(type_ string, iOPS int32, size int32, ) *V3ClusterStorage`

NewV3ClusterStorage instantiates a new V3ClusterStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ClusterStorageWithDefaults

`func NewV3ClusterStorageWithDefaults() *V3ClusterStorage`

NewV3ClusterStorageWithDefaults instantiates a new V3ClusterStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V3ClusterStorage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V3ClusterStorage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V3ClusterStorage) SetType(v string)`

SetType sets Type field to given value.


### GetIOPS

`func (o *V3ClusterStorage) GetIOPS() int32`

GetIOPS returns the IOPS field if non-nil, zero value otherwise.

### GetIOPSOk

`func (o *V3ClusterStorage) GetIOPSOk() (*int32, bool)`

GetIOPSOk returns a tuple with the IOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIOPS

`func (o *V3ClusterStorage) SetIOPS(v int32)`

SetIOPS sets IOPS field to given value.


### GetSize

`func (o *V3ClusterStorage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V3ClusterStorage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V3ClusterStorage) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


