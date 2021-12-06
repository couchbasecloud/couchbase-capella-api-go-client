# ClusterHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ClusterStatus**](ClusterStatus.md) |  | 
**Health** | [**ClusterHealth**](ClusterHealth.md) |  | 
**BucketStats** | Pointer to [**BucketStats**](BucketStats.md) |  | [optional] 
**NodeStats** | Pointer to [**NodeStats**](NodeStats.md) |  | [optional] 

## Methods

### NewClusterHealthResponse

`func NewClusterHealthResponse(status ClusterStatus, health ClusterHealth, ) *ClusterHealthResponse`

NewClusterHealthResponse instantiates a new ClusterHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHealthResponseWithDefaults

`func NewClusterHealthResponseWithDefaults() *ClusterHealthResponse`

NewClusterHealthResponseWithDefaults instantiates a new ClusterHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ClusterHealthResponse) GetStatus() ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterHealthResponse) GetStatusOk() (*ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterHealthResponse) SetStatus(v ClusterStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *ClusterHealthResponse) GetHealth() ClusterHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ClusterHealthResponse) GetHealthOk() (*ClusterHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ClusterHealthResponse) SetHealth(v ClusterHealth)`

SetHealth sets Health field to given value.


### GetBucketStats

`func (o *ClusterHealthResponse) GetBucketStats() BucketStats`

GetBucketStats returns the BucketStats field if non-nil, zero value otherwise.

### GetBucketStatsOk

`func (o *ClusterHealthResponse) GetBucketStatsOk() (*BucketStats, bool)`

GetBucketStatsOk returns a tuple with the BucketStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketStats

`func (o *ClusterHealthResponse) SetBucketStats(v BucketStats)`

SetBucketStats sets BucketStats field to given value.

### HasBucketStats

`func (o *ClusterHealthResponse) HasBucketStats() bool`

HasBucketStats returns a boolean if a field has been set.

### GetNodeStats

`func (o *ClusterHealthResponse) GetNodeStats() NodeStats`

GetNodeStats returns the NodeStats field if non-nil, zero value otherwise.

### GetNodeStatsOk

`func (o *ClusterHealthResponse) GetNodeStatsOk() (*NodeStats, bool)`

GetNodeStatsOk returns a tuple with the NodeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStats

`func (o *ClusterHealthResponse) SetNodeStats(v NodeStats)`

SetNodeStats sets NodeStats field to given value.

### HasNodeStats

`func (o *ClusterHealthResponse) HasNodeStats() bool`

HasNodeStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


