package checkr

import "testing"

// testKey is the key used to run integration tests against the checkr API.
const testKey = "83ebeabdec09f6670863766f792ead24d61fe3f9"

func TestSetAPIKey(t *testing.T) {
	SetAPIKey(testKey)
	if apiKey != testKey {
		t.Errorf("expected %q got %q", testKey, apiKey)
	}
}
