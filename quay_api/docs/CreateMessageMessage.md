# CreateMessageMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The actual message | 
**MediaType** | **string** | The media type of the message | 
**Severity** | **string** | The severity of the message | 

## Methods

### NewCreateMessageMessage

`func NewCreateMessageMessage(content string, mediaType string, severity string, ) *CreateMessageMessage`

NewCreateMessageMessage instantiates a new CreateMessageMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageMessageWithDefaults

`func NewCreateMessageMessageWithDefaults() *CreateMessageMessage`

NewCreateMessageMessageWithDefaults instantiates a new CreateMessageMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateMessageMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateMessageMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateMessageMessage) SetContent(v string)`

SetContent sets Content field to given value.


### GetMediaType

`func (o *CreateMessageMessage) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CreateMessageMessage) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CreateMessageMessage) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetSeverity

`func (o *CreateMessageMessage) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateMessageMessage) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateMessageMessage) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


