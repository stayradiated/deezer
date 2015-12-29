package deezer

import (
	"encoding/json"
	"fmt"
)

type Album struct {
	ID          int       `json:"id,omitempty"`           // The Deezer album id
	Title       string    `json:"title,omitempty"`        // The album title
	UPC         string    `json:"upc,omitempty"`          // The album UPC
	Link        string    `json:"link,omitempty"`         // The url of the album on Deezer url
	Share       string    `json:"share,omitempty"`        // The share link of the album on Deezer
	Cover       string    `json:"cover,omitempty"`        // The url of the album's cover.
	GenreID     int       `json:"genre_id,omitempty"`     // The album's first genre id. NB : -1 for not found
	Genres      GenreList `json:"genres,omitempty"`       // List of genre object
	Label       string    `json:"label,omitempty"`        // The album's label name
	NbTracks    int       `json:"nb_tracks,omitempty"`    // Number of tracks
	Duration    int       `json:"duration,omitempty"`     // The album's duration (seconds)
	Fans        int       `json:"fans,omitempty"`         // The number of album's Fans
	Rating      int       `json:"rating,omitempty"`       // The playlist's rate
	ReleaseDate string    `json:"release_date,omitempty"` // The album's release date
	RecordType  string    `json:"record_type,omitempty"`  // The record type of the album (EP / ALBUM / etc..)
	Available   bool      `json:"available,omitempty"`    // If the album is avaiable
	// Alternative    Album    `json:"alternative,omitempty"`     // Return an alternative album object if the current album is not available
	Tracklist      string    `json:"tracklist,omitempty"`       // API Link to the tracklist of this album
	ExplicitLyrics bool      `json:"explicit_lyrics,omitempty"` // Whether the album contains explicit lyrics
	Contributors   []Artist  `json:"contributors,omitempty"`    // Return a list of contributors on the album
	Artist         *Artist   `json:"artist,omitempty"`          // artist object containing : id, name, picture
	Tracks         TrackList `json:"tracks,omitempty"`          // list of track
}

type ExtendedAlbumList struct {
	Data  []Album `json:"data,omitempty"`
	Total int     `json:"total,omitempty"`
	Next  string  `json:"next,omitempty"`
}

type AlbumList []Album

func (a *AlbumList) UnmarshalJSON(data []byte) error {
	extendedAlbumList := ExtendedAlbumList{}
	if err := json.Unmarshal(data, &extendedAlbumList); err != nil {
		return err
	}

	*a = extendedAlbumList.Data

	return nil
}

func GetAlbum(id int) (Album, error) {
	path := fmt.Sprintf("/album/%d", id)
	result := Album{}
	err := get(path, nil, &result)
	return result, err
}

func GetAlbumComments(id, index, limit int) (CommentList, error) {
	path := fmt.Sprintf("/album/%d/comments", id)
	result := CommentList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetAlbumFans(id, index, limit int) (UserList, error) {
	path := fmt.Sprintf("/album/%d/fans", id)
	result := UserList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetAlbumTracks(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/album/%d/tracks", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}
