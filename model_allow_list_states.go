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

// AllowListStates the model 'AllowListStates'
type AllowListStates string

// List of allowListStates
const (
	ALLOWLISTSTATES_ACTIVE AllowListStates = "active"
	ALLOWLISTSTATES_PENDING AllowListStates = "pending"
	ALLOWLISTSTATES_FAILED AllowListStates = "failed"
	ALLOWLISTSTATES_EXPIRED AllowListStates = "expired"
)

var allowedAllowListStatesEnumValues = []AllowListStates{
	"active",
	"pending",
	"failed",
	"expired",
}

func (v *AllowListStates) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllowListStates(value)
	for _, existing := range allowedAllowListStatesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllowListStates", value)
}

// NewAllowListStatesFromValue returns a pointer to a valid AllowListStates
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllowListStatesFromValue(v string) (*AllowListStates, error) {
	ev := AllowListStates(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllowListStates: valid values are %v", v, allowedAllowListStatesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllowListStates) IsValid() bool {
	for _, existing := range allowedAllowListStatesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to allowListStates value
func (v AllowListStates) Ptr() *AllowListStates {
	return &v
}

type NullableAllowListStates struct {
	value *AllowListStates
	isSet bool
}

func (v NullableAllowListStates) Get() *AllowListStates {
	return v.value
}

func (v *NullableAllowListStates) Set(val *AllowListStates) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowListStates) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowListStates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowListStates(val *AllowListStates) *NullableAllowListStates {
	return &NullableAllowListStates{value: val, isSet: true}
}

func (v NullableAllowListStates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowListStates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

