# ServiceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | **[]string** |  | 
**NodeName** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewServiceStats

`func NewServiceStats(services []string, nodeName string, status string, ) *ServiceStats`

NewServiceStats instantiates a new ServiceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatsWithDefaults

`func NewServiceStatsWithDefaults() *ServiceStats`

NewServiceStatsWithDefaults instantiates a new ServiceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *ServiceStats) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServiceStats) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServiceStats) SetServices(v []string)`

SetServices sets Services field to given value.


### GetNodeName

`func (o *ServiceStats) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ServiceStats) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ServiceStats) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetStatus

`func (o *ServiceStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceStats) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


