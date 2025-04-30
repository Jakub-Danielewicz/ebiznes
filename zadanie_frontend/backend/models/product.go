package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name	string `json:"name"`
	Price	float64 `json:"price"`
	CategoryID	uint `json:"category_id"`
	Category Category `json:"category"`
}

func WithCategory() func(*gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Preload("Category")
		}
}
