package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Type string `gorm:"not null"`
}
