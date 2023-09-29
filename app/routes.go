package app

import (
	"GoECommerce/controllers"
	"fmt"
	"net/http"
)

func (a *App) Routes() {
	a.Router.HandleFunc("/", index())
	a.Router.HandleFunc("/login", controllers.Login(a.Db)).Methods("POST")
	a.Router.HandleFunc("/register", controllers.Register(a.Db)).Methods("POST")
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the go ecommerce")
	}
}
