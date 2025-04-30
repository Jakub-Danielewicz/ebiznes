package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
CartID	uint `json:"cartId"`
Cart	Cart `json:"cart"`
ProductID	uint `json:"productId"`
Product	Product `json:"product"`
Quantity	int `json:"quantity"`
}

