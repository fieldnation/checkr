package checkr

import "testing"

func TestCreateInvitation(t *testing.T) {
	SetAPIKey(testKey)
	i := Invitation{
		CandidateID: "e44aa283528e6fde7d542194",
		Package:     "driver_pro",
	}
	if err := i.Create(); err != nil {
		t.Error(err)
	}
}

func TestInvitationsIndex(t *testing.T) {
	SetAPIKey(testKey)
	i := Invitations{}
	if err := i.Index(); err != nil {
		t.Error(err)
	}
}
