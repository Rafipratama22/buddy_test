package entity

import "github.com/google/uuid"

// swagger:model Post
type Post struct {
	// Post ID
	// @PrimaryKey
	// @Column(type:int, unique: true, autoincrement: true)
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	// Post Title
	// @Column(type:varchar(400), nullable:false)
	Title string `json:"title" gorm:"type:varchar(400)"`
	// Post Content
	// @Column(type:varchar(400), nullable:false)
	Content string `json:"content" gorm:"type:text"`
	// Post Photo
	// @Column(type:text, nullable:false)
	Photo string `json:"photo" gorm:"type:text"`
	// Post Point
	// @Column(type:int, nullable:false, default:0)
	Point int `json:"point" gorm:"type:int;default:0"`
	// Post Hastag
	// @Column(type:varchar(400), nullable:true)
	Hastag string `json:"hastag" gorm:"type:varchar(400)"`
	// Post Thumbnaiil
	// @Column(type:text, nullable:false)
	Thumbnail string `json:"thumbnail" gorm:"type:text"`
	// Post Source
	// @Column(type:varchar(400), nullable:true)
	Source string `json:"source" gorm:"type:text"`
	// Post Author ID
	// @Column(type:uuid, nullable:false)
	AuthorID uuid.UUID `json:"author_id" gorm:"type:uuid"`
}