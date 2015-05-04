package deezer

import (
	"fmt"
	"net/url"
	"strconv"
)

type SearchRequest struct {
	Query  string
	Strict bool
	Order  string
	Index  int
}

func (s SearchRequest) Tracks() string {
	v := url.Values{}
	v.Set("q", s.Query)

	if s.Strict {
		v.Set("strict", strconv.FormatBool(s.Strict))
	}

	if s.Order != "" {
		v.Set("order", s.Order)
	}

	if s.Index != 0 {
		v.Set("index", strconv.Itoa(s.Index))
	}

	params := v.Encode()

	return fmt.Sprintf("search/track?%s", params)
}
