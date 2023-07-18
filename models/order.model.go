package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	TableId    uuid.UUID `gorm:"not null" json:"table,omitempty"`
	OrderItem  string    `gorm:"not null" json:"order_item,omitempty"`
	TotalPrice int64     `gorm:"not null" json:"total_price,omitempty"`
	User       uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateOrderRequest struct {
	OrderItem string    `json:"order_item"  binding:"required"`
	TableId   uuid.UUID `json:"table,omitempty"`
	User      uuid.UUID `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateOrder struct {
	OrderItem string    `json:"order_item"  binding:"required"`
	Category  uuid.UUID `json:"category_id,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
