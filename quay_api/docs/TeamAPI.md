# \TeamAPI

All URIs are relative to *https://quay.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationTeam**](TeamAPI.md#DeleteOrganizationTeam) | **Delete** /api/v1/organization/{orgname}/team/{teamname} | 
[**DeleteOrganizationTeamMember**](TeamAPI.md#DeleteOrganizationTeamMember) | **Delete** /api/v1/organization/{orgname}/team/{teamname}/members/{membername} | 
[**DeleteTeamMemberEmailInvite**](TeamAPI.md#DeleteTeamMemberEmailInvite) | **Delete** /api/v1/organization/{orgname}/team/{teamname}/invite/{email} | 
[**DisableOrganizationTeamSync**](TeamAPI.md#DisableOrganizationTeamSync) | **Delete** /api/v1/organization/{orgname}/team/{teamname}/syncing | 
[**EnableOrganizationTeamSync**](TeamAPI.md#EnableOrganizationTeamSync) | **Post** /api/v1/organization/{orgname}/team/{teamname}/syncing | 
[**GetOrganizationTeamMembers**](TeamAPI.md#GetOrganizationTeamMembers) | **Get** /api/v1/organization/{orgname}/team/{teamname}/members | 
[**GetOrganizationTeamPermissions**](TeamAPI.md#GetOrganizationTeamPermissions) | **Get** /api/v1/organization/{orgname}/team/{teamname}/permissions | 
[**InviteTeamMemberEmail**](TeamAPI.md#InviteTeamMemberEmail) | **Put** /api/v1/organization/{orgname}/team/{teamname}/invite/{email} | 
[**UpdateOrganizationTeam**](TeamAPI.md#UpdateOrganizationTeam) | **Put** /api/v1/organization/{orgname}/team/{teamname} | 
[**UpdateOrganizationTeamMember**](TeamAPI.md#UpdateOrganizationTeamMember) | **Put** /api/v1/organization/{orgname}/team/{teamname}/members/{membername} | 



## DeleteOrganizationTeam

> DeleteOrganizationTeam(ctx, orgname, teamname).Execute()





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
	teamname := "teamname_example" // string | The name of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteOrganizationTeam(context.Background(), orgname, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteOrganizationTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationTeamRequest struct via the builder pattern


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


## DeleteOrganizationTeamMember

> DeleteOrganizationTeamMember(ctx, orgname, membername, teamname).Execute()





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
	membername := "membername_example" // string | The username of the team member
	teamname := "teamname_example" // string | The name of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteOrganizationTeamMember(context.Background(), orgname, membername, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteOrganizationTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**membername** | **string** | The username of the team member | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationTeamMemberRequest struct via the builder pattern


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


## DeleteTeamMemberEmailInvite

> DeleteTeamMemberEmailInvite(ctx, orgname, email, teamname).Execute()





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
	email := "email_example" // string | 
	teamname := "teamname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteTeamMemberEmailInvite(context.Background(), orgname, email, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamMemberEmailInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**email** | **string** |  | 
**teamname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamMemberEmailInviteRequest struct via the builder pattern


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


## DisableOrganizationTeamSync

> DisableOrganizationTeamSync(ctx, orgname, teamname).Execute()



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
	teamname := "teamname_example" // string | The name of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DisableOrganizationTeamSync(context.Background(), orgname, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DisableOrganizationTeamSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableOrganizationTeamSyncRequest struct via the builder pattern


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


## EnableOrganizationTeamSync

> EnableOrganizationTeamSync(ctx, orgname, teamname).Execute()



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
	teamname := "teamname_example" // string | The name of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.EnableOrganizationTeamSync(context.Background(), orgname, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.EnableOrganizationTeamSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableOrganizationTeamSyncRequest struct via the builder pattern


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


## GetOrganizationTeamMembers

> GetOrganizationTeamMembers(ctx, orgname, teamname).IncludePending(includePending).Execute()





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
	teamname := "teamname_example" // string | The name of the team
	includePending := true // bool | Whether to include pending members (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.GetOrganizationTeamMembers(context.Background(), orgname, teamname).IncludePending(includePending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetOrganizationTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includePending** | **bool** | Whether to include pending members | 

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


## GetOrganizationTeamPermissions

> GetOrganizationTeamPermissions(ctx, orgname, teamname).Execute()





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
	teamname := "teamname_example" // string | The name of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.GetOrganizationTeamPermissions(context.Background(), orgname, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetOrganizationTeamPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationTeamPermissionsRequest struct via the builder pattern


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


## InviteTeamMemberEmail

> InviteTeamMemberEmail(ctx, orgname, email, teamname).Execute()





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
	email := "email_example" // string | 
	teamname := "teamname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.InviteTeamMemberEmail(context.Background(), orgname, email, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.InviteTeamMemberEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** |  | 
**email** | **string** |  | 
**teamname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteTeamMemberEmailRequest struct via the builder pattern


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


## UpdateOrganizationTeam

> UpdateOrganizationTeam(ctx, orgname, teamname).Body(body).Execute()





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
	teamname := "teamname_example" // string | The name of the team
	body := *openapiclient.NewTeamDescription("Role_example") // TeamDescription | Request body contents.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgname, teamname).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.UpdateOrganizationTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TeamDescription**](TeamDescription.md) | Request body contents. | 

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


## UpdateOrganizationTeamMember

> UpdateOrganizationTeamMember(ctx, orgname, membername, teamname).Execute()





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
	membername := "membername_example" // string | The username of the team member
	teamname := "teamname_example" // string | The name of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.UpdateOrganizationTeamMember(context.Background(), orgname, membername, teamname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.UpdateOrganizationTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgname** | **string** | The name of the organization | 
**membername** | **string** | The username of the team member | 
**teamname** | **string** | The name of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationTeamMemberRequest struct via the builder pattern


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

