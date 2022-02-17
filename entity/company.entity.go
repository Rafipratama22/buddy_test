package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;unique" json:"id,omitempty"`
	Name string `json:"name" gorm:"type:varchar(400)"`
	Email string `json:"email" gorm:"type:varchar(400)"`
	Password string `json:"password" gorm:"type:varchar(400)"`
	IsActive bool `json:"is_active" gorm:"type:boolean;default:true"`
}

func (c *Company) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}