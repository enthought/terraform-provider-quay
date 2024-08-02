# BuildTriggerActivateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]interface{}** | Arbitrary json. | 
**PullRobot** | Pointer to **string** | The name of the robot that will be used to pull images. | [optional] 

## Methods

### NewBuildTriggerActivateRequest

`func NewBuildTriggerActivateRequest(config map[string]interface{}, ) *BuildTriggerActivateRequest`

NewBuildTriggerActivateRequest instantiates a new BuildTriggerActivateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildTriggerActivateRequestWithDefaults

`func NewBuildTriggerActivateRequestWithDefaults() *BuildTriggerActivateRequest`

NewBuildTriggerActivateRequestWithDefaults instantiates a new BuildTriggerActivateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *BuildTriggerActivateRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BuildTriggerActivateRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BuildTriggerActivateRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetPullRobot

`func (o *BuildTriggerActivateRequest) GetPullRobot() string`

GetPullRobot returns the PullRobot field if non-nil, zero value otherwise.

### GetPullRobotOk

`func (o *BuildTriggerActivateRequest) GetPullRobotOk() (*string, bool)`

GetPullRobotOk returns a tuple with the PullRobot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRobot

`func (o *BuildTriggerActivateRequest) SetPullRobot(v string)`

SetPullRobot sets PullRobot field to given value.

### HasPullRobot

`func (o *BuildTriggerActivateRequest) HasPullRobot() bool`

HasPullRobot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


