package controllers

import (
	"encoding/json"
	"go-api-native/config"
	"go-api-native/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var author []models.Author

	if err := config.DB.Find(&author).Error; err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	:= json.Marshal(author)
}
