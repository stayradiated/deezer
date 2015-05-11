package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetGenres(t *testing.T) {
	result, err := GetGenres()
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetGenre(t *testing.T) {
	result, err := GetGenre(85)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetGenreArtists(t *testing.T) {
	result, err := GetGenreArtists(85)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetGenreRadios(t *testing.T) {
	result, err := GetGenreRadios(85)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
