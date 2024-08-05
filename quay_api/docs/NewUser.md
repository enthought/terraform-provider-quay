# NewUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The user&#39;s username | 
**Password** | **string** | The user&#39;s password | 
**Email** | Pointer to **string** | The user&#39;s email address | [optional] 
**InviteCode** | Pointer to **string** | The optional invite code | [optional] 
**RecaptchaResponse** | Pointer to **string** | The (may be disabled) recaptcha response code for verification | [optional] 

## Methods

### NewNewUser

`func NewNewUser(username string, password string, ) *NewUser`

NewNewUser instantiates a new NewUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUserWithDefaults

`func NewNewUserWithDefaults() *NewUser`

NewNewUserWithDefaults instantiates a new NewUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *NewUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NewUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NewUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *NewUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *NewUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NewUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInviteCode

`func (o *NewUser) GetInviteCode() string`

GetInviteCode returns the InviteCode field if non-nil, zero value otherwise.

### GetInviteCodeOk

`func (o *NewUser) GetInviteCodeOk() (*string, bool)`

GetInviteCodeOk returns a tuple with the InviteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteCode

`func (o *NewUser) SetInviteCode(v string)`

SetInviteCode sets InviteCode field to given value.

### HasInviteCode

`func (o *NewUser) HasInviteCode() bool`

HasInviteCode returns a boolean if a field has been set.

### GetRecaptchaResponse

`func (o *NewUser) GetRecaptchaResponse() string`

GetRecaptchaResponse returns the RecaptchaResponse field if non-nil, zero value otherwise.

### GetRecaptchaResponseOk

`func (o *NewUser) GetRecaptchaResponseOk() (*string, bool)`

GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaResponse

`func (o *NewUser) SetRecaptchaResponse(v string)`

SetRecaptchaResponse sets RecaptchaResponse field to given value.

### HasRecaptchaResponse

`func (o *NewUser) HasRecaptchaResponse() bool`

HasRecaptchaResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


