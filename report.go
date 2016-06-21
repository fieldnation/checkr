package checkr

import "time"

// Report represents a background check report. Depending on the selected
// package, a report can include the following screenings: SSN trace, sex
// offender search, national criminal search, county criminal searches and
// motor vehicle report.
type Report struct {
	ID                       string    `json:"id"`
	Object                   string    `json:"object"`
	URI                      string    `json:"uri"`
	Status                   string    `json:"status"`
	CreatedAt                time.Time `json:"created_at"`
	CompletedAt              time.Time `json:"completed_at"`
	TurnaroundTime           int       `json:"turnaround_time"`
	DueTime                  time.Time `json:"due_time"`
	Adjudication             string    `json:"adjudication"`
	Package                  string    `json:"package"`
	CandidateID              string    `json:"candidate_id"`
	SsnTraceID               string    `json:"ssn_trace_id"`
	SexOffenderSearchID      string    `json:"sex_offender_search_id"`
	NationalCriminalSearchID string    `json:"national_criminal_search_id"`
	CountyCriminalSearchIDs  []string  `json:"county_criminal_search_ids"`
	MotorVehicleReportID     string    `json:"motor_vehicle_report_id"`
	StateCriminalSearchIDs   []string  `json:"state_criminal_search_ids"`
}
