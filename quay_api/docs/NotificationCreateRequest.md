# NotificationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | The event on which the notification will respond | 
**Method** | **string** | The method of notification (such as email or web callback) | 
**Config** | **map[string]interface{}** | JSON config information for the specific method of notification | 
**EventConfig** | **map[string]interface{}** | JSON config information for the specific event of notification | 
**Title** | Pointer to **string** | The human-readable title of the notification | [optional] 

## Methods

### NewNotificationCreateRequest

`func NewNotificationCreateRequest(event string, method string, config map[string]interface{}, eventConfig map[string]interface{}, ) *NotificationCreateRequest`

NewNotificationCreateRequest instantiates a new NotificationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationCreateRequestWithDefaults

`func NewNotificationCreateRequestWithDefaults() *NotificationCreateRequest`

NewNotificationCreateRequestWithDefaults instantiates a new NotificationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NotificationCreateRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotificationCreateRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotificationCreateRequest) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetMethod

`func (o *NotificationCreateRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationCreateRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationCreateRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetConfig

`func (o *NotificationCreateRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NotificationCreateRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NotificationCreateRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetEventConfig

`func (o *NotificationCreateRequest) GetEventConfig() map[string]interface{}`

GetEventConfig returns the EventConfig field if non-nil, zero value otherwise.

### GetEventConfigOk

`func (o *NotificationCreateRequest) GetEventConfigOk() (*map[string]interface{}, bool)`

GetEventConfigOk returns a tuple with the EventConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventConfig

`func (o *NotificationCreateRequest) SetEventConfig(v map[string]interface{})`

SetEventConfig sets EventConfig field to given value.


### GetTitle

`func (o *NotificationCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NotificationCreateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


