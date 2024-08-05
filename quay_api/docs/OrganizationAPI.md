# \OrganizationAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOrganizationDetails**](OrganizationAPI.md#ChangeOrganizationDetails) | **Put** /api/v1/organization/{orgname} | 
[**CreateOrganization**](OrganizationAPI.md#CreateOrganization) | **Post** /api/v1/organization/ | 
[**CreateOrganizationApplication**](OrganizationAPI.md#CreateOrganizationApplication) | **Post** /api/v1/organization/{orgname}/applications | 
[**DeleteAdminedOrganization**](OrganizationAPI.md#DeleteAdminedOrganization) | **Delete** /api/v1/organization/{orgname} | 
[**DeleteOrganizationApplication**](OrganizationAPI.md#DeleteOrganizationApplication) | **Delete** /api/v1/organization/{orgname}/applications/{client_id} | 
[**GetApplicationInformation**](OrganizationAPI.md#GetApplicationInformation) | **Get** /api/v1/app/{client_id} | 
[**GetOrganization**](OrganizationAPI.md#GetOrganization) | **Get** /api/v1/organization/{orgname} | 
[**GetOrganizationApplication**](OrganizationAPI.md#GetOrganizationApplication) | **Get** /api/v1/organization/{orgname}/applications/{client_id} | 
[**GetOrganizationApplications**](OrganizationAPI.md#GetOrganizationApplications) | **Get** /api/v1/organization/{orgname}/applications | 
[**GetOrganizationCollaborators**](OrganizationAPI.md#GetOrganizationCollaborators) | **Get** /api/v1/organization/{orgname}/collaborators | 
[**GetOrganizationMember**](OrganizationAPI.md#GetOrganizationMember) | **Get** /api/v1/organization/{orgname}/members/{membername} | 
[**GetOrganizationMembers**](OrganizationAPI.md#GetOrganizationMembers) | **Get** /api/v1/organization/{orgname}/members | 
[**RemoveOrganizationMember**](OrganizationAPI.md#RemoveOrganizationMember) | **Delete** /api/v1/organization/{orgname}/members/{membername} | 
[**UpdateOrganizationApplication**](OrganizationAPI.md#UpdateOrganizationApplication) | **Put** /api/v1/organization/{orgname}/applications/{client_id} | 



## ChangeOrganizationDetails

> ChangeOrganizationDetails(ctx, orgname).Body(body).Execute()





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
	body := *openapiclient.NewUpdateOrg() // UpdateOrg | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.ChangeOrganizationDetails(context.Background(), orgname).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ChangeOrganizationDetails``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChangeOrganizationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateOrg**](UpdateOrg.md) | Request body contents. | 

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


## CreateOrganization

> CreateOrganization(ctx).Body(body).Execute()





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
	body := *openapiclient.NewNewOrg("Name_example", "Email_example") // NewOrg | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewOrg**](NewOrg.md) | Request body contents. | 

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


## CreateOrganizationApplication

> CreateOrganizationApplication(ctx, orgname).Body(body).Execute()





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
	body := *openapiclient.NewNewApp("Name_example") // NewApp | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.CreateOrganizationApplication(context.Background(), orgname).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOrganizationApplication``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateOrganizationApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewApp**](NewApp.md) | Request body contents. | 

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


## DeleteAdminedOrganization

> DeleteAdminedOrganization(ctx, orgname).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DeleteAdminedOrganization``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAdminedOrganizationRequest struct via the builder pattern


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


## DeleteOrganizationApplication

> DeleteOrganizationApplication(ctx, orgname, clientId).Execute()





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
	clientId := "clientId_example" // string | The OAuth client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.DeleteOrganizationApplication(context.Background(), orgname, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DeleteOrganizationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**clientId** | **string** | The OAuth client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplicationRequest struct via the builder pattern


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


## GetApplicationInformation

> GetApplicationInformation(ctx, clientId).Execute()





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
	clientId := "clientId_example" // string | The OAuth client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetApplicationInformation(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetApplicationInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The OAuth client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInformationRequest struct via the builder pattern


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


## GetOrganization

> GetOrganization(ctx, orgname).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetOrganization(context.Background(), orgname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganization``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


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


## GetOrganizationApplication

> GetOrganizationApplication(ctx, orgname, clientId).Execute()





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
	clientId := "clientId_example" // string | The OAuth client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetOrganizationApplication(context.Background(), orgname, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**clientId** | **string** | The OAuth client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplicationRequest struct via the builder pattern


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


## GetOrganizationApplications

> GetOrganizationApplications(ctx, orgname).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetOrganizationApplications(context.Background(), orgname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationApplications``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetOrganizationApplicationsRequest struct via the builder pattern


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


## GetOrganizationCollaborators

> GetOrganizationCollaborators(ctx, orgname).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetOrganizationCollaborators(context.Background(), orgname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationCollaborators``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetOrganizationCollaboratorsRequest struct via the builder pattern


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


## GetOrganizationMember

> GetOrganizationMember(ctx, orgname, membername).Execute()





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
	membername := "membername_example" // string | The username of the organization member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetOrganizationMember(context.Background(), orgname, membername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**membername** | **string** | The username of the organization member | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationMemberRequest struct via the builder pattern


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


## GetOrganizationMembers

> GetOrganizationMembers(ctx, orgname).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.GetOrganizationMembers(context.Background(), orgname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationMembers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetOrganizationMembersRequest struct via the builder pattern


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


## RemoveOrganizationMember

> RemoveOrganizationMember(ctx, orgname, membername).Execute()





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
	membername := "membername_example" // string | The username of the organization member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.RemoveOrganizationMember(context.Background(), orgname, membername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveOrganizationMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**membername** | **string** | The username of the organization member | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationMemberRequest struct via the builder pattern


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


## UpdateOrganizationApplication

> UpdateOrganizationApplication(ctx, orgname, clientId).Body(body).Execute()





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
	clientId := "clientId_example" // string | The OAuth client ID
	body := *openapiclient.NewUpdateApp("Name_example", "RedirectUri_example", "ApplicationUri_example") // UpdateApp | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAPI.UpdateOrganizationApplication(context.Background(), orgname, clientId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateOrganizationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**clientId** | **string** | The OAuth client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateApp**](UpdateApp.md) | Request body contents. | 

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

