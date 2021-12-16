# V3PlaceHosted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**V3Provider**](V3Provider.md) |  | 
**Region** | **string** | A valid region for the cloudProvider | 
**CIDR** | **string** | CIDR block. | 

## Methods

### NewV3PlaceHosted

`func NewV3PlaceHosted(provider V3Provider, region string, cIDR string, ) *V3PlaceHosted`

NewV3PlaceHosted instantiates a new V3PlaceHosted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3PlaceHostedWithDefaults

`func NewV3PlaceHostedWithDefaults() *V3PlaceHosted`

NewV3PlaceHostedWithDefaults instantiates a new V3PlaceHosted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *V3PlaceHosted) GetProvider() V3Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V3PlaceHosted) GetProviderOk() (*V3Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V3PlaceHosted) SetProvider(v V3Provider)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *V3PlaceHosted) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V3PlaceHosted) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V3PlaceHosted) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetCIDR

`func (o *V3PlaceHosted) GetCIDR() string`

GetCIDR returns the CIDR field if non-nil, zero value otherwise.

### GetCIDROk

`func (o *V3PlaceHosted) GetCIDROk() (*string, bool)`

GetCIDROk returns a tuple with the CIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIDR

`func (o *V3PlaceHosted) SetCIDR(v string)`

SetCIDR sets CIDR field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


