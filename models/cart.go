package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	User     User      `gorm:"foreignKey:ID"`
	Products []Product `gorm:"many2many:cart_product;"`
}
