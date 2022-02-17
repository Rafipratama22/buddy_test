package entity

import "github.com/google/uuid"

type Post struct {
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Title string `json:"title" gorm:"type:varchar(400)"`
	Content string `json:"content" gorm:"type:text"`
	Photo string `json:"photo" gorm:"type:text"`
	Point int `json:"point" gorm:"type:int;default:0"`
	Hastag string `json:"hastag" gorm:"type:varchar(400)"`
	Thumbnail string `json:"thumbnail" gorm:"type:text"`
	Source string `json:"source" gorm:"type:text"`
	AuthorID uuid.UUID `json:"author_id" gorm:"type:uuid"`
}