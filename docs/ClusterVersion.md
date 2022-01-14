# ClusterVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Components** | **map[string]string** |  | 

## Methods

### NewClusterVersion

`func NewClusterVersion(name string, components map[string]string, ) *ClusterVersion`

NewClusterVersion instantiates a new ClusterVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVersionWithDefaults

`func NewClusterVersionWithDefaults() *ClusterVersion`

NewClusterVersionWithDefaults instantiates a new ClusterVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterVersion) SetName(v string)`

SetName sets Name field to given value.


### GetComponents

`func (o *ClusterVersion) GetComponents() map[string]string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ClusterVersion) GetComponentsOk() (*map[string]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ClusterVersion) SetComponents(v map[string]string)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


