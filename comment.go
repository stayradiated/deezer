package deezer

import "fmt"

type Comment struct {
	Id     int      `json:"id,omitempty"`
	Text   string   `json:"text,omitempty"`
	Date   int      `json:"date,omitempty"`
	Object Playlist `json:"object,omitempty"`
	Author User     `json:"author,omitempty"`
}

type CommentList struct {
	Data  []Comment `json:"data,omitempty"`
	Total int       `json:"total,omitempty"`
	Next  string    `json:"next,omitempty"`
}

func GetComment(id int) (Comment, error) {
	path := fmt.Sprintf("comment/%d", id)
	result := Comment{}
	err := get(path, nil, &result)
	return result, err
}
