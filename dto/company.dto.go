package dto

type PostPointDTO struct {
	Point int `json:"point" binding:"required"`
}