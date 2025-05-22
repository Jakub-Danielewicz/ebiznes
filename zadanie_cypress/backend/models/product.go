package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string   `json:"name" gorm:"not null"`
	Price      float64  `json:"price" gorm:"not null"`
	CategoryID uint     `json:"category_id" gorm:"not null"`
	Category   Category `json:"category"`
}

func WithCategory() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Category")
	}
}
