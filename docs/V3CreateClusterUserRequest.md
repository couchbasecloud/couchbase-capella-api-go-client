# V3CreateClusterUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**AllBucketsAccess** | Pointer to [**V3BucketRoles**](V3BucketRoles.md) |  | [optional] 
**Buckets** | Pointer to [**[]V3CreateClusterUserRequestBuckets**](V3CreateClusterUserRequestBuckets.md) |  | [optional] 

## Methods

### NewV3CreateClusterUserRequest

`func NewV3CreateClusterUserRequest(username string, password string, ) *V3CreateClusterUserRequest`

NewV3CreateClusterUserRequest instantiates a new V3CreateClusterUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3CreateClusterUserRequestWithDefaults

`func NewV3CreateClusterUserRequestWithDefaults() *V3CreateClusterUserRequest`

NewV3CreateClusterUserRequestWithDefaults instantiates a new V3CreateClusterUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *V3CreateClusterUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V3CreateClusterUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V3CreateClusterUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *V3CreateClusterUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V3CreateClusterUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V3CreateClusterUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAllBucketsAccess

`func (o *V3CreateClusterUserRequest) GetAllBucketsAccess() V3BucketRoles`

GetAllBucketsAccess returns the AllBucketsAccess field if non-nil, zero value otherwise.

### GetAllBucketsAccessOk

`func (o *V3CreateClusterUserRequest) GetAllBucketsAccessOk() (*V3BucketRoles, bool)`

GetAllBucketsAccessOk returns a tuple with the AllBucketsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBucketsAccess

`func (o *V3CreateClusterUserRequest) SetAllBucketsAccess(v V3BucketRoles)`

SetAllBucketsAccess sets AllBucketsAccess field to given value.

### HasAllBucketsAccess

`func (o *V3CreateClusterUserRequest) HasAllBucketsAccess() bool`

HasAllBucketsAccess returns a boolean if a field has been set.

### GetBuckets

`func (o *V3CreateClusterUserRequest) GetBuckets() []V3CreateClusterUserRequestBuckets`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *V3CreateClusterUserRequest) GetBucketsOk() (*[]V3CreateClusterUserRequestBuckets, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *V3CreateClusterUserRequest) SetBuckets(v []V3CreateClusterUserRequestBuckets)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *V3CreateClusterUserRequest) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


