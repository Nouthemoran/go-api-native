package main

import (
	"fmt"
	"go-api-native/config"
	"go-api-native/routes"
	"net/http"

	log "github.com/sirupsen/logrus"
	"go-api-nativ"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	db.AutoMigrate(&Author{})

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}