# PutServiceKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The friendly name of a service key | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | The key/value pairs of this key&#39;s metadata | [optional] 
**Expiration** | Pointer to **map[string]interface{}** | The expiration date as a unix timestamp | [optional] 

## Methods

### NewPutServiceKey

`func NewPutServiceKey() *PutServiceKey`

NewPutServiceKey instantiates a new PutServiceKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutServiceKeyWithDefaults

`func NewPutServiceKeyWithDefaults() *PutServiceKey`

NewPutServiceKeyWithDefaults instantiates a new PutServiceKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutServiceKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutServiceKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutServiceKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutServiceKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *PutServiceKey) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PutServiceKey) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PutServiceKey) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PutServiceKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpiration

`func (o *PutServiceKey) GetExpiration() map[string]interface{}`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PutServiceKey) GetExpirationOk() (*map[string]interface{}, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PutServiceKey) SetExpiration(v map[string]interface{})`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *PutServiceKey) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


