package deezer

import (
	"encoding/json"
	"net/http"
)

type ApiRequest interface {
	Path() string
	ParseJSON(*json.Decoder) (interface{}, error)
}

func MakeRequest(a ApiRequest) interface{} {
	url := "https://api.deezer.com/" + a.Path()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(resp.Body)
	data, err := a.ParseJSON(d)
	if err != nil {
		panic(err)
	}

	return data
}
