package deezer

import (
	"encoding/json"
	"fmt"
)

type Artist struct {
	ID        int    `json:"id,omitempty"`        //	The artist's Deezer id
	Name      string `json:"name,omitempty"`      //	name	The artist's name
	Link      string `json:"link,omitempty"`      //	The url of the artist on Deezer
	Share     string `json:"share,omitempty"`     //	The share link of the artist on Deezer
	Picture   string `json:"picture,omitempty"`   //	The url of the artist picture.
	NbAlbum   int    `json:"nb_album,omitempty"`  //	The number of artist's albums
	NbFan     int    `json:"nb_fan,omitempty"`    //	The number of artist's fans
	Radio     bool   `json:"radio,omitempty"`     //	true if the artist has a smartradio
	Tracklist string `json:"tracklist,omitempty"` //	API Link to the top of this artist
	Role      string `json:"role,omitempty"`      //	The artist's role in a track or album
}

type ExtendedArtistList struct {
	Data  []Artist `json:"data,omitempty"`
	Total int      `json:"total,omitempty"`
	Next  string   `json:"next,omitempty"`
}

type ArtistList []Artist

func (a *ArtistList) UnmarshalJSON(data []byte) error {
	extendedArtistList := ExtendedArtistList{}
	if err := json.Unmarshal(data, &extendedArtistList); err != nil {
		return err
	}

	*a = extendedArtistList.Data

	return nil
}

func GetArtist(id int) (Artist, error) {
	path := fmt.Sprintf("/artist/%d", id)
	result := Artist{}
	err := get(path, nil, &result)
	return result, err
}

func GetArtistTop(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/artist/%d/top", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetArtistAlbums(id, index, limit int) (AlbumList, error) {
	path := fmt.Sprintf("/artist/%d/albums", id)
	result := AlbumList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetArtistComments(id, index, limit int) (CommentList, error) {
	path := fmt.Sprintf("/artist/%d/comments", id)
	result := CommentList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetArtistFans(id, index, limit int) (UserList, error) {
	path := fmt.Sprintf("/artist/%d/fans", id)
	result := UserList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetArtistRelated(id, index, limit int) (ArtistList, error) {
	path := fmt.Sprintf("/artist/%d/related", id)
	result := ArtistList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetArtistRadio(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/artist/%d/radio", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}
