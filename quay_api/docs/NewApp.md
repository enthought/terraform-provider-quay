# NewApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the application | 
**RedirectUri** | Pointer to **string** | The URI for the application&#39;s OAuth redirect | [optional] 
**ApplicationUri** | Pointer to **string** | The URI for the application&#39;s homepage | [optional] 
**Description** | Pointer to **string** | The human-readable description for the application | [optional] 
**AvatarEmail** | Pointer to **string** | The e-mail address of the avatar to use for the application | [optional] 

## Methods

### NewNewApp

`func NewNewApp(name string, ) *NewApp`

NewNewApp instantiates a new NewApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppWithDefaults

`func NewNewAppWithDefaults() *NewApp`

NewNewAppWithDefaults instantiates a new NewApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewApp) SetName(v string)`

SetName sets Name field to given value.


### GetRedirectUri

`func (o *NewApp) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *NewApp) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *NewApp) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *NewApp) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetApplicationUri

`func (o *NewApp) GetApplicationUri() string`

GetApplicationUri returns the ApplicationUri field if non-nil, zero value otherwise.

### GetApplicationUriOk

`func (o *NewApp) GetApplicationUriOk() (*string, bool)`

GetApplicationUriOk returns a tuple with the ApplicationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUri

`func (o *NewApp) SetApplicationUri(v string)`

SetApplicationUri sets ApplicationUri field to given value.

### HasApplicationUri

`func (o *NewApp) HasApplicationUri() bool`

HasApplicationUri returns a boolean if a field has been set.

### GetDescription

`func (o *NewApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvatarEmail

`func (o *NewApp) GetAvatarEmail() string`

GetAvatarEmail returns the AvatarEmail field if non-nil, zero value otherwise.

### GetAvatarEmailOk

`func (o *NewApp) GetAvatarEmailOk() (*string, bool)`

GetAvatarEmailOk returns a tuple with the AvatarEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarEmail

`func (o *NewApp) SetAvatarEmail(v string)`

SetAvatarEmail sets AvatarEmail field to given value.

### HasAvatarEmail

`func (o *NewApp) HasAvatarEmail() bool`

HasAvatarEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


