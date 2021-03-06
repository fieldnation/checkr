package checkr

import "testing"

func TestCandidateCreate(t *testing.T) {
	SetAPIKey(testKey)
	c := &Candidate{
		FirstName:           "John",
		MiddleName:          "Alfred",
		LastName:            "Smith",
		Email:               "john.smith@gmail.com",
		Phone:               "5555555555",
		Zipcode:             "90401",
		DOB:                 "1970-01-22",
		SSN:                 "543-43-4645",
		DriverLicenseNumber: "F211165",
		DriverLicenseState:  "CA",
	}
	if err := c.Create(); err != nil {
		t.Error(err)
	}
}

func TestCandidatesIndex(t *testing.T) {
	SetAPIKey(testKey)

	c := &Candidates{}
	if err := c.Index(); err != nil {
		t.Error(err)
	}

	prevID := c.Data[0].ID

	if err := c.Next(); err != nil {
		t.Error(err)
		return
	}

	if err := c.Index(); err != nil {
		t.Error(err)
	}

	if prevID == c.Data[0].ID {
		t.Errorf("expected %q and %q to be different", prevID, c.Data[0].ID)
	}
}

func TestCandidateShow(t *testing.T) {
	SetAPIKey(testKey)
	c := &Candidate{ID: "e44aa283528e6fde7d542194"}
	if err := c.Show(); err != nil {
		t.Error(err)
	}
}
