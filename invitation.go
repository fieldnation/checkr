package checkr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

// Invitation represents a background check invitation. The candidate will
// receive an email to submit their information.
type Invitation struct {
	ID            string     `json:"id,omitempty"`
	Status        string     `json:"status,omitempty"`
	URI           string     `json:"uri,omitempty"`
	InvitationURL string     `json:"invitation_url,omitempty"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
	Package       string     `json:"package,omitempty"`
	Object        string     `json:"object,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	CandidateID   string     `json:"candidate_id,omitempty"`
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
	body := bytes.NewBuffer(b)

	// create a new request
	url := URL.String() + invitations
	req, err := http.NewRequest(http.MethodPost, url, body)
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

// Invitations represents a listing of invitations.
type Invitations struct {
	Paginator
	Data []Invitation `json:"data"`
}

// Index shows the index list of Invitations.
func (i *Invitations) Index() error {

	// create a new request
	u, _ := url.Parse(URL.String())
	u.Path = path.Join(URL.Path, invitations)
	q := u.Query()

	// if page is set, encode it as a query parameter
	if i.Page() > 0 {
		q.Set("page", strconv.Itoa(i.Page()))
	}

	// if per page is set, encode it as a query parameter
	if i.PerPage() > 1 {
		q.Set("per_page", strconv.Itoa(i.PerPage()))
	}

	// clear page and per page values
	u.RawQuery = q.Encode()
	i.Clear()

	// create a new request
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
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

	// read the HTTP response body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// unmarshal the candidate list
	if err = json.Unmarshal(b, &i); err != nil {
		return err
	}

	// check the HTTP response status code is 200
	if resp.StatusCode != http.StatusOK {

		// return the HTTP response body as an error
		return errors.New(string(b))
	}

	return nil
}
