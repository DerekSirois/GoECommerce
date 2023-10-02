package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartItem []*CartItem `gorm:"many2many:cart_cartItem;"`
	UserID   uint
}

type CartJson struct {
	ProductId uint
	Quantity  int
}

func NewCart() *Cart {
	return &Cart{
		CartItem: make([]*CartItem, 0),
	}
}

func (u *User) AddToCart(db *gorm.DB, cartItem *CartItem) error {
	u.Cart.CartItem = append(u.Cart.CartItem, cartItem)
	err := u.UpdateCart(db)
	return err
}
