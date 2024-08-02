# NewOrgQuotaLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of quota limit: \&quot;Warning\&quot; or \&quot;Reject\&quot; | 
**ThresholdPercent** | **int32** | Quota threshold, in percent of quota | 

## Methods

### NewNewOrgQuotaLimit

`func NewNewOrgQuotaLimit(type_ string, thresholdPercent int32, ) *NewOrgQuotaLimit`

NewNewOrgQuotaLimit instantiates a new NewOrgQuotaLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewOrgQuotaLimitWithDefaults

`func NewNewOrgQuotaLimitWithDefaults() *NewOrgQuotaLimit`

NewNewOrgQuotaLimitWithDefaults instantiates a new NewOrgQuotaLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NewOrgQuotaLimit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewOrgQuotaLimit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewOrgQuotaLimit) SetType(v string)`

SetType sets Type field to given value.


### GetThresholdPercent

`func (o *NewOrgQuotaLimit) GetThresholdPercent() int32`

GetThresholdPercent returns the ThresholdPercent field if non-nil, zero value otherwise.

### GetThresholdPercentOk

`func (o *NewOrgQuotaLimit) GetThresholdPercentOk() (*int32, bool)`

GetThresholdPercentOk returns a tuple with the ThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercent

`func (o *NewOrgQuotaLimit) SetThresholdPercent(v int32)`

SetThresholdPercent sets ThresholdPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


