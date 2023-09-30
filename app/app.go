package app

import (
	"GoECommerce/models"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	Db     *gorm.DB
	Router *mux.Router
}

func New() (*App, error) {
	dsn := "host=localhost user=dev password=abcde dbname=GoECommerce sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	_ = db.AutoMigrate(&models.User{})
	_ = db.AutoMigrate(&models.Product{})
	_ = db.AutoMigrate(&models.Cart{})
	return &App{
		Router: mux.NewRouter(),
		Db:     db,
	}, nil
}

func (a *App) Run() {
	fmt.Println("Serving on port 8000")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
