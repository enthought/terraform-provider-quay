# NewOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Organization username | 
**Email** | Pointer to **string** | Organization contact email | [optional] 
**RecaptchaResponse** | Pointer to **string** | The (may be disabled) recaptcha response code for verification | [optional] 

## Methods

### NewNewOrg

`func NewNewOrg(name string, ) *NewOrg`

NewNewOrg instantiates a new NewOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewOrgWithDefaults

`func NewNewOrgWithDefaults() *NewOrg`

NewNewOrgWithDefaults instantiates a new NewOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewOrg) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *NewOrg) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewOrg) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewOrg) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NewOrg) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecaptchaResponse

`func (o *NewOrg) GetRecaptchaResponse() string`

GetRecaptchaResponse returns the RecaptchaResponse field if non-nil, zero value otherwise.

### GetRecaptchaResponseOk

`func (o *NewOrg) GetRecaptchaResponseOk() (*string, bool)`

GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaResponse

`func (o *NewOrg) SetRecaptchaResponse(v string)`

SetRecaptchaResponse sets RecaptchaResponse field to given value.

### HasRecaptchaResponse

`func (o *NewOrg) HasRecaptchaResponse() bool`

HasRecaptchaResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


