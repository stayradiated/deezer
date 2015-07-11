package deezer

import "fmt"

type Track struct {
	ID                 int      `json:"id,omitempty"`                  // The track's Deezer id
	Readable           bool     `json:"readable,omitempty"`            // true if the track is readable in the player for the current user
	Title              string   `json:"title,omitempty"`               // The track's title
	Isrc               string   `json:"isrc,omitempty"`                // The track isrc
	Link               string   `json:"link,omitempty"`                // The url of the track on Deezer
	Share              string   `json:"share,omitempty"`               // The share link of the track on Deezer
	Duration           int      `json:"duration,omitempty"`            // The track's duration in seconds
	TrackPosition      int      `json:"track_position,omitempty"`      // The position of the track in its album
	DiskNumber         int      `json:"disk_number,omitempty"`         // The track's album's disk number
	Rank               int      `json:"rank,omitempty"`                // The track's Deezer rank
	ReleaseDate        string   `json:"release_date,omitempty"`        // The track's release date
	ExplicitLyrics     bool     `json:"explicit_lyrics,omitempty"`     // Whether the track contains explicit lyrics
	Preview            string   `json:"preview,omitempty"`             // The url of track's preview file. This file contains the first 30 seconds of the track
	Bpm                float64  `json:"bpm,omitempty"`                 // Beats per minute
	Gain               float64  `json:"gain,omitempty"`                // Signal strength
	AvailableCountries []string `json:"available_countries,omitempty"` // List of countries where the track is available
	// Alternative         Track    `json:"alternative,omitempty"`         // Return an alternative readable track if the current track is not readable
	Contributors []Artist `json:"contributors,omitempty"` // Return a list of contributors on the track
	Artist       Artist   `json:"artist,omitempty"`       // artist object containing : id, name, link, share, picture, nb_album, nb_fan, radio, tracklist, role
	Album        Album    `json:"album,omitempty"`        // album object containing : id, title, link, cover, release_date
}

type TrackList struct {
	Data  []Track `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

func GetTrack(id int) (Track, error) {
	path := fmt.Sprintf("/track/%d", id)
	result := Track{}
	err := get(path, nil, &result)
	return result, err
}
