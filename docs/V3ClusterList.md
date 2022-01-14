# V3ClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | [**Cursor**](Cursor.md) |  | 
**Data** | [**V3ClusterListData**](V3ClusterListData.md) |  | 

## Methods

### NewV3ClusterList

`func NewV3ClusterList(cursor Cursor, data V3ClusterListData, ) *V3ClusterList`

NewV3ClusterList instantiates a new V3ClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ClusterListWithDefaults

`func NewV3ClusterListWithDefaults() *V3ClusterList`

NewV3ClusterListWithDefaults instantiates a new V3ClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *V3ClusterList) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *V3ClusterList) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *V3ClusterList) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.


### GetData

`func (o *V3ClusterList) GetData() V3ClusterListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V3ClusterList) GetDataOk() (*V3ClusterListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V3ClusterList) SetData(v V3ClusterListData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


