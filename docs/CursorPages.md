# CursorPages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Current page starting from 1. | [optional] 
**Next** | Pointer to **int32** | Next page number, it is not set on the last page.  | [optional] 
**Previous** | Pointer to **int32** | Previous page number, it is not set on the first page.  | [optional] 
**Last** | Pointer to **int32** | Last page number. | [optional] 
**PerPage** | Pointer to **int32** | How many items are displayed in the page. | [optional] 
**TotalItems** | **int32** | Total items found by the given query. | 

## Methods

### NewCursorPages

`func NewCursorPages(totalItems int32, ) *CursorPages`

NewCursorPages instantiates a new CursorPages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPagesWithDefaults

`func NewCursorPagesWithDefaults() *CursorPages`

NewCursorPagesWithDefaults instantiates a new CursorPages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CursorPages) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CursorPages) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CursorPages) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CursorPages) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetNext

`func (o *CursorPages) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CursorPages) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CursorPages) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *CursorPages) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *CursorPages) GetPrevious() int32`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *CursorPages) GetPreviousOk() (*int32, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *CursorPages) SetPrevious(v int32)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *CursorPages) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetLast

`func (o *CursorPages) GetLast() int32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *CursorPages) GetLastOk() (*int32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *CursorPages) SetLast(v int32)`

SetLast sets Last field to given value.

### HasLast

`func (o *CursorPages) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetPerPage

`func (o *CursorPages) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *CursorPages) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *CursorPages) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *CursorPages) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTotalItems

`func (o *CursorPages) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CursorPages) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CursorPages) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


