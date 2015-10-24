package deezer

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/jmcvetta/napping"
)

var BaseUrl = "http://api.deezer.com"

func listParams(index, limit int) *url.Values {
	v := &url.Values{}
	v.Set("index", strconv.Itoa(index))
	v.Set("limit", strconv.Itoa(limit))
	return v
}

func get(path string, params *url.Values, result interface{}) error {
	url := BaseUrl + path
	errMsg := ErrorResponse{}
	_, err := napping.Get(url, params, result, &errMsg)
	if err != nil {
		return err
	}
	if errMsg.Error.Code != 0 {
		return errors.New(errMsg.Error.String())
	}
	return nil
}

func Get(url string, result interface{}) error {
	errMsg := ErrorResponse{}
	_, err := napping.Get(url, nil, result, &errMsg)
	if err != nil {
		return err
	}
	if errMsg.Error.Code != 0 {
		return errors.New(errMsg.Error.String())
	}
	return nil
}
