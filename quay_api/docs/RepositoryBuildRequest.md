# RepositoryBuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** | The file id that was generated when the build spec was uploaded | [optional] 
**ArchiveUrl** | Pointer to **string** | The URL of the .tar.gz to build. Must start with \&quot;http\&quot; or \&quot;https\&quot;. | [optional] 
**Subdirectory** | Pointer to **string** | Subdirectory in which the Dockerfile can be found. You can only specify this or dockerfile_path | [optional] 
**DockerfilePath** | Pointer to **string** | Path to a dockerfile. You can only specify this or subdirectory. | [optional] 
**Context** | Pointer to **string** | Pass in the context for the dockerfile. This is optional. | [optional] 
**PullRobot** | Pointer to **string** | Username of a Quay robot account to use as pull credentials | [optional] 
**DockerTags** | Pointer to **[]string** | The tags to which the built images will be pushed. If none specified, \&quot;latest\&quot; is used. | [optional] 

## Methods

### NewRepositoryBuildRequest

`func NewRepositoryBuildRequest() *RepositoryBuildRequest`

NewRepositoryBuildRequest instantiates a new RepositoryBuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryBuildRequestWithDefaults

`func NewRepositoryBuildRequestWithDefaults() *RepositoryBuildRequest`

NewRepositoryBuildRequestWithDefaults instantiates a new RepositoryBuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *RepositoryBuildRequest) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *RepositoryBuildRequest) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *RepositoryBuildRequest) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *RepositoryBuildRequest) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetArchiveUrl

`func (o *RepositoryBuildRequest) GetArchiveUrl() string`

GetArchiveUrl returns the ArchiveUrl field if non-nil, zero value otherwise.

### GetArchiveUrlOk

`func (o *RepositoryBuildRequest) GetArchiveUrlOk() (*string, bool)`

GetArchiveUrlOk returns a tuple with the ArchiveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveUrl

`func (o *RepositoryBuildRequest) SetArchiveUrl(v string)`

SetArchiveUrl sets ArchiveUrl field to given value.

### HasArchiveUrl

`func (o *RepositoryBuildRequest) HasArchiveUrl() bool`

HasArchiveUrl returns a boolean if a field has been set.

### GetSubdirectory

`func (o *RepositoryBuildRequest) GetSubdirectory() string`

GetSubdirectory returns the Subdirectory field if non-nil, zero value otherwise.

### GetSubdirectoryOk

`func (o *RepositoryBuildRequest) GetSubdirectoryOk() (*string, bool)`

GetSubdirectoryOk returns a tuple with the Subdirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdirectory

`func (o *RepositoryBuildRequest) SetSubdirectory(v string)`

SetSubdirectory sets Subdirectory field to given value.

### HasSubdirectory

`func (o *RepositoryBuildRequest) HasSubdirectory() bool`

HasSubdirectory returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *RepositoryBuildRequest) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *RepositoryBuildRequest) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *RepositoryBuildRequest) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *RepositoryBuildRequest) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetContext

`func (o *RepositoryBuildRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RepositoryBuildRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RepositoryBuildRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *RepositoryBuildRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetPullRobot

`func (o *RepositoryBuildRequest) GetPullRobot() string`

GetPullRobot returns the PullRobot field if non-nil, zero value otherwise.

### GetPullRobotOk

`func (o *RepositoryBuildRequest) GetPullRobotOk() (*string, bool)`

GetPullRobotOk returns a tuple with the PullRobot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRobot

`func (o *RepositoryBuildRequest) SetPullRobot(v string)`

SetPullRobot sets PullRobot field to given value.

### HasPullRobot

`func (o *RepositoryBuildRequest) HasPullRobot() bool`

HasPullRobot returns a boolean if a field has been set.

### GetDockerTags

`func (o *RepositoryBuildRequest) GetDockerTags() []string`

GetDockerTags returns the DockerTags field if non-nil, zero value otherwise.

### GetDockerTagsOk

`func (o *RepositoryBuildRequest) GetDockerTagsOk() (*[]string, bool)`

GetDockerTagsOk returns a tuple with the DockerTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTags

`func (o *RepositoryBuildRequest) SetDockerTags(v []string)`

SetDockerTags sets DockerTags field to given value.

### HasDockerTags

`func (o *RepositoryBuildRequest) HasDockerTags() bool`

HasDockerTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


