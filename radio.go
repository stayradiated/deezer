package deezer

import "fmt"

type Radio struct {
	Id          int    `json:"id,omitempty"`          //	The radio deezeid
	Title       string `json:"title,omitempty"`       //	The radio title
	Description string `json:"description,omitempty"` //	The radio title
	Share       string `json:"share,omitempty"`       //	The share link of the radio on Deezer
	Picture     string `json:"picture,omitempty"`     //	The url of the radio picture.
	Tracklist   string `json:"tracklist,omitempty"`   //	API Link to the tracklist of this radio
}

type RadioList struct {
	Data []Radio `json:"data,omitempty"`
}

type GenreRadio struct {
	Id     int     `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Radios []Radio `json:"radios,omitempty"`
}

type GenreRadioList struct {
	Data []GenreRadio `json:"data,omitempty"`
}

func GetRadios() (RadioList, error) {
	path := "/radio"
	result := RadioList{}
	err := get(path, nil, &result)
	return result, err
}

func GetRadio(id int) (Radio, error) {
	path := fmt.Sprintf("/radio/%d", id)
	result := Radio{}
	err := get(path, nil, &result)
	return result, err
}

func GetRadioGenres() (GenreRadioList, error) {
	path := "/radio/genres"
	result := GenreRadioList{}
	err := get(path, nil, &result)
	return result, err
}

func GetRadioTop() (RadioList, error) {
	path := "/radio/top"
	result := RadioList{}
	err := get(path, nil, &result)
	return result, err
}

func GetRadioTracks(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/radio/%d/tracks", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}
