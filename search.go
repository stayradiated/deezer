package deezer

import (
	"strconv"

	"github.com/jmcvetta/napping"
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

func searchParams(query string, strict bool, order string, index, limit int) *napping.Params {
	return &napping.Params{
		"q":      query,
		"strict": strconv.FormatBool(strict),
		"order":  order,
		"index":  strconv.Itoa(index),
		"limit":  strconv.Itoa(limit),
	}
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
