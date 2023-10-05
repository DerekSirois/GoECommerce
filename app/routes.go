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
	a.CartRoutes()
}

func (a *App) ProductRoutes() {
	a.Router.HandleFunc("/api/product", auth.VerifyJWT(controllers.GetAllProduct(a.Db), true)).Methods("GET")
	a.Router.HandleFunc("/api/product/in-stock", auth.VerifyJWT(controllers.GetAllProductInStock(a.Db), false)).Methods("GET")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", auth.VerifyJWT(controllers.GetProductById(a.Db), false)).Methods("GET")
	a.Router.HandleFunc("/api/product", auth.VerifyJWT(controllers.CreateProduct(a.Db), true)).Methods("POST")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", auth.VerifyJWT(controllers.UpdateProduct(a.Db), true)).Methods("PUT")
	a.Router.HandleFunc("/api/product/{id:[0-9]+}", auth.VerifyJWT(controllers.DeleteProduct(a.Db), true)).Methods("DELETE")
}

func (a *App) CartRoutes() {
	a.Router.HandleFunc("/api/cart/{userId:[0-9]+}", auth.VerifyJWT(controllers.GetUserCart(a.Db), false)).Methods("GET")
	a.Router.HandleFunc("/api/cart/{userId:[0-9]+}", auth.VerifyJWT(controllers.AddItemToCard(a.Db), false)).Methods("POST")
	a.Router.HandleFunc("/api/cart/{userId:[0-9]+}/{itemId:[0-9]+}", auth.VerifyJWT(controllers.RemoveFromCart(a.Db), false)).Methods("DELETE")
	a.Router.HandleFunc("/api/cart/checkout/{id:[0-9]+}", auth.VerifyJWT(controllers.Checkout(a.Db), false)).Methods("POST")
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the go ecommerce")
	}
}
