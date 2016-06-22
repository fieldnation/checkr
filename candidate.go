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
	ID                          string    `json:"id"`
	Object                      string    `json:"object"`
	URI                         string    `json:"uri"`
	CreatedAt                   time.Time `json:"created_at"`
	FirstName                   string    `json:"first_name"`
	MiddleName                  string    `json:"middle_name"`
	NoMiddleName                bool      `json:"no_middle_name"`
	LastName                    string    `json:"last_name"`
	Email                       string    `json:"email"`
	Phone                       string    `json:"phone"`
	Zipcode                     string    `json:"zipcode"`
	DOB                         string    `json:"dob"`
	SSN                         string    `json:"ssn"`
	DriverLicenseNumber         string    `json:"driver_license_number"`
	DriverLicenseState          string    `json:"driver_license_state"`
	PreviousDriverLicenseNumber string    `json:"previous_driver_license_number"`
	PreviousDriverLicenseState  string    `json:"previous_driver_license_state"`
	CopyRequested               bool      `json:"copy_requested"`
	CustomID                    string    `json:"custom_id"`
	ReportIDs                   []string  `json:"report_ids"`
	GeoIDs                      []string  `json:"geo_ids"`
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
