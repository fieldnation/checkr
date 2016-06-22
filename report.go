package checkr

import "time"

// Report represents a background check report. Depending on the selected
// package, a report can include the following screenings: SSN trace, sex
// offender search, national criminal search, county criminal searches and
// motor vehicle report.
type Report struct {
	ID                       string    `json:"id,omitempty"`
	Object                   string    `json:"object,omitempty"`
	URI                      string    `json:"uri,omitempty"`
	Status                   string    `json:"status,omitempty"`
	CreatedAt                time.Time `json:"created_at,omitempty"`
	CompletedAt              time.Time `json:"completed_at,omitempty"`
	TurnaroundTime           int       `json:"turnaround_time,omitempty"`
	DueTime                  time.Time `json:"due_time,omitempty"`
	Adjudication             string    `json:"adjudication,omitempty"`
	Package                  string    `json:"package,omitempty"`
	CandidateID              string    `json:"candidate_id,omitempty"`
	SsnTraceID               string    `json:"ssn_trace_id,omitempty"`
	SexOffenderSearchID      string    `json:"sex_offender_search_id,omitempty"`
	NationalCriminalSearchID string    `json:"national_criminal_search_id,omitempty"`
	CountyCriminalSearchIDs  []string  `json:"county_criminal_search_ids,omitempty"`
	MotorVehicleReportID     string    `json:"motor_vehicle_report_id,omitempty"`
	StateCriminalSearchIDs   []string  `json:"state_criminal_search_ids,omitempty"`
}
