package routes

func AuthorRoutes(r *mux.Router) {
	router := PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authcontrol)
}