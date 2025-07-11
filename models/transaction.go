package models

import "time"

type Transaction struct {
	ID          uint `gorm:"primaryKey"`
	Amount      float64
	Category    string
	Description string
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
