package main

import (
	"net/http"

	"github.com/Drake0211/ORM_API_REST/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8080", r)
}
