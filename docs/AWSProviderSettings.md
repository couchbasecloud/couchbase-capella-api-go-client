# AWSProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to [**AwsRegions**](AwsRegions.md) |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**SupportBucket** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**VpcCidr** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**StackId** | Pointer to **string** |  | [optional] 
**PrivateSubnets** | Pointer to **[]string** |  | [optional] 
**AvailabilityZone** | Pointer to **[]string** |  | [optional] 
**AzCount** | Pointer to **int32** |  | [optional] 
**LastRotated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAWSProviderSettings

`func NewAWSProviderSettings() *AWSProviderSettings`

NewAWSProviderSettings instantiates a new AWSProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSProviderSettingsWithDefaults

`func NewAWSProviderSettingsWithDefaults() *AWSProviderSettings`

NewAWSProviderSettingsWithDefaults instantiates a new AWSProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AWSProviderSettings) GetRegion() AwsRegions`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSProviderSettings) GetRegionOk() (*AwsRegions, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSProviderSettings) SetRegion(v AwsRegions)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AWSProviderSettings) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetBucket

`func (o *AWSProviderSettings) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AWSProviderSettings) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AWSProviderSettings) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *AWSProviderSettings) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetSupportBucket

`func (o *AWSProviderSettings) GetSupportBucket() string`

GetSupportBucket returns the SupportBucket field if non-nil, zero value otherwise.

### GetSupportBucketOk

`func (o *AWSProviderSettings) GetSupportBucketOk() (*string, bool)`

GetSupportBucketOk returns a tuple with the SupportBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBucket

`func (o *AWSProviderSettings) SetSupportBucket(v string)`

SetSupportBucket sets SupportBucket field to given value.

### HasSupportBucket

`func (o *AWSProviderSettings) HasSupportBucket() bool`

HasSupportBucket returns a boolean if a field has been set.

### GetAccountId

`func (o *AWSProviderSettings) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSProviderSettings) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSProviderSettings) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AWSProviderSettings) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetVpcCidr

`func (o *AWSProviderSettings) GetVpcCidr() string`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *AWSProviderSettings) GetVpcCidrOk() (*string, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *AWSProviderSettings) SetVpcCidr(v string)`

SetVpcCidr sets VpcCidr field to given value.

### HasVpcCidr

`func (o *AWSProviderSettings) HasVpcCidr() bool`

HasVpcCidr returns a boolean if a field has been set.

### GetVpcId

`func (o *AWSProviderSettings) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSProviderSettings) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSProviderSettings) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AWSProviderSettings) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetStackId

`func (o *AWSProviderSettings) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *AWSProviderSettings) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *AWSProviderSettings) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *AWSProviderSettings) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *AWSProviderSettings) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *AWSProviderSettings) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *AWSProviderSettings) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *AWSProviderSettings) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AWSProviderSettings) GetAvailabilityZone() []string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AWSProviderSettings) GetAvailabilityZoneOk() (*[]string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AWSProviderSettings) SetAvailabilityZone(v []string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AWSProviderSettings) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzCount

`func (o *AWSProviderSettings) GetAzCount() int32`

GetAzCount returns the AzCount field if non-nil, zero value otherwise.

### GetAzCountOk

`func (o *AWSProviderSettings) GetAzCountOk() (*int32, bool)`

GetAzCountOk returns a tuple with the AzCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCount

`func (o *AWSProviderSettings) SetAzCount(v int32)`

SetAzCount sets AzCount field to given value.

### HasAzCount

`func (o *AWSProviderSettings) HasAzCount() bool`

HasAzCount returns a boolean if a field has been set.

### GetLastRotated

`func (o *AWSProviderSettings) GetLastRotated() time.Time`

GetLastRotated returns the LastRotated field if non-nil, zero value otherwise.

### GetLastRotatedOk

`func (o *AWSProviderSettings) GetLastRotatedOk() (*time.Time, bool)`

GetLastRotatedOk returns a tuple with the LastRotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotated

`func (o *AWSProviderSettings) SetLastRotated(v time.Time)`

SetLastRotated sets LastRotated field to given value.

### HasLastRotated

`func (o *AWSProviderSettings) HasLastRotated() bool`

HasLastRotated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


