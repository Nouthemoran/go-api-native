package routes

import (
"go-api-native/controllers/authorcontroller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", controllers.Index).Methods("GET") // Menggunakan "author" yang diimpor
}
