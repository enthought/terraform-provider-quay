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

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

// Subtract two slices, c = a - b
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
