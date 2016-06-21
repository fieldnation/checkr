package checkr

import "time"

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
