package deezer

import "fmt"

type Artist struct {
	Id        int    `json:"id"`        //	The artist's Deezer id
	Name      string `json:"name"`      //	name	The artist's name
	Link      string `json:"link"`      //	The url of the artist on Deezer
	Share     string `json:"share"`     //	The share link of the artist on Deezer
	Picture   string `json:"picture"`   //	The url of the artist picture.
	NbAlbum   int    `json:"nb_album"`  //	The number of artist's albums
	NbFan     int    `json:"nb_fan"`    //	The number of artist's fans
	Radio     bool   `json:"radio"`     //	true if the artist has a smartradio
	Tracklist string `json:"tracklist"` //	API Link to the top of this artist
	Role      string `json:"role"`      //	The artist's role in a track or album
}

type ArtistList struct {
	Data  []Artist `json:"data"`
	Total int      `json:"total"`
}

type ArtistRequest struct {
	Id int
}

func (a ArtistRequest) Get() string {
	return fmt.Sprintf("artist/%d", a.Id) // Artist
}

func (a ArtistRequest) Top() string {
	return fmt.Sprintf("artist/%d/top", a.Id) // TrackList
}

func (a ArtistRequest) Albums() string {
	return fmt.Sprintf("artist/%d/albums", a.Id) // AlbumList
}

func (a ArtistRequest) Comments() string {
	return fmt.Sprintf("artist/%d/comments", a.Id) // CommentList
}

func (a ArtistRequest) Fans() string {
	return fmt.Sprintf("artist/%d/fans", a.Id) // UserList
}

func (a ArtistRequest) Related() string {
	return fmt.Sprintf("artist/%d/related", a.Id) // ArtistList
}

func (a ArtistRequest) Aadio() string {
	return fmt.Sprintf("artist/%d/radio", a.Id) // TrackList
}
