# ListProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | [**Cursor**](Cursor.md) |  | 
**Data** | [**[]Project**](Project.md) |  | 

## Methods

### NewListProjectsResponse

`func NewListProjectsResponse(cursor Cursor, data []Project, ) *ListProjectsResponse`

NewListProjectsResponse instantiates a new ListProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProjectsResponseWithDefaults

`func NewListProjectsResponseWithDefaults() *ListProjectsResponse`

NewListProjectsResponseWithDefaults instantiates a new ListProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListProjectsResponse) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListProjectsResponse) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListProjectsResponse) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.


### GetData

`func (o *ListProjectsResponse) GetData() []Project`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListProjectsResponse) GetDataOk() (*[]Project, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListProjectsResponse) SetData(v []Project)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


