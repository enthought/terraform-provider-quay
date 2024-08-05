# CreateRobot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional text description for the robot | [optional] 
**UnstructuredMetadata** | Pointer to **map[string]interface{}** | Optional unstructured metadata for the robot | [optional] 

## Methods

### NewCreateRobot

`func NewCreateRobot() *CreateRobot`

NewCreateRobot instantiates a new CreateRobot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRobotWithDefaults

`func NewCreateRobotWithDefaults() *CreateRobot`

NewCreateRobotWithDefaults instantiates a new CreateRobot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateRobot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRobot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRobot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRobot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnstructuredMetadata

`func (o *CreateRobot) GetUnstructuredMetadata() map[string]interface{}`

GetUnstructuredMetadata returns the UnstructuredMetadata field if non-nil, zero value otherwise.

### GetUnstructuredMetadataOk

`func (o *CreateRobot) GetUnstructuredMetadataOk() (*map[string]interface{}, bool)`

GetUnstructuredMetadataOk returns a tuple with the UnstructuredMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstructuredMetadata

`func (o *CreateRobot) SetUnstructuredMetadata(v map[string]interface{})`

SetUnstructuredMetadata sets UnstructuredMetadata field to given value.

### HasUnstructuredMetadata

`func (o *CreateRobot) HasUnstructuredMetadata() bool`

HasUnstructuredMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


