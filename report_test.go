package checkr

import "testing"

func TestReportCreate(t *testing.T) {
	SetAPIKey(testKey)
	r := &Report{
		Package:     "driver_pro",
		CandidateID: "e44aa283528e6fde7d542194",
	}
	if err := r.Create(); err != nil {
		t.Error(err)
	}
}

func TestReportShow(t *testing.T) {
	SetAPIKey(testKey)
	r := Report{ID: "4722c07dd9a10c3985ae432a"}
	if err := r.Show(); err != nil {
		t.Error(err)
	}
}
