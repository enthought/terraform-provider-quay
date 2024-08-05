# \LogsAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportOrgLogs**](LogsAPI.md#ExportOrgLogs) | **Post** /api/v1/organization/{orgname}/exportlogs | 
[**ExportRepoLogs**](LogsAPI.md#ExportRepoLogs) | **Post** /api/v1/repository/{repository}/exportlogs | 
[**ExportUserLogs**](LogsAPI.md#ExportUserLogs) | **Post** /api/v1/user/exportlogs | 
[**GetAggregateOrgLogs**](LogsAPI.md#GetAggregateOrgLogs) | **Get** /api/v1/organization/{orgname}/aggregatelogs | 
[**GetAggregateRepoLogs**](LogsAPI.md#GetAggregateRepoLogs) | **Get** /api/v1/repository/{repository}/aggregatelogs | 
[**GetAggregateUserLogs**](LogsAPI.md#GetAggregateUserLogs) | **Get** /api/v1/user/aggregatelogs | 
[**ListOrgLogs**](LogsAPI.md#ListOrgLogs) | **Get** /api/v1/organization/{orgname}/logs | 
[**ListRepoLogs**](LogsAPI.md#ListRepoLogs) | **Get** /api/v1/repository/{repository}/logs | 
[**ListUserLogs**](LogsAPI.md#ListUserLogs) | **Get** /api/v1/user/logs | 



## ExportOrgLogs

> ExportOrgLogs(ctx, orgname).Body(body).Endtime(endtime).Starttime(starttime).Execute()





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
	orgname := "orgname_example" // string | The name of the organization
	body := *openapiclient.NewExportLogs() // ExportLogs | Request body contents.
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.ExportOrgLogs(context.Background(), orgname).Body(body).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ExportOrgLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportOrgLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ExportLogs**](ExportLogs.md) | Request body contents. | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## ExportRepoLogs

> ExportRepoLogs(ctx, repository).Body(body).Endtime(endtime).Starttime(starttime).Execute()





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
	body := *openapiclient.NewExportLogs() // ExportLogs | Request body contents.
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.ExportRepoLogs(context.Background(), repository).Body(body).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ExportRepoLogs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportRepoLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ExportLogs**](ExportLogs.md) | Request body contents. | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## ExportUserLogs

> ExportUserLogs(ctx).Body(body).Endtime(endtime).Starttime(starttime).Execute()





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
	body := *openapiclient.NewExportLogs() // ExportLogs | Request body contents.
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.ExportUserLogs(context.Background()).Body(body).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ExportUserLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportUserLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExportLogs**](ExportLogs.md) | Request body contents. | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## GetAggregateOrgLogs

> GetAggregateOrgLogs(ctx, orgname).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()





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
	orgname := "orgname_example" // string | The name of the organization
	performer := "performer_example" // string | Username for which to filter logs. (optional)
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetAggregateOrgLogs(context.Background(), orgname).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetAggregateOrgLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregateOrgLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **performer** | **string** | Username for which to filter logs. | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## GetAggregateRepoLogs

> GetAggregateRepoLogs(ctx, repository).Endtime(endtime).Starttime(starttime).Execute()





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
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetAggregateRepoLogs(context.Background(), repository).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetAggregateRepoLogs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAggregateRepoLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## GetAggregateUserLogs

> GetAggregateUserLogs(ctx).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()





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
	performer := "performer_example" // string | Username for which to filter logs. (optional)
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.GetAggregateUserLogs(context.Background()).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetAggregateUserLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregateUserLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **performer** | **string** | Username for which to filter logs. | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## ListOrgLogs

> ListOrgLogs(ctx, orgname).NextPage(nextPage).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()





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
	orgname := "orgname_example" // string | The name of the organization
	nextPage := "nextPage_example" // string | The page token for the next page (optional)
	performer := "performer_example" // string | Username for which to filter logs. (optional)
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.ListOrgLogs(context.Background(), orgname).NextPage(nextPage).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ListOrgLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPage** | **string** | The page token for the next page | 
 **performer** | **string** | Username for which to filter logs. | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## ListRepoLogs

> ListRepoLogs(ctx, repository).NextPage(nextPage).Endtime(endtime).Starttime(starttime).Execute()





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
	nextPage := "nextPage_example" // string | The page token for the next page (optional)
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.ListRepoLogs(context.Background(), repository).NextPage(nextPage).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ListRepoLogs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListRepoLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPage** | **string** | The page token for the next page | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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


## ListUserLogs

> ListUserLogs(ctx).NextPage(nextPage).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()





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
	nextPage := "nextPage_example" // string | The page token for the next page (optional)
	performer := "performer_example" // string | Username for which to filter logs. (optional)
	endtime := "endtime_example" // string | Latest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)
	starttime := "starttime_example" // string | Earliest time for logs. Format: \"%m/%d/%Y\" in UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.ListUserLogs(context.Background()).NextPage(nextPage).Performer(performer).Endtime(endtime).Starttime(starttime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ListUserLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPage** | **string** | The page token for the next page | 
 **performer** | **string** | Username for which to filter logs. | 
 **endtime** | **string** | Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **string** | Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

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

