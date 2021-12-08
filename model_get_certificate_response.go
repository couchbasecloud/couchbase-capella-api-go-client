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

// GetCertificateResponse struct for GetCertificateResponse
type GetCertificateResponse struct {
	Certificate string `json:"certificate"`
}

// NewGetCertificateResponse instantiates a new GetCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCertificateResponse(certificate string) *GetCertificateResponse {
	this := GetCertificateResponse{}
	this.Certificate = certificate
	return &this
}

// NewGetCertificateResponseWithDefaults instantiates a new GetCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCertificateResponseWithDefaults() *GetCertificateResponse {
	this := GetCertificateResponse{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *GetCertificateResponse) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *GetCertificateResponse) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *GetCertificateResponse) SetCertificate(v string) {
	o.Certificate = v
}

func (o GetCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableGetCertificateResponse struct {
	value *GetCertificateResponse
	isSet bool
}

func (v NullableGetCertificateResponse) Get() *GetCertificateResponse {
	return v.value
}

func (v *NullableGetCertificateResponse) Set(val *GetCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCertificateResponse(val *GetCertificateResponse) *NullableGetCertificateResponse {
	return &NullableGetCertificateResponse{value: val, isSet: true}
}

func (v NullableGetCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

