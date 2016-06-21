package checkr

import "time"

// Subscription represents a background check subscription.
type Subscription struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	URI           string    `json:"uri"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	CanceledAt    time.Time `json:"canceled_at"`
	Package       string    `json:"package"`
	IntervalCount int       `json:"interval_count"`
	IntervalUnit  string    `json:"interval_unit"`
	StartDate     string    `json:"start_date"`
	CandidateID   string    `json:"candidate_id"`
}
