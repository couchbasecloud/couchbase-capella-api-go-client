# AzureProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupName** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**AzureRegions**](AzureRegions.md) |  | [optional] 
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

## Methods

### NewAzureProviderSettings

`func NewAzureProviderSettings() *AzureProviderSettings`

NewAzureProviderSettings instantiates a new AzureProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProviderSettingsWithDefaults

`func NewAzureProviderSettingsWithDefaults() *AzureProviderSettings`

NewAzureProviderSettingsWithDefaults instantiates a new AzureProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupName

`func (o *AzureProviderSettings) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureProviderSettings) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureProviderSettings) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureProviderSettings) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetRegion

`func (o *AzureProviderSettings) GetRegion() AzureRegions`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureProviderSettings) GetRegionOk() (*AzureRegions, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureProviderSettings) SetRegion(v AzureRegions)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AzureProviderSettings) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAksClusterName

`func (o *AzureProviderSettings) GetAksClusterName() string`

GetAksClusterName returns the AksClusterName field if non-nil, zero value otherwise.

### GetAksClusterNameOk

`func (o *AzureProviderSettings) GetAksClusterNameOk() (*string, bool)`

GetAksClusterNameOk returns a tuple with the AksClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAksClusterName

`func (o *AzureProviderSettings) SetAksClusterName(v string)`

SetAksClusterName sets AksClusterName field to given value.

### HasAksClusterName

`func (o *AzureProviderSettings) HasAksClusterName() bool`

HasAksClusterName returns a boolean if a field has been set.

### GetDnsPrefix

`func (o *AzureProviderSettings) GetDnsPrefix() string`

GetDnsPrefix returns the DnsPrefix field if non-nil, zero value otherwise.

### GetDnsPrefixOk

`func (o *AzureProviderSettings) GetDnsPrefixOk() (*string, bool)`

GetDnsPrefixOk returns a tuple with the DnsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrefix

`func (o *AzureProviderSettings) SetDnsPrefix(v string)`

SetDnsPrefix sets DnsPrefix field to given value.

### HasDnsPrefix

`func (o *AzureProviderSettings) HasDnsPrefix() bool`

HasDnsPrefix returns a boolean if a field has been set.

### GetDefaultNodePoolName

`func (o *AzureProviderSettings) GetDefaultNodePoolName() string`

GetDefaultNodePoolName returns the DefaultNodePoolName field if non-nil, zero value otherwise.

### GetDefaultNodePoolNameOk

`func (o *AzureProviderSettings) GetDefaultNodePoolNameOk() (*string, bool)`

GetDefaultNodePoolNameOk returns a tuple with the DefaultNodePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNodePoolName

`func (o *AzureProviderSettings) SetDefaultNodePoolName(v string)`

SetDefaultNodePoolName sets DefaultNodePoolName field to given value.

### HasDefaultNodePoolName

`func (o *AzureProviderSettings) HasDefaultNodePoolName() bool`

HasDefaultNodePoolName returns a boolean if a field has been set.

### GetTags

`func (o *AzureProviderSettings) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AzureProviderSettings) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AzureProviderSettings) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AzureProviderSettings) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVNetCidr

`func (o *AzureProviderSettings) GetVNetCidr() string`

GetVNetCidr returns the VNetCidr field if non-nil, zero value otherwise.

### GetVNetCidrOk

`func (o *AzureProviderSettings) GetVNetCidrOk() (*string, bool)`

GetVNetCidrOk returns a tuple with the VNetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetCidr

`func (o *AzureProviderSettings) SetVNetCidr(v string)`

SetVNetCidr sets VNetCidr field to given value.

### HasVNetCidr

`func (o *AzureProviderSettings) HasVNetCidr() bool`

HasVNetCidr returns a boolean if a field has been set.

### GetVNetName

`func (o *AzureProviderSettings) GetVNetName() string`

GetVNetName returns the VNetName field if non-nil, zero value otherwise.

### GetVNetNameOk

`func (o *AzureProviderSettings) GetVNetNameOk() (*string, bool)`

GetVNetNameOk returns a tuple with the VNetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNetName

`func (o *AzureProviderSettings) SetVNetName(v string)`

SetVNetName sets VNetName field to given value.

### HasVNetName

`func (o *AzureProviderSettings) HasVNetName() bool`

HasVNetName returns a boolean if a field has been set.

### GetStorageAccountName

`func (o *AzureProviderSettings) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *AzureProviderSettings) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *AzureProviderSettings) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *AzureProviderSettings) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### GetConfigContainerName

`func (o *AzureProviderSettings) GetConfigContainerName() string`

GetConfigContainerName returns the ConfigContainerName field if non-nil, zero value otherwise.

### GetConfigContainerNameOk

`func (o *AzureProviderSettings) GetConfigContainerNameOk() (*string, bool)`

GetConfigContainerNameOk returns a tuple with the ConfigContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContainerName

`func (o *AzureProviderSettings) SetConfigContainerName(v string)`

SetConfigContainerName sets ConfigContainerName field to given value.

### HasConfigContainerName

`func (o *AzureProviderSettings) HasConfigContainerName() bool`

HasConfigContainerName returns a boolean if a field has been set.

### GetImportContainerName

`func (o *AzureProviderSettings) GetImportContainerName() string`

GetImportContainerName returns the ImportContainerName field if non-nil, zero value otherwise.

### GetImportContainerNameOk

`func (o *AzureProviderSettings) GetImportContainerNameOk() (*string, bool)`

GetImportContainerNameOk returns a tuple with the ImportContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportContainerName

`func (o *AzureProviderSettings) SetImportContainerName(v string)`

SetImportContainerName sets ImportContainerName field to given value.

### HasImportContainerName

`func (o *AzureProviderSettings) HasImportContainerName() bool`

HasImportContainerName returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *AzureProviderSettings) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *AzureProviderSettings) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *AzureProviderSettings) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *AzureProviderSettings) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetAzCount

`func (o *AzureProviderSettings) GetAzCount() int32`

GetAzCount returns the AzCount field if non-nil, zero value otherwise.

### GetAzCountOk

`func (o *AzureProviderSettings) GetAzCountOk() (*int32, bool)`

GetAzCountOk returns a tuple with the AzCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCount

`func (o *AzureProviderSettings) SetAzCount(v int32)`

SetAzCount sets AzCount field to given value.

### HasAzCount

`func (o *AzureProviderSettings) HasAzCount() bool`

HasAzCount returns a boolean if a field has been set.

### GetDeploymentName

`func (o *AzureProviderSettings) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *AzureProviderSettings) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *AzureProviderSettings) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *AzureProviderSettings) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### GetMsiName

`func (o *AzureProviderSettings) GetMsiName() string`

GetMsiName returns the MsiName field if non-nil, zero value otherwise.

### GetMsiNameOk

`func (o *AzureProviderSettings) GetMsiNameOk() (*string, bool)`

GetMsiNameOk returns a tuple with the MsiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsiName

`func (o *AzureProviderSettings) SetMsiName(v string)`

SetMsiName sets MsiName field to given value.

### HasMsiName

`func (o *AzureProviderSettings) HasMsiName() bool`

HasMsiName returns a boolean if a field has been set.

### GetTenantId

`func (o *AzureProviderSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureProviderSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureProviderSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureProviderSettings) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureProviderSettings) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureProviderSettings) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureProviderSettings) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureProviderSettings) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


