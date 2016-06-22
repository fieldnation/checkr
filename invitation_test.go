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

func TestInvitationShow(t *testing.T) {
	SetAPIKey(testKey)
	list := Invitations{}
	if err := list.Index(); err != nil {
		t.Error(err)
	}
	i := &Invitation{ID: list.Data[0].ID}
	if err := i.Show(); err != nil {
		t.Error(err)
	}
}

func TestInvitationsIndex(t *testing.T) {

	SetAPIKey(testKey)

	i := Invitations{}
	if err := i.Index(); err != nil {
		t.Error(err)
	}

	prevID := i.Data[0].ID

	if err := i.Next(); err != nil {
		t.Error(err)
		return
	}

	if err := i.Index(); err != nil {
		t.Error(err)
	}

	if prevID == i.Data[0].ID {
		t.Errorf("expected %q and %q to be different", prevID, i.Data[0].ID)
	}
}
