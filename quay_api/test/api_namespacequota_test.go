/*
Quay Frontend

Testing NamespacequotaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package quay_api

import (
	"context"
	openapiclient "github.com/enthought/terraform-provider-quay/quay_api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_quay_api_NamespacequotaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamespacequotaAPIService ChangeOrganizationQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.ChangeOrganizationQuota(context.Background(), orgname, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService ChangeOrganizationQuotaLimit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var limitId string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.ChangeOrganizationQuotaLimit(context.Background(), orgname, limitId, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService CreateOrganizationQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.NamespacequotaAPI.CreateOrganizationQuota(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService CreateOrganizationQuotaLimit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.CreateOrganizationQuotaLimit(context.Background(), orgname, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService DeleteOrganizationQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.DeleteOrganizationQuota(context.Background(), orgname, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService DeleteOrganizationQuotaLimit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var limitId string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.DeleteOrganizationQuotaLimit(context.Background(), orgname, limitId, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService GetOrganizationQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.GetOrganizationQuota(context.Background(), orgname, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService GetOrganizationQuotaLimit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var limitId string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.GetOrganizationQuotaLimit(context.Background(), orgname, limitId, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService GetUserQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.GetUserQuota(context.Background(), quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService GetUserQuotaLimit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var limitId string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.GetUserQuotaLimit(context.Background(), limitId, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService ListOrganizationQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.NamespacequotaAPI.ListOrganizationQuota(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService ListOrganizationQuotaLimit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string
		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.ListOrganizationQuotaLimit(context.Background(), orgname, quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService ListUserQuota", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.NamespacequotaAPI.ListUserQuota(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamespacequotaAPIService ListUserQuotaLimit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var quotaId string

		httpRes, err := apiClient.NamespacequotaAPI.ListUserQuotaLimit(context.Background(), quotaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
