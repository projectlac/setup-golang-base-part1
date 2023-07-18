package models

import (
	"time"

	"github.com/google/uuid"
)

type TableOrder struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	TableName string    `gorm:"uniqueIndex;not null" json:"table_name,omitempty"`
	User      uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateTableOrderRequest struct {
	TableName string    `json:"table_name"  binding:"required"`
	User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateTableOrder struct {
	TableName string    `json:"table_name,omitempty"`
	User      string    `json:"user,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
