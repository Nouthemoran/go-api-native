package config

import (
	"fmt"
	"go-api-native/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connection := fmt.Sprintf("%v:%v@/%v?parseTime=true", ENV.DB_USER, ENV.DB_PASSW, ENV.DB_DATABASE)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Author{})

	DB = db
	log.Println("Database connected")
}
