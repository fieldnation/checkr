package checkr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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

// CandidateList represents a listing of candidates.
type CandidateList struct {
	Pagination
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

// Candidates sends a request to create a new Candidate.
//
// https://api.checkr.com/v1/candidates?page=2&per_page=25
//
func Candidates() (*CandidateList, error) {

	// create a new request
	url := URL.String() + candidates
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// set API key for authentication and authorization
	req.SetBasicAuth(apiKey, "")

	// send the HTTP request with the default Go client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read the HTTP response body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal the candidate list
	var list CandidateList
	if err = json.Unmarshal(b, &list); err != nil {
		return nil, err
	}

	// check the HTTP response status code is 200
	if resp.StatusCode != http.StatusOK {

		// return the HTTP response body as an error
		return nil, errors.New(string(b))
	}

	return &list, nil
}
