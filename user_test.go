package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func NoTestGetUser(t *testing.T) {
	result, err := GetUser(374224155)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetUserAlbums(t *testing.T) {
	result, err := GetUserAlbums(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetUserArtists(t *testing.T) {
	result, err := GetUserArtists(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetUserCharts(t *testing.T) {
	result, err := GetUserCharts(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetUserFlow(t *testing.T) {
	result, err := GetUserFlow(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetUserFollowings(t *testing.T) {
	result, err := GetUserFollowings(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetUserPlaylists(t *testing.T) {
	result, err := GetUserPlaylists(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func NoTestGetUserRadios(t *testing.T) {
	result, err := GetUserRadios(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserTracks(t *testing.T) {
	result, err := GetUserTracks(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
