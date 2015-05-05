package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func NoTestGetPlaylist(t *testing.T) {
	result, err := GetPlaylist(785141981)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetPlaylistComments(t *testing.T) {
	result, err := GetPlaylistComments(785141981, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetPlaylistFans(t *testing.T) {
	result, err := GetPlaylistFans(785141981, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetPlaylistTracks(t *testing.T) {
	result, err := GetPlaylistTracks(785141981, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
