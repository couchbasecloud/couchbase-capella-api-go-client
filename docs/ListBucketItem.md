# ListBucketItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MemoryQuota** | **int32** |  | 
**Replicas** | **int32** |  | 
**ConflictResolution** | [**ConflictResolution**](ConflictResolution.md) |  | [default to SEQNO]
**Status** | **string** |  | 

## Methods

### NewListBucketItem

`func NewListBucketItem(name string, memoryQuota int32, replicas int32, conflictResolution ConflictResolution, status string, ) *ListBucketItem`

NewListBucketItem instantiates a new ListBucketItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBucketItemWithDefaults

`func NewListBucketItemWithDefaults() *ListBucketItem`

NewListBucketItemWithDefaults instantiates a new ListBucketItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListBucketItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBucketItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBucketItem) SetName(v string)`

SetName sets Name field to given value.


### GetMemoryQuota

`func (o *ListBucketItem) GetMemoryQuota() int32`

GetMemoryQuota returns the MemoryQuota field if non-nil, zero value otherwise.

### GetMemoryQuotaOk

`func (o *ListBucketItem) GetMemoryQuotaOk() (*int32, bool)`

GetMemoryQuotaOk returns a tuple with the MemoryQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryQuota

`func (o *ListBucketItem) SetMemoryQuota(v int32)`

SetMemoryQuota sets MemoryQuota field to given value.


### GetReplicas

`func (o *ListBucketItem) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ListBucketItem) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ListBucketItem) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetConflictResolution

`func (o *ListBucketItem) GetConflictResolution() ConflictResolution`

GetConflictResolution returns the ConflictResolution field if non-nil, zero value otherwise.

### GetConflictResolutionOk

`func (o *ListBucketItem) GetConflictResolutionOk() (*ConflictResolution, bool)`

GetConflictResolutionOk returns a tuple with the ConflictResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolution

`func (o *ListBucketItem) SetConflictResolution(v ConflictResolution)`

SetConflictResolution sets ConflictResolution field to given value.


### GetStatus

`func (o *ListBucketItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBucketItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBucketItem) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


