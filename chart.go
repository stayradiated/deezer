package deezer

import "fmt"

type Chart struct {
	Tracks    TrackList    `json:"tracks,omitempty"`
	Albums    AlbumList    `json:"albums,omitempty"`
	Artists   ArtistList   `json:"artists,omitempty"`
	Playlists PlaylistList `json:"playlists,omitempty"`
}

func GetCharts(index, limit int) (Chart, error) {
	path := "/chart"
	result := Chart{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetChartTracks(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/chart/%d/tracks", id)
	result := TrackList{}
	err := get(path, nil, &result)
	return result, err
}

func GetChartAlbums(id, index, limit int) (AlbumList, error) {
	path := fmt.Sprintf("/chart/%d/albums", id)
	result := AlbumList{}
	err := get(path, nil, &result)
	return result, err
}

func GetChartArtists(id, index, limit int) (ArtistList, error) {
	path := fmt.Sprintf("/chart/%d/artists", id)
	result := ArtistList{}
	err := get(path, nil, &result)
	return result, err
}

func GetChartPlaylists(id, index, limit int) (PlaylistList, error) {
	path := fmt.Sprintf("/chart/%d/playlists", id)
	result := PlaylistList{}
	err := get(path, nil, &result)
	return result, err
}
