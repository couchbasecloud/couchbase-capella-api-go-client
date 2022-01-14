# CloudSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Status** | [**CloudStatus**](CloudStatus.md) |  | 
**Provider** | [**Provider**](Provider.md) |  | 
**Region** | [**Regions**](Regions.md) |  | 
**VirtualNetworkID** | **string** |  | 
**VirtualNetworkCIDR** | **string** |  | 

## Methods

### NewCloudSummary

`func NewCloudSummary(id string, name string, status CloudStatus, provider Provider, region Regions, virtualNetworkID string, virtualNetworkCIDR string, ) *CloudSummary`

NewCloudSummary instantiates a new CloudSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSummaryWithDefaults

`func NewCloudSummaryWithDefaults() *CloudSummary`

NewCloudSummaryWithDefaults instantiates a new CloudSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CloudSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSummary) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *CloudSummary) GetStatus() CloudStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSummary) GetStatusOk() (*CloudStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSummary) SetStatus(v CloudStatus)`

SetStatus sets Status field to given value.


### GetProvider

`func (o *CloudSummary) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudSummary) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudSummary) SetProvider(v Provider)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *CloudSummary) GetRegion() Regions`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudSummary) GetRegionOk() (*Regions, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudSummary) SetRegion(v Regions)`

SetRegion sets Region field to given value.


### GetVirtualNetworkID

`func (o *CloudSummary) GetVirtualNetworkID() string`

GetVirtualNetworkID returns the VirtualNetworkID field if non-nil, zero value otherwise.

### GetVirtualNetworkIDOk

`func (o *CloudSummary) GetVirtualNetworkIDOk() (*string, bool)`

GetVirtualNetworkIDOk returns a tuple with the VirtualNetworkID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkID

`func (o *CloudSummary) SetVirtualNetworkID(v string)`

SetVirtualNetworkID sets VirtualNetworkID field to given value.


### GetVirtualNetworkCIDR

`func (o *CloudSummary) GetVirtualNetworkCIDR() string`

GetVirtualNetworkCIDR returns the VirtualNetworkCIDR field if non-nil, zero value otherwise.

### GetVirtualNetworkCIDROk

`func (o *CloudSummary) GetVirtualNetworkCIDROk() (*string, bool)`

GetVirtualNetworkCIDROk returns a tuple with the VirtualNetworkCIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkCIDR

`func (o *CloudSummary) SetVirtualNetworkCIDR(v string)`

SetVirtualNetworkCIDR sets VirtualNetworkCIDR field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


