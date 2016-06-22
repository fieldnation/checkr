package checkr

import "testing"

func TestReportShow(t *testing.T) {
	SetAPIKey(testKey)
	r := Report{ID: "4722c07dd9a10c3985ae432a"}
	if err := r.Show(); err != nil {
		t.Error(err)
	}
}
