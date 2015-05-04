package deezer

import (
	"encoding/json"
	"fmt"
)

type Track struct {
	Id             int    `json:"id"`
	Readable       bool   `json:"readable"`
	Title          string `json:"title"`
	Link           string `json:"link"`
	Duration       int    `json:"duration"`
	Rank           int    `json:"rank"`
	ExplicitLyrics bool   `json:"explicit_lyrics"`
	Preview        string `json:"preview"`
	Artist         Artist `json:"artist"`
	Album          Album  `json:"album"`
}

type TrackList struct {
	Data  []Track `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

type TrackRequest struct {
	Id int
}

func (t TrackRequest) Path() string {
	return fmt.Sprintf("track/%d", t.Id)
}

func (t TrackRequest) ParseJSON(d *json.Decoder) (interface{}, error) {
	data := Track{}
	err := d.Decode(&data)
	return data, err
}
