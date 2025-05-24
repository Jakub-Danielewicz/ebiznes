package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint    `json:"userId"`
	Total  float64 `json:"total"`
}
