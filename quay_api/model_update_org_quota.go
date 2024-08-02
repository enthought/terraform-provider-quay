/*
Quay Frontend

This API allows you to perform many of the operations required to work with Quay repositories, users, and organizations.

API version: v1
Contact: admin@example.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quay_api

import (
	"encoding/json"
)

// checks if the UpdateOrgQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrgQuota{}

// UpdateOrgQuota Description of a new organization quota
type UpdateOrgQuota struct {
	// Number of bytes the organization is allowed
	LimitBytes *int32 `json:"limit_bytes,omitempty"`
}

// NewUpdateOrgQuota instantiates a new UpdateOrgQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrgQuota() *UpdateOrgQuota {
	this := UpdateOrgQuota{}
	return &this
}

// NewUpdateOrgQuotaWithDefaults instantiates a new UpdateOrgQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgQuotaWithDefaults() *UpdateOrgQuota {
	this := UpdateOrgQuota{}
	return &this
}

// GetLimitBytes returns the LimitBytes field value if set, zero value otherwise.
func (o *UpdateOrgQuota) GetLimitBytes() int32 {
	if o == nil || IsNil(o.LimitBytes) {
		var ret int32
		return ret
	}
	return *o.LimitBytes
}

// GetLimitBytesOk returns a tuple with the LimitBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgQuota) GetLimitBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitBytes) {
		return nil, false
	}
	return o.LimitBytes, true
}

// HasLimitBytes returns a boolean if a field has been set.
func (o *UpdateOrgQuota) HasLimitBytes() bool {
	if o != nil && !IsNil(o.LimitBytes) {
		return true
	}

	return false
}

// SetLimitBytes gets a reference to the given int32 and assigns it to the LimitBytes field.
func (o *UpdateOrgQuota) SetLimitBytes(v int32) {
	o.LimitBytes = &v
}

func (o UpdateOrgQuota) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrgQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitBytes) {
		toSerialize["limit_bytes"] = o.LimitBytes
	}
	return toSerialize, nil
}

type NullableUpdateOrgQuota struct {
	value *UpdateOrgQuota
	isSet bool
}

func (v NullableUpdateOrgQuota) Get() *UpdateOrgQuota {
	return v.value
}

func (v *NullableUpdateOrgQuota) Set(val *UpdateOrgQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrgQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrgQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrgQuota(val *UpdateOrgQuota) *NullableUpdateOrgQuota {
	return &NullableUpdateOrgQuota{value: val, isSet: true}
}

func (v NullableUpdateOrgQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrgQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
