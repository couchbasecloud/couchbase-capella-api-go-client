# Cloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TenantId** | **string** |  | 
**Status** | [**CloudStatus**](CloudStatus.md) |  | 
**Provider** | [**Provider**](Provider.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**PreflightStartedAt** | Pointer to **time.Time** |  | [optional] 
**PreflightFailedAt** | Pointer to **time.Time** |  | [optional] 
**PreflightCompletedAt** | Pointer to **time.Time** |  | [optional] 
**BootstrappedAt** | Pointer to **time.Time** |  | [optional] 
**DeployedAt** | Pointer to **time.Time** |  | [optional] 
**DestroyedAt** | Pointer to **time.Time** |  | [optional] 
**ProviderSettings** | [**CloudProviderSettings**](CloudProviderSettings.md) |  | 
**Version** | [**CloudVersion**](CloudVersion.md) |  | 

## Methods

### NewCloud

`func NewCloud(id string, name string, tenantId string, status CloudStatus, provider Provider, createdAt time.Time, providerSettings CloudProviderSettings, version CloudVersion, ) *Cloud`

NewCloud instantiates a new Cloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWithDefaults

`func NewCloudWithDefaults() *Cloud`

NewCloudWithDefaults instantiates a new Cloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cloud) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cloud) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cloud) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Cloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cloud) SetName(v string)`

SetName sets Name field to given value.


### GetTenantId

`func (o *Cloud) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Cloud) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Cloud) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetStatus

`func (o *Cloud) GetStatus() CloudStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cloud) GetStatusOk() (*CloudStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cloud) SetStatus(v CloudStatus)`

SetStatus sets Status field to given value.


### GetProvider

`func (o *Cloud) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Cloud) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Cloud) SetProvider(v Provider)`

SetProvider sets Provider field to given value.


### GetCreatedAt

`func (o *Cloud) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cloud) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cloud) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cloud) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cloud) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cloud) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Cloud) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPreflightStartedAt

`func (o *Cloud) GetPreflightStartedAt() time.Time`

GetPreflightStartedAt returns the PreflightStartedAt field if non-nil, zero value otherwise.

### GetPreflightStartedAtOk

`func (o *Cloud) GetPreflightStartedAtOk() (*time.Time, bool)`

GetPreflightStartedAtOk returns a tuple with the PreflightStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreflightStartedAt

`func (o *Cloud) SetPreflightStartedAt(v time.Time)`

SetPreflightStartedAt sets PreflightStartedAt field to given value.

### HasPreflightStartedAt

`func (o *Cloud) HasPreflightStartedAt() bool`

HasPreflightStartedAt returns a boolean if a field has been set.

### GetPreflightFailedAt

`func (o *Cloud) GetPreflightFailedAt() time.Time`

GetPreflightFailedAt returns the PreflightFailedAt field if non-nil, zero value otherwise.

### GetPreflightFailedAtOk

`func (o *Cloud) GetPreflightFailedAtOk() (*time.Time, bool)`

GetPreflightFailedAtOk returns a tuple with the PreflightFailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreflightFailedAt

`func (o *Cloud) SetPreflightFailedAt(v time.Time)`

SetPreflightFailedAt sets PreflightFailedAt field to given value.

### HasPreflightFailedAt

`func (o *Cloud) HasPreflightFailedAt() bool`

HasPreflightFailedAt returns a boolean if a field has been set.

### GetPreflightCompletedAt

`func (o *Cloud) GetPreflightCompletedAt() time.Time`

GetPreflightCompletedAt returns the PreflightCompletedAt field if non-nil, zero value otherwise.

### GetPreflightCompletedAtOk

`func (o *Cloud) GetPreflightCompletedAtOk() (*time.Time, bool)`

GetPreflightCompletedAtOk returns a tuple with the PreflightCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreflightCompletedAt

`func (o *Cloud) SetPreflightCompletedAt(v time.Time)`

SetPreflightCompletedAt sets PreflightCompletedAt field to given value.

### HasPreflightCompletedAt

`func (o *Cloud) HasPreflightCompletedAt() bool`

HasPreflightCompletedAt returns a boolean if a field has been set.

### GetBootstrappedAt

`func (o *Cloud) GetBootstrappedAt() time.Time`

GetBootstrappedAt returns the BootstrappedAt field if non-nil, zero value otherwise.

### GetBootstrappedAtOk

`func (o *Cloud) GetBootstrappedAtOk() (*time.Time, bool)`

GetBootstrappedAtOk returns a tuple with the BootstrappedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrappedAt

`func (o *Cloud) SetBootstrappedAt(v time.Time)`

SetBootstrappedAt sets BootstrappedAt field to given value.

### HasBootstrappedAt

`func (o *Cloud) HasBootstrappedAt() bool`

HasBootstrappedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *Cloud) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *Cloud) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *Cloud) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *Cloud) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetDestroyedAt

`func (o *Cloud) GetDestroyedAt() time.Time`

GetDestroyedAt returns the DestroyedAt field if non-nil, zero value otherwise.

### GetDestroyedAtOk

`func (o *Cloud) GetDestroyedAtOk() (*time.Time, bool)`

GetDestroyedAtOk returns a tuple with the DestroyedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedAt

`func (o *Cloud) SetDestroyedAt(v time.Time)`

SetDestroyedAt sets DestroyedAt field to given value.

### HasDestroyedAt

`func (o *Cloud) HasDestroyedAt() bool`

HasDestroyedAt returns a boolean if a field has been set.

### GetProviderSettings

`func (o *Cloud) GetProviderSettings() CloudProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *Cloud) GetProviderSettingsOk() (*CloudProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *Cloud) SetProviderSettings(v CloudProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.


### GetVersion

`func (o *Cloud) GetVersion() CloudVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cloud) GetVersionOk() (*CloudVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cloud) SetVersion(v CloudVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


