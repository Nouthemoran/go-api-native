package routes

import "github.com/gorilla/mux"

func AuthorRoutes(r *mux.Router) {
	router := PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authcontroller.Index).Methods("GET")

}
