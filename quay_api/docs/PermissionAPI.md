# \PermissionAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTeamPermissions**](PermissionAPI.md#ChangeTeamPermissions) | **Put** /api/v1/repository/{repository}/permissions/team/{teamname} | 
[**ChangeUserPermissions**](PermissionAPI.md#ChangeUserPermissions) | **Put** /api/v1/repository/{repository}/permissions/user/{username} | 
[**DeleteTeamPermissions**](PermissionAPI.md#DeleteTeamPermissions) | **Delete** /api/v1/repository/{repository}/permissions/team/{teamname} | 
[**DeleteUserPermissions**](PermissionAPI.md#DeleteUserPermissions) | **Delete** /api/v1/repository/{repository}/permissions/user/{username} | 
[**GetTeamPermissions**](PermissionAPI.md#GetTeamPermissions) | **Get** /api/v1/repository/{repository}/permissions/team/{teamname} | 
[**GetUserPermissions**](PermissionAPI.md#GetUserPermissions) | **Get** /api/v1/repository/{repository}/permissions/user/{username} | 
[**GetUserTransitivePermission**](PermissionAPI.md#GetUserTransitivePermission) | **Get** /api/v1/repository/{repository}/permissions/user/{username}/transitive | 
[**ListRepoTeamPermissions**](PermissionAPI.md#ListRepoTeamPermissions) | **Get** /api/v1/repository/{repository}/permissions/team/ | 
[**ListRepoUserPermissions**](PermissionAPI.md#ListRepoUserPermissions) | **Get** /api/v1/repository/{repository}/permissions/user/ | 



## ChangeTeamPermissions

> ChangeTeamPermissions(ctx, repository, teamname).Body(body).Execute()





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
	teamname := "teamname_example" // string | The name of the team to which the permission applies
	body := *openapiclient.NewTeamPermission("Role_example") // TeamPermission | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionAPI.ChangeTeamPermissions(context.Background(), repository, teamname).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.ChangeTeamPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**teamname** | **string** | The name of the team to which the permission applies | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTeamPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TeamPermission**](TeamPermission.md) | Request body contents. | 

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


## ChangeUserPermissions

> ChangeUserPermissions(ctx, repository, username).Body(body).Execute()





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
	username := "username_example" // string | The username of the user to which the permission applies
	body := *openapiclient.NewUserPermission("Role_example") // UserPermission | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionAPI.ChangeUserPermissions(context.Background(), repository, username).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.ChangeUserPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**username** | **string** | The username of the user to which the permission applies | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UserPermission**](UserPermission.md) | Request body contents. | 

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


## DeleteTeamPermissions

> DeleteTeamPermissions(ctx, repository, teamname).Execute()





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
	teamname := "teamname_example" // string | The name of the team to which the permission applies

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionAPI.DeleteTeamPermissions(context.Background(), repository, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.DeleteTeamPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**teamname** | **string** | The name of the team to which the permission applies | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamPermissionsRequest struct via the builder pattern


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


## DeleteUserPermissions

> DeleteUserPermissions(ctx, repository, username).Execute()





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
	username := "username_example" // string | The username of the user to which the permission applies

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionAPI.DeleteUserPermissions(context.Background(), repository, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.DeleteUserPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**username** | **string** | The username of the user to which the permission applies | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPermissionsRequest struct via the builder pattern


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


## GetTeamPermissions

> GetTeamPermissions(ctx, repository, teamname).Execute()





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
	teamname := "teamname_example" // string | The name of the team to which the permission applies

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionAPI.GetTeamPermissions(context.Background(), repository, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.GetTeamPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**teamname** | **string** | The name of the team to which the permission applies | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamPermissionsRequest struct via the builder pattern


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


## GetUserPermissions

> GetUserPermissions(ctx, repository, username).Execute()





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
	username := "username_example" // string | The username of the user to which the permission applies

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionAPI.GetUserPermissions(context.Background(), repository, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.GetUserPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**username** | **string** | The username of the user to which the permission applies | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionsRequest struct via the builder pattern


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


## GetUserTransitivePermission

> GetUserTransitivePermission(ctx, repository, username).Execute()





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
	username := "username_example" // string | The username of the user to which the permissions apply

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionAPI.GetUserTransitivePermission(context.Background(), repository, username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.GetUserTransitivePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string** | The full path of the repository. e.g. namespace/name | 
**username** | **string** | The username of the user to which the permissions apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTransitivePermissionRequest struct via the builder pattern


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


## ListRepoTeamPermissions

> ListRepoTeamPermissions(ctx, repository).Execute()





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
	r, err := apiClient.PermissionAPI.ListRepoTeamPermissions(context.Background(), repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.ListRepoTeamPermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListRepoTeamPermissionsRequest struct via the builder pattern


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


## ListRepoUserPermissions

> ListRepoUserPermissions(ctx, repository).Execute()





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
	r, err := apiClient.PermissionAPI.ListRepoUserPermissions(context.Background(), repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.ListRepoUserPermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiListRepoUserPermissionsRequest struct via the builder pattern


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

