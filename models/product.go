package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Quantity    int
	Price       float32
}

func GetAllProduct(db *gorm.DB) (*[]Product, error) {
	var p []Product
	result := db.Find(&p)
	return &p, result.Error
}

func GetAllProductInStock(db *gorm.DB) (*[]Product, error) {
	var p []Product
	result := db.Find(&p, "quantity > 0")
	return &p, result.Error
}

func (p *Product) GetById(db *gorm.DB, id uint) error {
	result := db.First(p, "id = ?", id)
	return result.Error
}

func (p *Product) Create(db *gorm.DB) error {
	result := db.Create(p)
	return result.Error
}

func (p *Product) Update(db *gorm.DB) error {
	pDb := &Product{}
	err := pDb.GetById(db, p.ID)
	if err != nil {
		return err
	}
	result := db.Model(pDb).Select("name", "description", "quantity", "price").Updates(p)
	return result.Error
}

func (p *Product) Delete(db *gorm.DB) error {
	result := db.Delete(p)
	return result.Error
}
