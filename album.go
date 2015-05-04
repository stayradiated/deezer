package deezer

import "fmt"

type Album struct {
	Id             int      `json:"id"`              // The Deezer album id
	Title          string   `json:"title"`           // The album title
	UPC            string   `json:"upc"`             // The album UPC
	Link           string   `json:"link"`            // The url of the album on Deezer url
	Share          string   `json:"share"`           // The share link of the album on Deezer
	Cover          string   `json:"cover"`           // The url of the album's cover.
	GenreId        int      `json:"genre_id"`        // The album's first genre id. NB : -1 for not found
	Genres         []Genre  `json:"genres"`          // List of genre object
	Label          string   `json:"label"`           // The album's label name
	NbTracks       int      `json:"nb_tracks"`       // Number of tracks
	Duration       int      `json:"duration"`        // The album's duration (seconds)
	Fans           int      `json:"fans"`            // The number of album's Fans
	Rating         int      `json:"rating"`          // The playlist's rate
	ReleaseDate    int      `json:"release_date"`    // The album's release date
	Record_type    string   `json:"record_type"`     // The record type of the album (EP / ALBUM / etc..)
	Available      bool     `json:"available"`       // If the album is avaiable
	Alternative    Album    `json:"alternative"`     // Return an alternative album object if the current album is not available
	Tracklist      string   `json:"tracklist"`       // API Link to the tracklist of this album
	ExplicitLyrics bool     `json:"explicit_lyrics"` // Whether the album contains explicit lyrics
	Contributors   []Artist `json:"contributors"`    // Return a list of contributors on the album
	Artist         Artist   `json:"artist"`          // artist object containing : id, name, picture
	Tracks         []Track  `json:"tracks"`          // list of track
}

type AlbumRequest struct {
	Id int
}

func (a AlbumRequest) Get() string {
	return fmt.Sprintf("album/%d", t.Id) // Album
}

func (a AlbumRequest) Comments() string {
	return fmt.Sprintf("album/%d/comments", t.Id) // CommentList
}

func (a AlbumRequest) Fans() string {
	return fmt.Sprintf("album/%d/fans", t.Id) // UserList
}

func (a AlbumRequest) Tracks() string {
	return fmt.Sprintf("album/%d/tracks", t.Id) // TrackList
}
