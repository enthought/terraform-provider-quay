# \AppspecifictokensAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppToken**](AppspecifictokensAPI.md#CreateAppToken) | **Post** /api/v1/user/apptoken | 
[**GetAppToken**](AppspecifictokensAPI.md#GetAppToken) | **Get** /api/v1/user/apptoken/{token_uuid} | 
[**ListAppTokens**](AppspecifictokensAPI.md#ListAppTokens) | **Get** /api/v1/user/apptoken | 
[**RevokeAppToken**](AppspecifictokensAPI.md#RevokeAppToken) | **Delete** /api/v1/user/apptoken/{token_uuid} | 



## CreateAppToken

> CreateAppToken(ctx).Body(body).Execute()





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
	body := *openapiclient.NewNewToken("FriendlyName_example") // NewToken | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppspecifictokensAPI.CreateAppToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppspecifictokensAPI.CreateAppToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppTokenRequest struct via the builder pattern


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


## GetAppToken

> GetAppToken(ctx, tokenUuid).Execute()





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
	tokenUuid := "tokenUuid_example" // string | The uuid of the app specific token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppspecifictokensAPI.GetAppToken(context.Background(), tokenUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppspecifictokensAPI.GetAppToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenUuid** | **string** | The uuid of the app specific token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppTokenRequest struct via the builder pattern


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


## ListAppTokens

> ListAppTokens(ctx).Expiring(expiring).Execute()





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
	expiring := true // bool | If true, only returns those tokens expiring soon (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppspecifictokensAPI.ListAppTokens(context.Background()).Expiring(expiring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppspecifictokensAPI.ListAppTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expiring** | **bool** | If true, only returns those tokens expiring soon | 

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


## RevokeAppToken

> RevokeAppToken(ctx, tokenUuid).Execute()





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
	tokenUuid := "tokenUuid_example" // string | The uuid of the app specific token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppspecifictokensAPI.RevokeAppToken(context.Background(), tokenUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppspecifictokensAPI.RevokeAppToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenUuid** | **string** | The uuid of the app specific token | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAppTokenRequest struct via the builder pattern


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

