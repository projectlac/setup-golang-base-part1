package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	CategoryName string    `gorm:"uniqueIndex;not null" json:"category_name,omitempty"`
	User         uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateCategoryRequest struct {
	CategoryName string    `json:"title"  binding:"required"`
	User         string    `json:"user,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type UpdateCategory struct {
	CategoryName string    `json:"title,omitempty"`
	User         string    `json:"user,omitempty"`
	CreateAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
