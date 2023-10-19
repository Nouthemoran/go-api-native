package routes

import "github.com/gorilla/mux"

func AuthorRoutes(r *mux.Router) {
	router := r("/authors").Subrouter()

	router.HandleFunc("", authcontroller.Index).Methods("GET")

}
