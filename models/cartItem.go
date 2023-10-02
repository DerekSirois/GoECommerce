package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ProductId uint
	Quantity  int
}

func NewCartItem(p uint, q int) *CartItem {
	return &CartItem{
		ProductId: p,
		Quantity:  q,
	}
}
