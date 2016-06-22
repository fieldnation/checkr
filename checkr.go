package checkr

import (
	"net/url"
)

// v1 is the checkr v1 api URL
const (
	scheme      = "https"
	host        = "api.checkr.com"
	version     = "v1"
	candidates  = "/candidates"
	invitations = "/invitations"
)

var (
	// key is the key to use to authenticate and authorize HTTP requests
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
