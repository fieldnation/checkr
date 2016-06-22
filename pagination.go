package checkr

// Pagination represents a list of paginated results.
type Pagination struct {
	Object       string `json:"object"`
	NextHref     string `json:"next_href"`
	PreviousHref string `json:"previous_href"`
	Count        int    `json:"count"`
}
