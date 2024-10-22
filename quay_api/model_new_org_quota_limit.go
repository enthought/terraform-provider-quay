/*
Quay Frontend

This API allows you to perform many of the operations required to work with Quay repositories, users, and organizations.

API version: v1
Contact: admin@example.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quay_api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NewOrgQuotaLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewOrgQuotaLimit{}

// NewOrgQuotaLimit Description of a new organization quota limit
type NewOrgQuotaLimit struct {
	// Type of quota limit: \"Warning\" or \"Reject\"
	Type string `json:"type"`
	// Quota threshold, in percent of quota
	ThresholdPercent int32 `json:"threshold_percent"`
}

type _NewOrgQuotaLimit NewOrgQuotaLimit

// NewNewOrgQuotaLimit instantiates a new NewOrgQuotaLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewOrgQuotaLimit(type_ string, thresholdPercent int32) *NewOrgQuotaLimit {
	this := NewOrgQuotaLimit{}
	this.Type = type_
	this.ThresholdPercent = thresholdPercent
	return &this
}

// NewNewOrgQuotaLimitWithDefaults instantiates a new NewOrgQuotaLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewOrgQuotaLimitWithDefaults() *NewOrgQuotaLimit {
	this := NewOrgQuotaLimit{}
	return &this
}

// GetType returns the Type field value
func (o *NewOrgQuotaLimit) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NewOrgQuotaLimit) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NewOrgQuotaLimit) SetType(v string) {
	o.Type = v
}

// GetThresholdPercent returns the ThresholdPercent field value
func (o *NewOrgQuotaLimit) GetThresholdPercent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ThresholdPercent
}

// GetThresholdPercentOk returns a tuple with the ThresholdPercent field value
// and a boolean to check if the value has been set.
func (o *NewOrgQuotaLimit) GetThresholdPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdPercent, true
}

// SetThresholdPercent sets field value
func (o *NewOrgQuotaLimit) SetThresholdPercent(v int32) {
	o.ThresholdPercent = v
}

func (o NewOrgQuotaLimit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewOrgQuotaLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["threshold_percent"] = o.ThresholdPercent
	return toSerialize, nil
}

func (o *NewOrgQuotaLimit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"threshold_percent",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNewOrgQuotaLimit := _NewOrgQuotaLimit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNewOrgQuotaLimit)

	if err != nil {
		return err
	}

	*o = NewOrgQuotaLimit(varNewOrgQuotaLimit)

	return err
}

type NullableNewOrgQuotaLimit struct {
	value *NewOrgQuotaLimit
	isSet bool
}

func (v NullableNewOrgQuotaLimit) Get() *NewOrgQuotaLimit {
	return v.value
}

func (v *NullableNewOrgQuotaLimit) Set(val *NewOrgQuotaLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrgQuotaLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrgQuotaLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrgQuotaLimit(val *NewOrgQuotaLimit) *NullableNewOrgQuotaLimit {
	return &NullableNewOrgQuotaLimit{value: val, isSet: true}
}

func (v NullableNewOrgQuotaLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrgQuotaLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
