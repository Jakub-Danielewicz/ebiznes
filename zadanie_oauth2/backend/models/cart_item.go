package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Cart      Cart    `json:"cart"`
	CartID    uint    `json:"cartId"`
	Product   Product `json:"product"`
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
}
