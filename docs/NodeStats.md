# NodeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**TotalCount** | **int32** |  | 
**FailedOverCount** | **int32** |  | 
**UnhealthyCount** | **int32** |  | 
**ServiceStats** | [**[]ServiceStats**](ServiceStats.md) |  | 

## Methods

### NewNodeStats

`func NewNodeStats(status string, totalCount int32, failedOverCount int32, unhealthyCount int32, serviceStats []ServiceStats, ) *NodeStats`

NewNodeStats instantiates a new NodeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStatsWithDefaults

`func NewNodeStatsWithDefaults() *NodeStats`

NewNodeStatsWithDefaults instantiates a new NodeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NodeStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeStats) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalCount

`func (o *NodeStats) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *NodeStats) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *NodeStats) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetFailedOverCount

`func (o *NodeStats) GetFailedOverCount() int32`

GetFailedOverCount returns the FailedOverCount field if non-nil, zero value otherwise.

### GetFailedOverCountOk

`func (o *NodeStats) GetFailedOverCountOk() (*int32, bool)`

GetFailedOverCountOk returns a tuple with the FailedOverCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOverCount

`func (o *NodeStats) SetFailedOverCount(v int32)`

SetFailedOverCount sets FailedOverCount field to given value.


### GetUnhealthyCount

`func (o *NodeStats) GetUnhealthyCount() int32`

GetUnhealthyCount returns the UnhealthyCount field if non-nil, zero value otherwise.

### GetUnhealthyCountOk

`func (o *NodeStats) GetUnhealthyCountOk() (*int32, bool)`

GetUnhealthyCountOk returns a tuple with the UnhealthyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyCount

`func (o *NodeStats) SetUnhealthyCount(v int32)`

SetUnhealthyCount sets UnhealthyCount field to given value.


### GetServiceStats

`func (o *NodeStats) GetServiceStats() []ServiceStats`

GetServiceStats returns the ServiceStats field if non-nil, zero value otherwise.

### GetServiceStatsOk

`func (o *NodeStats) GetServiceStatsOk() (*[]ServiceStats, bool)`

GetServiceStatsOk returns a tuple with the ServiceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStats

`func (o *NodeStats) SetServiceStats(v []ServiceStats)`

SetServiceStats sets ServiceStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


