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

// ClusterHealth * `healthy` - denotes that the resource is in a healthy state. * `unhealthy` - denotes that the resource is in an uhealthy state. * `degraded` - is reported for the overall health of a cluster or overall health for all nodes if all unhealthy nodes in the cluster are failed over. * `failedOver` - enotes that a couchbase server node failed over to a new pod. This status is attached only to a node. * `inactiveFailed` - is the value of the cluster membership as reported for a couchbase server node when it is in failed over state. * `N/A` - is the value of the cluster membership if we are unable to determine the health of the cluster 
type ClusterHealth string

// List of clusterHealth
const (
	CLUSTERHEALTH_HEALTHY ClusterHealth = "healthy"
	CLUSTERHEALTH_UNHEALTHY ClusterHealth = "unhealthy"
	CLUSTERHEALTH_DEGRADED ClusterHealth = "degraded"
	CLUSTERHEALTH_FAILED_OVER ClusterHealth = "failedOver"
	CLUSTERHEALTH_INACTIVE_FAILED ClusterHealth = "inactiveFailed"
	CLUSTERHEALTH_N_A ClusterHealth = "N/A"
)

var allowedClusterHealthEnumValues = []ClusterHealth{
	"healthy",
	"unhealthy",
	"degraded",
	"failedOver",
	"inactiveFailed",
	"N/A",
}

func (v *ClusterHealth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterHealth(value)
	for _, existing := range allowedClusterHealthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterHealth", value)
}

// NewClusterHealthFromValue returns a pointer to a valid ClusterHealth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterHealthFromValue(v string) (*ClusterHealth, error) {
	ev := ClusterHealth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterHealth: valid values are %v", v, allowedClusterHealthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterHealth) IsValid() bool {
	for _, existing := range allowedClusterHealthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clusterHealth value
func (v ClusterHealth) Ptr() *ClusterHealth {
	return &v
}

type NullableClusterHealth struct {
	value *ClusterHealth
	isSet bool
}

func (v NullableClusterHealth) Get() *ClusterHealth {
	return v.value
}

func (v *NullableClusterHealth) Set(val *ClusterHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterHealth(val *ClusterHealth) *NullableClusterHealth {
	return &NullableClusterHealth{value: val, isSet: true}
}

func (v NullableClusterHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

