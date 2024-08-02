package provider

import "net/url"

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
