package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	product  Product
	quantity int
}
