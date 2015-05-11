package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSearchTrack(t *testing.T) {
	result, err := SearchTrack("Stornoway Zorbing", false, RANKING, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestSearchArtist(t *testing.T) {
	result, err := SearchArtist("Squirrel Nut Zippers", false, RANKING, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestSearchAlbum(t *testing.T) {
	result, err := SearchAlbum("Fleetwood Mac Rumours", false, RANKING, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestSearchUser(t *testing.T) {
	result, err := SearchUser("stayrad", false, RANKING, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestSearchPlaylist(t *testing.T) {
	result, err := SearchPlaylist("Super Hits", false, RANKING, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
