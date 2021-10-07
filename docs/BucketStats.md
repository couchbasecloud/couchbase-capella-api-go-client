# BucketStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**TotalCount** | **int32** |  | 
**UnhealthyCount** | **int32** |  | 
**HealthStats** | **map[string]string** |  | 

## Methods

### NewBucketStats

`func NewBucketStats(status string, totalCount int32, unhealthyCount int32, healthStats map[string]string, ) *BucketStats`

NewBucketStats instantiates a new BucketStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketStatsWithDefaults

`func NewBucketStatsWithDefaults() *BucketStats`

NewBucketStatsWithDefaults instantiates a new BucketStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BucketStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BucketStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BucketStats) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalCount

`func (o *BucketStats) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BucketStats) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BucketStats) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetUnhealthyCount

`func (o *BucketStats) GetUnhealthyCount() int32`

GetUnhealthyCount returns the UnhealthyCount field if non-nil, zero value otherwise.

### GetUnhealthyCountOk

`func (o *BucketStats) GetUnhealthyCountOk() (*int32, bool)`

GetUnhealthyCountOk returns a tuple with the UnhealthyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyCount

`func (o *BucketStats) SetUnhealthyCount(v int32)`

SetUnhealthyCount sets UnhealthyCount field to given value.


### GetHealthStats

`func (o *BucketStats) GetHealthStats() map[string]string`

GetHealthStats returns the HealthStats field if non-nil, zero value otherwise.

### GetHealthStatsOk

`func (o *BucketStats) GetHealthStatsOk() (*map[string]string, bool)`

GetHealthStatsOk returns a tuple with the HealthStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStats

`func (o *BucketStats) SetHealthStats(v map[string]string)`

SetHealthStats sets HealthStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


