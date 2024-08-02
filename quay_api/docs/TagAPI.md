# \TagAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTag**](TagAPI.md#ChangeTag) | **Put** /api/v1/repository/{repository}/tag/{tag} | 
[**DeleteFullTag**](TagAPI.md#DeleteFullTag) | **Delete** /api/v1/repository/{repository}/tag/{tag} | 
[**ListRepoTags**](TagAPI.md#ListRepoTags) | **Get** /api/v1/repository/{repository}/tag/ | 
[**RestoreTag**](TagAPI.md#RestoreTag) | **Post** /api/v1/repository/{repository}/tag/{tag}/restore | 



## ChangeTag

> ChangeTag(ctx, repository, tag).Body(body).Execute()





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
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	tag := "tag_example" // string | The name of the tag
	body := map[string]interface{}{ ... } // map[string]interface{} | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAPI.ChangeTag(context.Background(), repository, tag).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.ChangeTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**tag** | **string** | The name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFullTag

> DeleteFullTag(ctx, repository, tag).Execute()





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
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	tag := "tag_example" // string | The name of the tag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAPI.DeleteFullTag(context.Background(), repository, tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.DeleteFullTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**tag** | **string** | The name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFullTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepoTags

> ListRepoTags(ctx, repository).OnlyActiveTags(onlyActiveTags).Page(page).Limit(limit).FilterTagName(filterTagName).SpecificTag(specificTag).Execute()



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
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	onlyActiveTags := true // bool | Filter to only active tags. (optional)
	page := int32(56) // int32 | Page index for the results. Default 1. (optional)
	limit := int32(56) // int32 | Limit to the number of results to return per page. Max 100. (optional)
	filterTagName := "filterTagName_example" // string | Syntax: <op>:<name> Filters the tag names based on the operation.<op> can be 'like' or 'eq'. (optional)
	specificTag := "specificTag_example" // string | Filters the tags to the specific tag. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAPI.ListRepoTags(context.Background(), repository).OnlyActiveTags(onlyActiveTags).Page(page).Limit(limit).FilterTagName(filterTagName).SpecificTag(specificTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.ListRepoTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRepoTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onlyActiveTags** | **bool** | Filter to only active tags. | 
 **page** | **int32** | Page index for the results. Default 1. | 
 **limit** | **int32** | Limit to the number of results to return per page. Max 100. | 
 **filterTagName** | **string** | Syntax: &lt;op&gt;:&lt;name&gt; Filters the tag names based on the operation.&lt;op&gt; can be &#39;like&#39; or &#39;eq&#39;. | 
 **specificTag** | **string** | Filters the tags to the specific tag. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreTag

> RestoreTag(ctx, repository, tag).Body(body).Execute()





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
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	tag := "tag_example" // string | The name of the tag
	body := *openapiclient.NewRestoreTag() // RestoreTag | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAPI.RestoreTag(context.Background(), repository, tag).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.RestoreTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**tag** | **string** | The name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RestoreTag**](RestoreTag.md) | Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

