package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.ListenAndServe(fmt.Sprintf(":%v", ":9000"), r)
}
