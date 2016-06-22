package checkr

import (
	"errors"
	"net/url"
	"strconv"
)

// Paginater describes pagination
type Paginater interface {
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

// Pagination represents a list of paginated results.
type Pagination struct {
	Object       string `json:"object,omitempty"`
	NextHref     string `json:"next_href,omitempty"`
	PreviousHref string `json:"previous_href,omitempty"`
	Count        int    `json:"count,omitempty"`
	page         int
	perPage      int
}

func (p Pagination) First() bool {
	return p.PreviousHref != ""
}

func (p Pagination) Last() bool {
	return p.NextHref != ""
}

func (p *Pagination) Next() error {

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

func (p *Pagination) Previous() error {

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

func (p *Pagination) SetPage(page int) {
	p.page = page
}

func (p *Pagination) SetPerPage(perPage int) {
	p.perPage = perPage
}

func (p Pagination) Page() int {
	return p.page
}

func (p Pagination) PerPage() int {
	return p.perPage
}

func (p Pagination) Total() int {
	return p.Count
}

func (p *Pagination) Clear() {
	p.page = -1
	p.perPage = -1
}

// http://api.checkr.com/v1/candidates?page=1&per_page=25
