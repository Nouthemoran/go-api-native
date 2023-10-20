package controllers

import (
	"encoding/json"
	"go-api-native/config"
	"go-api-native/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []models.Author

	err := config.DB.Find(&users).Error
	if err != nil {
		w.WriteHeader(500)
		res, _ := json.Marshal(map[string]string{"status": "failed"})
		w.Write(res)
	}

	res, _ := json.Marshal(models.Author)
	w.Write(res)
}
