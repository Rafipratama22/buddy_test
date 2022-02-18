package dto

type PostDto struct {
	Content string `json:"content"`
	Hastag string `json:"hashtag"`
	Photo string `json:"photo"`
	Source string `json:"source"`
	Thumbnail string `json:"thumbnail"`
	Title string `json:"title"`
}