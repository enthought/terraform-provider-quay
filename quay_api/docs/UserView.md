# UserView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | Pointer to **bool** | Whether the user&#39;s email address has been verified | [optional] 
**Anonymous** | **bool** | true if this user data represents a guest user | 
**Email** | Pointer to **string** | The user&#39;s email address | [optional] 
**Avatar** | **map[string]interface{}** | Avatar data representing the user&#39;s icon | 
**Organizations** | Pointer to **[]map[string]interface{}** | Information about the organizations in which the user is a member | [optional] 
**Logins** | Pointer to **[]map[string]interface{}** | The list of external login providers against which the user has authenticated | [optional] 
**CanCreateRepo** | Pointer to **bool** | Whether the user has permission to create repositories | [optional] 
**PreferredNamespace** | Pointer to **bool** | If true, the user&#39;s namespace is the preferred namespace to display | [optional] 

## Methods

### NewUserView

`func NewUserView(anonymous bool, avatar map[string]interface{}, ) *UserView`

NewUserView instantiates a new UserView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserViewWithDefaults

`func NewUserViewWithDefaults() *UserView`

NewUserViewWithDefaults instantiates a new UserView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerified

`func (o *UserView) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserView) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserView) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserView) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetAnonymous

`func (o *UserView) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *UserView) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *UserView) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.


### GetEmail

`func (o *UserView) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserView) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserView) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserView) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *UserView) GetAvatar() map[string]interface{}`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserView) GetAvatarOk() (*map[string]interface{}, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserView) SetAvatar(v map[string]interface{})`

SetAvatar sets Avatar field to given value.


### GetOrganizations

`func (o *UserView) GetOrganizations() []map[string]interface{}`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *UserView) GetOrganizationsOk() (*[]map[string]interface{}, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *UserView) SetOrganizations(v []map[string]interface{})`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *UserView) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetLogins

`func (o *UserView) GetLogins() []map[string]interface{}`

GetLogins returns the Logins field if non-nil, zero value otherwise.

### GetLoginsOk

`func (o *UserView) GetLoginsOk() (*[]map[string]interface{}, bool)`

GetLoginsOk returns a tuple with the Logins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogins

`func (o *UserView) SetLogins(v []map[string]interface{})`

SetLogins sets Logins field to given value.

### HasLogins

`func (o *UserView) HasLogins() bool`

HasLogins returns a boolean if a field has been set.

### GetCanCreateRepo

`func (o *UserView) GetCanCreateRepo() bool`

GetCanCreateRepo returns the CanCreateRepo field if non-nil, zero value otherwise.

### GetCanCreateRepoOk

`func (o *UserView) GetCanCreateRepoOk() (*bool, bool)`

GetCanCreateRepoOk returns a tuple with the CanCreateRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateRepo

`func (o *UserView) SetCanCreateRepo(v bool)`

SetCanCreateRepo sets CanCreateRepo field to given value.

### HasCanCreateRepo

`func (o *UserView) HasCanCreateRepo() bool`

HasCanCreateRepo returns a boolean if a field has been set.

### GetPreferredNamespace

`func (o *UserView) GetPreferredNamespace() bool`

GetPreferredNamespace returns the PreferredNamespace field if non-nil, zero value otherwise.

### GetPreferredNamespaceOk

`func (o *UserView) GetPreferredNamespaceOk() (*bool, bool)`

GetPreferredNamespaceOk returns a tuple with the PreferredNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredNamespace

`func (o *UserView) SetPreferredNamespace(v bool)`

SetPreferredNamespace sets PreferredNamespace field to given value.

### HasPreferredNamespace

`func (o *UserView) HasPreferredNamespace() bool`

HasPreferredNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


