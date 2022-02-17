package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;unique" json:"id,omitempty"`
	Name string `json:"name" gorm:"type:varchar(400)"`
	Email string `json:"email" gorm:"type:varchar(400);unique"`
	Password string `json:"password" gorm:"type:varchar(400)"`
	IsActive bool `json:"is_active" gorm:"type:boolean;default:true"`
}

func (u *User) BeforeCreate(scope *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}