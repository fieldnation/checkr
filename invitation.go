package checkr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Invitation represents a background check invitation. The candidate will
// receive an email to submit their information.
type Invitation struct {
	ID            string    `json:"id"`
	Status        string    `json:"status"`
	URI           string    `json:"uri"`
	InvitationURL string    `json:"invitation_url"`
	CompletedAt   time.Time `json:"completed_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	ExpiresAt     time.Time `json:"expires_at"`
	Package       string    `json:"package"`
	Object        string    `json:"object"`
	CreatedAt     time.Time `json:"created_at"`
	CandidateID   string    `json:"candidate_id"`
}

// Create sends an HTTP request to create a new Invitation.
// The Invitation attributes Package and CandidateID are required.
func (i Invitation) Create() error {

	// check that Invitation.Package and Invitation.CandidateID are set
	if i.Package == "" && i.CandidateID == "" {
		return errors.New("the checkr.Invitation attributes Package and CandidateID are requred to create an Invitation")
	}

	// marshal the Invitation to buffered bytes representing JSON
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)

	// create a new request
	url := v1 + "invitations"
	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}

	// set API key for authentication and authorization
	req.SetBasicAuth(apiKey, "")

	// send the HTTP request with the default Go client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// check the HTTP response status code is 201
	if resp.StatusCode != http.StatusCreated {

		// read the HTTP response body
		defer resp.Body.Close()
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// return the HTTP response body as an error
		return errors.New(string(b))
	}

	return nil
}
