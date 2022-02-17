package repository

import (
	"github.com/Rafipratama22/mnc_test.git/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(newpost entity.Post) (entity.Post, error)
	DetailPost(id int) (entity.Post, error)
	UpdatePost(id int, newpost entity.Post) (entity.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (c *postRepository) CreatePost(newpost entity.Post) (entity.Post, error) {
	result := c.db.Create(&newpost)
	if result.Error != nil {
		return newpost, result.Error
	} else {
		return newpost, nil
	}
}

func (c *postRepository) DetailPost(id int) (entity.Post, error){
	var post entity.Post
	result := c.db.Model(&post).Where("id = ?", id).First(&post)
	if result.Error != nil {
		return post, result.Error
	} else {
		return post, nil
	}
}

func (c *postRepository) UpdatePost(id int, newpost entity.Post) (entity.Post, error) {
	var post entity.Post
	result := c.db.Model(&post).Where("id = ?", id).Updates(&newpost)
	if result.Error != nil {
		return post, result.Error
	} else {
		return newpost, nil
	}
}