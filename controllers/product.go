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

func GetAllProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := models.GetAllProduct(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, p, http.StatusOK)
	}
}

func GetAllProductInStock(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := models.GetAllProductInStock(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, p, http.StatusOK)
	}
}

func GetProductById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		p := &models.Product{}
		err = p.GetById(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, p, http.StatusOK)
	}
}

func CreateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &models.Product{}
		err := json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		err = p.Create(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, &utils.Response{Msg: "Product created successfully"}, http.StatusCreated)
	}
}

func UpdateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		p := &models.Product{}
		err = json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}
		p.ID = uint(id)

		err = p.Update(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, &utils.Response{Msg: "Product updated successfully"}, http.StatusCreated)
	}
}

func DeleteProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		p := &models.Product{}
		err = p.GetById(db, uint(id))
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}

		err = p.Delete(db)
		if err != nil {
			utils.Respond(w, r, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.Respond(w, r, &utils.Response{Msg: "Product deleted successfully"}, http.StatusCreated)
	}
}
