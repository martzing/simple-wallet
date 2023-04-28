package models

import "time"

type Wallet struct {
	ID        int       `gorm:"primarykey;autoIncrement" json:"id"`
	Balance   float64   `json:"balance"`
	TokenID   int       `json:"token_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Token     Token     `gorm:"foreignKey:TokenID" json:"token"`
}
