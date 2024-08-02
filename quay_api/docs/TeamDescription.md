# TeamDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | Org wide permissions that should apply to the team | 
**Description** | Pointer to **string** | Markdown description for the team | [optional] 

## Methods

### NewTeamDescription

`func NewTeamDescription(role string, ) *TeamDescription`

NewTeamDescription instantiates a new TeamDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamDescriptionWithDefaults

`func NewTeamDescriptionWithDefaults() *TeamDescription`

NewTeamDescriptionWithDefaults instantiates a new TeamDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *TeamDescription) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TeamDescription) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TeamDescription) SetRole(v string)`

SetRole sets Role field to given value.


### GetDescription

`func (o *TeamDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


