# CreateInstallUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username of the user being created | 
**Email** | Pointer to **string** | The email address of the user being created | [optional] 

## Methods

### NewCreateInstallUser

`func NewCreateInstallUser(username string, ) *CreateInstallUser`

NewCreateInstallUser instantiates a new CreateInstallUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstallUserWithDefaults

`func NewCreateInstallUserWithDefaults() *CreateInstallUser`

NewCreateInstallUserWithDefaults instantiates a new CreateInstallUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateInstallUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateInstallUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateInstallUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *CreateInstallUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateInstallUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateInstallUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateInstallUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


