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

// checks if the TokenPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenPermission{}

// TokenPermission Description of a token permission
type TokenPermission struct {
	// Role to use for the token
	Role string `json:"role"`
}

type _TokenPermission TokenPermission

// NewTokenPermission instantiates a new TokenPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPermission(role string) *TokenPermission {
	this := TokenPermission{}
	this.Role = role
	return &this
}

// NewTokenPermissionWithDefaults instantiates a new TokenPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPermissionWithDefaults() *TokenPermission {
	this := TokenPermission{}
	return &this
}

// GetRole returns the Role field value
func (o *TokenPermission) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *TokenPermission) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *TokenPermission) SetRole(v string) {
	o.Role = v
}

func (o TokenPermission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *TokenPermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
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

	varTokenPermission := _TokenPermission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenPermission)

	if err != nil {
		return err
	}

	*o = TokenPermission(varTokenPermission)

	return err
}

type NullableTokenPermission struct {
	value *TokenPermission
	isSet bool
}

func (v NullableTokenPermission) Get() *TokenPermission {
	return v.value
}

func (v *NullableTokenPermission) Set(val *TokenPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPermission(val *TokenPermission) *NullableTokenPermission {
	return &NullableTokenPermission{value: val, isSet: true}
}

func (v NullableTokenPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}