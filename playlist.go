package deezer

import "fmt"

type Playlist struct {
	Id            int       `json:"id,omitempty"`             // The playlist's Deezer id
	Title         string    `json:"title,omitempty"`          // The playlist's title
	Description   string    `json:"description,omitempty"`    // The playlist description
	Duration      int       `json:"duration,omitempty"`       // The playlist's duration (seconds)
	Public        bool      `json:"public,omitempty"`         // If the playlist is public or not
	IsLovedTrack  bool      `json:"is_loved_track,omitempty"` // If the playlist is the love tracks playlist
	Collaborative bool      `json:"collaborative,omitempty"`  // If the playlist is collaborative or not
	Rating        int       `json:"rating,omitempty"`         // The playlist's rate
	Fans          int       `json:"fans,omitempty"`           // The number of playlist's fans
	Link          string    `json:"link,omitempty"`           // The url of the playlist on Deezer
	Share         string    `json:"share,omitempty"`          // The share link of the playlist on Deezer
	Picture       string    `json:"picture,omitempty"`        // The url of the playlist's cover.
	Checksum      string    `json:"hecksum,omitempty"`        // The checksum for the track list
	Creator       User      `json:"creator,omitempty"`        // User object containing : id, name
	Tracks        TrackList `json:"tracks,omitempty"`         // List of tracks
}

type PlaylistList struct {
	Data  []Playlist `json:"data,omitempty"`
	Total int        `json:"total,omitempty"`
	Next  string     `json:"next,omitempty"`
}

func GetPlaylist(id int) (Playlist, error) {
	path := fmt.Sprintf("/playlist/%d", id)
	result := Playlist{}
	err := get(path, nil, &result)
	return result, err
}

func GetPlaylistComments(id, index, limit int) (CommentList, error) {
	path := fmt.Sprintf("/playlist/%d/comments", id)
	result := CommentList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetPlaylistFans(id, index, limit int) (UserList, error) {
	path := fmt.Sprintf("/playlist/%d/fans", id)
	result := UserList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetPlaylistTracks(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/playlist/%d/tracks", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}
