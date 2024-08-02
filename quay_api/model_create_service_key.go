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

// checks if the CreateServiceKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceKey{}

// CreateServiceKey Description of creation of a service key
type CreateServiceKey struct {
	// The service authenticating with this key
	Service string `json:"service"`
	// The friendly name of a service key
	Name *string `json:"name,omitempty"`
	// The key/value pairs of this key's metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// If specified, the extra notes for the key
	Notes *string `json:"notes,omitempty"`
	// The expiration date as a unix timestamp
	Expiration map[string]interface{} `json:"expiration"`
}

type _CreateServiceKey CreateServiceKey

// NewCreateServiceKey instantiates a new CreateServiceKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceKey(service string, expiration map[string]interface{}) *CreateServiceKey {
	this := CreateServiceKey{}
	this.Service = service
	this.Expiration = expiration
	return &this
}

// NewCreateServiceKeyWithDefaults instantiates a new CreateServiceKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceKeyWithDefaults() *CreateServiceKey {
	this := CreateServiceKey{}
	return &this
}

// GetService returns the Service field value
func (o *CreateServiceKey) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *CreateServiceKey) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *CreateServiceKey) SetService(v string) {
	o.Service = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateServiceKey) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceKey) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateServiceKey) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateServiceKey) SetName(v string) {
	o.Name = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateServiceKey) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceKey) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateServiceKey) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreateServiceKey) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateServiceKey) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceKey) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateServiceKey) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateServiceKey) SetNotes(v string) {
	o.Notes = &v
}

// GetExpiration returns the Expiration field value
func (o *CreateServiceKey) GetExpiration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *CreateServiceKey) GetExpirationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Expiration, true
}

// SetExpiration sets field value
func (o *CreateServiceKey) SetExpiration(v map[string]interface{}) {
	o.Expiration = v
}

func (o CreateServiceKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service"] = o.Service
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["expiration"] = o.Expiration
	return toSerialize, nil
}

func (o *CreateServiceKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service",
		"expiration",
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

	varCreateServiceKey := _CreateServiceKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateServiceKey)

	if err != nil {
		return err
	}

	*o = CreateServiceKey(varCreateServiceKey)

	return err
}

type NullableCreateServiceKey struct {
	value *CreateServiceKey
	isSet bool
}

func (v NullableCreateServiceKey) Get() *CreateServiceKey {
	return v.value
}

func (v *NullableCreateServiceKey) Set(val *CreateServiceKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceKey(val *CreateServiceKey) *NullableCreateServiceKey {
	return &NullableCreateServiceKey{value: val, isSet: true}
}

func (v NullableCreateServiceKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}