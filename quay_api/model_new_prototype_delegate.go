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

// checks if the NewPrototypeDelegate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewPrototypeDelegate{}

// NewPrototypeDelegate Information about the user or team to which the rule grants access
type NewPrototypeDelegate struct {
	// The name for the delegate team or user
	Name string `json:"name"`
	// Whether the delegate is a user or a team
	Kind string `json:"kind"`
}

type _NewPrototypeDelegate NewPrototypeDelegate

// NewNewPrototypeDelegate instantiates a new NewPrototypeDelegate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewPrototypeDelegate(name string, kind string) *NewPrototypeDelegate {
	this := NewPrototypeDelegate{}
	this.Name = name
	this.Kind = kind
	return &this
}

// NewNewPrototypeDelegateWithDefaults instantiates a new NewPrototypeDelegate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewPrototypeDelegateWithDefaults() *NewPrototypeDelegate {
	this := NewPrototypeDelegate{}
	return &this
}

// GetName returns the Name field value
func (o *NewPrototypeDelegate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewPrototypeDelegate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewPrototypeDelegate) SetName(v string) {
	o.Name = v
}

// GetKind returns the Kind field value
func (o *NewPrototypeDelegate) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NewPrototypeDelegate) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NewPrototypeDelegate) SetKind(v string) {
	o.Kind = v
}

func (o NewPrototypeDelegate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewPrototypeDelegate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["kind"] = o.Kind
	return toSerialize, nil
}

func (o *NewPrototypeDelegate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"kind",
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

	varNewPrototypeDelegate := _NewPrototypeDelegate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNewPrototypeDelegate)

	if err != nil {
		return err
	}

	*o = NewPrototypeDelegate(varNewPrototypeDelegate)

	return err
}

type NullableNewPrototypeDelegate struct {
	value *NewPrototypeDelegate
	isSet bool
}

func (v NullableNewPrototypeDelegate) Get() *NewPrototypeDelegate {
	return v.value
}

func (v *NullableNewPrototypeDelegate) Set(val *NewPrototypeDelegate) {
	v.value = val
	v.isSet = true
}

func (v NullableNewPrototypeDelegate) IsSet() bool {
	return v.isSet
}

func (v *NullableNewPrototypeDelegate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewPrototypeDelegate(val *NewPrototypeDelegate) *NullableNewPrototypeDelegate {
	return &NullableNewPrototypeDelegate{value: val, isSet: true}
}

func (v NullableNewPrototypeDelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewPrototypeDelegate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}