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
	config.ConnectDB()

	r := mux.NewRouter()

	log.Println("Server running on port")
	http.ListenAndServe(fmt.Sprintf(":%v", ":9000"), r)
}
