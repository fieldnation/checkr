package checkr

import "time"

// Subscription represents a background check subscription.
type Subscription struct {
	ID            string    `json:"id,omitempty"`
	Object        string    `json:"object,omitempty"`
	URI           string    `json:"uri,omitempty"`
	Status        string    `json:"status,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	CanceledAt    time.Time `json:"canceled_at,omitempty"`
	Package       string    `json:"package,omitempty"`
	IntervalCount int       `json:"interval_count,omitempty"`
	IntervalUnit  string    `json:"interval_unit,omitempty"`
	StartDate     string    `json:"start_date,omitempty"`
	CandidateID   string    `json:"candidate_id,omitempty"`
}
