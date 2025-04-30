package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
ID uint `json:"id"`
UserID	uint `json:"userId"`
CartItems []CartItem `json:"userId"`
	}
