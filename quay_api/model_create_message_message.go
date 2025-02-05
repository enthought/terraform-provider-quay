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

// checks if the CreateMessageMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessageMessage{}

// CreateMessageMessage A single message
type CreateMessageMessage struct {
	// The actual message
	Content string `json:"content"`
	// The media type of the message
	MediaType string `json:"media_type"`
	// The severity of the message
	Severity string `json:"severity"`
}

type _CreateMessageMessage CreateMessageMessage

// NewCreateMessageMessage instantiates a new CreateMessageMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageMessage(content string, mediaType string, severity string) *CreateMessageMessage {
	this := CreateMessageMessage{}
	this.Content = content
	this.MediaType = mediaType
	this.Severity = severity
	return &this
}

// NewCreateMessageMessageWithDefaults instantiates a new CreateMessageMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageMessageWithDefaults() *CreateMessageMessage {
	this := CreateMessageMessage{}
	return &this
}

// GetContent returns the Content field value
func (o *CreateMessageMessage) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateMessageMessage) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateMessageMessage) SetContent(v string) {
	o.Content = v
}

// GetMediaType returns the MediaType field value
func (o *CreateMessageMessage) GetMediaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *CreateMessageMessage) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *CreateMessageMessage) SetMediaType(v string) {
	o.MediaType = v
}

// GetSeverity returns the Severity field value
func (o *CreateMessageMessage) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *CreateMessageMessage) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *CreateMessageMessage) SetSeverity(v string) {
	o.Severity = v
}

func (o CreateMessageMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["media_type"] = o.MediaType
	toSerialize["severity"] = o.Severity
	return toSerialize, nil
}

func (o *CreateMessageMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"media_type",
		"severity",
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

	varCreateMessageMessage := _CreateMessageMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMessageMessage)

	if err != nil {
		return err
	}

	*o = CreateMessageMessage(varCreateMessageMessage)

	return err
}

type NullableCreateMessageMessage struct {
	value *CreateMessageMessage
	isSet bool
}

func (v NullableCreateMessageMessage) Get() *CreateMessageMessage {
	return v.value
}

func (v *NullableCreateMessageMessage) Set(val *CreateMessageMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageMessage(val *CreateMessageMessage) *NullableCreateMessageMessage {
	return &NullableCreateMessageMessage{value: val, isSet: true}
}

func (v NullableCreateMessageMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
