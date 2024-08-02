# CreateServiceKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | The service authenticating with this key | 
**Name** | Pointer to **string** | The friendly name of a service key | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | The key/value pairs of this key&#39;s metadata | [optional] 
**Notes** | Pointer to **string** | If specified, the extra notes for the key | [optional] 
**Expiration** | **map[string]interface{}** | The expiration date as a unix timestamp | 

## Methods

### NewCreateServiceKey

`func NewCreateServiceKey(service string, expiration map[string]interface{}, ) *CreateServiceKey`

NewCreateServiceKey instantiates a new CreateServiceKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceKeyWithDefaults

`func NewCreateServiceKeyWithDefaults() *CreateServiceKey`

NewCreateServiceKeyWithDefaults instantiates a new CreateServiceKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CreateServiceKey) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateServiceKey) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateServiceKey) SetService(v string)`

SetService sets Service field to given value.


### GetName

`func (o *CreateServiceKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateServiceKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateServiceKey) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateServiceKey) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateServiceKey) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateServiceKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotes

`func (o *CreateServiceKey) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateServiceKey) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateServiceKey) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateServiceKey) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetExpiration

`func (o *CreateServiceKey) GetExpiration() map[string]interface{}`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CreateServiceKey) GetExpirationOk() (*map[string]interface{}, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CreateServiceKey) SetExpiration(v map[string]interface{})`

SetExpiration sets Expiration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


