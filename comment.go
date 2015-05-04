package deezer

import "fmt"

type Comment struct {
	Id     int      `json:"id"`
	Text   string   `json:"text"`
	Date   int      `json:"date"`
	Object Playlist `json:"object"`
	Author User     `json:"author"`
}

type CommentRequest struct {
	Id int
}

func (c CommentRequest) Get() string {
	return fmt.Sprintf("comment/%d", c.Id)
}
