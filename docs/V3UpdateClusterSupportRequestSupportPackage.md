# V3UpdateClusterSupportRequestSupportPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to [**V3SupportPackageTimezones**](V3SupportPackageTimezones.md) |  | [optional] 
**Type** | [**V3SupportPackageType**](V3SupportPackageType.md) |  | 

## Methods

### NewV3UpdateClusterSupportRequestSupportPackage

`func NewV3UpdateClusterSupportRequestSupportPackage(type_ V3SupportPackageType, ) *V3UpdateClusterSupportRequestSupportPackage`

NewV3UpdateClusterSupportRequestSupportPackage instantiates a new V3UpdateClusterSupportRequestSupportPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3UpdateClusterSupportRequestSupportPackageWithDefaults

`func NewV3UpdateClusterSupportRequestSupportPackageWithDefaults() *V3UpdateClusterSupportRequestSupportPackage`

NewV3UpdateClusterSupportRequestSupportPackageWithDefaults instantiates a new V3UpdateClusterSupportRequestSupportPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *V3UpdateClusterSupportRequestSupportPackage) GetTimezone() V3SupportPackageTimezones`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V3UpdateClusterSupportRequestSupportPackage) GetTimezoneOk() (*V3SupportPackageTimezones, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V3UpdateClusterSupportRequestSupportPackage) SetTimezone(v V3SupportPackageTimezones)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V3UpdateClusterSupportRequestSupportPackage) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetType

`func (o *V3UpdateClusterSupportRequestSupportPackage) GetType() V3SupportPackageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V3UpdateClusterSupportRequestSupportPackage) GetTypeOk() (*V3SupportPackageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V3UpdateClusterSupportRequestSupportPackage) SetType(v V3SupportPackageType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


