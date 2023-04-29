package models

import (
	"time"
)

type TransferTransaction struct {
	ID              []byte    `gorm:"type:binary(16);primaryKey" json:"id"`
	FromUserID      int       `json:"from_user_id"`
	ToUserID        int       `json:"to_user_id"`
	FromTokenID     int       `json:"from_token_id"`
	ToTokenID       int       `json:"to_token_id"`
	FromTokenAmount float64   `json:"from_token_amount"`
	ToTokenAmount   float64   `json:"to_token_amount"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	FromToken       Token     `gorm:"foreignKey:FromTokenID" json:"from_token"`
	ToToken         Token     `gorm:"foreignKey:ToTokenID" json:"to_token"`
	FromUser        User      `gorm:"foreignKey:FromUserID" json:"from_user"`
	ToUser          User      `gorm:"foreignKey:ToUserID" json:"to_user"`
}
