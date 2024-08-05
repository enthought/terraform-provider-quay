# \RepotokenAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeToken**](RepotokenAPI.md#ChangeToken) | **Put** /api/v1/repository/{repository}/tokens/{code} | 
[**CreateToken**](RepotokenAPI.md#CreateToken) | **Post** /api/v1/repository/{repository}/tokens/ | 
[**DeleteToken**](RepotokenAPI.md#DeleteToken) | **Delete** /api/v1/repository/{repository}/tokens/{code} | 
[**GetTokens**](RepotokenAPI.md#GetTokens) | **Get** /api/v1/repository/{repository}/tokens/{code} | 
[**ListRepoTokens**](RepotokenAPI.md#ListRepoTokens) | **Get** /api/v1/repository/{repository}/tokens/ | 



## ChangeToken

> ChangeToken(ctx, repository, code).Body(body).Execute()





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
	code := "code_example" // string | The token code
	body := *openapiclient.NewTokenPermission("Role_example") // TokenPermission | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepotokenAPI.ChangeToken(context.Background(), repository, code).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepotokenAPI.ChangeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**code** | **string** | The token code | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TokenPermission**](TokenPermission.md) | Request body contents. | 

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


## CreateToken

> CreateToken(ctx, repository).Body(body).Execute()





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
	body := *openapiclient.NewNewToken("FriendlyName_example") // NewToken | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepotokenAPI.CreateToken(context.Background(), repository).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepotokenAPI.CreateToken``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewToken**](NewToken.md) | Request body contents. | 

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


## DeleteToken

> DeleteToken(ctx, repository, code).Execute()





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
	code := "code_example" // string | The token code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepotokenAPI.DeleteToken(context.Background(), repository, code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepotokenAPI.DeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**code** | **string** | The token code | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


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


## GetTokens

> GetTokens(ctx, repository, code).Execute()





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
	code := "code_example" // string | The token code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepotokenAPI.GetTokens(context.Background(), repository, code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepotokenAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**code** | **string** | The token code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


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


## ListRepoTokens

> ListRepoTokens(ctx, repository).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepotokenAPI.ListRepoTokens(context.Background(), repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepotokenAPI.ListRepoTokens``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListRepoTokensRequest struct via the builder pattern


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

