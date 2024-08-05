package quay_api

import (
	"encoding/json"
	"errors"
	openapiclient "github.com/enthought/terraform-provider-quay/quay_api"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"os"
	"testing"
)

func handleQuayAPIError(t *testing.T, err error) {
	var apiErr *openapiclient.GenericOpenAPIError
	if err != nil {
		if errors.As(err, &apiErr) {
			require.Nil(t, err, string(apiErr.Body()))
		} else {
			require.Nil(t, err)
		}
	}
}

func newConfiguration() *openapiclient.Configuration {
	quayURL := os.Getenv("QUAY_URL")
	configuration := &openapiclient.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: openapiclient.ServerConfigurations{
			{
				URL:         quayURL,
				Description: "No description provided",
			},
		},
		OperationServers: map[string]openapiclient.ServerConfigurations{},
	}

	authToken := os.Getenv("QUAY_TOKEN")
	configuration.AddDefaultHeader("Authorization", "Bearer "+authToken)

	return configuration
}

func unmarshallArbitraryJSON(httpRes *http.Response) (map[string]interface{}, error) {
	var data map[string]interface{}

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
