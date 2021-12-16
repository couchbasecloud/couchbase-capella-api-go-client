# V3ClusterVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Components** | [**V3ClusterVersionComponents**](V3ClusterVersionComponents.md) |  | 

## Methods

### NewV3ClusterVersion

`func NewV3ClusterVersion(name string, components V3ClusterVersionComponents, ) *V3ClusterVersion`

NewV3ClusterVersion instantiates a new V3ClusterVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ClusterVersionWithDefaults

`func NewV3ClusterVersionWithDefaults() *V3ClusterVersion`

NewV3ClusterVersionWithDefaults instantiates a new V3ClusterVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V3ClusterVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3ClusterVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3ClusterVersion) SetName(v string)`

SetName sets Name field to given value.


### GetComponents

`func (o *V3ClusterVersion) GetComponents() V3ClusterVersionComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *V3ClusterVersion) GetComponentsOk() (*V3ClusterVersionComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *V3ClusterVersion) SetComponents(v V3ClusterVersionComponents)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


