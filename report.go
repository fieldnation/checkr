package checkr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Report represents a background check report. Depending on the selected
// package, a report can include the following screenings: SSN trace, sex
// offender search, national criminal search, county criminal searches and
// motor vehicle report.
type Report struct {
	ID                       string     `json:"id,omitempty"`
	Object                   string     `json:"object,omitempty"`
	URI                      string     `json:"uri,omitempty"`
	Status                   string     `json:"status,omitempty"`
	CreatedAt                *time.Time `json:"created_at,omitempty"`
	CompletedAt              *time.Time `json:"completed_at,omitempty"`
	TurnaroundTime           int        `json:"turnaround_time,omitempty"`
	DueTime                  *time.Time `json:"due_time,omitempty"`
	Adjudication             string     `json:"adjudication,omitempty"`
	Package                  string     `json:"package,omitempty"`
	CandidateID              string     `json:"candidate_id,omitempty"`
	SsnTraceID               string     `json:"ssn_trace_id,omitempty"`
	SexOffenderSearchID      string     `json:"sex_offender_search_id,omitempty"`
	NationalCriminalSearchID string     `json:"national_criminal_search_id,omitempty"`
	CountyCriminalSearchIDs  []string   `json:"county_criminal_search_ids,omitempty"`
	MotorVehicleReportID     string     `json:"motor_vehicle_report_id,omitempty"`
	StateCriminalSearchIDs   []string   `json:"state_criminal_search_ids,omitempty"`
}

// Show shows one report.
func (r *Report) Show() error {

	if r.ID == "" {
		return errors.New("an id is needed to show a report")
	}

	// create a new request
	u, _ := url.Parse(URL.String())
	u.Path = path.Join(URL.Path, reports, r.ID)

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

	// unmarshal the candidate
	if err = json.Unmarshal(b, &r); err != nil {
		return err
	}

	// check the HTTP response status code is 200
	if resp.StatusCode != http.StatusOK {

		// return the HTTP response body as an error
		return errors.New(string(b))
	}

	return nil
}
