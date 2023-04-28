package models

import "time"

type Token struct {
	ID        int       `gorm:"primarykey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	Image     string    `json:"image"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
