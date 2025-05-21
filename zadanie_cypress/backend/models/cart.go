package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
UserID	uint `json:"userId"`
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cartItems"`
	}
