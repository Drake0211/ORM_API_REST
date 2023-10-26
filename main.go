package main

import (
	"net/http"

	"github.com/Drake0211/ORM_API_REST/db"
	"github.com/Drake0211/ORM_API_REST/models"
	"github.com/Drake0211/ORM_API_REST/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBconection()

	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8080", r)
}
