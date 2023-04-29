package models

import "time"

type User struct {
	ID                  int                   `gorm:"primarykey;autoIncrement" json:"id"`
	Username            string                `json:"username"`
	Password            string                `json:"password"`
	Email               string                `json:"email"`
	IsActive            bool                  `json:"is_active"`
	Role                string                `json:"role"`
	CreatedAt           time.Time             `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time             `gorm:"column:updated_at" json:"updated_at"`
	Wallet              []Wallet              `gorm:"foreignKey:UserID" json:"wallet"`
	TransferTransaction []TransferTransaction `gorm:"foreignKey:FromUserID" json:"transfer_transaction"`
}
