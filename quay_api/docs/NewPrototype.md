# NewPrototype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | Role that should be applied to the delegate | 
**ActivatingUser** | Pointer to [**NewPrototypeActivatingUser**](NewPrototypeActivatingUser.md) |  | [optional] 
**Delegate** | [**NewPrototypeDelegate**](NewPrototypeDelegate.md) |  | 

## Methods

### NewNewPrototype

`func NewNewPrototype(role string, delegate NewPrototypeDelegate, ) *NewPrototype`

NewNewPrototype instantiates a new NewPrototype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewPrototypeWithDefaults

`func NewNewPrototypeWithDefaults() *NewPrototype`

NewNewPrototypeWithDefaults instantiates a new NewPrototype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *NewPrototype) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NewPrototype) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NewPrototype) SetRole(v string)`

SetRole sets Role field to given value.


### GetActivatingUser

`func (o *NewPrototype) GetActivatingUser() NewPrototypeActivatingUser`

GetActivatingUser returns the ActivatingUser field if non-nil, zero value otherwise.

### GetActivatingUserOk

`func (o *NewPrototype) GetActivatingUserOk() (*NewPrototypeActivatingUser, bool)`

GetActivatingUserOk returns a tuple with the ActivatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatingUser

`func (o *NewPrototype) SetActivatingUser(v NewPrototypeActivatingUser)`

SetActivatingUser sets ActivatingUser field to given value.

### HasActivatingUser

`func (o *NewPrototype) HasActivatingUser() bool`

HasActivatingUser returns a boolean if a field has been set.

### GetDelegate

`func (o *NewPrototype) GetDelegate() NewPrototypeDelegate`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *NewPrototype) GetDelegateOk() (*NewPrototypeDelegate, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *NewPrototype) SetDelegate(v NewPrototypeDelegate)`

SetDelegate sets Delegate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


