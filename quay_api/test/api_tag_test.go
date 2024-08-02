/*
Quay Frontend

Testing TagAPIService

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

func Test_quay_api_TagAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagAPIService ChangeTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var tag string

		httpRes, err := apiClient.TagAPI.ChangeTag(context.Background(), repository, tag).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAPIService DeleteFullTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var tag string

		httpRes, err := apiClient.TagAPI.DeleteFullTag(context.Background(), repository, tag).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAPIService ListRepoTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string

		httpRes, err := apiClient.TagAPI.ListRepoTags(context.Background(), repository).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAPIService RestoreTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var repository string
		var tag string

		httpRes, err := apiClient.TagAPI.RestoreTag(context.Background(), repository, tag).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}