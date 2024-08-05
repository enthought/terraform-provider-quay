# \TriggerAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateBuildTrigger**](TriggerAPI.md#ActivateBuildTrigger) | **Post** /api/v1/repository/{repository}/trigger/{trigger_uuid}/activate | 
[**DeleteBuildTrigger**](TriggerAPI.md#DeleteBuildTrigger) | **Delete** /api/v1/repository/{repository}/trigger/{trigger_uuid} | 
[**GetBuildTrigger**](TriggerAPI.md#GetBuildTrigger) | **Get** /api/v1/repository/{repository}/trigger/{trigger_uuid} | 
[**ListBuildTriggers**](TriggerAPI.md#ListBuildTriggers) | **Get** /api/v1/repository/{repository}/trigger/ | 
[**ListTriggerRecentBuilds**](TriggerAPI.md#ListTriggerRecentBuilds) | **Get** /api/v1/repository/{repository}/trigger/{trigger_uuid}/builds | 
[**ManuallyStartBuildTrigger**](TriggerAPI.md#ManuallyStartBuildTrigger) | **Post** /api/v1/repository/{repository}/trigger/{trigger_uuid}/start | 
[**UpdateBuildTrigger**](TriggerAPI.md#UpdateBuildTrigger) | **Put** /api/v1/repository/{repository}/trigger/{trigger_uuid} | 



## ActivateBuildTrigger

> ActivateBuildTrigger(ctx, triggerUuid, repository).Body(body).Execute()





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
	triggerUuid := "triggerUuid_example" // string | The UUID of the build trigger
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	body := *openapiclient.NewBuildTriggerActivateRequest(map[string]interface{}(123)) // BuildTriggerActivateRequest | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.ActivateBuildTrigger(context.Background(), triggerUuid, repository).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.ActivateBuildTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerUuid** | **string** | The UUID of the build trigger | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateBuildTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**BuildTriggerActivateRequest**](BuildTriggerActivateRequest.md) | Request body contents. | 

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


## DeleteBuildTrigger

> DeleteBuildTrigger(ctx, triggerUuid, repository).Execute()





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
	triggerUuid := "triggerUuid_example" // string | The UUID of the build trigger
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.DeleteBuildTrigger(context.Background(), triggerUuid, repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.DeleteBuildTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerUuid** | **string** | The UUID of the build trigger | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuildTriggerRequest struct via the builder pattern


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


## GetBuildTrigger

> GetBuildTrigger(ctx, triggerUuid, repository).Execute()





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
	triggerUuid := "triggerUuid_example" // string | The UUID of the build trigger
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.GetBuildTrigger(context.Background(), triggerUuid, repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetBuildTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerUuid** | **string** | The UUID of the build trigger | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildTriggerRequest struct via the builder pattern


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


## ListBuildTriggers

> ListBuildTriggers(ctx, repository).Execute()





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
	r, err := apiClient.TriggerAPI.ListBuildTriggers(context.Background(), repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.ListBuildTriggers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListBuildTriggersRequest struct via the builder pattern


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


## ListTriggerRecentBuilds

> ListTriggerRecentBuilds(ctx, triggerUuid, repository).Limit(limit).Execute()





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
	triggerUuid := "triggerUuid_example" // string | The UUID of the build trigger
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	limit := int32(56) // int32 | The maximum number of builds to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.ListTriggerRecentBuilds(context.Background(), triggerUuid, repository).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.ListTriggerRecentBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerUuid** | **string** | The UUID of the build trigger | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTriggerRecentBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The maximum number of builds to return | 

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


## ManuallyStartBuildTrigger

> ManuallyStartBuildTrigger(ctx, triggerUuid, repository).Body(body).Execute()





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
	triggerUuid := "triggerUuid_example" // string | The UUID of the build trigger
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	body := *openapiclient.NewRunParameters() // RunParameters | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.ManuallyStartBuildTrigger(context.Background(), triggerUuid, repository).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.ManuallyStartBuildTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerUuid** | **string** | The UUID of the build trigger | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiManuallyStartBuildTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**RunParameters**](RunParameters.md) | Request body contents. | 

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


## UpdateBuildTrigger

> UpdateBuildTrigger(ctx, triggerUuid, repository).Body(body).Execute()





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
	triggerUuid := "triggerUuid_example" // string | The UUID of the build trigger
	repository := "repository_example" // string | The full path of the repository. e.g. namespace/name
	body := *openapiclient.NewUpdateTrigger(false) // UpdateTrigger | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.UpdateBuildTrigger(context.Background(), triggerUuid, repository).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.UpdateBuildTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerUuid** | **string** | The UUID of the build trigger | 
**repository** | **string** | The full path of the repository. e.g. namespace/name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuildTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateTrigger**](UpdateTrigger.md) | Request body contents. | 

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

