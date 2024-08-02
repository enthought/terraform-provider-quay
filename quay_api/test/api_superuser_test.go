package quay_api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

func Test_quay_api_SuperuserAPIService(t *testing.T) {
	configuration := newConfiguration()
	apiClient := quay_api.NewAPIClient(configuration)

	t.Run("Test SuperuserAPIService ApproveServiceKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var kid string

		httpRes, err := apiClient.SuperuserAPI.ApproveServiceKey(context.Background(), kid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService ChangeOrganization", func(t *testing.T) {
		// We can only change the org name with this function
		orgList := []string{
			"superuser-change-org-0",
			"superuser-change-org-1",
		}

		// Ensure resource is destroyed
		defer func(organization quay_api.ApiDeleteAdminedOrganizationRequest) {
			_, err := organization.Execute()
			handleQuayAPIError(t, err)
		}(apiClient.OrganizationAPI.DeleteAdminedOrganization(context.Background(), orgList[0]))

		// Create resource
		newOrg := quay_api.NewNewOrg(orgList[0], orgList[0]+"@example.com")
		httpRes, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Body(*newOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

		// Change resource
		updateOrg := quay_api.UpdateOrg{Name: &orgList[1]}
		httpRes, err = apiClient.SuperuserAPI.ChangeOrganization(context.Background(), orgList[0]).Body(updateOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		// Verify resource
		httpRes, err = apiClient.OrganizationAPI.GetOrganization(context.Background(), orgList[1]).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		data, err := unmarshallArbitraryJSON(httpRes)
		require.Nil(t, err)
		assert.Equal(t, orgList[1], data["name"])

		// Change resource back
		updateOrg = quay_api.UpdateOrg{Name: &orgList[0]}
		httpRes, err = apiClient.SuperuserAPI.ChangeOrganization(context.Background(), orgList[1]).Body(updateOrg).Execute()
		handleQuayAPIError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SuperuserAPIService ChangeOrganizationQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var quotaId string
		var namespace string

		httpRes, err := apiClient.SuperuserAPI.ChangeOrganizationQuotaSuperUser(context.Background(), quotaId, namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService ChangeUserQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var quotaId string
		var namespace string

		httpRes, err := apiClient.SuperuserAPI.ChangeUserQuotaSuperUser(context.Background(), quotaId, namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService CreateInstallUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SuperuserAPI.CreateInstallUser(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService CreateOrganizationQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		httpRes, err := apiClient.SuperuserAPI.CreateOrganizationQuotaSuperUser(context.Background(), namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService CreateServiceKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SuperuserAPI.CreateServiceKey(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService CreateUserQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		httpRes, err := apiClient.SuperuserAPI.CreateUserQuotaSuperUser(context.Background(), namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService DeleteOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.SuperuserAPI.DeleteOrganization(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService DeleteOrganizationQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var quotaId string
		var namespace string

		httpRes, err := apiClient.SuperuserAPI.DeleteOrganizationQuotaSuperUser(context.Background(), quotaId, namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService DeleteServiceKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var kid string

		httpRes, err := apiClient.SuperuserAPI.DeleteServiceKey(context.Background(), kid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService DeleteUserQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var quotaId string
		var namespace string

		httpRes, err := apiClient.SuperuserAPI.DeleteUserQuotaSuperUser(context.Background(), quotaId, namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService GetRepoBuildLogsSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var buildUuid string

		httpRes, err := apiClient.SuperuserAPI.GetRepoBuildLogsSuperUser(context.Background(), buildUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService GetRepoBuildStatusSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var buildUuid string

		httpRes, err := apiClient.SuperuserAPI.GetRepoBuildStatusSuperUser(context.Background(), repository, buildUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService GetRepoBuildSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var buildUuid string

		httpRes, err := apiClient.SuperuserAPI.GetRepoBuildSuperUser(context.Background(), repository, buildUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService GetServiceKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var kid string

		httpRes, err := apiClient.SuperuserAPI.GetServiceKey(context.Background(), kid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService ListAllLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SuperuserAPI.ListAllLogs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService ListAllUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SuperuserAPI.ListAllUsers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService ListOrganizationQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		httpRes, err := apiClient.SuperuserAPI.ListOrganizationQuotaSuperUser(context.Background(), namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService ListServiceKeys", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SuperuserAPI.ListServiceKeys(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService ListUserQuotaSuperUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		httpRes, err := apiClient.SuperuserAPI.ListUserQuotaSuperUser(context.Background(), namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuperuserAPIService UpdateServiceKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var kid string

		httpRes, err := apiClient.SuperuserAPI.UpdateServiceKey(context.Background(), kid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
