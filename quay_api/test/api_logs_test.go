/*
Quay Frontend

Testing LogsAPIService

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

func Test_quay_api_LogsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LogsAPIService ExportOrgLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.LogsAPI.ExportOrgLogs(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService ExportRepoLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.LogsAPI.ExportRepoLogs(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService ExportUserLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.LogsAPI.ExportUserLogs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService GetAggregateOrgLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.LogsAPI.GetAggregateOrgLogs(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService GetAggregateRepoLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.LogsAPI.GetAggregateRepoLogs(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService GetAggregateUserLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.LogsAPI.GetAggregateUserLogs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService ListOrgLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgname string

		httpRes, err := apiClient.LogsAPI.ListOrgLogs(context.Background(), orgname).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService ListRepoLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.LogsAPI.ListRepoLogs(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsAPIService ListUserLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.LogsAPI.ListUserLogs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
