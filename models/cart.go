package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartItem []*CartItem `gorm:"many2many:cart_cartItem;"`
	UserID   uint
}

func NewCart() *Cart {
	return &Cart{
		CartItem: make([]*CartItem, 0),
	}
}

func (u *User) AddToCart(db *gorm.DB, cartItem *CartItem) {
	u.Cart.CartItem = append(u.Cart.CartItem, cartItem)

}
