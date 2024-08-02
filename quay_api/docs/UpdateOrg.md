# UpdateOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name for the organization | [optional] 
**Email** | Pointer to **string** | Organization contact email | [optional] 

## Methods

### NewUpdateOrg

`func NewUpdateOrg() *UpdateOrg`

NewUpdateOrg instantiates a new UpdateOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgWithDefaults

`func NewUpdateOrgWithDefaults() *UpdateOrg`

NewUpdateOrgWithDefaults instantiates a new UpdateOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateOrg) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateOrg) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateOrg) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateOrg) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


