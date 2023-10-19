package routes

import (
	"github.com/gorilla/mux"

	"go-api-native/controllers/authorcontroller"

)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", author).Methods("GET")

}
