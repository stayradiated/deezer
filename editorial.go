package deezer

import "fmt"

type Editorial struct {
	Id   int    `json:"id"`   // The editorial's Deezer id
	Name string `json:"name"` // The editorial's name
}

type EditorialList struct {
	Data []Editorial
}

type Chart struct {
	Tracks    TrackList    `json:"tracks"`
	Albums    AlbumList    `json:"albums"`
	Artists   ArtistList   `json:"artists"`
	Playlists PlaylistList `json:"playlists"`
}

type EditorialRequest struct {
	Id int
}

func (e EditorialRequest) All() string {
	return "editorial" // EditorialList
}

func (e EditorialRequest) Get() string {
	return fmt.Sprintf("editorial/%d", e.Id) // Editorial
}

func (e EditorialRequest) Selection() string {
	return fmt.Sprintf("editorial/%d/selection", e.Id) // AlbumList
}

func (e EditorialRequest) Charts() string {
	return fmt.Sprintf("editorial/%d/charts", e.Id) // ChartList
}

func (e EditorialRequest) Releases() string {
	return fmt.Sprintf("editorial/%d/releases", e.Id) // AlbumList
}
