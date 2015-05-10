package deezer

import (
	"errors"
	"strconv"

	"github.com/jmcvetta/napping"
)

var BaseUrl = "http://api.deezer.com"

func listParams(index, limit int) *napping.Params {
	return &napping.Params{
		"index": strconv.Itoa(index),
		"limit": strconv.Itoa(limit),
	}
}

func get(path string, params *napping.Params, result interface{}) error {
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
