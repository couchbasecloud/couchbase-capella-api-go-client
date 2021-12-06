/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecloud

import (
	"encoding/json"
)

// CouchbaseBucketSpec struct for CouchbaseBucketSpec
type CouchbaseBucketSpec struct {
	Name string `json:"name"`
	MemoryQuota int32 `json:"memoryQuota"`
	Replicas *int32 `json:"replicas,omitempty"`
	ConflictResolution *ConflictResolution `json:"conflictResolution,omitempty"`
}

// NewCouchbaseBucketSpec instantiates a new CouchbaseBucketSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouchbaseBucketSpec(name string, memoryQuota int32) *CouchbaseBucketSpec {
	this := CouchbaseBucketSpec{}
	this.Name = name
	this.MemoryQuota = memoryQuota
	var replicas int32 = 1
	this.Replicas = &replicas
	var conflictResolution ConflictResolution = SEQNO
	this.ConflictResolution = &conflictResolution
	return &this
}

// NewCouchbaseBucketSpecWithDefaults instantiates a new CouchbaseBucketSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouchbaseBucketSpecWithDefaults() *CouchbaseBucketSpec {
	this := CouchbaseBucketSpec{}
	var replicas int32 = 1
	this.Replicas = &replicas
	var conflictResolution ConflictResolution = SEQNO
	this.ConflictResolution = &conflictResolution
	return &this
}

// GetName returns the Name field value
func (o *CouchbaseBucketSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CouchbaseBucketSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CouchbaseBucketSpec) SetName(v string) {
	o.Name = v
}

// GetMemoryQuota returns the MemoryQuota field value
func (o *CouchbaseBucketSpec) GetMemoryQuota() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryQuota
}

// GetMemoryQuotaOk returns a tuple with the MemoryQuota field value
// and a boolean to check if the value has been set.
func (o *CouchbaseBucketSpec) GetMemoryQuotaOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MemoryQuota, true
}

// SetMemoryQuota sets field value
func (o *CouchbaseBucketSpec) SetMemoryQuota(v int32) {
	o.MemoryQuota = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *CouchbaseBucketSpec) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouchbaseBucketSpec) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *CouchbaseBucketSpec) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *CouchbaseBucketSpec) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetConflictResolution returns the ConflictResolution field value if set, zero value otherwise.
func (o *CouchbaseBucketSpec) GetConflictResolution() ConflictResolution {
	if o == nil || o.ConflictResolution == nil {
		var ret ConflictResolution
		return ret
	}
	return *o.ConflictResolution
}

// GetConflictResolutionOk returns a tuple with the ConflictResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouchbaseBucketSpec) GetConflictResolutionOk() (*ConflictResolution, bool) {
	if o == nil || o.ConflictResolution == nil {
		return nil, false
	}
	return o.ConflictResolution, true
}

// HasConflictResolution returns a boolean if a field has been set.
func (o *CouchbaseBucketSpec) HasConflictResolution() bool {
	if o != nil && o.ConflictResolution != nil {
		return true
	}

	return false
}

// SetConflictResolution gets a reference to the given ConflictResolution and assigns it to the ConflictResolution field.
func (o *CouchbaseBucketSpec) SetConflictResolution(v ConflictResolution) {
	o.ConflictResolution = &v
}

func (o CouchbaseBucketSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["memoryQuota"] = o.MemoryQuota
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.ConflictResolution != nil {
		toSerialize["conflictResolution"] = o.ConflictResolution
	}
	return json.Marshal(toSerialize)
}

type NullableCouchbaseBucketSpec struct {
	value *CouchbaseBucketSpec
	isSet bool
}

func (v NullableCouchbaseBucketSpec) Get() *CouchbaseBucketSpec {
	return v.value
}

func (v *NullableCouchbaseBucketSpec) Set(val *CouchbaseBucketSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCouchbaseBucketSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCouchbaseBucketSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouchbaseBucketSpec(val *CouchbaseBucketSpec) *NullableCouchbaseBucketSpec {
	return &NullableCouchbaseBucketSpec{value: val, isSet: true}
}

func (v NullableCouchbaseBucketSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouchbaseBucketSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


