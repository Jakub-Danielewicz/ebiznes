package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID	uint
	Cart	Cart
	ProductID	uint
	Product	Product
	Quantity	int
}

