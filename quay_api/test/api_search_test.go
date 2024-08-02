/*
Quay Frontend

Testing SearchAPIService

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

func Test_quay_api_SearchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SearchAPIService ConductRepoSearch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SearchAPI.ConductRepoSearch(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService ConductSearch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SearchAPI.ConductSearch(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService GetMatchingEntities", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var prefix string

		httpRes, err := apiClient.SearchAPI.GetMatchingEntities(context.Background(), prefix).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}