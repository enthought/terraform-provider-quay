# \SearchAPI

All URIs are relative to *https://quay.enthought.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConductRepoSearch**](SearchAPI.md#ConductRepoSearch) | **Get** /api/v1/find/repositories | 
[**ConductSearch**](SearchAPI.md#ConductSearch) | **Get** /api/v1/find/all | 
[**GetMatchingEntities**](SearchAPI.md#GetMatchingEntities) | **Get** /api/v1/entities/{prefix} | 



## ConductRepoSearch

> ConductRepoSearch(ctx).IncludeUsage(includeUsage).Page(page).Query(query).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/enthought/terraform-provider-quay/quay_api"
)

func main() {
	includeUsage := true // bool | Whether to include usage metadata (optional)
	page := int32(56) // int32 | The page. (optional)
	query := "query_example" // string | The search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchAPI.ConductRepoSearch(context.Background()).IncludeUsage(includeUsage).Page(page).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.ConductRepoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConductRepoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUsage** | **bool** | Whether to include usage metadata | 
 **page** | **int32** | The page. | 
 **query** | **string** | The search query. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConductSearch

> ConductSearch(ctx).Query(query).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/enthought/terraform-provider-quay/quay_api"
)

func main() {
	query := "query_example" // string | The search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchAPI.ConductSearch(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.ConductSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConductSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchingEntities

> GetMatchingEntities(ctx, prefix).IncludeOrgs(includeOrgs).IncludeTeams(includeTeams).Namespace(namespace).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/enthought/terraform-provider-quay/quay_api"
)

func main() {
	prefix := "prefix_example" // string | 
	includeOrgs := true // bool | Whether to include orgs names. (optional)
	includeTeams := true // bool | Whether to include team names. (optional)
	namespace := "namespace_example" // string | Namespace to use when querying for org entities. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchAPI.GetMatchingEntities(context.Background(), prefix).IncludeOrgs(includeOrgs).IncludeTeams(includeTeams).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetMatchingEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchingEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeOrgs** | **bool** | Whether to include orgs names. | 
 **includeTeams** | **bool** | Whether to include team names. | 
 **namespace** | **string** | Namespace to use when querying for org entities. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

