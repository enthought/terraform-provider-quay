package quay_api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

func Test_quay_api_RepositoryAPIService(t *testing.T) {
	configuration := newConfiguration()
	apiClient := quay_api.NewAPIClient(configuration)

	t.Run("Test RepositoryAPIService ChangeRepoVisibility", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.RepositoryAPI.ChangeRepoVisibility(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService CreateRepo", func(t *testing.T) {
		orgName := "create-repo"
		repoName := "test"

		// Ensure org is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName))

		// Create org
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Create repo
		newRepo := quay_api.NewRepo{
			Repository:  repoName,
			Visibility:  "private",
			Namespace:   &orgName,
			Description: "test",
		}
		httpRes, err = apiClient.RepositoryAPI.CreateRepo(context.Background()).Body(newRepo).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)
	})

	t.Run("Test RepositoryAPIService DeleteRepository", func(t *testing.T) {
		orgName := "delete-repo"
		repoName := "test"

		// Ensure org is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName))

		// Create org
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Create repo
		newRepo := quay_api.NewRepo{
			Repository:  repoName,
			Visibility:  "private",
			Namespace:   &orgName,
			Description: "test",
		}
		httpRes, err = apiClient.RepositoryAPI.CreateRepo(context.Background()).Body(newRepo).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Delete repo
		httpRes, err = apiClient.RepositoryAPI.DeleteRepository(context.Background(), orgName+"/"+repoName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test RepositoryAPIService GetRepo", func(t *testing.T) {
		orgName := "get-repo"
		repoName := "test"

		// Ensure org is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName))

		// Create org
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Create repo
		newRepo := quay_api.NewRepo{
			Repository:  repoName,
			Visibility:  "private",
			Namespace:   &orgName,
			Description: "test",
		}
		httpRes, err = apiClient.RepositoryAPI.CreateRepo(context.Background()).Body(newRepo).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Get repo
		httpRes, err = apiClient.RepositoryAPI.GetRepo(context.Background(), orgName+"/"+repoName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RepositoryAPIService ListRepos", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RepositoryAPI.ListRepos(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService UpdateRepo", func(t *testing.T) {
		orgName := "update-repo"
		repoName := "test"

		// Ensure org is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgName))

		// Create org
		newOrg := quay_api.NewNewOrg(orgName, orgName+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Create repo
		newRepo := quay_api.NewRepo{
			Repository:  repoName,
			Visibility:  "private",
			Namespace:   &orgName,
			Description: "test",
		}
		httpRes, err = apiClient.RepositoryAPI.CreateRepo(context.Background()).Body(newRepo).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Update repo
		newRepoUpdate := quay_api.NewRepoUpdate("test2")
		httpRes, err = apiClient.RepositoryAPI.UpdateRepo(context.Background(), orgName+"/"+repoName).Body(*newRepoUpdate).Execute()
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
