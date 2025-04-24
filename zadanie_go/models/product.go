package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name	string
	Price	float64
	CategoryID	uint
}

func WithCategory() func(*gorm.DB) *gorm.DB {
		return func(db *gorm.DB) *gorm.DB {
			return db.Preload("Category")
		}
}
