# UpdateOrgQuotaLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of quota limit: \&quot;Warning\&quot; or \&quot;Reject\&quot; | [optional] 
**ThresholdPercent** | Pointer to **int32** | Quota threshold, in percent of quota | [optional] 

## Methods

### NewUpdateOrgQuotaLimit

`func NewUpdateOrgQuotaLimit() *UpdateOrgQuotaLimit`

NewUpdateOrgQuotaLimit instantiates a new UpdateOrgQuotaLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgQuotaLimitWithDefaults

`func NewUpdateOrgQuotaLimitWithDefaults() *UpdateOrgQuotaLimit`

NewUpdateOrgQuotaLimitWithDefaults instantiates a new UpdateOrgQuotaLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateOrgQuotaLimit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOrgQuotaLimit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOrgQuotaLimit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateOrgQuotaLimit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetThresholdPercent

`func (o *UpdateOrgQuotaLimit) GetThresholdPercent() int32`

GetThresholdPercent returns the ThresholdPercent field if non-nil, zero value otherwise.

### GetThresholdPercentOk

`func (o *UpdateOrgQuotaLimit) GetThresholdPercentOk() (*int32, bool)`

GetThresholdPercentOk returns a tuple with the ThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercent

`func (o *UpdateOrgQuotaLimit) SetThresholdPercent(v int32)`

SetThresholdPercent sets ThresholdPercent field to given value.

### HasThresholdPercent

`func (o *UpdateOrgQuotaLimit) HasThresholdPercent() bool`

HasThresholdPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


