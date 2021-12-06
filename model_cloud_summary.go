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

// CloudSummary struct for CloudSummary
type CloudSummary struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Status CloudStatus `json:"status"`
	Provider Provider `json:"provider"`
	Region Regions `json:"region"`
	VirtualNetworkID string `json:"virtualNetworkID"`
	VirtualNetworkCIDR string `json:"virtualNetworkCIDR"`
}

// NewCloudSummary instantiates a new CloudSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSummary(id string, name string, status CloudStatus, provider Provider, region Regions, virtualNetworkID string, virtualNetworkCIDR string) *CloudSummary {
	this := CloudSummary{}
	this.Id = id
	this.Name = name
	this.Status = status
	this.Provider = provider
	this.Region = region
	this.VirtualNetworkID = virtualNetworkID
	this.VirtualNetworkCIDR = virtualNetworkCIDR
	return &this
}

// NewCloudSummaryWithDefaults instantiates a new CloudSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSummaryWithDefaults() *CloudSummary {
	this := CloudSummary{}
	return &this
}

// GetId returns the Id field value
func (o *CloudSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CloudSummary) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CloudSummary) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CloudSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudSummary) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *CloudSummary) GetStatus() CloudStatus {
	if o == nil {
		var ret CloudStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudSummary) GetStatusOk() (*CloudStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloudSummary) SetStatus(v CloudStatus) {
	o.Status = v
}

// GetProvider returns the Provider field value
func (o *CloudSummary) GetProvider() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CloudSummary) GetProviderOk() (*Provider, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CloudSummary) SetProvider(v Provider) {
	o.Provider = v
}

// GetRegion returns the Region field value
func (o *CloudSummary) GetRegion() Regions {
	if o == nil {
		var ret Regions
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CloudSummary) GetRegionOk() (*Regions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CloudSummary) SetRegion(v Regions) {
	o.Region = v
}

// GetVirtualNetworkID returns the VirtualNetworkID field value
func (o *CloudSummary) GetVirtualNetworkID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualNetworkID
}

// GetVirtualNetworkIDOk returns a tuple with the VirtualNetworkID field value
// and a boolean to check if the value has been set.
func (o *CloudSummary) GetVirtualNetworkIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualNetworkID, true
}

// SetVirtualNetworkID sets field value
func (o *CloudSummary) SetVirtualNetworkID(v string) {
	o.VirtualNetworkID = v
}

// GetVirtualNetworkCIDR returns the VirtualNetworkCIDR field value
func (o *CloudSummary) GetVirtualNetworkCIDR() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualNetworkCIDR
}

// GetVirtualNetworkCIDROk returns a tuple with the VirtualNetworkCIDR field value
// and a boolean to check if the value has been set.
func (o *CloudSummary) GetVirtualNetworkCIDROk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualNetworkCIDR, true
}

// SetVirtualNetworkCIDR sets field value
func (o *CloudSummary) SetVirtualNetworkCIDR(v string) {
	o.VirtualNetworkCIDR = v
}

func (o CloudSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["virtualNetworkID"] = o.VirtualNetworkID
	}
	if true {
		toSerialize["virtualNetworkCIDR"] = o.VirtualNetworkCIDR
	}
	return json.Marshal(toSerialize)
}

type NullableCloudSummary struct {
	value *CloudSummary
	isSet bool
}

func (v NullableCloudSummary) Get() *CloudSummary {
	return v.value
}

func (v *NullableCloudSummary) Set(val *CloudSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSummary(val *CloudSummary) *NullableCloudSummary {
	return &NullableCloudSummary{value: val, isSet: true}
}

func (v NullableCloudSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


