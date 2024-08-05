# \NamespacequotaAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOrganizationQuota**](NamespacequotaAPI.md#ChangeOrganizationQuota) | **Put** /api/v1/organization/{orgname}/quota/{quota_id} | 
[**ChangeOrganizationQuotaLimit**](NamespacequotaAPI.md#ChangeOrganizationQuotaLimit) | **Put** /api/v1/organization/{orgname}/quota/{quota_id}/limit/{limit_id} | 
[**CreateOrganizationQuota**](NamespacequotaAPI.md#CreateOrganizationQuota) | **Post** /api/v1/organization/{orgname}/quota | 
[**CreateOrganizationQuotaLimit**](NamespacequotaAPI.md#CreateOrganizationQuotaLimit) | **Post** /api/v1/organization/{orgname}/quota/{quota_id}/limit | 
[**DeleteOrganizationQuota**](NamespacequotaAPI.md#DeleteOrganizationQuota) | **Delete** /api/v1/organization/{orgname}/quota/{quota_id} | 
[**DeleteOrganizationQuotaLimit**](NamespacequotaAPI.md#DeleteOrganizationQuotaLimit) | **Delete** /api/v1/organization/{orgname}/quota/{quota_id}/limit/{limit_id} | 
[**GetOrganizationQuota**](NamespacequotaAPI.md#GetOrganizationQuota) | **Get** /api/v1/organization/{orgname}/quota/{quota_id} | 
[**GetOrganizationQuotaLimit**](NamespacequotaAPI.md#GetOrganizationQuotaLimit) | **Get** /api/v1/organization/{orgname}/quota/{quota_id}/limit/{limit_id} | 
[**GetUserQuota**](NamespacequotaAPI.md#GetUserQuota) | **Get** /api/v1/user/quota/{quota_id} | 
[**GetUserQuotaLimit**](NamespacequotaAPI.md#GetUserQuotaLimit) | **Get** /api/v1/user/quota/{quota_id}/limit/{limit_id} | 
[**ListOrganizationQuota**](NamespacequotaAPI.md#ListOrganizationQuota) | **Get** /api/v1/organization/{orgname}/quota | 
[**ListOrganizationQuotaLimit**](NamespacequotaAPI.md#ListOrganizationQuotaLimit) | **Get** /api/v1/organization/{orgname}/quota/{quota_id}/limit | 
[**ListUserQuota**](NamespacequotaAPI.md#ListUserQuota) | **Get** /api/v1/user/quota | 
[**ListUserQuotaLimit**](NamespacequotaAPI.md#ListUserQuotaLimit) | **Get** /api/v1/user/quota/{quota_id}/limit | 



## ChangeOrganizationQuota

> ChangeOrganizationQuota(ctx, orgname, quotaId).Body(body).Execute()



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
	orgname := "orgname_example" // string | 
	quotaId := "quotaId_example" // string | 
	body := *openapiclient.NewUpdateOrgQuota() // UpdateOrgQuota | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.ChangeOrganizationQuota(context.Background(), orgname, quotaId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.ChangeOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateOrgQuota**](UpdateOrgQuota.md) | Request body contents. | 

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


## ChangeOrganizationQuotaLimit

> ChangeOrganizationQuotaLimit(ctx, orgname, limitId, quotaId).Body(body).Execute()



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
	orgname := "orgname_example" // string | 
	limitId := "limitId_example" // string | 
	quotaId := "quotaId_example" // string | 
	body := *openapiclient.NewUpdateOrgQuotaLimit() // UpdateOrgQuotaLimit | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.ChangeOrganizationQuotaLimit(context.Background(), orgname, limitId, quotaId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.ChangeOrganizationQuotaLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**limitId** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeOrganizationQuotaLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateOrgQuotaLimit**](UpdateOrgQuotaLimit.md) | Request body contents. | 

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


## CreateOrganizationQuota

> CreateOrganizationQuota(ctx, orgname).Body(body).Execute()





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
	orgname := "orgname_example" // string | 
	body := *openapiclient.NewNewOrgQuota(int32(123)) // NewOrgQuota | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.CreateOrganizationQuota(context.Background(), orgname).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.CreateOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewOrgQuota**](NewOrgQuota.md) | Request body contents. | 

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


## CreateOrganizationQuotaLimit

> CreateOrganizationQuotaLimit(ctx, orgname, quotaId).Body(body).Execute()



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
	orgname := "orgname_example" // string | 
	quotaId := "quotaId_example" // string | 
	body := *openapiclient.NewNewOrgQuotaLimit("Type_example", int32(123)) // NewOrgQuotaLimit | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.CreateOrganizationQuotaLimit(context.Background(), orgname, quotaId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.CreateOrganizationQuotaLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationQuotaLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewOrgQuotaLimit**](NewOrgQuotaLimit.md) | Request body contents. | 

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


## DeleteOrganizationQuota

> DeleteOrganizationQuota(ctx, orgname, quotaId).Execute()



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
	orgname := "orgname_example" // string | 
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.DeleteOrganizationQuota(context.Background(), orgname, quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.DeleteOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationQuotaRequest struct via the builder pattern


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


## DeleteOrganizationQuotaLimit

> DeleteOrganizationQuotaLimit(ctx, orgname, limitId, quotaId).Execute()



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
	orgname := "orgname_example" // string | 
	limitId := "limitId_example" // string | 
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.DeleteOrganizationQuotaLimit(context.Background(), orgname, limitId, quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.DeleteOrganizationQuotaLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**limitId** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationQuotaLimitRequest struct via the builder pattern


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


## GetOrganizationQuota

> GetOrganizationQuota(ctx, orgname, quotaId).Execute()



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
	orgname := "orgname_example" // string | 
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.GetOrganizationQuota(context.Background(), orgname, quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.GetOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationQuotaRequest struct via the builder pattern


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


## GetOrganizationQuotaLimit

> GetOrganizationQuotaLimit(ctx, orgname, limitId, quotaId).Execute()



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
	orgname := "orgname_example" // string | 
	limitId := "limitId_example" // string | 
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.GetOrganizationQuotaLimit(context.Background(), orgname, limitId, quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.GetOrganizationQuotaLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**limitId** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationQuotaLimitRequest struct via the builder pattern


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


## GetUserQuota

> GetUserQuota(ctx, quotaId).Execute()



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
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.GetUserQuota(context.Background(), quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.GetUserQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserQuotaRequest struct via the builder pattern


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


## GetUserQuotaLimit

> GetUserQuotaLimit(ctx, limitId, quotaId).Execute()



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
	limitId := "limitId_example" // string | 
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.GetUserQuotaLimit(context.Background(), limitId, quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.GetUserQuotaLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserQuotaLimitRequest struct via the builder pattern


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


## ListOrganizationQuota

> ListOrganizationQuota(ctx, orgname).Execute()



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
	orgname := "orgname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.ListOrganizationQuota(context.Background(), orgname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.ListOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationQuotaRequest struct via the builder pattern


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


## ListOrganizationQuotaLimit

> ListOrganizationQuotaLimit(ctx, orgname, quotaId).Execute()



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
	orgname := "orgname_example" // string | 
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.ListOrganizationQuotaLimit(context.Background(), orgname, quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.ListOrganizationQuotaLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationQuotaLimitRequest struct via the builder pattern


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


## ListUserQuota

> ListUserQuota(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.ListUserQuota(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.ListUserQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserQuotaRequest struct via the builder pattern


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


## ListUserQuotaLimit

> ListUserQuotaLimit(ctx, quotaId).Execute()



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
	quotaId := "quotaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamespacequotaAPI.ListUserQuotaLimit(context.Background(), quotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamespacequotaAPI.ListUserQuotaLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserQuotaLimitRequest struct via the builder pattern


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

