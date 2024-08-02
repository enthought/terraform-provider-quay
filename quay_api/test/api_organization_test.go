package quay_api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

func Test_quay_api_OrganizationAPIService(t *testing.T) {
	configuration := newConfiguration()
	apiClient := quay_api.NewAPIClient(configuration)

	t.Run("Test OrganizationAPIService ChangeOrganizationDetails", func(t *testing.T) {
		// We can only change the org email with this function
		orgName := "change-org-details"

		// Ensure resource is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName))

		// Create resource
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Change resource
		newOrgEmail := "neworgname@example.com"
		updateOrg := quay_api.UpdateOrg{Email: &newOrgEmail}
		httpRes, err = apiClient.OrganizationAPI.ChangeOrganizationDetails(context.Background(), orgName).Body(updateOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		// Verify resource
		httpRes, err = apiClient.OrganizationAPI.GetOrganization(context.Background(), orgName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		data, err := unmarshallArbitraryJSON(httpRes)
		require.Nil(t, err)
		assert.Equal(t, newOrgEmail, data["email"])
	})

	t.Run("Test OrganizationAPIService CreateOrganization", func(t *testing.T) {
		orgName := "create-org"

		// Ensure resource is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName))

		// Create resource
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)
	})

	t.Run("Test OrganizationAPIService CreateOrganizationApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.OrganizationAPI.CreateOrganizationApplication(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService DeleteAdminedOrganization", func(t *testing.T) {
		orgName := "delete-org"

		// Create resource
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Delete resource
		httpRes, err = apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test OrganizationAPIService DeleteOrganizationApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var clientId string

		httpRes, err := apiClient.OrganizationAPI.DeleteOrganizationApplication(context.Background(), orgname, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService GetApplicationInformation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string

		httpRes, err := apiClient.OrganizationAPI.GetApplicationInformation(context.Background(), clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService GetOrganization", func(t *testing.T) {
		orgName := "get-org"

		// Ensure resource is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName))

		// Create resource
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Get resource
		httpRes, err = apiClient.OrganizationAPI.GetOrganization(context.Background(), orgName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test OrganizationAPIService GetOrganizationApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var clientId string

		httpRes, err := apiClient.OrganizationAPI.GetOrganizationApplication(context.Background(), orgname, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService GetOrganizationApplications", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.OrganizationAPI.GetOrganizationApplications(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService GetOrganizationCollaborators", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.OrganizationAPI.GetOrganizationCollaborators(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService GetOrganizationMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var membername string

		httpRes, err := apiClient.OrganizationAPI.GetOrganizationMember(context.Background(), orgname, membername).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService GetOrganizationMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.OrganizationAPI.GetOrganizationMembers(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService RemoveOrganizationMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var membername string

		httpRes, err := apiClient.OrganizationAPI.RemoveOrganizationMember(context.Background(), orgname, membername).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationAPIService UpdateOrganizationApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var clientId string

		httpRes, err := apiClient.OrganizationAPI.UpdateOrganizationApplication(context.Background(), orgname, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
