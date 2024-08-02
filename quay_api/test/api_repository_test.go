/*
Quay Frontend

Testing RepositoryAPIService

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

func Test_quay_api_RepositoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoryAPIService ChangeRepoVisibility", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.RepositoryAPI.ChangeRepoVisibility(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService CreateRepo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RepositoryAPI.CreateRepo(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService DeleteRepository", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.RepositoryAPI.DeleteRepository(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService GetRepo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.RepositoryAPI.GetRepo(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService ListRepos", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RepositoryAPI.ListRepos(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService UpdateRepo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.RepositoryAPI.UpdateRepo(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}