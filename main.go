package main

import (
	"GoECommerce/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	a.Routes()
	a.Run()
}
