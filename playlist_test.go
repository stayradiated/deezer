package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetPlaylist(t *testing.T) {
	result, err := GetPlaylist(785141981)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetPlaylistComments(t *testing.T) {
	result, err := GetPlaylistComments(785141981, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetPlaylistFans(t *testing.T) {
	result, err := GetPlaylistFans(785141981, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetPlaylistTracks(t *testing.T) {
	result, err := GetPlaylistTracks(785141981, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
