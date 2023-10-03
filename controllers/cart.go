package controllers

import (
	"GoECommerce/models"
	"GoECommerce/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetUserCart(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, err := strconv.ParseUint(vars["userId"], 10, 32)
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		u := &models.User{}
		err = u.GetById(db, uint(userId))
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, u.Cart, http.StatusOK)
	}
}

func AddItemToCard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, err := strconv.ParseUint(vars["userId"], 10, 32)
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		u := &models.User{}
		err = u.GetById(db, uint(userId))
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		c := &models.CartJson{}
		err = json.NewDecoder(r.Body).Decode(c)
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		ci := models.NewCartItem(c.ProductId, c.Quantity)
		err = u.AddToCart(db, ci)
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, &utils.Response{Msg: "Item added to cart"}, http.StatusOK)
	}
}

func RemoveFromCart(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, err := strconv.ParseUint(vars["userId"], 10, 32)
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		u := &models.User{}
		err = u.GetById(db, uint(userId))
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		itemId, err := strconv.ParseUint(vars["itemId"], 10, 32)
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		item := &models.CartItem{}
		err = item.GetById(db, uint(itemId))
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		err = u.Cart.RemoveFromCart(db, item)
		if err != nil {
			utils.Respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, &utils.Response{Msg: "Item removed successfully"}, http.StatusOK)
	}
}
