package deezer

type Chart struct {
	Tracks    TrackList    `json:"tracks"`
	Albums    AlbumList    `json:"albums"`
	Artists   ArtistList   `json:"artists"`
	Playlists PlaylistList `json:"playlists"`
}

type ChartRequest struct {
}

func (c ChartRequest) All() string {
	return "chart"
}
