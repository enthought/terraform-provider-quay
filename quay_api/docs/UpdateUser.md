# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The user&#39;s password | [optional] 
**InvoiceEmail** | Pointer to **bool** | Whether the user desires to receive an invoice email. | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address | [optional] 
**TagExpirationS** | Pointer to **int32** | The number of seconds for tag expiration | [optional] 
**Username** | Pointer to **string** | The user&#39;s username | [optional] 

## Methods

### NewUpdateUser

`func NewUpdateUser() *UpdateUser`

NewUpdateUser instantiates a new UpdateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserWithDefaults

`func NewUpdateUserWithDefaults() *UpdateUser`

NewUpdateUserWithDefaults instantiates a new UpdateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetInvoiceEmail

`func (o *UpdateUser) GetInvoiceEmail() bool`

GetInvoiceEmail returns the InvoiceEmail field if non-nil, zero value otherwise.

### GetInvoiceEmailOk

`func (o *UpdateUser) GetInvoiceEmailOk() (*bool, bool)`

GetInvoiceEmailOk returns a tuple with the InvoiceEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceEmail

`func (o *UpdateUser) SetInvoiceEmail(v bool)`

SetInvoiceEmail sets InvoiceEmail field to given value.

### HasInvoiceEmail

`func (o *UpdateUser) HasInvoiceEmail() bool`

HasInvoiceEmail returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTagExpirationS

`func (o *UpdateUser) GetTagExpirationS() int32`

GetTagExpirationS returns the TagExpirationS field if non-nil, zero value otherwise.

### GetTagExpirationSOk

`func (o *UpdateUser) GetTagExpirationSOk() (*int32, bool)`

GetTagExpirationSOk returns a tuple with the TagExpirationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagExpirationS

`func (o *UpdateUser) SetTagExpirationS(v int32)`

SetTagExpirationS sets TagExpirationS field to given value.

### HasTagExpirationS

`func (o *UpdateUser) HasTagExpirationS() bool`

HasTagExpirationS returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


