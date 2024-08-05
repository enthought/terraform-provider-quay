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

// checks if the CreateInstallUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstallUser{}

// CreateInstallUser Data for creating a user
type CreateInstallUser struct {
	// The username of the user being created
	Username string `json:"username"`
	// The email address of the user being created
	Email *string `json:"email,omitempty"`
}

type _CreateInstallUser CreateInstallUser

// NewCreateInstallUser instantiates a new CreateInstallUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstallUser(username string) *CreateInstallUser {
	this := CreateInstallUser{}
	this.Username = username
	return &this
}

// NewCreateInstallUserWithDefaults instantiates a new CreateInstallUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstallUserWithDefaults() *CreateInstallUser {
	this := CreateInstallUser{}
	return &this
}

// GetUsername returns the Username field value
func (o *CreateInstallUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateInstallUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateInstallUser) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateInstallUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstallUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateInstallUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateInstallUser) SetEmail(v string) {
	o.Email = &v
}

func (o CreateInstallUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInstallUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

func (o *CreateInstallUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varCreateInstallUser := _CreateInstallUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateInstallUser)

	if err != nil {
		return err
	}

	*o = CreateInstallUser(varCreateInstallUser)

	return err
}

type NullableCreateInstallUser struct {
	value *CreateInstallUser
	isSet bool
}

func (v NullableCreateInstallUser) Get() *CreateInstallUser {
	return v.value
}

func (v *NullableCreateInstallUser) Set(val *CreateInstallUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstallUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstallUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstallUser(val *CreateInstallUser) *NullableCreateInstallUser {
	return &NullableCreateInstallUser{value: val, isSet: true}
}

func (v NullableCreateInstallUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstallUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
