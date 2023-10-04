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

func (c *Cart) RemoveFromCart(db *gorm.DB, item *CartItem) error {
	err := db.Model(c).Association("CartItem").Delete(item)
	return err
}

func (c *Cart) Checkout(db *gorm.DB) error {
	for _, element := range c.CartItem {
		i := &Product{}
		err := i.GetById(db, element.ProductId)
		if err != nil {
			return err
		}

		i.Quantity -= element.Quantity
		err = i.Update(db)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cart) EmptyCart(db *gorm.DB) error {
	err := db.Model(c).Association("CartItem").Clear()
	return err
}
