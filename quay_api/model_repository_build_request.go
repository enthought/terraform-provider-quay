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

// checks if the RepositoryBuildRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryBuildRequest{}

// RepositoryBuildRequest Description of a new repository build.
type RepositoryBuildRequest struct {
	// The file id that was generated when the build spec was uploaded
	FileId *string `json:"file_id,omitempty"`
	// The URL of the .tar.gz to build. Must start with \"http\" or \"https\".
	ArchiveUrl *string `json:"archive_url,omitempty"`
	// Subdirectory in which the Dockerfile can be found. You can only specify this or dockerfile_path
	Subdirectory *string `json:"subdirectory,omitempty"`
	// Path to a dockerfile. You can only specify this or subdirectory.
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
	// Pass in the context for the dockerfile. This is optional.
	Context *string `json:"context,omitempty"`
	// Username of a Quay robot account to use as pull credentials
	PullRobot *string `json:"pull_robot,omitempty"`
	// The tags to which the built images will be pushed. If none specified, \"latest\" is used.
	DockerTags []string `json:"docker_tags,omitempty"`
}

// NewRepositoryBuildRequest instantiates a new RepositoryBuildRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryBuildRequest() *RepositoryBuildRequest {
	this := RepositoryBuildRequest{}
	return &this
}

// NewRepositoryBuildRequestWithDefaults instantiates a new RepositoryBuildRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryBuildRequestWithDefaults() *RepositoryBuildRequest {
	this := RepositoryBuildRequest{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *RepositoryBuildRequest) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryBuildRequest) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *RepositoryBuildRequest) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *RepositoryBuildRequest) SetFileId(v string) {
	o.FileId = &v
}

// GetArchiveUrl returns the ArchiveUrl field value if set, zero value otherwise.
func (o *RepositoryBuildRequest) GetArchiveUrl() string {
	if o == nil || IsNil(o.ArchiveUrl) {
		var ret string
		return ret
	}
	return *o.ArchiveUrl
}

// GetArchiveUrlOk returns a tuple with the ArchiveUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryBuildRequest) GetArchiveUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ArchiveUrl) {
		return nil, false
	}
	return o.ArchiveUrl, true
}

// HasArchiveUrl returns a boolean if a field has been set.
func (o *RepositoryBuildRequest) HasArchiveUrl() bool {
	if o != nil && !IsNil(o.ArchiveUrl) {
		return true
	}

	return false
}

// SetArchiveUrl gets a reference to the given string and assigns it to the ArchiveUrl field.
func (o *RepositoryBuildRequest) SetArchiveUrl(v string) {
	o.ArchiveUrl = &v
}

// GetSubdirectory returns the Subdirectory field value if set, zero value otherwise.
func (o *RepositoryBuildRequest) GetSubdirectory() string {
	if o == nil || IsNil(o.Subdirectory) {
		var ret string
		return ret
	}
	return *o.Subdirectory
}

// GetSubdirectoryOk returns a tuple with the Subdirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryBuildRequest) GetSubdirectoryOk() (*string, bool) {
	if o == nil || IsNil(o.Subdirectory) {
		return nil, false
	}
	return o.Subdirectory, true
}

// HasSubdirectory returns a boolean if a field has been set.
func (o *RepositoryBuildRequest) HasSubdirectory() bool {
	if o != nil && !IsNil(o.Subdirectory) {
		return true
	}

	return false
}

// SetSubdirectory gets a reference to the given string and assigns it to the Subdirectory field.
func (o *RepositoryBuildRequest) SetSubdirectory(v string) {
	o.Subdirectory = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *RepositoryBuildRequest) GetDockerfilePath() string {
	if o == nil || IsNil(o.DockerfilePath) {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryBuildRequest) GetDockerfilePathOk() (*string, bool) {
	if o == nil || IsNil(o.DockerfilePath) {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *RepositoryBuildRequest) HasDockerfilePath() bool {
	if o != nil && !IsNil(o.DockerfilePath) {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *RepositoryBuildRequest) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RepositoryBuildRequest) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryBuildRequest) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RepositoryBuildRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *RepositoryBuildRequest) SetContext(v string) {
	o.Context = &v
}

// GetPullRobot returns the PullRobot field value if set, zero value otherwise.
func (o *RepositoryBuildRequest) GetPullRobot() string {
	if o == nil || IsNil(o.PullRobot) {
		var ret string
		return ret
	}
	return *o.PullRobot
}

// GetPullRobotOk returns a tuple with the PullRobot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryBuildRequest) GetPullRobotOk() (*string, bool) {
	if o == nil || IsNil(o.PullRobot) {
		return nil, false
	}
	return o.PullRobot, true
}

// HasPullRobot returns a boolean if a field has been set.
func (o *RepositoryBuildRequest) HasPullRobot() bool {
	if o != nil && !IsNil(o.PullRobot) {
		return true
	}

	return false
}

// SetPullRobot gets a reference to the given string and assigns it to the PullRobot field.
func (o *RepositoryBuildRequest) SetPullRobot(v string) {
	o.PullRobot = &v
}

// GetDockerTags returns the DockerTags field value if set, zero value otherwise.
func (o *RepositoryBuildRequest) GetDockerTags() []string {
	if o == nil || IsNil(o.DockerTags) {
		var ret []string
		return ret
	}
	return o.DockerTags
}

// GetDockerTagsOk returns a tuple with the DockerTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryBuildRequest) GetDockerTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.DockerTags) {
		return nil, false
	}
	return o.DockerTags, true
}

// HasDockerTags returns a boolean if a field has been set.
func (o *RepositoryBuildRequest) HasDockerTags() bool {
	if o != nil && !IsNil(o.DockerTags) {
		return true
	}

	return false
}

// SetDockerTags gets a reference to the given []string and assigns it to the DockerTags field.
func (o *RepositoryBuildRequest) SetDockerTags(v []string) {
	o.DockerTags = v
}

func (o RepositoryBuildRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryBuildRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileId) {
		toSerialize["file_id"] = o.FileId
	}
	if !IsNil(o.ArchiveUrl) {
		toSerialize["archive_url"] = o.ArchiveUrl
	}
	if !IsNil(o.Subdirectory) {
		toSerialize["subdirectory"] = o.Subdirectory
	}
	if !IsNil(o.DockerfilePath) {
		toSerialize["dockerfile_path"] = o.DockerfilePath
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.PullRobot) {
		toSerialize["pull_robot"] = o.PullRobot
	}
	if !IsNil(o.DockerTags) {
		toSerialize["docker_tags"] = o.DockerTags
	}
	return toSerialize, nil
}

type NullableRepositoryBuildRequest struct {
	value *RepositoryBuildRequest
	isSet bool
}

func (v NullableRepositoryBuildRequest) Get() *RepositoryBuildRequest {
	return v.value
}

func (v *NullableRepositoryBuildRequest) Set(val *RepositoryBuildRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryBuildRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryBuildRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryBuildRequest(val *RepositoryBuildRequest) *NullableRepositoryBuildRequest {
	return &NullableRepositoryBuildRequest{value: val, isSet: true}
}

func (v NullableRepositoryBuildRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryBuildRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
