# ApiErrorDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A reference to the error type resource | 
**Title** | **string** | The title of the error. Can be used to uniquely identify the kind of error. | 
**Description** | **string** | A more detailed description of the error that may include help for fixing the issue. | 

## Methods

### NewApiErrorDescription

`func NewApiErrorDescription(type_ string, title string, description string, ) *ApiErrorDescription`

NewApiErrorDescription instantiates a new ApiErrorDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorDescriptionWithDefaults

`func NewApiErrorDescriptionWithDefaults() *ApiErrorDescription`

NewApiErrorDescriptionWithDefaults instantiates a new ApiErrorDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiErrorDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiErrorDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiErrorDescription) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *ApiErrorDescription) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiErrorDescription) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiErrorDescription) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ApiErrorDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiErrorDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiErrorDescription) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


