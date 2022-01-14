# UpdateDatabaseUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Organization user to update which is assigned to the database user.  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Buckets** | Pointer to [**[]BucketRole**](BucketRole.md) |  | [optional] 
**AllBucketsAccess** | Pointer to [**BucketRoleTypes**](BucketRoleTypes.md) |  | [optional] 

## Methods

### NewUpdateDatabaseUserRequest

`func NewUpdateDatabaseUserRequest() *UpdateDatabaseUserRequest`

NewUpdateDatabaseUserRequest instantiates a new UpdateDatabaseUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatabaseUserRequestWithDefaults

`func NewUpdateDatabaseUserRequestWithDefaults() *UpdateDatabaseUserRequest`

NewUpdateDatabaseUserRequestWithDefaults instantiates a new UpdateDatabaseUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UpdateDatabaseUserRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateDatabaseUserRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateDatabaseUserRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UpdateDatabaseUserRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateDatabaseUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateDatabaseUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateDatabaseUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateDatabaseUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateDatabaseUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateDatabaseUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateDatabaseUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateDatabaseUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetBuckets

`func (o *UpdateDatabaseUserRequest) GetBuckets() []BucketRole`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *UpdateDatabaseUserRequest) GetBucketsOk() (*[]BucketRole, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *UpdateDatabaseUserRequest) SetBuckets(v []BucketRole)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *UpdateDatabaseUserRequest) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetAllBucketsAccess

`func (o *UpdateDatabaseUserRequest) GetAllBucketsAccess() BucketRoleTypes`

GetAllBucketsAccess returns the AllBucketsAccess field if non-nil, zero value otherwise.

### GetAllBucketsAccessOk

`func (o *UpdateDatabaseUserRequest) GetAllBucketsAccessOk() (*BucketRoleTypes, bool)`

GetAllBucketsAccessOk returns a tuple with the AllBucketsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBucketsAccess

`func (o *UpdateDatabaseUserRequest) SetAllBucketsAccess(v BucketRoleTypes)`

SetAllBucketsAccess sets AllBucketsAccess field to given value.

### HasAllBucketsAccess

`func (o *UpdateDatabaseUserRequest) HasAllBucketsAccess() bool`

HasAllBucketsAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


