package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	log.Println("Ser")
	http.ListenAndServe(fmt.Sprintf(":%v", ":9000"), r)
}
