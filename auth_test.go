package deezer

import (
	"testing"
)

func TestAuth(t *testing.T) {

	err := Login("328122", "20bda80ea77d46bd59bb5457449b1312")
	if err != nil {
		t.Fatal(err)
	}
}
