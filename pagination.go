package checkr

// Pagination represents a list of paginated results.
type Pagination struct {
	Object       string `json:"object,omitempty"`
	NextHref     string `json:"next_href,omitempty"`
	PreviousHref string `json:"previous_href,omitempty"`
	Count        int    `json:"count,omitempty"`
}
