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

// V3ClusterVersion struct for V3ClusterVersion
type V3ClusterVersion struct {
	Name string `json:"name"`
	Components V3ClusterVersionComponents `json:"components"`
}

// NewV3ClusterVersion instantiates a new V3ClusterVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3ClusterVersion(name string, components V3ClusterVersionComponents) *V3ClusterVersion {
	this := V3ClusterVersion{}
	this.Name = name
	this.Components = components
	return &this
}

// NewV3ClusterVersionWithDefaults instantiates a new V3ClusterVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3ClusterVersionWithDefaults() *V3ClusterVersion {
	this := V3ClusterVersion{}
	return &this
}

// GetName returns the Name field value
func (o *V3ClusterVersion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V3ClusterVersion) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V3ClusterVersion) SetName(v string) {
	o.Name = v
}

// GetComponents returns the Components field value
func (o *V3ClusterVersion) GetComponents() V3ClusterVersionComponents {
	if o == nil {
		var ret V3ClusterVersionComponents
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *V3ClusterVersion) GetComponentsOk() (*V3ClusterVersionComponents, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value
func (o *V3ClusterVersion) SetComponents(v V3ClusterVersionComponents) {
	o.Components = v
}

func (o V3ClusterVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableV3ClusterVersion struct {
	value *V3ClusterVersion
	isSet bool
}

func (v NullableV3ClusterVersion) Get() *V3ClusterVersion {
	return v.value
}

func (v *NullableV3ClusterVersion) Set(val *V3ClusterVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableV3ClusterVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableV3ClusterVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3ClusterVersion(val *V3ClusterVersion) *NullableV3ClusterVersion {
	return &NullableV3ClusterVersion{value: val, isSet: true}
}

func (v NullableV3ClusterVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3ClusterVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


