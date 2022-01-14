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

// AzureVolumeTypes the model 'AzureVolumeTypes'
type AzureVolumeTypes string

// List of azureVolumeTypes
const (
	AZUREVOLUMETYPES_P4 AzureVolumeTypes = "P4"
	AZUREVOLUMETYPES_P6 AzureVolumeTypes = "P6"
	AZUREVOLUMETYPES_P10 AzureVolumeTypes = "P10"
	AZUREVOLUMETYPES_P15 AzureVolumeTypes = "P15"
	AZUREVOLUMETYPES_P20 AzureVolumeTypes = "P20"
	AZUREVOLUMETYPES_P30 AzureVolumeTypes = "P30"
	AZUREVOLUMETYPES_P40 AzureVolumeTypes = "P40"
	AZUREVOLUMETYPES_P50 AzureVolumeTypes = "P50"
	AZUREVOLUMETYPES_P60 AzureVolumeTypes = "P60"
	AZUREVOLUMETYPES_P70 AzureVolumeTypes = "P70"
)

var allowedAzureVolumeTypesEnumValues = []AzureVolumeTypes{
	"P4",
	"P6",
	"P10",
	"P15",
	"P20",
	"P30",
	"P40",
	"P50",
	"P60",
	"P70",
}

func (v *AzureVolumeTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AzureVolumeTypes(value)
	for _, existing := range allowedAzureVolumeTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AzureVolumeTypes", value)
}

// NewAzureVolumeTypesFromValue returns a pointer to a valid AzureVolumeTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAzureVolumeTypesFromValue(v string) (*AzureVolumeTypes, error) {
	ev := AzureVolumeTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AzureVolumeTypes: valid values are %v", v, allowedAzureVolumeTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AzureVolumeTypes) IsValid() bool {
	for _, existing := range allowedAzureVolumeTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to azureVolumeTypes value
func (v AzureVolumeTypes) Ptr() *AzureVolumeTypes {
	return &v
}

type NullableAzureVolumeTypes struct {
	value *AzureVolumeTypes
	isSet bool
}

func (v NullableAzureVolumeTypes) Get() *AzureVolumeTypes {
	return v.value
}

func (v *NullableAzureVolumeTypes) Set(val *AzureVolumeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureVolumeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureVolumeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureVolumeTypes(val *AzureVolumeTypes) *NullableAzureVolumeTypes {
	return &NullableAzureVolumeTypes{value: val, isSet: true}
}

func (v NullableAzureVolumeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureVolumeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

