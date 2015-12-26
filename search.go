package deezer

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	RANKING       = "RANKING"
	TRACK_ASC     = "TRACK_ASC"
	TRACK_DESC    = "TRACK_DESC"
	ARTIST_ASC    = "ARTIST_ASC"
	ARTIST_DESC   = "ARTIST_DESC"
	ALBUM_ASC     = "ALBUM_ASC"
	ALBUM_DESC    = "ALBUM_DESC"
	RATING_ASC    = "RATING_ASC"
	RATING_DESC   = "RATING_DESC"
	DURATION_ASC  = "DURATION_ASC"
	DURATION_DESC = "DURATION_DESC"
)

func searchParams(query string, strict bool, order string, index, limit int) *url.Values {
	v := &url.Values{}
	v.Set("q", query)
	v.Set("strict", strconv.FormatBool(strict))
	v.Set("order", order)
	v.Set("index", strconv.Itoa(index))
	v.Set("limit", strconv.Itoa(limit))
	return v
}

func SearchTrack(query string, strict bool, order string, index, limit int) (TrackList, error) {
	path := "/search/track"
	result := TrackList{}
	err := get(path, searchParams(query, strict, order, index, limit), &result)
	return result, err
}

func SearchAlbum(query string, strict bool, order string, index, limit int) (AlbumList, error) {
	path := "/search/album"
	result := AlbumList{}
	err := get(path, searchParams(query, strict, order, index, limit), &result)
	return result, err
}

func SearchArtist(query string, strict bool, order string, index, limit int) (ArtistList, error) {
	path := "/search/artist"
	result := ArtistList{}
	err := get(path, searchParams(query, strict, order, index, limit), &result)
	return result, err
}

func SearchUser(query string, strict bool, order string, index, limit int) (UserList, error) {
	path := "/search/user"
	result := UserList{}
	err := get(path, searchParams(query, strict, order, index, limit), &result)
	return result, err
}

func SearchPlaylist(query string, strict bool, order string, index, limit int) (PlaylistList, error) {
	path := "/search/playlist"
	result := PlaylistList{}
	err := get(path, searchParams(query, strict, order, index, limit), &result)
	return result, err
}

type AdvancedSearch struct {
	Artist string
	Album  string
	Track  string
	Label  string
	DurMin int
	DurMax int
	BPMMin int
	BPMMax int
}

func (a AdvancedSearch) String() (s string) {
	if a.Artist != "" {
		s += fmt.Sprintf("artist:\"%s\" ", a.Artist)
	}
	if a.Album != "" {
		s += fmt.Sprintf("album:\"%s\" ", a.Album)
	}
	if a.Track != "" {
		s += fmt.Sprintf("track:\"%s\" ", a.Track)
	}
	if a.Label != "" {
		s += fmt.Sprintf("label:\"%s\" ", a.Track)
	}
	if a.DurMin > 0 {
		s += fmt.Sprintf("dur_min:\"%d\" ", a.DurMin)
	}
	if a.DurMax > 0 {
		s += fmt.Sprintf("dur_max:\"%d\" ", a.DurMax)
	}
	if a.BPMMin > 0 {
		s += fmt.Sprintf("bpm_max:\"%d\" ", a.BPMMin)
	}
	if a.BPMMax > 0 {
		s += fmt.Sprintf("bpm_max:\"%d\" ", a.BPMMax)
	}
	return s
}
