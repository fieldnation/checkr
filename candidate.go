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

// Candidate represents a candidate to be screened.
type Candidate struct {
	ID                          string     `json:"id,omitempty"`
	Object                      string     `json:"object,omitempty"`
	URI                         string     `json:"uri,omitempty"`
	CreatedAt                   *time.Time `json:"created_at,omitempty"`
	FirstName                   string     `json:"first_name,omitempty"`
	MiddleName                  string     `json:"middle_name,omitempty"`
	NoMiddleName                bool       `json:"no_middle_name,omitempty"`
	LastName                    string     `json:"last_name,omitempty"`
	Email                       string     `json:"email,omitempty"`
	Phone                       string     `json:"phone,omitempty"`
	Zipcode                     string     `json:"zipcode,omitempty"`
	DOB                         string     `json:"dob,omitempty"`
	SSN                         string     `json:"ssn,omitempty"`
	DriverLicenseNumber         string     `json:"driver_license_number,omitempty"`
	DriverLicenseState          string     `json:"driver_license_state,omitempty"`
	PreviousDriverLicenseNumber string     `json:"previous_driver_license_number,omitempty"`
	PreviousDriverLicenseState  string     `json:"previous_driver_license_state,omitempty"`
	CopyRequested               bool       `json:"copy_requested,omitempty"`
	CustomID                    string     `json:"custom_id,omitempty"`
	ReportIDs                   []string   `json:"report_ids,omitempty"`
	GeoIDs                      []string   `json:"geo_ids,omitempty"`
}

// Candidates represents a listing of candidates.
type Candidates struct {
	Paginator
	Data []Candidate `json:"data"`
}

// Create sends a request to create a new Candidate.
func (c Candidate) Create() error {

	// marshal the candidate to buffered bytes representing JSON
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(b)

	// create a new request
	url := URL.String() + candidates
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

// Index shows the index endpoint list of Candidates.
func (c *Candidates) Index() error {

	// create a new request
	u, _ := url.Parse(URL.String())
	u.Path = path.Join(URL.Path, candidates)
	q := u.Query()

	// if page is set, encode it as a query parameter
	if c.Page() > 0 {
		q.Set("page", strconv.Itoa(c.Page()))
	}

	// if per page is set, encode it as a query parameter
	if c.PerPage() > 1 {
		q.Set("per_page", strconv.Itoa(c.PerPage()))
	}

	// clear page and per page values
	u.RawQuery = q.Encode()
	c.Clear()

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
	if err = json.Unmarshal(b, &c); err != nil {
		return err
	}

	// check the HTTP response status code is 200
	if resp.StatusCode != http.StatusOK {

		// return the HTTP response body as an error
		return errors.New(string(b))
	}

	return nil
}
