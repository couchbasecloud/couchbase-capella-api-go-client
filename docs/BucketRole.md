# BucketRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Bucket Name | 
**Roles** | [**[]BucketRoleTypes**](BucketRoleTypes.md) |  | 

## Methods

### NewBucketRole

`func NewBucketRole(name string, roles []BucketRoleTypes, ) *BucketRole`

NewBucketRole instantiates a new BucketRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketRoleWithDefaults

`func NewBucketRoleWithDefaults() *BucketRole`

NewBucketRoleWithDefaults instantiates a new BucketRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BucketRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BucketRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BucketRole) SetName(v string)`

SetName sets Name field to given value.


### GetRoles

`func (o *BucketRole) GetRoles() []BucketRoleTypes`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BucketRole) GetRolesOk() (*[]BucketRoleTypes, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BucketRole) SetRoles(v []BucketRoleTypes)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


