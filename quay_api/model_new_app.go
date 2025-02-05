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

// checks if the NewApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewApp{}

// NewApp Description of a new organization application.
type NewApp struct {
	// The name of the application
	Name string `json:"name"`
	// The URI for the application's OAuth redirect
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The URI for the application's homepage
	ApplicationUri *string `json:"application_uri,omitempty"`
	// The human-readable description for the application
	Description *string `json:"description,omitempty"`
	// The e-mail address of the avatar to use for the application
	AvatarEmail *string `json:"avatar_email,omitempty"`
}

type _NewApp NewApp

// NewNewApp instantiates a new NewApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewApp(name string) *NewApp {
	this := NewApp{}
	this.Name = name
	return &this
}

// NewNewAppWithDefaults instantiates a new NewApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewAppWithDefaults() *NewApp {
	this := NewApp{}
	return &this
}

// GetName returns the Name field value
func (o *NewApp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewApp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewApp) SetName(v string) {
	o.Name = v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *NewApp) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri) {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewApp) GetRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *NewApp) HasRedirectUri() bool {
	if o != nil && !IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *NewApp) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetApplicationUri returns the ApplicationUri field value if set, zero value otherwise.
func (o *NewApp) GetApplicationUri() string {
	if o == nil || IsNil(o.ApplicationUri) {
		var ret string
		return ret
	}
	return *o.ApplicationUri
}

// GetApplicationUriOk returns a tuple with the ApplicationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewApp) GetApplicationUriOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationUri) {
		return nil, false
	}
	return o.ApplicationUri, true
}

// HasApplicationUri returns a boolean if a field has been set.
func (o *NewApp) HasApplicationUri() bool {
	if o != nil && !IsNil(o.ApplicationUri) {
		return true
	}

	return false
}

// SetApplicationUri gets a reference to the given string and assigns it to the ApplicationUri field.
func (o *NewApp) SetApplicationUri(v string) {
	o.ApplicationUri = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewApp) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewApp) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewApp) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewApp) SetDescription(v string) {
	o.Description = &v
}

// GetAvatarEmail returns the AvatarEmail field value if set, zero value otherwise.
func (o *NewApp) GetAvatarEmail() string {
	if o == nil || IsNil(o.AvatarEmail) {
		var ret string
		return ret
	}
	return *o.AvatarEmail
}

// GetAvatarEmailOk returns a tuple with the AvatarEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewApp) GetAvatarEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarEmail) {
		return nil, false
	}
	return o.AvatarEmail, true
}

// HasAvatarEmail returns a boolean if a field has been set.
func (o *NewApp) HasAvatarEmail() bool {
	if o != nil && !IsNil(o.AvatarEmail) {
		return true
	}

	return false
}

// SetAvatarEmail gets a reference to the given string and assigns it to the AvatarEmail field.
func (o *NewApp) SetAvatarEmail(v string) {
	o.AvatarEmail = &v
}

func (o NewApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.RedirectUri) {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if !IsNil(o.ApplicationUri) {
		toSerialize["application_uri"] = o.ApplicationUri
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AvatarEmail) {
		toSerialize["avatar_email"] = o.AvatarEmail
	}
	return toSerialize, nil
}

func (o *NewApp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varNewApp := _NewApp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNewApp)

	if err != nil {
		return err
	}

	*o = NewApp(varNewApp)

	return err
}

type NullableNewApp struct {
	value *NewApp
	isSet bool
}

func (v NullableNewApp) Get() *NewApp {
	return v.value
}

func (v *NullableNewApp) Set(val *NewApp) {
	v.value = val
	v.isSet = true
}

func (v NullableNewApp) IsSet() bool {
	return v.isSet
}

func (v *NullableNewApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewApp(val *NewApp) *NullableNewApp {
	return &NullableNewApp{value: val, isSet: true}
}

func (v NullableNewApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
