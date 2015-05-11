package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetTrack(t *testing.T) {
	result, err := GetArtist(34796791)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
