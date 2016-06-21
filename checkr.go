package checkr

// v1 is the checkr v1 api URL
const v1 = "https://api.checkr.com/v1/"

// apiKey is the key to use to authenticate and authorize HTTP requests
var apiKey string

// SetAPIKey sets an API key.
func SetAPIKey(key string) {
	apiKey = key
}
