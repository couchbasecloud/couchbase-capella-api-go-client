# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **int32** |  | 
**Services** | [**[]CouchbaseServices**](CouchbaseServices.md) |  | 
**Aws** | Pointer to [**ServerAws**](ServerAws.md) |  | [optional] 
**Azure** | Pointer to [**ServerAzure**](ServerAzure.md) |  | [optional] 

## Methods

### NewServer

`func NewServer(size int32, services []CouchbaseServices, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *Server) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Server) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Server) SetSize(v int32)`

SetSize sets Size field to given value.


### GetServices

`func (o *Server) GetServices() []CouchbaseServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Server) GetServicesOk() (*[]CouchbaseServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Server) SetServices(v []CouchbaseServices)`

SetServices sets Services field to given value.


### GetAws

`func (o *Server) GetAws() ServerAws`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *Server) GetAwsOk() (*ServerAws, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *Server) SetAws(v ServerAws)`

SetAws sets Aws field to given value.

### HasAws

`func (o *Server) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *Server) GetAzure() ServerAzure`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *Server) GetAzureOk() (*ServerAzure, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *Server) SetAzure(v ServerAzure)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *Server) HasAzure() bool`

HasAzure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


