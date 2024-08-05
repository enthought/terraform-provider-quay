# ExportLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** | The callback URL to invoke with a link to the exported logs | [optional] 
**CallbackEmail** | Pointer to **string** | The e-mail address at which to e-mail a link to the exported logs | [optional] 

## Methods

### NewExportLogs

`func NewExportLogs() *ExportLogs`

NewExportLogs instantiates a new ExportLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportLogsWithDefaults

`func NewExportLogsWithDefaults() *ExportLogs`

NewExportLogsWithDefaults instantiates a new ExportLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *ExportLogs) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ExportLogs) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ExportLogs) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *ExportLogs) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCallbackEmail

`func (o *ExportLogs) GetCallbackEmail() string`

GetCallbackEmail returns the CallbackEmail field if non-nil, zero value otherwise.

### GetCallbackEmailOk

`func (o *ExportLogs) GetCallbackEmailOk() (*string, bool)`

GetCallbackEmailOk returns a tuple with the CallbackEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackEmail

`func (o *ExportLogs) SetCallbackEmail(v string)`

SetCallbackEmail sets CallbackEmail field to given value.

### HasCallbackEmail

`func (o *ExportLogs) HasCallbackEmail() bool`

HasCallbackEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


