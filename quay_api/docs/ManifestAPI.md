# \ManifestAPI

All URIs are relative to *https://quay.enthought.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddManifestLabel**](ManifestAPI.md#AddManifestLabel) | **Post** /api/v1/repository/{repository}/manifest/{manifestref}/labels | 
[**DeleteManifestLabel**](ManifestAPI.md#DeleteManifestLabel) | **Delete** /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid} | 
[**GetManifestLabel**](ManifestAPI.md#GetManifestLabel) | **Get** /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid} | 
[**GetRepoManifest**](ManifestAPI.md#GetRepoManifest) | **Get** /api/v1/repository/{repository}/manifest/{manifestref} | 
[**ListManifestLabels**](ManifestAPI.md#ListManifestLabels) | **Get** /api/v1/repository/{repository}/manifest/{manifestref}/labels | 



## AddManifestLabel

> AddManifestLabel(ctx, repository, manifestref).Body(body).Execute()





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
	manifestref := "manifestref_example" // string | The digest of the manifest
	body := *openapiclient.NewAddLabel("Key_example", "Value_example") // AddLabel | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.AddManifestLabel(context.Background(), repository, manifestref).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestAPI.AddManifestLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**manifestref** | **string** | The digest of the manifest | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddManifestLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddLabel**](AddLabel.md) | Request body contents. | 

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


## DeleteManifestLabel

> DeleteManifestLabel(ctx, labelid, repository, manifestref).Execute()





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
	labelid := "labelid_example" // string | The ID of the label
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	manifestref := "manifestref_example" // string | The digest of the manifest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.DeleteManifestLabel(context.Background(), labelid, repository, manifestref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestAPI.DeleteManifestLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelid** | **string** | The ID of the label | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**manifestref** | **string** | The digest of the manifest | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManifestLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetManifestLabel

> GetManifestLabel(ctx, labelid, repository, manifestref).Execute()





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
	labelid := "labelid_example" // string | The ID of the label
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	manifestref := "manifestref_example" // string | The digest of the manifest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.GetManifestLabel(context.Background(), labelid, repository, manifestref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestAPI.GetManifestLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelid** | **string** | The ID of the label | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**manifestref** | **string** | The digest of the manifest | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManifestLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetRepoManifest

> GetRepoManifest(ctx, repository, manifestref).IncludeModelcard(includeModelcard).Execute()



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
	manifestref := "manifestref_example" // string | The digest of the manifest
	includeModelcard := true // bool | If specified, include modelcard markdown from image, if any (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.GetRepoManifest(context.Background(), repository, manifestref).IncludeModelcard(includeModelcard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestAPI.GetRepoManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**manifestref** | **string** | The digest of the manifest | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepoManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeModelcard** | **bool** | If specified, include modelcard markdown from image, if any | 

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


## ListManifestLabels

> ListManifestLabels(ctx, repository, manifestref).Filter(filter).Execute()



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
	manifestref := "manifestref_example" // string | The digest of the manifest
	filter := "filter_example" // string | If specified, only labels matching the given prefix will be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.ListManifestLabels(context.Background(), repository, manifestref).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestAPI.ListManifestLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**manifestref** | **string** | The digest of the manifest | 

### Other Parameters

Other parameters are passed through a pointer to a apiListManifestLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | If specified, only labels matching the given prefix will be returned | 

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

