# CreateDatabaseUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Organisation user which is assigned to the database user.  | [optional] 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Access** | [**[]BucketRole**](BucketRole.md) |  | 

## Methods

### NewCreateDatabaseUserRequest

`func NewCreateDatabaseUserRequest(username string, password string, access []BucketRole, ) *CreateDatabaseUserRequest`

NewCreateDatabaseUserRequest instantiates a new CreateDatabaseUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseUserRequestWithDefaults

`func NewCreateDatabaseUserRequestWithDefaults() *CreateDatabaseUserRequest`

NewCreateDatabaseUserRequestWithDefaults instantiates a new CreateDatabaseUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateDatabaseUserRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateDatabaseUserRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateDatabaseUserRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateDatabaseUserRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *CreateDatabaseUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateDatabaseUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateDatabaseUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CreateDatabaseUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDatabaseUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDatabaseUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAccess

`func (o *CreateDatabaseUserRequest) GetAccess() []BucketRole`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CreateDatabaseUserRequest) GetAccessOk() (*[]BucketRole, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CreateDatabaseUserRequest) SetAccess(v []BucketRole)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


