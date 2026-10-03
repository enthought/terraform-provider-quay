package provider

import (
	"errors"
	"net/url"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

func handleQuayAPIError(err error) string {
	var apiErr *quay_api.GenericOpenAPIError
	errDetail := ""
	if errors.As(err, &apiErr) {
		errDetail = string(apiErr.Body())
	} else {
		errDetail = err.Error()
	}
	return errDetail
}

// apiErrorModel extracts the parsed ApiError from a client error. The generated
// client stores the decoded error as either an ApiError value or an *ApiError
// pointer, so both forms are handled.
func apiErrorModel(err error) (*quay_api.ApiError, bool) {
	var apiErr *quay_api.GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		return nil, false
	}
	switch m := apiErr.Model().(type) {
	case quay_api.ApiError:
		return &m, true
	case *quay_api.ApiError:
		return m, true
	}
	return nil, false
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

// Subtract two slices, c = a - b.
func subtractStringSlice(a []string, b []string) []string {
	bMap := make(map[string]struct{})
	for _, v := range b {
		bMap[v] = struct{}{}
	}

	var c []string
	for _, v := range a {
		if _, found := bMap[v]; !found {
			c = append(c, v)
		}
	}

	return c
}
