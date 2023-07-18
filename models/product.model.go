package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name,omitempty"`
	Image     string    `gorm:"not null" json:"image,omitempty"`
	Price     int64     `gorm:"not null" json:"price,omitempty"`
	IsSoldOut bool      `gorm:"default:false" json:"isSoldOut,omitempty"`
	Category  uuid.UUID `gorm:"not null" json:"category_id,omitempty"`
	User      uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateProductRequest struct {
	Name      string    `json:"name"  binding:"required"`
	Image     string    `json:"image" binding:"required"`
	User      string    `json:"user,omitempty"`
	Price     int64     `json:"price,omitempty"`
	Category  uuid.UUID `json:"category_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateProduct struct {
	Name      string    `json:"name,omitempty"`
	Image     string    `json:"image,omitempty"`
	User      string    `json:"user,omitempty"`
	Price     int64     `json:"price,omitempty"`
	IsSoldOut bool      `json:"isSoldOut,omitempty"`
	Category  uuid.UUID `json:"category_id,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
