# BucketRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Bucket Name | 
**BucketAccess** | [**[]BucketRoleTypes**](BucketRoleTypes.md) |  | 

## Methods

### NewBucketRole

`func NewBucketRole(bucketName string, bucketAccess []BucketRoleTypes, ) *BucketRole`

NewBucketRole instantiates a new BucketRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketRoleWithDefaults

`func NewBucketRoleWithDefaults() *BucketRole`

NewBucketRoleWithDefaults instantiates a new BucketRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *BucketRole) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *BucketRole) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *BucketRole) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetBucketAccess

`func (o *BucketRole) GetBucketAccess() []BucketRoleTypes`

GetBucketAccess returns the BucketAccess field if non-nil, zero value otherwise.

### GetBucketAccessOk

`func (o *BucketRole) GetBucketAccessOk() (*[]BucketRoleTypes, bool)`

GetBucketAccessOk returns a tuple with the BucketAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketAccess

`func (o *BucketRole) SetBucketAccess(v []BucketRoleTypes)`

SetBucketAccess sets BucketAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


