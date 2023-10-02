package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password []byte
	IsAdmin  bool
	Cart     Cart
}

type UserJson struct {
	Username string
	Password string
	IsAdmin  bool
}

func (u *User) Create(db *gorm.DB) error {
	u.Cart = *NewCart()
	result := db.Create(u)
	return result.Error
}

func (u *User) GetByUsername(db *gorm.DB, username string) error {
	result := db.First(u, "username = ?", username)
	return result.Error
}

func (u *User) GetById(db *gorm.DB, id uint) error {
	result := db.First(u, "id = ?", id)
	return result.Error
}

func (u *User) Update(db *gorm.DB) error {
	uDb := &User{}
	err := uDb.GetById(db, u.ID)
	if err != nil {
		return err
	}
	result := db.Model(uDb).Updates(u)
	return result.Error
}
