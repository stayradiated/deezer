package deezer

import "fmt"

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GenreRequest struct {
	Id int
}

func (g GenreRequest) All() string {
	return "genre"
}

func (g GenreRequest) Get() string {
	return fmt.Sprintf("genre/%d", g.Id) // Genre
}

func (g GenreRequest) Artists() string {
	return fmt.Sprintf("genre/%d/artists", g.Id) // ArtistList
}

func (g GenreRequest) Radios() string {
	return fmt.Sprintf("genre/%d/radios", g.Id) // RadioList
}
