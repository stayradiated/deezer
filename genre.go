package deezer

import "fmt"

type Genre struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GenreList struct {
	Data []Genre `json:"data,omitempty"`
}

func GetGenres() (GenreList, error) {
	path := "/genre"
	result := GenreList{}
	err := get(path, nil, &result)
	return result, err
}

func GetGenre(id int) (Genre, error) {
	path := fmt.Sprintf("/genre/%d", id) // Genre
	result := Genre{}
	err := get(path, nil, &result)
	return result, err
}

func GetGenreArtists(id int) (ArtistList, error) {
	path := fmt.Sprintf("/genre/%d/artists", id)
	result := ArtistList{}
	err := get(path, nil, &result)
	return result, err
}

func GetGenreRadios(id int) (RadioList, error) {
	path := fmt.Sprintf("/genre/%d/radios", id)
	result := RadioList{}
	err := get(path, nil, &result)
	return result, err
}
