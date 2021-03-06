/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecapella

import (
	"encoding/json"
	"fmt"
)

// V3Provider the model 'V3Provider'
type V3Provider string

// List of v3Provider
const (
	V3PROVIDER_AWS V3Provider = "aws"
	V3PROVIDER_AZURE V3Provider = "azure"
	V3PROVIDER_GCP V3Provider = "gcp"
)

var allowedV3ProviderEnumValues = []V3Provider{
	"aws",
	"azure",
	"gcp",
}

func (v *V3Provider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V3Provider(value)
	for _, existing := range allowedV3ProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V3Provider", value)
}

// NewV3ProviderFromValue returns a pointer to a valid V3Provider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV3ProviderFromValue(v string) (*V3Provider, error) {
	ev := V3Provider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V3Provider: valid values are %v", v, allowedV3ProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V3Provider) IsValid() bool {
	for _, existing := range allowedV3ProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v3Provider value
func (v V3Provider) Ptr() *V3Provider {
	return &v
}

type NullableV3Provider struct {
	value *V3Provider
	isSet bool
}

func (v NullableV3Provider) Get() *V3Provider {
	return v.value
}

func (v *NullableV3Provider) Set(val *V3Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableV3Provider) IsSet() bool {
	return v.isSet
}

func (v *NullableV3Provider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3Provider(val *V3Provider) *NullableV3Provider {
	return &NullableV3Provider{value: val, isSet: true}
}

func (v NullableV3Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3Provider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

