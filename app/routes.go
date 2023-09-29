package app

import (
	"fmt"
	"net/http"
)

func (a *App) Routes() {
	a.Router.HandleFunc("/", index())
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Welcome to the go ecommerce")
	}
}
