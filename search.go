package deezer

import (
	"encoding/json"
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

func (s SearchRequest) Path() string {
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

func (s SearchRequest) ParseJSON(d *json.Decoder) (interface{}, error) {
	data := SearchResponse{}
	err := d.Decode(&data)
	return data, err
}
