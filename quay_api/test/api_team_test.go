package quay_api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

func Test_quay_api_TeamAPIService(t *testing.T) {
	configuration := newConfiguration()
	apiClient := quay_api.NewAPIClient(configuration)

	t.Run("Test TeamAPIService DeleteOrganizationTeam", func(t *testing.T) {
		orgName := "delete-org-team"
		teamName := "delete"

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
			Role:        "admin",
			Description: nil,
		}
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		// Delete team
		httpRes, err = apiClient.TeamAPI.DeleteOrganizationTeam(context.Background(), orgName, teamName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService DeleteOrganizationTeamMember", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var membername string
		var teamname string

		httpRes, err := apiClient.TeamAPI.DeleteOrganizationTeamMember(context.Background(), orgname, membername, teamname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamAPIService DeleteTeamMemberEmailInvite", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var email string
		var teamname string

		httpRes, err := apiClient.TeamAPI.DeleteTeamMemberEmailInvite(context.Background(), orgname, email, teamname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamAPIService DisableOrganizationTeamSync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var teamname string

		httpRes, err := apiClient.TeamAPI.DisableOrganizationTeamSync(context.Background(), orgname, teamname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamAPIService EnableOrganizationTeamSync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var teamname string

		httpRes, err := apiClient.TeamAPI.EnableOrganizationTeamSync(context.Background(), orgname, teamname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamAPIService GetOrganizationTeamMembers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var teamname string

		httpRes, err := apiClient.TeamAPI.GetOrganizationTeamMembers(context.Background(), orgname, teamname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamAPIService GetOrganizationTeamPermissions", func(t *testing.T) {
		orgName := "get-org-team-perm"
		teamName := "getorgteamperm"

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
			Role:        "admin",
			Description: nil,
		}
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		// Get team permissions
		httpRes, err = apiClient.TeamAPI.GetOrganizationTeamPermissions(context.Background(), orgName, teamName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService InviteTeamMemberEmail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var email string
		var teamname string

		httpRes, err := apiClient.TeamAPI.InviteTeamMemberEmail(context.Background(), orgname, email, teamname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamAPIService UpdateOrganizationTeam", func(t *testing.T) {
		orgName := "update-org-team"
		teamName := "update"

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
			Role:        "admin",
			Description: nil,
		}
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService UpdateOrganizationTeamMember", func(t *testing.T) {
		orgName := "add-org-team-member"
		teamName := "addorgteammember"
		robotShortname := "robot"
		memberName := orgName + "+" + robotShortname

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
			Role:        "admin",
			Description: nil,
		}
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeam(context.Background(), orgName, teamName).Body(newTeam).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		// Create org robot
		newCreateRobot := quay_api.NewCreateRobot()
		httpRes, err = apiClient.RobotAPI.CreateOrgRobot(context.Background(), orgName, robotShortname).Body(*newCreateRobot).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Add org team member
		httpRes, err = apiClient.TeamAPI.UpdateOrganizationTeamMember(context.Background(), orgName, memberName, teamName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

}
