package quay_api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

func Test_quay_api_RobotAPIService(t *testing.T) {
	configuration := newConfiguration()
	apiClient := quay_api.NewAPIClient(configuration)

	t.Run("Test RobotAPIService CreateOrgRobot", func(t *testing.T) {
		orgName := "create-org-robot"
		robotName := "create"
		robotDescription := "test"

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

		// Create robot
		newRobot := quay_api.CreateRobot{
			Description:          &robotDescription,
			UnstructuredMetadata: nil,
		}
		httpRes, err = apiClient.RobotAPI.CreateOrgRobot(context.Background(), orgName, robotName).Body(newRobot).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)
	})

	t.Run("Test RobotAPIService CreateUserRobot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var robotShortname string

		httpRes, err := apiClient.RobotAPI.CreateUserRobot(context.Background(), robotShortname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService DeleteOrgRobot", func(t *testing.T) {
		orgName := "delete-org-robot"
		robotName := "delete"
		robotDescription := "test"

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

		// Create robot
		newRobot := quay_api.CreateRobot{
			Description:          &robotDescription,
			UnstructuredMetadata: nil,
		}
		httpRes, err = apiClient.RobotAPI.CreateOrgRobot(context.Background(), orgName, robotName).Body(newRobot).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Delete robot
		httpRes, err = apiClient.RobotAPI.DeleteOrgRobot(context.Background(), orgName, robotName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test RobotAPIService DeleteUserRobot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var robotShortname string

		httpRes, err := apiClient.RobotAPI.DeleteUserRobot(context.Background(), robotShortname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService GetOrgRobot", func(t *testing.T) {
		orgName := "get-org-robot"
		robotName := "get"
		robotDescription := "test"

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

		// Create robot
		newRobot := quay_api.CreateRobot{
			Description:          &robotDescription,
			UnstructuredMetadata: nil,
		}
		httpRes, err = apiClient.RobotAPI.CreateOrgRobot(context.Background(), orgName, robotName).Body(newRobot).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Get robot
		httpRes, err = apiClient.RobotAPI.GetOrgRobot(context.Background(), orgName, robotName).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RobotAPIService GetOrgRobotPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var robotShortname string

		httpRes, err := apiClient.RobotAPI.GetOrgRobotPermissions(context.Background(), orgname, robotShortname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService GetOrgRobots", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.RobotAPI.GetOrgRobots(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService GetUserRobot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var robotShortname string

		httpRes, err := apiClient.RobotAPI.GetUserRobot(context.Background(), robotShortname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService GetUserRobotPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var robotShortname string

		httpRes, err := apiClient.RobotAPI.GetUserRobotPermissions(context.Background(), robotShortname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService GetUserRobots", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RobotAPI.GetUserRobots(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService RegenerateOrgRobotToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var robotShortname string

		httpRes, err := apiClient.RobotAPI.RegenerateOrgRobotToken(context.Background(), orgname, robotShortname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RobotAPIService RegenerateUserRobotToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var robotShortname string

		httpRes, err := apiClient.RobotAPI.RegenerateUserRobotToken(context.Background(), robotShortname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
