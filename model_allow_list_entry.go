/*
Couchbase Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchbasecloud

import (
	"encoding/json"
	"time"
)

// AllowListEntry struct for AllowListEntry
type AllowListEntry struct {
	CidrBlock string `json:"cidrBlock"`
	RuleType AllowListRules `json:"ruleType"`
	Duration *string `json:"duration,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Comment *string `json:"comment,omitempty"`
	State *AllowListStates `json:"state,omitempty"`
}

// NewAllowListEntry instantiates a new AllowListEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowListEntry(cidrBlock string, ruleType AllowListRules) *AllowListEntry {
	this := AllowListEntry{}
	this.CidrBlock = cidrBlock
	this.RuleType = ruleType
	return &this
}

// NewAllowListEntryWithDefaults instantiates a new AllowListEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowListEntryWithDefaults() *AllowListEntry {
	this := AllowListEntry{}
	return &this
}

// GetCidrBlock returns the CidrBlock field value
func (o *AllowListEntry) GetCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetCidrBlockOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CidrBlock, true
}

// SetCidrBlock sets field value
func (o *AllowListEntry) SetCidrBlock(v string) {
	o.CidrBlock = v
}

// GetRuleType returns the RuleType field value
func (o *AllowListEntry) GetRuleType() AllowListRules {
	if o == nil {
		var ret AllowListRules
		return ret
	}

	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetRuleTypeOk() (*AllowListRules, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value
func (o *AllowListEntry) SetRuleType(v AllowListRules) {
	o.RuleType = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AllowListEntry) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AllowListEntry) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *AllowListEntry) SetDuration(v string) {
	o.Duration = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AllowListEntry) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AllowListEntry) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AllowListEntry) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AllowListEntry) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AllowListEntry) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AllowListEntry) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AllowListEntry) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AllowListEntry) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AllowListEntry) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AllowListEntry) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AllowListEntry) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AllowListEntry) SetComment(v string) {
	o.Comment = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AllowListEntry) GetState() AllowListStates {
	if o == nil || o.State == nil {
		var ret AllowListStates
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowListEntry) GetStateOk() (*AllowListStates, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AllowListEntry) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AllowListStates and assigns it to the State field.
func (o *AllowListEntry) SetState(v AllowListStates) {
	o.State = &v
}

func (o AllowListEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cidrBlock"] = o.CidrBlock
	}
	if true {
		toSerialize["ruleType"] = o.RuleType
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableAllowListEntry struct {
	value *AllowListEntry
	isSet bool
}

func (v NullableAllowListEntry) Get() *AllowListEntry {
	return v.value
}

func (v *NullableAllowListEntry) Set(val *AllowListEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowListEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowListEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowListEntry(val *AllowListEntry) *NullableAllowListEntry {
	return &NullableAllowListEntry{value: val, isSet: true}
}

func (v NullableAllowListEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowListEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

