package routes

import (
	"github.com/gorilla/mux"

	"go-api-native/controllers/author"

)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods("GET")

}
