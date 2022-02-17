package repository

import (
	"github.com/Rafipratama22/mnc_test.git/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllPoint(user_id uuid.UUID) (int, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (c *userRepository) GetAllPoint(user_id uuid.UUID) (int, error){
	var post []entity.Post
	result := c.db.Model(&post).Where("author_id = ?", user_id).Find(&post)
	if result.Error != nil {
		return 0, result.Error
	} else {
		var total int
		for _, value := range post {
			total += value.Point
		}
		return total, nil
	}
}