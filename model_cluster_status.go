/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecloud

import (
	"encoding/json"
	"fmt"
)

// ClusterStatus * `draft` - indicates that all fields under database model are mutable. * `needs_deploy` - indicates that the new configuration changes have not been applied to the cluster * `job_scheduled` - indicates that the new configuration changes have been scheduled to be applied. * `deploying` - indicates that database is currently being deployed/new configurations being applied. * `deploy_failed` - indicates that the last deployment failed. * `deploy_succeeded` - indicates that the last deployment succeeded. * `ready` - indicates that the database is ready. * `destroying` - indicates that the database is being destroyed. * `destroy_failed` - indicates that the database destroy job failed. * `destroy_succeeded` - indicates that the database destroy job succeeded. * `metrics_failed` - indicates that the database is not serving metrics. * `preflight_started` - indicates a preflight check has started on the database * `preflight_failed` - indicates a preflight check has failed on the database * `preflight_succeeded` - indicates a preflight check has succeeded on the database * `management_blocked` - indicates the IAM credentiald for the cloud hosting the databse have been deleted. * `upgrading` - indicates the database is being upgrading.
type ClusterStatus string

// List of clusterStatus
const (
	DRAFT                       ClusterStatus = "draft"
	NEEDS_DEPLOY                ClusterStatus = "needs_deploy"
	JOB_SCHEDULED               ClusterStatus = "job_scheduled"
	CLUSTER_DEPLOYING           ClusterStatus = "deploying"
	CLUSTER_DEPLOY_FAILED       ClusterStatus = "deploy_failed"
	CLUSTER_DEPLOY_SUCCEEDED    ClusterStatus = "deploy_succeeded"
	CLUSTER_READY               ClusterStatus = "ready"
	CLUSTER_DESTROYING          ClusterStatus = "destroying"
	CLUSTER_DESTROY_FAILED      ClusterStatus = "destroy_failed"
	CLUSTER_DESTROY_SUCCEEDED   ClusterStatus = "destroy_succeeded"
	METRICS_FAILED              ClusterStatus = "metrics_failed"
	CLUSTER_PREFLIGHT_STARTED   ClusterStatus = "preflight_started"
	CLUSTER_PREFLIGHT_FAILED    ClusterStatus = "preflight_failed"
	CLUSTER_PREFLIGHT_SUCCEEDED ClusterStatus = "preflight_succeeded"
	CLUSTER_MANAGEMENT_BLOCKED  ClusterStatus = "management_blocked"
	UPGRADING                   ClusterStatus = "upgrading"
)

var allowedClusterStatusEnumValues = []ClusterStatus{
	"draft",
	"needs_deploy",
	"job_scheduled",
	"deploying",
	"deploy_failed",
	"deploy_succeeded",
	"ready",
	"destroying",
	"destroy_failed",
	"destroy_succeeded",
	"metrics_failed",
	"preflight_started",
	"preflight_failed",
	"preflight_succeeded",
	"management_blocked",
	"upgrading",
}

func (v *ClusterStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterStatus(value)
	for _, existing := range allowedClusterStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterStatus", value)
}

// NewClusterStatusFromValue returns a pointer to a valid ClusterStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterStatusFromValue(v string) (*ClusterStatus, error) {
	ev := ClusterStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterStatus: valid values are %v", v, allowedClusterStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterStatus) IsValid() bool {
	for _, existing := range allowedClusterStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clusterStatus value
func (v ClusterStatus) Ptr() *ClusterStatus {
	return &v
}

type NullableClusterStatus struct {
	value *ClusterStatus
	isSet bool
}

func (v NullableClusterStatus) Get() *ClusterStatus {
	return v.value
}

func (v *NullableClusterStatus) Set(val *ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatus(val *ClusterStatus) *NullableClusterStatus {
	return &NullableClusterStatus{value: val, isSet: true}
}

func (v NullableClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
