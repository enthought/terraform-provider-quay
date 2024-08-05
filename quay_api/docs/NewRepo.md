# NewRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | **string** | Repository name | 
**Visibility** | **string** | Visibility which the repository will start with | 
**Namespace** | Pointer to **string** | Namespace in which the repository should be created. If omitted, the username of the caller is used | [optional] 
**Description** | **string** | Markdown encoded description for the repository | 

## Methods

### NewNewRepo

`func NewNewRepo(repository string, visibility string, description string, ) *NewRepo`

NewNewRepo instantiates a new NewRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRepoWithDefaults

`func NewNewRepoWithDefaults() *NewRepo`

NewNewRepoWithDefaults instantiates a new NewRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *NewRepo) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NewRepo) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NewRepo) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetVisibility

`func (o *NewRepo) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NewRepo) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NewRepo) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetNamespace

`func (o *NewRepo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NewRepo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NewRepo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *NewRepo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetDescription

`func (o *NewRepo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewRepo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewRepo) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


