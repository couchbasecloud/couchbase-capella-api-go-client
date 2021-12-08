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

// AWSProviderSettings struct for AWSProviderSettings
type AWSProviderSettings struct {
	Region *AwsRegions `json:"region,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
	SupportBucket *string `json:"supportBucket,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
	VpcCidr *string `json:"vpcCidr,omitempty"`
	VpcId *string `json:"vpcId,omitempty"`
	StackId *string `json:"stackId,omitempty"`
	PrivateSubnets *[]string `json:"privateSubnets,omitempty"`
	AvailabilityZone *[]string `json:"availabilityZone,omitempty"`
	AzCount *int32 `json:"azCount,omitempty"`
	LastRotated *time.Time `json:"lastRotated,omitempty"`
}

// NewAWSProviderSettings instantiates a new AWSProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSProviderSettings() *AWSProviderSettings {
	this := AWSProviderSettings{}
	return &this
}

// NewAWSProviderSettingsWithDefaults instantiates a new AWSProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSProviderSettingsWithDefaults() *AWSProviderSettings {
	this := AWSProviderSettings{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetRegion() AwsRegions {
	if o == nil || o.Region == nil {
		var ret AwsRegions
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetRegionOk() (*AwsRegions, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given AwsRegions and assigns it to the Region field.
func (o *AWSProviderSettings) SetRegion(v AwsRegions) {
	o.Region = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *AWSProviderSettings) SetBucket(v string) {
	o.Bucket = &v
}

// GetSupportBucket returns the SupportBucket field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetSupportBucket() string {
	if o == nil || o.SupportBucket == nil {
		var ret string
		return ret
	}
	return *o.SupportBucket
}

// GetSupportBucketOk returns a tuple with the SupportBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetSupportBucketOk() (*string, bool) {
	if o == nil || o.SupportBucket == nil {
		return nil, false
	}
	return o.SupportBucket, true
}

// HasSupportBucket returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasSupportBucket() bool {
	if o != nil && o.SupportBucket != nil {
		return true
	}

	return false
}

// SetSupportBucket gets a reference to the given string and assigns it to the SupportBucket field.
func (o *AWSProviderSettings) SetSupportBucket(v string) {
	o.SupportBucket = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AWSProviderSettings) SetAccountId(v string) {
	o.AccountId = &v
}

// GetVpcCidr returns the VpcCidr field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetVpcCidr() string {
	if o == nil || o.VpcCidr == nil {
		var ret string
		return ret
	}
	return *o.VpcCidr
}

// GetVpcCidrOk returns a tuple with the VpcCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetVpcCidrOk() (*string, bool) {
	if o == nil || o.VpcCidr == nil {
		return nil, false
	}
	return o.VpcCidr, true
}

// HasVpcCidr returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasVpcCidr() bool {
	if o != nil && o.VpcCidr != nil {
		return true
	}

	return false
}

// SetVpcCidr gets a reference to the given string and assigns it to the VpcCidr field.
func (o *AWSProviderSettings) SetVpcCidr(v string) {
	o.VpcCidr = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetVpcId() string {
	if o == nil || o.VpcId == nil {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetVpcIdOk() (*string, bool) {
	if o == nil || o.VpcId == nil {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasVpcId() bool {
	if o != nil && o.VpcId != nil {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *AWSProviderSettings) SetVpcId(v string) {
	o.VpcId = &v
}

// GetStackId returns the StackId field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetStackId() string {
	if o == nil || o.StackId == nil {
		var ret string
		return ret
	}
	return *o.StackId
}

// GetStackIdOk returns a tuple with the StackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetStackIdOk() (*string, bool) {
	if o == nil || o.StackId == nil {
		return nil, false
	}
	return o.StackId, true
}

// HasStackId returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasStackId() bool {
	if o != nil && o.StackId != nil {
		return true
	}

	return false
}

// SetStackId gets a reference to the given string and assigns it to the StackId field.
func (o *AWSProviderSettings) SetStackId(v string) {
	o.StackId = &v
}

// GetPrivateSubnets returns the PrivateSubnets field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetPrivateSubnets() []string {
	if o == nil || o.PrivateSubnets == nil {
		var ret []string
		return ret
	}
	return *o.PrivateSubnets
}

// GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetPrivateSubnetsOk() (*[]string, bool) {
	if o == nil || o.PrivateSubnets == nil {
		return nil, false
	}
	return o.PrivateSubnets, true
}

// HasPrivateSubnets returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasPrivateSubnets() bool {
	if o != nil && o.PrivateSubnets != nil {
		return true
	}

	return false
}

// SetPrivateSubnets gets a reference to the given []string and assigns it to the PrivateSubnets field.
func (o *AWSProviderSettings) SetPrivateSubnets(v []string) {
	o.PrivateSubnets = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetAvailabilityZone() []string {
	if o == nil || o.AvailabilityZone == nil {
		var ret []string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetAvailabilityZoneOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZone == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone != nil {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given []string and assigns it to the AvailabilityZone field.
func (o *AWSProviderSettings) SetAvailabilityZone(v []string) {
	o.AvailabilityZone = &v
}

// GetAzCount returns the AzCount field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetAzCount() int32 {
	if o == nil || o.AzCount == nil {
		var ret int32
		return ret
	}
	return *o.AzCount
}

// GetAzCountOk returns a tuple with the AzCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetAzCountOk() (*int32, bool) {
	if o == nil || o.AzCount == nil {
		return nil, false
	}
	return o.AzCount, true
}

// HasAzCount returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasAzCount() bool {
	if o != nil && o.AzCount != nil {
		return true
	}

	return false
}

// SetAzCount gets a reference to the given int32 and assigns it to the AzCount field.
func (o *AWSProviderSettings) SetAzCount(v int32) {
	o.AzCount = &v
}

// GetLastRotated returns the LastRotated field value if set, zero value otherwise.
func (o *AWSProviderSettings) GetLastRotated() time.Time {
	if o == nil || o.LastRotated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRotated
}

// GetLastRotatedOk returns a tuple with the LastRotated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSProviderSettings) GetLastRotatedOk() (*time.Time, bool) {
	if o == nil || o.LastRotated == nil {
		return nil, false
	}
	return o.LastRotated, true
}

// HasLastRotated returns a boolean if a field has been set.
func (o *AWSProviderSettings) HasLastRotated() bool {
	if o != nil && o.LastRotated != nil {
		return true
	}

	return false
}

// SetLastRotated gets a reference to the given time.Time and assigns it to the LastRotated field.
func (o *AWSProviderSettings) SetLastRotated(v time.Time) {
	o.LastRotated = &v
}

func (o AWSProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.SupportBucket != nil {
		toSerialize["supportBucket"] = o.SupportBucket
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.VpcCidr != nil {
		toSerialize["vpcCidr"] = o.VpcCidr
	}
	if o.VpcId != nil {
		toSerialize["vpcId"] = o.VpcId
	}
	if o.StackId != nil {
		toSerialize["stackId"] = o.StackId
	}
	if o.PrivateSubnets != nil {
		toSerialize["privateSubnets"] = o.PrivateSubnets
	}
	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if o.AzCount != nil {
		toSerialize["azCount"] = o.AzCount
	}
	if o.LastRotated != nil {
		toSerialize["lastRotated"] = o.LastRotated
	}
	return json.Marshal(toSerialize)
}

type NullableAWSProviderSettings struct {
	value *AWSProviderSettings
	isSet bool
}

func (v NullableAWSProviderSettings) Get() *AWSProviderSettings {
	return v.value
}

func (v *NullableAWSProviderSettings) Set(val *AWSProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSProviderSettings(val *AWSProviderSettings) *NullableAWSProviderSettings {
	return &NullableAWSProviderSettings{value: val, isSet: true}
}

func (v NullableAWSProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

