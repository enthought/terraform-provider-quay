package quay_api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

func Test_quay_api_PermissionAPIService(t *testing.T) {
	configuration := newConfiguration()
	apiClient := quay_api.NewAPIClient(configuration)

	t.Run("Test PermissionAPIService ChangeTeamPermissions", func(t *testing.T) {
		orgName := "change-team-perm"
		teamName := "test"
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

		// Create team
		newTeam := quay_api.TeamDescription{
			Role:        "member",
			Description: nil,
		}
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

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

		// Change team permissions
		newTeamPermission := quay_api.NewTeamPermission("read")
		httpRes, err = apiClient.PermissionAPI.ChangeTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Body(*newTeamPermission).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PermissionAPIService ChangeUserPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var username string

		httpRes, err := apiClient.PermissionAPI.ChangeUserPermissions(context.Background(), repository, username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionAPIService DeleteTeamPermissions", func(t *testing.T) {
		orgName := "delete-team-perm"
		teamName := "test"
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

		// Create team
		newTeam := quay_api.TeamDescription{
			Role:        "member",
			Description: nil,
		}
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

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

		// Change team permissions
		newTeamPermission := quay_api.NewTeamPermission("read")
		httpRes, err = apiClient.PermissionAPI.ChangeTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Body(*newTeamPermission).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		// Delete team permissions
		httpRes, err = apiClient.PermissionAPI.DeleteTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test PermissionAPIService DeleteUserPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var username string

		httpRes, err := apiClient.PermissionAPI.DeleteUserPermissions(context.Background(), repository, username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionAPIService GetTeamPermissions", func(t *testing.T) {
		orgName := "get-team-perm"
		teamName := "test"
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

		// Create team
		newTeam := quay_api.TeamDescription{
			Role:        "member",
			Description: nil,
		}
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

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

		// Change team permissions
		newTeamPermission := quay_api.NewTeamPermission("read")
		httpRes, err = apiClient.PermissionAPI.ChangeTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Body(*newTeamPermission).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		// Get team permissions
		httpRes, err = apiClient.PermissionAPI.GetTeamPermissions(context.Background(), orgName+"/"+repoName, teamName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PermissionAPIService GetUserPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var username string

		httpRes, err := apiClient.PermissionAPI.GetUserPermissions(context.Background(), repository, username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionAPIService GetUserTransitivePermission", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var username string

		httpRes, err := apiClient.PermissionAPI.GetUserTransitivePermission(context.Background(), repository, username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionAPIService ListRepoTeamPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.PermissionAPI.ListRepoTeamPermissions(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionAPIService ListRepoUserPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.PermissionAPI.ListRepoUserPermissions(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
