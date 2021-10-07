# ListDatabaseUsersResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Organization user ID Database user is attached to  | [optional] 
**Username** | **string** |  | 
**Access** | [**[]BucketRole**](BucketRole.md) |  | 

## Methods

### NewListDatabaseUsersResponseItem

`func NewListDatabaseUsersResponseItem(username string, access []BucketRole, ) *ListDatabaseUsersResponseItem`

NewListDatabaseUsersResponseItem instantiates a new ListDatabaseUsersResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatabaseUsersResponseItemWithDefaults

`func NewListDatabaseUsersResponseItemWithDefaults() *ListDatabaseUsersResponseItem`

NewListDatabaseUsersResponseItemWithDefaults instantiates a new ListDatabaseUsersResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ListDatabaseUsersResponseItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListDatabaseUsersResponseItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListDatabaseUsersResponseItem) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListDatabaseUsersResponseItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *ListDatabaseUsersResponseItem) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListDatabaseUsersResponseItem) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListDatabaseUsersResponseItem) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAccess

`func (o *ListDatabaseUsersResponseItem) GetAccess() []BucketRole`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ListDatabaseUsersResponseItem) GetAccessOk() (*[]BucketRole, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ListDatabaseUsersResponseItem) SetAccess(v []BucketRole)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


