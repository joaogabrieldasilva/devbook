package dto

type CreatePost struct {
	Title string `json:"title"`
	Content string `json:"content"`
}