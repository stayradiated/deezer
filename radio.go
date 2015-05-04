package deezer

import "fmt"

type Radio struct {
	Id          int    `json:"id"`          //	The radio deezer ID
	Title       string `json:"title"`       //	The radio title
	Description string `json:"description"` //	The radio title
	Share       string `json:"share"`       //	The share link of the radio on Deezer
	Picture     string `json:"picture"`     //	The url of the radio picture.
	Tracklist   string `json:"tracklist"`   //	API Link to the tracklist of this radio
}

type RadioList struct {
	Data []Radio `json:"data"`
}

type GenreRadio struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Radios []Radio `json:"radios"`
}

type GenreRadioList struct {
	Data []GenreRadio `json:"data"`
}

type RadioRequest struct {
	Id int
}

func (r RadioRequest) All() string {
	return "radio" // RadioList
}

func (r RadioRequest) Get() string {
	return fmt.Sprintf("radio/%d", r.Id) // Radio
}

func (r RadioRequest) Genres() string {
	return fmt.Sprintf("radio/%d/genres", r.Id) // GenreRadioList
}

func (r RadioRequest) Top() string {
	return fmt.Sprintf("radio/%d/top", r.Id) // RadioList
}

func (r RadioRequest) Tracks() string {
	return fmt.Sprintf("radio/%d/tracks", r.Id) // TrackList
}
