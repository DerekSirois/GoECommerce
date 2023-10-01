package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Products []*Product `gorm:"many2many:cart_product;"`
	UserID   uint
}

func NewCart() *Cart {
	return &Cart{
		Products: make([]*Product, 0),
	}
}