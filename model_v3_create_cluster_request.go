/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecapella

import (
	"encoding/json"
)

// V3CreateClusterRequest struct for V3CreateClusterRequest
type V3CreateClusterRequest struct {
	Environment V3Environment `json:"environment"`
	// Name of the cluster
	ClusterName string `json:"clusterName"`
	// Uniquely identifiable projectId
	ProjectId string `json:"projectId"`
	// A short description that describes the cluster
	Description *string `json:"description,omitempty"`
	Place V3Place `json:"place"`
	// Server specifications for creating the cluster
	Servers []V3Servers `json:"servers"`
	SupportPackage V3SupportPackage `json:"supportPackage"`
}

// NewV3CreateClusterRequest instantiates a new V3CreateClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3CreateClusterRequest(environment V3Environment, clusterName string, projectId string, place V3Place, servers []V3Servers, supportPackage V3SupportPackage) *V3CreateClusterRequest {
	this := V3CreateClusterRequest{}
	this.Environment = environment
	this.ClusterName = clusterName
	this.ProjectId = projectId
	this.Place = place
	this.Servers = servers
	this.SupportPackage = supportPackage
	return &this
}

// NewV3CreateClusterRequestWithDefaults instantiates a new V3CreateClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3CreateClusterRequestWithDefaults() *V3CreateClusterRequest {
	this := V3CreateClusterRequest{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *V3CreateClusterRequest) GetEnvironment() V3Environment {
	if o == nil {
		var ret V3Environment
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterRequest) GetEnvironmentOk() (*V3Environment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *V3CreateClusterRequest) SetEnvironment(v V3Environment) {
	o.Environment = v
}

// GetClusterName returns the ClusterName field value
func (o *V3CreateClusterRequest) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterRequest) GetClusterNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *V3CreateClusterRequest) SetClusterName(v string) {
	o.ClusterName = v
}

// GetProjectId returns the ProjectId field value
func (o *V3CreateClusterRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterRequest) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *V3CreateClusterRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V3CreateClusterRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3CreateClusterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V3CreateClusterRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V3CreateClusterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPlace returns the Place field value
func (o *V3CreateClusterRequest) GetPlace() V3Place {
	if o == nil {
		var ret V3Place
		return ret
	}

	return o.Place
}

// GetPlaceOk returns a tuple with the Place field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterRequest) GetPlaceOk() (*V3Place, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Place, true
}

// SetPlace sets field value
func (o *V3CreateClusterRequest) SetPlace(v V3Place) {
	o.Place = v
}

// GetServers returns the Servers field value
func (o *V3CreateClusterRequest) GetServers() []V3Servers {
	if o == nil {
		var ret []V3Servers
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterRequest) GetServersOk() (*[]V3Servers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Servers, true
}

// SetServers sets field value
func (o *V3CreateClusterRequest) SetServers(v []V3Servers) {
	o.Servers = v
}

// GetSupportPackage returns the SupportPackage field value
func (o *V3CreateClusterRequest) GetSupportPackage() V3SupportPackage {
	if o == nil {
		var ret V3SupportPackage
		return ret
	}

	return o.SupportPackage
}

// GetSupportPackageOk returns a tuple with the SupportPackage field value
// and a boolean to check if the value has been set.
func (o *V3CreateClusterRequest) GetSupportPackageOk() (*V3SupportPackage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportPackage, true
}

// SetSupportPackage sets field value
func (o *V3CreateClusterRequest) SetSupportPackage(v V3SupportPackage) {
	o.SupportPackage = v
}

func (o V3CreateClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["clusterName"] = o.ClusterName
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["place"] = o.Place
	}
	if true {
		toSerialize["servers"] = o.Servers
	}
	if true {
		toSerialize["supportPackage"] = o.SupportPackage
	}
	return json.Marshal(toSerialize)
}

type NullableV3CreateClusterRequest struct {
	value *V3CreateClusterRequest
	isSet bool
}

func (v NullableV3CreateClusterRequest) Get() *V3CreateClusterRequest {
	return v.value
}

func (v *NullableV3CreateClusterRequest) Set(val *V3CreateClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV3CreateClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV3CreateClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3CreateClusterRequest(val *V3CreateClusterRequest) *NullableV3CreateClusterRequest {
	return &NullableV3CreateClusterRequest{value: val, isSet: true}
}

func (v NullableV3CreateClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3CreateClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


