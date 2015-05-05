package deezer

import "fmt"

type User struct {
	Id              int    `json:"id,omitempty"`               //	The user's Deezer ID
	Name            string `json:"name,omitempty"`             //	The user's Deezer nickname
	Lastname        string `json:"lastname,omitempty"`         //	The user's last name
	Firstname       string `json:"firstname,omitempty"`        //	The user's first name
	Email           string `json:"email,omitempty"`            //	The user's email
	Status          int    `json:"status,omitempty"`           //	The user's status
	Birthday        int    `json:"birthday,omitempty"`         //	The user's birthday
	InscriptionDate int    `json:"inscription_date,omitempty"` //	The user's inscription date
	Gender          string `json:"gender,omitempty"`           //	The user's gender : F or M
	Link            string `json:"link,omitempty"`             //	The url of the profil for the user on Deezer
	Picture         string `json:"picture,omitempty"`          //	The url of the user's profil picture.
	Country         string `json:"country,omitempty"`          //	The user's country
	Lang            string `json:"lang,omitempty"`             //	The user's language
	Tracklist       string `json:"tracklist,omitempty"`        //	API Link to the flow of this user
}

type UserList struct {
	Data  []User `json:"data,omitempty"`
	Total int    `json:"total,omitempty"`
	Next  string `json:"next,omitempty"`
}

func GetUser(id int) (User, error) {
	path := fmt.Sprintf("/user/%d", id)
	result := User{}
	err := get(path, nil, &result)
	return result, err
}

func GetUserAlbums(id, index, limit int) (AlbumList, error) {
	path := fmt.Sprintf("/user/%d/albums", id)
	result := AlbumList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetUserArtists(id, index, limit int) (ArtistList, error) {
	path := fmt.Sprintf("/user/%d/artists", id)
	result := ArtistList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetUserCharts(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/user/%d/charts", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetUserFlow(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/user/%d/flow", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetUserFollowings(id, index, limit int) (UserList, error) {
	path := fmt.Sprintf("/user/%d/followings", id)
	result := UserList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetUserPlaylists(id, index, limit int) (PlaylistList, error) {
	path := fmt.Sprintf("/user/%d/playlists", id)
	result := PlaylistList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetUserRadios(id, index, limit int) (RadioList, error) {
	path := fmt.Sprintf("/user/%d/radios", id)
	result := RadioList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetUserTracks(id, index, limit int) (TrackList, error) {
	path := fmt.Sprintf("/user/%d/tracks", id)
	result := TrackList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}
