# V3CreateClusterUserRequestBuckets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Scope** | **string** | enter &#39;*&#39; for getting access to all scopes within a bucket or the scope name for getting access to a single scope | 
**Access** | [**V3BucketRoles**](V3BucketRoles.md) |  | 

## Methods

### NewV3CreateClusterUserRequestBuckets

`func NewV3CreateClusterUserRequestBuckets(name string, scope string, access V3BucketRoles, ) *V3CreateClusterUserRequestBuckets`

NewV3CreateClusterUserRequestBuckets instantiates a new V3CreateClusterUserRequestBuckets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3CreateClusterUserRequestBucketsWithDefaults

`func NewV3CreateClusterUserRequestBucketsWithDefaults() *V3CreateClusterUserRequestBuckets`

NewV3CreateClusterUserRequestBucketsWithDefaults instantiates a new V3CreateClusterUserRequestBuckets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V3CreateClusterUserRequestBuckets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3CreateClusterUserRequestBuckets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3CreateClusterUserRequestBuckets) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *V3CreateClusterUserRequestBuckets) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *V3CreateClusterUserRequestBuckets) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *V3CreateClusterUserRequestBuckets) SetScope(v string)`

SetScope sets Scope field to given value.


### GetAccess

`func (o *V3CreateClusterUserRequestBuckets) GetAccess() V3BucketRoles`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *V3CreateClusterUserRequestBuckets) GetAccessOk() (*V3BucketRoles, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *V3CreateClusterUserRequestBuckets) SetAccess(v V3BucketRoles)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


