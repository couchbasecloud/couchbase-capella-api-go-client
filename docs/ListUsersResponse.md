# ListUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | [**Cursor**](Cursor.md) |  | 
**Data** | [**[]ListUsersResponseItem**](ListUsersResponseItem.md) |  | 

## Methods

### NewListUsersResponse

`func NewListUsersResponse(cursor Cursor, data []ListUsersResponseItem, ) *ListUsersResponse`

NewListUsersResponse instantiates a new ListUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersResponseWithDefaults

`func NewListUsersResponseWithDefaults() *ListUsersResponse`

NewListUsersResponseWithDefaults instantiates a new ListUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListUsersResponse) GetCursor() Cursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListUsersResponse) GetCursorOk() (*Cursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListUsersResponse) SetCursor(v Cursor)`

SetCursor sets Cursor field to given value.


### GetData

`func (o *ListUsersResponse) GetData() []ListUsersResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListUsersResponse) GetDataOk() (*[]ListUsersResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListUsersResponse) SetData(v []ListUsersResponseItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


