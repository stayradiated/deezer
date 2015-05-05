package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetCharts(t *testing.T) {
	result, err := GetCharts(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
