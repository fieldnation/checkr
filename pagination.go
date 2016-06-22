package checkr

import (
	"errors"
	"net/url"
	"strconv"
)

// Pagination describes pagination
type Pagination interface {
	First() bool
	Last() bool
	Next() error
	Previous() error
	SetPage(page int)
	SetPerPage(perPage int)
	Page() int
	PerPage() int
	Total() int
	Clear()
}

// Paginator represents a list of paginated results.
type Paginator struct {
	Object       string `json:"object,omitempty"`
	NextHref     string `json:"next_href,omitempty"`
	PreviousHref string `json:"previous_href,omitempty"`
	Count        int    `json:"count,omitempty"`
	page         int
	perPage      int
}

// First checks if the current iterator is on the first page.
func (p Paginator) First() bool {
	return p.PreviousHref == ""
}

// Last checks if the current iterator is on the last page.
func (p Paginator) Last() bool {
	return p.NextHref == ""
}

// Next primes the iterator to use the next page.
func (p *Paginator) Next() error {

	if p.Last() {
		return errors.New("cannot advance to next page")
	}

	u, err := url.Parse(p.NextHref)
	if err != nil {
		return err
	}

	q := u.Query()
	page := q.Get("page")
	if page == "" {
		return errors.New("cannot find 'page' query param")
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	perPage := q.Get("per_page")
	if perPage == "" {
		return errors.New("cannot find 'per_page' query param")
	}

	perPageNum, err := strconv.Atoi(perPage)
	if err != nil {
		return err
	}

	p.page = pageNum
	p.perPage = perPageNum

	return nil
}

// Previous primes the iterator to use the previous page.
func (p *Paginator) Previous() error {

	if p.First() {
		return errors.New("cannot advance to previous page")
	}

	u, err := url.Parse(p.PreviousHref)
	if err != nil {
		return err
	}

	q := u.Query()
	page := q.Get("page")
	if page == "" {
		return errors.New("cannot find 'page' query param")
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	perPage := q.Get("per_page")
	if perPage == "" {
		return errors.New("cannot find 'per_page' query param")
	}

	perPageNum, err := strconv.Atoi(perPage)
	if err != nil {
		return err
	}

	p.page = pageNum
	p.perPage = perPageNum

	return nil
}

// SetPage sets the page iterator.
func (p *Paginator) SetPage(page int) {
	p.page = page
}

// SetPerPage sets the number of results to show per page.
func (p *Paginator) SetPerPage(perPage int) {
	p.perPage = perPage
}

// Page returns the current iterator page value.
func (p Paginator) Page() int {
	return p.page
}

// PerPage returns the current per page value.
func (p Paginator) PerPage() int {
	return p.perPage
}

// Total returns the total number of results.
func (p Paginator) Total() int {
	return p.Count
}

// Clear clears the page and per page values.
// This is helpful for setting default values.
func (p *Paginator) Clear() {
	p.page = -1
	p.perPage = -1
}
