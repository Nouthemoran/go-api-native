package controllers

import (
	"encoding/json"
	"go-api-native/config"
	"go-api-native/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var authors []models.Author

	err := config.DB.Find(&authors).Error
	if err != nil {
		w.WriteHeader(500)
		res, _ := json.Marshal(map[sstring]string{"status": "failed"})
		w.Write(res)
	}

	res, _ := json.Marshal(authors)
	w.Write(res)
}
