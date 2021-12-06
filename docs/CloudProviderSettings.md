# CloudProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupName** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**AwsRegions**](AwsRegions.md) |  | [optional] 
**AksClusterName** | Pointer to **string** |  | [optional] 
**DnsPrefix** | Pointer to **string** |  | [optional] 
**DefaultNodePoolName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**VNetCidr** | Pointer to **string** |  | [optional] 
**VNetName** | Pointer to **string** |  | [optional] 
**StorageAccountName** | Pointer to **string** |  | [optional] 
**ConfigContainerName** | Pointer to **string** |  | [optional] 
**ImportContainerName** | Pointer to **string** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**AzCount** | Pointer to **int32** |  | [optional] 
**DeploymentName** | Pointer to **string** |  | [optional] 
**MsiName** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**SupportBucket** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**VpcCidr** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**StackId** | Pointer to **string** |  | [optional] 
**PrivateSubnets** | Pointer to **[]string** |  | [optional] 
**AvailabilityZone** | Pointer to **[]string** |  | [optional] 
**LastRotated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCloudProviderSettings

`func NewCloudProviderSettings() *CloudProviderSettings`

NewCloudProviderSettings instantiates a new CloudProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderSettingsWithDefaults

`func NewCloudProviderSettingsWithDefaults() *CloudProviderSettings`

NewCloudProviderSettingsWithDefaults instantiates a new CloudProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupName

`func (o *CloudProviderSettings) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *CloudProviderSettings) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *CloudProviderSettings) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *CloudProviderSettings) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetRegion

`func (o *CloudProviderSettings) GetRegion() AwsRegions`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudProviderSettings) GetRegionOk() (*AwsRegions, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudProviderSettings) SetRegion(v AwsRegions)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudProviderSettings) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAksClusterName

`func (o *CloudProviderSettings) GetAksClusterName() string`

GetAksClusterName returns the AksClusterName field if non-nil, zero value otherwise.

### GetAksClusterNameOk

`func (o *CloudProviderSettings) GetAksClusterNameOk() (*string, bool)`

GetAksClusterNameOk returns a tuple with the AksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAksClusterName

`func (o *CloudProviderSettings) SetAksClusterName(v string)`

SetAksClusterName sets AksClusterName field to given value.

### HasAksClusterName

`func (o *CloudProviderSettings) HasAksClusterName() bool`

HasAksClusterName returns a boolean if a field has been set.

### GetDnsPrefix

`func (o *CloudProviderSettings) GetDnsPrefix() string`

GetDnsPrefix returns the DnsPrefix field if non-nil, zero value otherwise.

### GetDnsPrefixOk

`func (o *CloudProviderSettings) GetDnsPrefixOk() (*string, bool)`

GetDnsPrefixOk returns a tuple with the DnsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrefix

`func (o *CloudProviderSettings) SetDnsPrefix(v string)`

SetDnsPrefix sets DnsPrefix field to given value.

### HasDnsPrefix

`func (o *CloudProviderSettings) HasDnsPrefix() bool`

HasDnsPrefix returns a boolean if a field has been set.

### GetDefaultNodePoolName

`func (o *CloudProviderSettings) GetDefaultNodePoolName() string`

GetDefaultNodePoolName returns the DefaultNodePoolName field if non-nil, zero value otherwise.

### GetDefaultNodePoolNameOk

`func (o *CloudProviderSettings) GetDefaultNodePoolNameOk() (*string, bool)`

GetDefaultNodePoolNameOk returns a tuple with the DefaultNodePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNodePoolName

`func (o *CloudProviderSettings) SetDefaultNodePoolName(v string)`

SetDefaultNodePoolName sets DefaultNodePoolName field to given value.

### HasDefaultNodePoolName

`func (o *CloudProviderSettings) HasDefaultNodePoolName() bool`

HasDefaultNodePoolName returns a boolean if a field has been set.

### GetTags

`func (o *CloudProviderSettings) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudProviderSettings) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudProviderSettings) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudProviderSettings) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVNetCidr

`func (o *CloudProviderSettings) GetVNetCidr() string`

GetVNetCidr returns the VNetCidr field if non-nil, zero value otherwise.

### GetVNetCidrOk

`func (o *CloudProviderSettings) GetVNetCidrOk() (*string, bool)`

GetVNetCidrOk returns a tuple with the VNetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetCidr

`func (o *CloudProviderSettings) SetVNetCidr(v string)`

SetVNetCidr sets VNetCidr field to given value.

### HasVNetCidr

`func (o *CloudProviderSettings) HasVNetCidr() bool`

HasVNetCidr returns a boolean if a field has been set.

### GetVNetName

`func (o *CloudProviderSettings) GetVNetName() string`

GetVNetName returns the VNetName field if non-nil, zero value otherwise.

### GetVNetNameOk

`func (o *CloudProviderSettings) GetVNetNameOk() (*string, bool)`

GetVNetNameOk returns a tuple with the VNetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetName

`func (o *CloudProviderSettings) SetVNetName(v string)`

SetVNetName sets VNetName field to given value.

### HasVNetName

`func (o *CloudProviderSettings) HasVNetName() bool`

HasVNetName returns a boolean if a field has been set.

### GetStorageAccountName

`func (o *CloudProviderSettings) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *CloudProviderSettings) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *CloudProviderSettings) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *CloudProviderSettings) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### GetConfigContainerName

`func (o *CloudProviderSettings) GetConfigContainerName() string`

GetConfigContainerName returns the ConfigContainerName field if non-nil, zero value otherwise.

### GetConfigContainerNameOk

`func (o *CloudProviderSettings) GetConfigContainerNameOk() (*string, bool)`

GetConfigContainerNameOk returns a tuple with the ConfigContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContainerName

`func (o *CloudProviderSettings) SetConfigContainerName(v string)`

SetConfigContainerName sets ConfigContainerName field to given value.

### HasConfigContainerName

`func (o *CloudProviderSettings) HasConfigContainerName() bool`

HasConfigContainerName returns a boolean if a field has been set.

### GetImportContainerName

`func (o *CloudProviderSettings) GetImportContainerName() string`

GetImportContainerName returns the ImportContainerName field if non-nil, zero value otherwise.

### GetImportContainerNameOk

`func (o *CloudProviderSettings) GetImportContainerNameOk() (*string, bool)`

GetImportContainerNameOk returns a tuple with the ImportContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportContainerName

`func (o *CloudProviderSettings) SetImportContainerName(v string)`

SetImportContainerName sets ImportContainerName field to given value.

### HasImportContainerName

`func (o *CloudProviderSettings) HasImportContainerName() bool`

HasImportContainerName returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *CloudProviderSettings) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *CloudProviderSettings) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *CloudProviderSettings) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *CloudProviderSettings) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetAzCount

`func (o *CloudProviderSettings) GetAzCount() int32`

GetAzCount returns the AzCount field if non-nil, zero value otherwise.

### GetAzCountOk

`func (o *CloudProviderSettings) GetAzCountOk() (*int32, bool)`

GetAzCountOk returns a tuple with the AzCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCount

`func (o *CloudProviderSettings) SetAzCount(v int32)`

SetAzCount sets AzCount field to given value.

### HasAzCount

`func (o *CloudProviderSettings) HasAzCount() bool`

HasAzCount returns a boolean if a field has been set.

### GetDeploymentName

`func (o *CloudProviderSettings) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *CloudProviderSettings) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *CloudProviderSettings) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *CloudProviderSettings) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### GetMsiName

`func (o *CloudProviderSettings) GetMsiName() string`

GetMsiName returns the MsiName field if non-nil, zero value otherwise.

### GetMsiNameOk

`func (o *CloudProviderSettings) GetMsiNameOk() (*string, bool)`

GetMsiNameOk returns a tuple with the MsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsiName

`func (o *CloudProviderSettings) SetMsiName(v string)`

SetMsiName sets MsiName field to given value.

### HasMsiName

`func (o *CloudProviderSettings) HasMsiName() bool`

HasMsiName returns a boolean if a field has been set.

### GetTenantId

`func (o *CloudProviderSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CloudProviderSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CloudProviderSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CloudProviderSettings) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CloudProviderSettings) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CloudProviderSettings) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CloudProviderSettings) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CloudProviderSettings) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetBucket

`func (o *CloudProviderSettings) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CloudProviderSettings) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CloudProviderSettings) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CloudProviderSettings) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetSupportBucket

`func (o *CloudProviderSettings) GetSupportBucket() string`

GetSupportBucket returns the SupportBucket field if non-nil, zero value otherwise.

### GetSupportBucketOk

`func (o *CloudProviderSettings) GetSupportBucketOk() (*string, bool)`

GetSupportBucketOk returns a tuple with the SupportBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBucket

`func (o *CloudProviderSettings) SetSupportBucket(v string)`

SetSupportBucket sets SupportBucket field to given value.

### HasSupportBucket

`func (o *CloudProviderSettings) HasSupportBucket() bool`

HasSupportBucket returns a boolean if a field has been set.

### GetAccountId

`func (o *CloudProviderSettings) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CloudProviderSettings) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CloudProviderSettings) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CloudProviderSettings) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetVpcCidr

`func (o *CloudProviderSettings) GetVpcCidr() string`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *CloudProviderSettings) GetVpcCidrOk() (*string, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *CloudProviderSettings) SetVpcCidr(v string)`

SetVpcCidr sets VpcCidr field to given value.

### HasVpcCidr

`func (o *CloudProviderSettings) HasVpcCidr() bool`

HasVpcCidr returns a boolean if a field has been set.

### GetVpcId

`func (o *CloudProviderSettings) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CloudProviderSettings) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CloudProviderSettings) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *CloudProviderSettings) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetStackId

`func (o *CloudProviderSettings) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CloudProviderSettings) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CloudProviderSettings) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CloudProviderSettings) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *CloudProviderSettings) GetPrivateSubnets() []string`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *CloudProviderSettings) GetPrivateSubnetsOk() (*[]string, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *CloudProviderSettings) SetPrivateSubnets(v []string)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *CloudProviderSettings) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CloudProviderSettings) GetAvailabilityZone() []string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CloudProviderSettings) GetAvailabilityZoneOk() (*[]string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CloudProviderSettings) SetAvailabilityZone(v []string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CloudProviderSettings) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetLastRotated

`func (o *CloudProviderSettings) GetLastRotated() time.Time`

GetLastRotated returns the LastRotated field if non-nil, zero value otherwise.

### GetLastRotatedOk

`func (o *CloudProviderSettings) GetLastRotatedOk() (*time.Time, bool)`

GetLastRotatedOk returns a tuple with the LastRotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotated

`func (o *CloudProviderSettings) SetLastRotated(v time.Time)`

SetLastRotated sets LastRotated field to given value.

### HasLastRotated

`func (o *CloudProviderSettings) HasLastRotated() bool`

HasLastRotated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


