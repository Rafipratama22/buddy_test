package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// swagger:model Company
type Company struct {
	// Company ID
	// @PrimaryKey
	// @Column(type:uuid, unique: true)
	ID uuid.UUID `gorm:"primary_key;type:uuid;unique" json:"id,omitempty"`
	// Company Name
	// @Column(type:varchar(400), nullable:false)
	Name string `json:"name" binding:"required" gorm:"type:varchar(400)"`
	// Company Email
	// @Column(type:varchar(400), nullable:false)
	Email string `json:"email" binding:"required" gorm:"type:varchar(400)"`
	// Company Phone
	// @Column(type:varchar(400), nullable:false)
	Password string `json:"password" binding:"required" gorm:"type:varchar(400)"`
	// Company IsActive
	// @Column(type:boolean, default:true)
	IsActive bool `json:"is_active" gorm:"type:boolean;default:true"`
}

func (c *Company) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}