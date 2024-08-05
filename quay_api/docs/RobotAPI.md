# \RobotAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgRobot**](RobotAPI.md#CreateOrgRobot) | **Put** /api/v1/organization/{orgname}/robots/{robot_shortname} | 
[**CreateUserRobot**](RobotAPI.md#CreateUserRobot) | **Put** /api/v1/user/robots/{robot_shortname} | 
[**DeleteOrgRobot**](RobotAPI.md#DeleteOrgRobot) | **Delete** /api/v1/organization/{orgname}/robots/{robot_shortname} | 
[**DeleteUserRobot**](RobotAPI.md#DeleteUserRobot) | **Delete** /api/v1/user/robots/{robot_shortname} | 
[**GetOrgRobot**](RobotAPI.md#GetOrgRobot) | **Get** /api/v1/organization/{orgname}/robots/{robot_shortname} | 
[**GetOrgRobotPermissions**](RobotAPI.md#GetOrgRobotPermissions) | **Get** /api/v1/organization/{orgname}/robots/{robot_shortname}/permissions | 
[**GetOrgRobots**](RobotAPI.md#GetOrgRobots) | **Get** /api/v1/organization/{orgname}/robots | 
[**GetUserRobot**](RobotAPI.md#GetUserRobot) | **Get** /api/v1/user/robots/{robot_shortname} | 
[**GetUserRobotPermissions**](RobotAPI.md#GetUserRobotPermissions) | **Get** /api/v1/user/robots/{robot_shortname}/permissions | 
[**GetUserRobots**](RobotAPI.md#GetUserRobots) | **Get** /api/v1/user/robots | 
[**RegenerateOrgRobotToken**](RobotAPI.md#RegenerateOrgRobotToken) | **Post** /api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate | 
[**RegenerateUserRobotToken**](RobotAPI.md#RegenerateUserRobotToken) | **Post** /api/v1/user/robots/{robot_shortname}/regenerate | 



## CreateOrgRobot

> CreateOrgRobot(ctx, orgname, robotShortname).Body(body).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix
	body := *openapiclient.NewCreateRobot() // CreateRobot | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.CreateOrgRobot(context.Background(), orgname, robotShortname).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.CreateOrgRobot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgRobotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CreateRobot**](CreateRobot.md) | Request body contents. | 

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


## CreateUserRobot

> CreateUserRobot(ctx, robotShortname).Body(body).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix
	body := *openapiclient.NewCreateRobot() // CreateRobot | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.CreateUserRobot(context.Background(), robotShortname).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.CreateUserRobot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRobotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateRobot**](CreateRobot.md) | Request body contents. | 

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


## DeleteOrgRobot

> DeleteOrgRobot(ctx, orgname, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.DeleteOrgRobot(context.Background(), orgname, robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.DeleteOrgRobot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgRobotRequest struct via the builder pattern


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


## DeleteUserRobot

> DeleteUserRobot(ctx, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.DeleteUserRobot(context.Background(), robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.DeleteUserRobot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRobotRequest struct via the builder pattern


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


## GetOrgRobot

> GetOrgRobot(ctx, orgname, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.GetOrgRobot(context.Background(), orgname, robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.GetOrgRobot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRobotRequest struct via the builder pattern


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


## GetOrgRobotPermissions

> GetOrgRobotPermissions(ctx, orgname, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.GetOrgRobotPermissions(context.Background(), orgname, robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.GetOrgRobotPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRobotPermissionsRequest struct via the builder pattern


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


## GetOrgRobots

> GetOrgRobots(ctx, orgname).Limit(limit).Token(token).Permissions(permissions).Execute()





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
	limit := int32(56) // int32 | If specified, the number of robots to return. (optional)
	token := true // bool | If false, the robot's token is not returned. (optional)
	permissions := true // bool | Whether to include repostories and teams in which the robots have permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.GetOrgRobots(context.Background(), orgname).Limit(limit).Token(token).Permissions(permissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.GetOrgRobots``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetOrgRobotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | If specified, the number of robots to return. | 
 **token** | **bool** | If false, the robot&#39;s token is not returned. | 
 **permissions** | **bool** | Whether to include repostories and teams in which the robots have permission. | 

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


## GetUserRobot

> GetUserRobot(ctx, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.GetUserRobot(context.Background(), robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.GetUserRobot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRobotRequest struct via the builder pattern


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


## GetUserRobotPermissions

> GetUserRobotPermissions(ctx, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.GetUserRobotPermissions(context.Background(), robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.GetUserRobotPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRobotPermissionsRequest struct via the builder pattern


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


## GetUserRobots

> GetUserRobots(ctx).Limit(limit).Token(token).Permissions(permissions).Execute()





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
	limit := int32(56) // int32 | If specified, the number of robots to return. (optional)
	token := true // bool | If false, the robot's token is not returned. (optional)
	permissions := true // bool | Whether to include repositories and teams in which the robots have permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.GetUserRobots(context.Background()).Limit(limit).Token(token).Permissions(permissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.GetUserRobots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRobotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | If specified, the number of robots to return. | 
 **token** | **bool** | If false, the robot&#39;s token is not returned. | 
 **permissions** | **bool** | Whether to include repositories and teams in which the robots have permission. | 

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


## RegenerateOrgRobotToken

> RegenerateOrgRobotToken(ctx, orgname, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.RegenerateOrgRobotToken(context.Background(), orgname, robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RegenerateOrgRobotToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateOrgRobotTokenRequest struct via the builder pattern


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


## RegenerateUserRobotToken

> RegenerateUserRobotToken(ctx, robotShortname).Execute()





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
	robotShortname := "robotShortname_example" // string | The short name for the robot, without any user or organization prefix

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RobotAPI.RegenerateUserRobotToken(context.Background(), robotShortname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RobotAPI.RegenerateUserRobotToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**robotShortname** | **string** | The short name for the robot, without any user or organization prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateUserRobotTokenRequest struct via the builder pattern


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

