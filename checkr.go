// Package checkr provides a client for the checkr API.
package checkr

import (
	"net/url"
)

const (
	// defaults used to build the checkr API URL, these settings are
	// customizable and can be overridden via checkr.URL.
	scheme  = "https"
	host    = "api.checkr.com"
	version = "v1"

	// endpoints used by different checkr API versions
	candidates  = "/candidates"
	invitations = "/invitations"
	reports     = "/reports"
)

var (
	// apiKey is the key to use to authenticate and authorize HTTP requests
	apiKey string

	// URL is the url for the specific checkr API version chosen
	URL *url.URL
)

// the init method defaults to the latest checkr API version
func init() {
	URL = &url.URL{Scheme: scheme, Host: host, Path: version}
}

// SetAPIKey sets an API key.
func SetAPIKey(key string) {
	apiKey = key
}
