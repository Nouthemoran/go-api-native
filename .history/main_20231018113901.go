package main

import (
	"fmt"
	"go-api-native/config"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.LoadConfig()
	

	r := mux.NewRouter()

	log.Println("Server running on port 9000")
	http.ListenAndServe(fmt.Sprintf(":%v", ":9000"), r)
}
