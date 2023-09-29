package controllers

import (
	"GoECommerce/auth"
	"GoECommerce/models"
	"GoECommerce/utils"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &models.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		u := &models.User{}
		err = u.GetByUsername(db, uJson.Username)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		if !CheckPasswordHash(uJson.Password, u.Password) {
			utils.Respond(w, r, &utils.Response{Msg: "Wrong password or username"}, http.StatusUnauthorized)
			return
		}
		token, err := auth.CreateJWTToken(int(u.ID), u.Username, u.IsAdmin)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.ResponseToken{Token: token}, http.StatusOK)
	}
}

func Register(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uJson := &models.UserJson{}
		err := json.NewDecoder(r.Body).Decode(uJson)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		if uJson.Username == "" || uJson.Password == "" {
			utils.Respond(w, r, &utils.Response{Msg: "Username and password are required"}, http.StatusBadRequest)
			return
		}
		hash, err := HashPassword(uJson.Password)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		u := &models.User{
			Username: uJson.Username,
			Password: hash,
			IsAdmin:  uJson.IsAdmin,
		}
		err = u.Create(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.Respond(w, r, &utils.Response{Msg: "User created successfully"}, http.StatusOK)
	}
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
