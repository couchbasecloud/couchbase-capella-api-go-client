# V3Place

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleAZ** | **bool** | Has to be true if the supportPackage.type is equal to \&quot;Basic\&quot; | 
**Hosted** | Pointer to [**V3PlaceHosted**](V3PlaceHosted.md) |  | [optional] 

## Methods

### NewV3Place

`func NewV3Place(singleAZ bool, ) *V3Place`

NewV3Place instantiates a new V3Place object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3PlaceWithDefaults

`func NewV3PlaceWithDefaults() *V3Place`

NewV3PlaceWithDefaults instantiates a new V3Place object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleAZ

`func (o *V3Place) GetSingleAZ() bool`

GetSingleAZ returns the SingleAZ field if non-nil, zero value otherwise.

### GetSingleAZOk

`func (o *V3Place) GetSingleAZOk() (*bool, bool)`

GetSingleAZOk returns a tuple with the SingleAZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAZ

`func (o *V3Place) SetSingleAZ(v bool)`

SetSingleAZ sets SingleAZ field to given value.


### GetHosted

`func (o *V3Place) GetHosted() V3PlaceHosted`

GetHosted returns the Hosted field if non-nil, zero value otherwise.

### GetHostedOk

`func (o *V3Place) GetHostedOk() (*V3PlaceHosted, bool)`

GetHostedOk returns a tuple with the Hosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosted

`func (o *V3Place) SetHosted(v V3PlaceHosted)`

SetHosted sets Hosted field to given value.

### HasHosted

`func (o *V3Place) HasHosted() bool`

HasHosted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


