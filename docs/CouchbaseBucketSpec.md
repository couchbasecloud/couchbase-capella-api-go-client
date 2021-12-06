# CouchbaseBucketSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MemoryQuota** | **int32** |  | 
**Replicas** | Pointer to **int32** |  | [optional] [default to 1]
**ConflictResolution** | Pointer to [**ConflictResolution**](ConflictResolution.md) |  | [optional] [default to SEQNO]

## Methods

### NewCouchbaseBucketSpec

`func NewCouchbaseBucketSpec(name string, memoryQuota int32, ) *CouchbaseBucketSpec`

NewCouchbaseBucketSpec instantiates a new CouchbaseBucketSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseBucketSpecWithDefaults

`func NewCouchbaseBucketSpecWithDefaults() *CouchbaseBucketSpec`

NewCouchbaseBucketSpecWithDefaults instantiates a new CouchbaseBucketSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CouchbaseBucketSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CouchbaseBucketSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CouchbaseBucketSpec) SetName(v string)`

SetName sets Name field to given value.


### GetMemoryQuota

`func (o *CouchbaseBucketSpec) GetMemoryQuota() int32`

GetMemoryQuota returns the MemoryQuota field if non-nil, zero value otherwise.

### GetMemoryQuotaOk

`func (o *CouchbaseBucketSpec) GetMemoryQuotaOk() (*int32, bool)`

GetMemoryQuotaOk returns a tuple with the MemoryQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryQuota

`func (o *CouchbaseBucketSpec) SetMemoryQuota(v int32)`

SetMemoryQuota sets MemoryQuota field to given value.


### GetReplicas

`func (o *CouchbaseBucketSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CouchbaseBucketSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CouchbaseBucketSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *CouchbaseBucketSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetConflictResolution

`func (o *CouchbaseBucketSpec) GetConflictResolution() ConflictResolution`

GetConflictResolution returns the ConflictResolution field if non-nil, zero value otherwise.

### GetConflictResolutionOk

`func (o *CouchbaseBucketSpec) GetConflictResolutionOk() (*ConflictResolution, bool)`

GetConflictResolutionOk returns a tuple with the ConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolution

`func (o *CouchbaseBucketSpec) SetConflictResolution(v ConflictResolution)`

SetConflictResolution sets ConflictResolution field to given value.

### HasConflictResolution

`func (o *CouchbaseBucketSpec) HasConflictResolution() bool`

HasConflictResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


