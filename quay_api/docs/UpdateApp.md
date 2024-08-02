# UpdateApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the application | 
**RedirectUri** | **string** | The URI for the application&#39;s OAuth redirect | 
**ApplicationUri** | **string** | The URI for the application&#39;s homepage | 
**Description** | Pointer to **string** | The human-readable description for the application | [optional] 
**AvatarEmail** | Pointer to **string** | The e-mail address of the avatar to use for the application | [optional] 

## Methods

### NewUpdateApp

`func NewUpdateApp(name string, redirectUri string, applicationUri string, ) *UpdateApp`

NewUpdateApp instantiates a new UpdateApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppWithDefaults

`func NewUpdateAppWithDefaults() *UpdateApp`

NewUpdateAppWithDefaults instantiates a new UpdateApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApp) SetName(v string)`

SetName sets Name field to given value.


### GetRedirectUri

`func (o *UpdateApp) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *UpdateApp) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *UpdateApp) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetApplicationUri

`func (o *UpdateApp) GetApplicationUri() string`

GetApplicationUri returns the ApplicationUri field if non-nil, zero value otherwise.

### GetApplicationUriOk

`func (o *UpdateApp) GetApplicationUriOk() (*string, bool)`

GetApplicationUriOk returns a tuple with the ApplicationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUri

`func (o *UpdateApp) SetApplicationUri(v string)`

SetApplicationUri sets ApplicationUri field to given value.


### GetDescription

`func (o *UpdateApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvatarEmail

`func (o *UpdateApp) GetAvatarEmail() string`

GetAvatarEmail returns the AvatarEmail field if non-nil, zero value otherwise.

### GetAvatarEmailOk

`func (o *UpdateApp) GetAvatarEmailOk() (*string, bool)`

GetAvatarEmailOk returns a tuple with the AvatarEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarEmail

`func (o *UpdateApp) SetAvatarEmail(v string)`

SetAvatarEmail sets AvatarEmail field to given value.

### HasAvatarEmail

`func (o *UpdateApp) HasAvatarEmail() bool`

HasAvatarEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


