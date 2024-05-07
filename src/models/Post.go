package models

import "time"

type Post struct {
	ID uint64 `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	AuthorID uint64 `json:"authorId,omitempty"`
	AuthorUsername string `json:"authorUsername,omitempty"`
	Likes uint64 `json:"likes"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}