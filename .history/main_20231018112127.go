package main

import (
	"fmt"
	log
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	log.Println("Server running on port 9000")
	http.ListenAndServe(fmt.Sprintf(":%v", ":9000"), r)
}
