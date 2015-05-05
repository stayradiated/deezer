package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func NoTestGetGenres(t *testing.T) {
	result, err := GetGenres()
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetGenre(t *testing.T) {
	result, err := GetGenre(85)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetGenreArtists(t *testing.T) {
	result, err := GetGenreArtists(85)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetGenreRadios(t *testing.T) {
	result, err := GetGenreRadios(85)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
