# RunParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | Pointer to **string** | (SCM only) If specified, the name of the branch to build. | [optional] 
**CommitSha** | Pointer to **string** | (Custom Only) If specified, the ref/SHA1 used to checkout a git repository. | [optional] 

## Methods

### NewRunParameters

`func NewRunParameters() *RunParameters`

NewRunParameters instantiates a new RunParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunParametersWithDefaults

`func NewRunParametersWithDefaults() *RunParameters`

NewRunParametersWithDefaults instantiates a new RunParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *RunParameters) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *RunParameters) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *RunParameters) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *RunParameters) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetCommitSha

`func (o *RunParameters) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *RunParameters) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *RunParameters) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *RunParameters) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


