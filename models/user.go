package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password []byte
	IsAdmin  bool
}

type UserJson struct {
	Username string
	Password string
	IsAdmin  bool
}

func (u *User) Create(db *gorm.DB) error {
	result := db.Create(u)
	return result.Error
}

func (u *User) GetByUsername(db *gorm.DB, username string) error {
	result := db.First(u, "username = ?", username)
	return result.Error
}
