# \RepositorynotificationAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepoNotification**](RepositorynotificationAPI.md#CreateRepoNotification) | **Post** /api/v1/repository/{repository}/notification/ | 
[**DeleteRepoNotification**](RepositorynotificationAPI.md#DeleteRepoNotification) | **Delete** /api/v1/repository/{repository}/notification/{uuid} | 
[**GetRepoNotification**](RepositorynotificationAPI.md#GetRepoNotification) | **Get** /api/v1/repository/{repository}/notification/{uuid} | 
[**ListRepoNotifications**](RepositorynotificationAPI.md#ListRepoNotifications) | **Get** /api/v1/repository/{repository}/notification/ | 
[**ResetRepositoryNotificationFailures**](RepositorynotificationAPI.md#ResetRepositoryNotificationFailures) | **Post** /api/v1/repository/{repository}/notification/{uuid} | 
[**TestRepoNotification**](RepositorynotificationAPI.md#TestRepoNotification) | **Post** /api/v1/repository/{repository}/notification/{uuid}/test | 



## CreateRepoNotification

> CreateRepoNotification(ctx, repository).Body(body).Execute()



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
	body := *openapiclient.NewNotificationCreateRequest("Event_example", "Method_example", map[string]interface{}(123), map[string]interface{}(123)) // NotificationCreateRequest | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositorynotificationAPI.CreateRepoNotification(context.Background(), repository).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositorynotificationAPI.CreateRepoNotification``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateRepoNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NotificationCreateRequest**](NotificationCreateRequest.md) | Request body contents. | 

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


## DeleteRepoNotification

> DeleteRepoNotification(ctx, repository, uuid).Execute()





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
	uuid := "uuid_example" // string | The UUID of the notification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositorynotificationAPI.DeleteRepoNotification(context.Background(), repository, uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositorynotificationAPI.DeleteRepoNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**uuid** | **string** | The UUID of the notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepoNotificationRequest struct via the builder pattern


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


## GetRepoNotification

> GetRepoNotification(ctx, repository, uuid).Execute()





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
	uuid := "uuid_example" // string | The UUID of the notification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositorynotificationAPI.GetRepoNotification(context.Background(), repository, uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositorynotificationAPI.GetRepoNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**uuid** | **string** | The UUID of the notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepoNotificationRequest struct via the builder pattern


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


## ListRepoNotifications

> ListRepoNotifications(ctx, repository).Execute()





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
	r, err := apiClient.RepositorynotificationAPI.ListRepoNotifications(context.Background(), repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositorynotificationAPI.ListRepoNotifications``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListRepoNotificationsRequest struct via the builder pattern


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


## ResetRepositoryNotificationFailures

> ResetRepositoryNotificationFailures(ctx, repository, uuid).Execute()





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
	uuid := "uuid_example" // string | The UUID of the notification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositorynotificationAPI.ResetRepositoryNotificationFailures(context.Background(), repository, uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositorynotificationAPI.ResetRepositoryNotificationFailures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**uuid** | **string** | The UUID of the notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetRepositoryNotificationFailuresRequest struct via the builder pattern


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


## TestRepoNotification

> TestRepoNotification(ctx, repository, uuid).Execute()





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
	uuid := "uuid_example" // string | The UUID of the notification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositorynotificationAPI.TestRepoNotification(context.Background(), repository, uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositorynotificationAPI.TestRepoNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**uuid** | **string** | The UUID of the notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRepoNotificationRequest struct via the builder pattern


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

