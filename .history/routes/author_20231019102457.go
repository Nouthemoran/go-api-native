package routes

import (
    "go-api-native/controllers/author"
	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authorontroller.Index).Methods("GET") // Menggunakan "author" yang diimpor
}

