package deezer

import "fmt"

type Playlist struct {
	Id            int       `json:"id"`             // The playlist's Deezer id
	Title         string    `json:"title"`          // The playlist's title
	Description   string    `json:"description"`    // The playlist description
	Duration      int       `json:"duration"`       // The playlist's duration (seconds)
	Public        bool      `json:"public"`         // If the playlist is public or not
	IsLovedTrack  bool      `json:"is_loved_track"` // If the playlist is the love tracks playlist
	Collaborative bool      `json:"collaborative"`  // If the playlist is collaborative or not
	Rating        int       `json:"rating"`         // The playlist's rate
	Fans          int       `json:"fans"`           // The number of playlist's fans
	Link          string    `json:"link"`           // The url of the playlist on Deezer
	Share         string    `json:"share"`          // The share link of the playlist on Deezer
	Picture       string    `json:"picture"`        // The url of the playlist's cover.
	Checksum      string    `json:"hecksum"`        // The checksum for the track list
	Creator       User      `json:"creator"`        // User object containing : id, name
	Tracks        TrackList `json:"tracks"`         // List of tracks
}

type PlaylistList struct {
	Data []Playlist
}

type PlaylistRequest struct {
	Id int
}

func (p PlaylistRequest) Get() string {
	return fmt.Sprintf("playlist/%d", p.Id) // Playlist
}

func (p PlaylistRequest) Comments() string {
	return fmt.Sprintf("playlist/%d/comments", p.Id) // CommentList
}

func (p PlaylistRequest) Fans() string {
	return fmt.Sprintf("playlist/%d/fans", p.Id) // UserList
}

func (p PlaylistRequest) Tracks() string {
	return fmt.Sprintf("playlist/%d/tracks", p.Id) // TrackList
}
