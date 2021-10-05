# ListClustersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | [**Cursor**](Cursor.md) |  | 
**Data** | [**[]ClusterSummary**](ClusterSummary.md) |  | 

## Methods

### NewListClustersResponse

`func NewListClustersResponse(cursor Cursor, data []ClusterSummary, ) *ListClustersResponse`

NewListClustersResponse instantiates a new ListClustersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClustersResponseWithDefaults

`func NewListClustersResponseWithDefaults() *ListClustersResponse`

NewListClustersResponseWithDefaults instantiates a new ListClustersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListClustersResponse) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListClustersResponse) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListClustersResponse) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.


### GetData

`func (o *ListClustersResponse) GetData() []ClusterSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListClustersResponse) GetDataOk() (*[]ClusterSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListClustersResponse) SetData(v []ClusterSummary)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


