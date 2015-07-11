package deezer

import "fmt"

type Editorial struct {
	ID   int    `json:"id,omitempty"`   // The editorial's Deezer id
	Name string `json:"name,omitempty"` // The editorial's name
}

type EditorialList struct {
	Data  []Editorial `json:"data,omitempty"`
	Total int         `json:"total,omitempty"`
	Next  string      `json:"next,omitempty"`
}

func GetEditorials() (EditorialList, error) {
	path := "/editorial"
	result := EditorialList{}
	err := get(path, nil, &result)
	return result, err
}

func GetEditorial(id int) (Editorial, error) {
	path := fmt.Sprintf("editorial/%d", id)
	result := Editorial{}
	err := get(path, nil, &result)
	return result, err
}

func GetEditorialSelection(id int, date string, index, limit int) (AlbumList, error) {
	path := fmt.Sprintf("editorial/%d/selection", id)
	result := AlbumList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetEditorialCharts(id, index, limit int) (Chart, error) {
	path := fmt.Sprintf("editorial/%d/charts", id)
	result := Chart{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}

func GetEditorialReleases(id, index, limit int) (AlbumList, error) {
	path := fmt.Sprintf("editorial/%d/releases", id)
	result := AlbumList{}
	err := get(path, listParams(index, limit), &result)
	return result, err
}
