package app

import (
	"GoECommerce/auth"
	"GoECommerce/controllers"
	"fmt"
	"net/http"
)

func (a *App) Routes() {
	a.Router.HandleFunc("/", index())
	a.Router.HandleFunc("/login", controllers.Login(a.Db)).Methods("POST")
	a.Router.HandleFunc("/register", controllers.Register(a.Db)).Methods("POST")

	a.ProductRoutes()
}

func (a *App) ProductRoutes() {
	a.Router.HandleFunc("/api/product", auth.VerifyJWT(controllers.GetAllProduct(a.Db))).Methods("GET")
	a.Router.HandleFunc("/api/product/in-stock", auth.VerifyJWT(controllers.GetAllProductInStock(a.Db))).Methods("GET")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", auth.VerifyJWT(controllers.GetProductById(a.Db))).Methods("GET")
	a.Router.HandleFunc("/api/product", auth.VerifyJWT(controllers.CreateProduct(a.Db))).Methods("POST")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", auth.VerifyJWT(controllers.UpdateProduct(a.Db))).Methods("PUT")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", auth.VerifyJWT(controllers.DeleteProduct(a.Db))).Methods("DELETE")
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the go ecommerce")
	}
}
