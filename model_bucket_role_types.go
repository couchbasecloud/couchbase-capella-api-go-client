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

// BucketRoleTypes * `data_writer` - Gives the user permission to read and write data for the bucket. * `data_reader` - Gives the user read-only permission for the bucket. 
type BucketRoleTypes string

// List of bucketRoleTypes
const (
	BUCKETROLETYPES_WRITER BucketRoleTypes = "data_writer"
	BUCKETROLETYPES_READER BucketRoleTypes = "data_reader"
)

var allowedBucketRoleTypesEnumValues = []BucketRoleTypes{
	"data_writer",
	"data_reader",
}

func (v *BucketRoleTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BucketRoleTypes(value)
	for _, existing := range allowedBucketRoleTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BucketRoleTypes", value)
}

// NewBucketRoleTypesFromValue returns a pointer to a valid BucketRoleTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBucketRoleTypesFromValue(v string) (*BucketRoleTypes, error) {
	ev := BucketRoleTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BucketRoleTypes: valid values are %v", v, allowedBucketRoleTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BucketRoleTypes) IsValid() bool {
	for _, existing := range allowedBucketRoleTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bucketRoleTypes value
func (v BucketRoleTypes) Ptr() *BucketRoleTypes {
	return &v
}

type NullableBucketRoleTypes struct {
	value *BucketRoleTypes
	isSet bool
}

func (v NullableBucketRoleTypes) Get() *BucketRoleTypes {
	return v.value
}

func (v *NullableBucketRoleTypes) Set(val *BucketRoleTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketRoleTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketRoleTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketRoleTypes(val *BucketRoleTypes) *NullableBucketRoleTypes {
	return &NullableBucketRoleTypes{value: val, isSet: true}
}

func (v NullableBucketRoleTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketRoleTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

