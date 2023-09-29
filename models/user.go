package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password []byte
}

type UserJson struct {
	Username string
	Password string
}

func (u *User) Create(db *gorm.DB) error {
	result := db.Create(u)
	return result.Error
}

func (u *User) GetByUsername(db *gorm.DB, username string) error {
	result := db.First(u, "username = ?", username)
	return result.Error
}