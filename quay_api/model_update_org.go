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

// checks if the UpdateOrg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrg{}

// UpdateOrg Description of updates for an organization
type UpdateOrg struct {
	// The new name for the organization
	Name *string `json:"name,omitempty"`
	// Organization contact email
	Email *string `json:"email,omitempty"`
}

// NewUpdateOrg instantiates a new UpdateOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrg() *UpdateOrg {
	this := UpdateOrg{}
	return &this
}

// NewUpdateOrgWithDefaults instantiates a new UpdateOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgWithDefaults() *UpdateOrg {
	this := UpdateOrg{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrg) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrg) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrg) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrg) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateOrg) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrg) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateOrg) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateOrg) SetEmail(v string) {
	o.Email = &v
}

func (o UpdateOrg) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableUpdateOrg struct {
	value *UpdateOrg
	isSet bool
}

func (v NullableUpdateOrg) Get() *UpdateOrg {
	return v.value
}

func (v *NullableUpdateOrg) Set(val *UpdateOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrg(val *UpdateOrg) *NullableUpdateOrg {
	return &NullableUpdateOrg{value: val, isSet: true}
}

func (v NullableUpdateOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
