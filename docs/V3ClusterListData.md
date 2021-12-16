# V3ClusterListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]V3ClusterListDataItems**](V3ClusterListDataItems.md) |  | [optional] 

## Methods

### NewV3ClusterListData

`func NewV3ClusterListData() *V3ClusterListData`

NewV3ClusterListData instantiates a new V3ClusterListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ClusterListDataWithDefaults

`func NewV3ClusterListDataWithDefaults() *V3ClusterListData`

NewV3ClusterListDataWithDefaults instantiates a new V3ClusterListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *V3ClusterListData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *V3ClusterListData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *V3ClusterListData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *V3ClusterListData) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetItems

`func (o *V3ClusterListData) GetItems() []V3ClusterListDataItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V3ClusterListData) GetItemsOk() (*[]V3ClusterListDataItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V3ClusterListData) SetItems(v []V3ClusterListDataItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *V3ClusterListData) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


