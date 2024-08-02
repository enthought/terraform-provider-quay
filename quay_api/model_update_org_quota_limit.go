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

// checks if the UpdateOrgQuotaLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrgQuotaLimit{}

// UpdateOrgQuotaLimit Description of changing organization quota limit
type UpdateOrgQuotaLimit struct {
	// Type of quota limit: \"Warning\" or \"Reject\"
	Type *string `json:"type,omitempty"`
	// Quota threshold, in percent of quota
	ThresholdPercent *int32 `json:"threshold_percent,omitempty"`
}

// NewUpdateOrgQuotaLimit instantiates a new UpdateOrgQuotaLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrgQuotaLimit() *UpdateOrgQuotaLimit {
	this := UpdateOrgQuotaLimit{}
	return &this
}

// NewUpdateOrgQuotaLimitWithDefaults instantiates a new UpdateOrgQuotaLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgQuotaLimitWithDefaults() *UpdateOrgQuotaLimit {
	this := UpdateOrgQuotaLimit{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateOrgQuotaLimit) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgQuotaLimit) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateOrgQuotaLimit) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateOrgQuotaLimit) SetType(v string) {
	o.Type = &v
}

// GetThresholdPercent returns the ThresholdPercent field value if set, zero value otherwise.
func (o *UpdateOrgQuotaLimit) GetThresholdPercent() int32 {
	if o == nil || IsNil(o.ThresholdPercent) {
		var ret int32
		return ret
	}
	return *o.ThresholdPercent
}

// GetThresholdPercentOk returns a tuple with the ThresholdPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgQuotaLimit) GetThresholdPercentOk() (*int32, bool) {
	if o == nil || IsNil(o.ThresholdPercent) {
		return nil, false
	}
	return o.ThresholdPercent, true
}

// HasThresholdPercent returns a boolean if a field has been set.
func (o *UpdateOrgQuotaLimit) HasThresholdPercent() bool {
	if o != nil && !IsNil(o.ThresholdPercent) {
		return true
	}

	return false
}

// SetThresholdPercent gets a reference to the given int32 and assigns it to the ThresholdPercent field.
func (o *UpdateOrgQuotaLimit) SetThresholdPercent(v int32) {
	o.ThresholdPercent = &v
}

func (o UpdateOrgQuotaLimit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrgQuotaLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ThresholdPercent) {
		toSerialize["threshold_percent"] = o.ThresholdPercent
	}
	return toSerialize, nil
}

type NullableUpdateOrgQuotaLimit struct {
	value *UpdateOrgQuotaLimit
	isSet bool
}

func (v NullableUpdateOrgQuotaLimit) Get() *UpdateOrgQuotaLimit {
	return v.value
}

func (v *NullableUpdateOrgQuotaLimit) Set(val *UpdateOrgQuotaLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrgQuotaLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrgQuotaLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrgQuotaLimit(val *UpdateOrgQuotaLimit) *NullableUpdateOrgQuotaLimit {
	return &NullableUpdateOrgQuotaLimit{value: val, isSet: true}
}

func (v NullableUpdateOrgQuotaLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrgQuotaLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
