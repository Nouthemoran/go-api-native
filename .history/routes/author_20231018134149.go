package routes

import (
	"go-api-native/controllers/author/auth"
	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authcontroller.Index).Methods("GET")

}
