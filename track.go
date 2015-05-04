package deezer

import "fmt"

type Track struct {
	Id                  int      `json:"id"`                  // The track's Deezer id
	Readable            bool     `json:"readable"`            // true if the track is readable in the player for the current user
	Title               string   `json:"title"`               // The track's title
	Isrc                string   `json:"isrc"`                // The track isrc
	Link                string   `json:"link"`                // The url of the track on Deezer
	Share               string   `json:"share"`               // The share link of the track on Deezer
	Duration            int      `json:"duration"`            // The track's duration in seconds
	Track_position      int      `json:"track_position"`      // The position of the track in its album
	Disk_number         int      `json:"disk_number"`         // The track's album's disk number
	Rank                int      `json:"rank"`                // The track's Deezer rank
	Release_date        int      `json:"release_date"`        // The track's release date
	Explicit_lyrics     bool     `json:"explicit_lyrics"`     // Whether the track contains explicit lyrics
	Preview             string   `json:"preview"`             // The url of track's preview file. This file contains the first 30 seconds of the track
	Bpm                 float64  `json:"bpm"`                 // Beats per minute
	Gain                float64  `json:"gain"`                // Signal strength
	Available_countries []string `json:"available_countries"` // List of countries where the track is available
	// Alternative         Track    `json:"alternative"`         // Return an alternative readable track if the current track is not readable
	Contributors []Artist `json:"contributors"` // Return a list of contributors on the track
	Artist       Artist   `json:"artist"`       // artist object containing : id, name, link, share, picture, nb_album, nb_fan, radio, tracklist, role
	Album        Album    `json:"album"`        // album object containing : id, title, link, cover, release_date
}

type TrackList struct {
	Data  []Track `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

type TrackRequest struct {
	Id int
}

func (t TrackRequest) Get() string {
	return fmt.Sprintf("track/%d", t.Id)
}
