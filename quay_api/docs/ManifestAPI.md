# \ManifestAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddManifestLabel**](ManifestAPI.md#AddManifestLabel) | **Post** /api/v1/repository/{repository}/manifest/{manifestref}/labels | 
[**DeleteManifestLabel**](ManifestAPI.md#DeleteManifestLabel) | **Delete** /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid} | 
[**GetManifestLabel**](ManifestAPI.md#GetManifestLabel) | **Get** /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid} | 
[**GetRepoManifest**](ManifestAPI.md#GetRepoManifest) | **Get** /api/v1/repository/{repository}/manifest/{manifestref} | 
[**ListManifestLabels**](ManifestAPI.md#ListManifestLabels) | **Get** /api/v1/repository/{repository}/manifest/{manifestref}/labels | 



## AddManifestLabel

> AddManifestLabel(ctx, manifestref, repository).Body(body).Execute()





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
	manifestref := "manifestref_example" // string | The digest of the manifest
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	body := *openapiclient.NewAddLabel("Key_example", "Value_example") // AddLabel | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.AddManifestLabel(context.Background(), manifestref, repository).Body(body).Execute()
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
**manifestref** | **string** | The digest of the manifest | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

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

> DeleteManifestLabel(ctx, manifestref, repository, labelid).Execute()





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
	manifestref := "manifestref_example" // string | The digest of the manifest
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	labelid := "labelid_example" // string | The ID of the label

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.DeleteManifestLabel(context.Background(), manifestref, repository, labelid).Execute()
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
**manifestref** | **string** | The digest of the manifest | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**labelid** | **string** | The ID of the label | 

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

> GetManifestLabel(ctx, manifestref, repository, labelid).Execute()





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
	manifestref := "manifestref_example" // string | The digest of the manifest
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	labelid := "labelid_example" // string | The ID of the label

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.GetManifestLabel(context.Background(), manifestref, repository, labelid).Execute()
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
**manifestref** | **string** | The digest of the manifest | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**labelid** | **string** | The ID of the label | 

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

> GetRepoManifest(ctx, manifestref, repository).Execute()



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
	manifestref := "manifestref_example" // string | The digest of the manifest
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.GetRepoManifest(context.Background(), manifestref, repository).Execute()
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
**manifestref** | **string** | The digest of the manifest | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepoManifestRequest struct via the builder pattern


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


## ListManifestLabels

> ListManifestLabels(ctx, manifestref, repository).Filter(filter).Execute()



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
	manifestref := "manifestref_example" // string | The digest of the manifest
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	filter := "filter_example" // string | If specified, only labels matching the given prefix will be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManifestAPI.ListManifestLabels(context.Background(), manifestref, repository).Filter(filter).Execute()
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
**manifestref** | **string** | The digest of the manifest | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

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

