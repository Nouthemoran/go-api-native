package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	
	r := mux.NewRouter()

	log.Println("Server running on port 9000")
	http.ListenAndServe(fmt.Sprintf(":%v", ":9000"), r)
}
