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

// checks if the ChangeVisibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeVisibility{}

// ChangeVisibility Change the visibility for the repository.
type ChangeVisibility struct {
	// Visibility which the repository will start with
	Visibility string `json:"visibility"`
}

type _ChangeVisibility ChangeVisibility

// NewChangeVisibility instantiates a new ChangeVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeVisibility(visibility string) *ChangeVisibility {
	this := ChangeVisibility{}
	this.Visibility = visibility
	return &this
}

// NewChangeVisibilityWithDefaults instantiates a new ChangeVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeVisibilityWithDefaults() *ChangeVisibility {
	this := ChangeVisibility{}
	return &this
}

// GetVisibility returns the Visibility field value
func (o *ChangeVisibility) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *ChangeVisibility) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *ChangeVisibility) SetVisibility(v string) {
	o.Visibility = v
}

func (o ChangeVisibility) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeVisibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visibility"] = o.Visibility
	return toSerialize, nil
}

func (o *ChangeVisibility) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"visibility",
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

	varChangeVisibility := _ChangeVisibility{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChangeVisibility)

	if err != nil {
		return err
	}

	*o = ChangeVisibility(varChangeVisibility)

	return err
}

type NullableChangeVisibility struct {
	value *ChangeVisibility
	isSet bool
}

func (v NullableChangeVisibility) Get() *ChangeVisibility {
	return v.value
}

func (v *NullableChangeVisibility) Set(val *ChangeVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeVisibility(val *ChangeVisibility) *NullableChangeVisibility {
	return &NullableChangeVisibility{value: val, isSet: true}
}

func (v NullableChangeVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
