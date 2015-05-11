package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAlbumGet(t *testing.T) {
	result, err := GetAlbum(6575789)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetAlbumTracks(t *testing.T) {
	result, err := GetAlbumTracks(6575789, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetAlbumComments(t *testing.T) {
	result, err := GetAlbumComments(6575789, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetAlbumFans(t *testing.T) {
	result, err := GetAlbumFans(6575789, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
