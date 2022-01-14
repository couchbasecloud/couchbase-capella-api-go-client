# Cursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pages** | [**CursorPages**](CursorPages.md) |  | 
**Hrefs** | [**CursorHrefs**](CursorHrefs.md) |  | 

## Methods

### NewCursor

`func NewCursor(pages CursorPages, hrefs CursorHrefs, ) *Cursor`

NewCursor instantiates a new Cursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorWithDefaults

`func NewCursorWithDefaults() *Cursor`

NewCursorWithDefaults instantiates a new Cursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPages

`func (o *Cursor) GetPages() CursorPages`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *Cursor) GetPagesOk() (*CursorPages, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *Cursor) SetPages(v CursorPages)`

SetPages sets Pages field to given value.


### GetHrefs

`func (o *Cursor) GetHrefs() CursorHrefs`

GetHrefs returns the Hrefs field if non-nil, zero value otherwise.

### GetHrefsOk

`func (o *Cursor) GetHrefsOk() (*CursorHrefs, bool)`

GetHrefsOk returns a tuple with the Hrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHrefs

`func (o *Cursor) SetHrefs(v CursorHrefs)`

SetHrefs sets Hrefs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


